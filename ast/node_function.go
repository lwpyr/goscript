package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/instruction"
)

type FunctionDefNode struct {
	Node
	FuncName  string
	Block     ASTNode
	InParam   []ASTNode
	OutParam  []ASTNode
	TailArray bool

	Function *common.Instruction
}

type ParamNode struct {
	Node
	Symbol   string
	TypeNode ASTNode
}

type TypeNodeType int

const (
	SimpleType = iota
	SliceType
	MapType
	ChanType
	FuncType
)

type TypeNode struct {
	Node
	SimpleTypeName string
	TypeNodeType   TypeNodeType
	Key            *common.DataType // map
	Value          ASTNode          // map slice chan
	InParam        []ASTNode        // func
	OutParam       []ASTNode        // func
	TailArray      bool
}

type ConvertNode struct {
	Node
	Type  ASTNode
	Value ASTNode
}

func (b *ConvertNode) CheckIsConstant() {
	b.Variadic = true
	if !b.Value.IsVariadic() {
		b.Variadic = false
	}
}

func (b *ConvertNode) Compile(c *CompileContext) {
	b.Type.Compile(c)
	b.DataType = b.Type.GetDataType()
	b.Value.Compile(c)
	convertInstruction := instruction.TypeConvert(b.Value.GetDataType(), b.DataType)
	if common.IsError(convertInstruction) {
		panic(common.NewCompileErr(b.ErrorWithSource("type convert error")))
	}
	b.AppendInstruction(b.Value.GetInstructions()...)
	if convertInstruction != nil {
		b.AppendInstruction(convertInstruction)
	}

	b.CheckIsConstant()
	b.PostProcess()
	b.StackIncrement = 1
}

func (p *ParamNode) Compile(c *CompileContext) {
	p.TypeNode.Compile(c)
	p.DataType = p.TypeNode.GetDataType()
}

func (t *TypeNode) Compile(c *CompileContext) {
	var dType *common.DataType
	switch t.TypeNodeType {
	case SimpleType:
		dType = c.FindType(t.SimpleTypeName)
	case SliceType:
		t.Value.Compile(c)
		ItemType := t.Value.GetDataType()
		dType = c.TypeRegistry.FindSliceType(ItemType.Type)
	case ChanType:
		t.Value.Compile(c)
		ItemType := t.Value.GetDataType()
		dType = c.TypeRegistry.FindChanType(ItemType.Type)
	case MapType:
		t.Value.Compile(c)
		ValType := t.Value.GetDataType()
		dType = c.TypeRegistry.FindMapType(t.Key.Type, ValType.Type)
	case FuncType:
		meta := &common.FunctionMeta{
			TailArray: t.TailArray,
			ConstExpr: false,
		}
		for _, in := range t.InParam {
			in.Compile(c)
			meta.In = append(meta.In, in.GetDataType())
		}
		for _, out := range t.OutParam {
			out.Compile(c)
			meta.Out = append(meta.Out, out.GetDataType())
		}
		dType = c.TypeRegistry.FindFuncType(meta)
	default:
		panic(common.NewCompileErr(t.ErrorWithSource("bad node type, this should never happened because of the enum limitation")))
	}
	if dType == nil {
		panic(common.NewTypeErr(t.ErrorWithSource("type not found")))
	}
	t.DataType = dType
}

func (f *FunctionDefNode) Compile(c *CompileContext) {
	c.MakeFunctionScope()
	defer c.ReturnParentScope()

	f.DataType = &common.DataType{
		LambdaMeta: &common.FunctionMeta{
			TailArray: f.TailArray,
		},
	}
	funcPlaceHolder := common.Instruction(nil)
	f.Function = &funcPlaceHolder

	meta := f.DataType.LambdaMeta

	for idx, in := range f.InParam {
		in.Compile(c)
		inNode := in.(*ParamNode)
		meta.In = append(meta.In, in.GetDataType())

		if idx == len(f.InParam)-1 && meta.TailArray {
			// symbol's type should be array if there is a tail array sign
			inNode.DataType = c.TypeRegistry.FindSliceType(inNode.DataType.Type)
		}
		c.Scope.AddLocalVariable(&common.Symbol{
			Symbol:   inNode.Symbol,
			DataType: inNode.DataType,
		})
	}
	for _, out := range f.OutParam {
		out.Compile(c)
		if outNode, ok := out.(*ParamNode); ok {
			c.Scope.AddReturnVariable(&common.Symbol{
				Symbol:   outNode.Symbol,
				DataType: outNode.DataType,
			})
		}
		meta.Out = append(meta.Out, out.GetDataType())
	}

	f.DataType = c.FindFuncType(f.DataType.LambdaMeta)

	c.PushFunctionMeta(f.DataType.LambdaMeta)
	defer c.PopFunctionMeta()

	if f.FuncName != "" {
		f.Variadic = false
		// named function
		c.Scope.Outer.AddConstant(f.FuncName, &common.Symbol{
			Symbol:   f.FuncName,
			DataType: f.DataType,
			Data:     f.Function,
		})

		f.Block.Compile(c)
		blockInst := f.Block.GetInstructions()
		blockInst = append(blockInst, func(m *common.Memory, stk *common.Stack) {
			stk.Pc = -1
		})
		*f.Function = func(m *common.Memory, stk *common.Stack) {
			for stk.Pc != -1 {
				blockInst[stk.Pc](m, stk)
			}
		}
	} else {
		// lambda function
		c.Scope.AddLocalVariable(&common.Symbol{
			Symbol:     "#",
			SymbolType: common.Local,
			Scope:      c.Scope,
		})

		// let the hunt begin!
		c.Scope.SetCaptureMode()

		f.Block.Compile(c)
		capturedVariables := c.Scope.Capture
		closureVariable := c.Scope.GetSymbol("#")
		if len(capturedVariables) > 0 {
			// oh, we capture some outer local variables
			f.Variadic = true
			rawFunc := instruction.ConnectInstructions(f.Block.GetInstructions())
			f.AppendInstruction(instruction.PushLambdaWithCapture(rawFunc, capturedVariables, closureVariable))
		} else {
			// sad, nothing captured
			f.Variadic = false
			lambda := instruction.ConnectInstructions(f.Block.GetInstructions())
			f.AppendInstruction(instruction.PushLambda(lambda))
		}
		f.StackIncrement = 1
	}
}

type FunctionCallNode struct {
	Node
	Function ASTNode
	Params   []ASTNode
	Meta     *common.FunctionMeta
}

func (f *FunctionCallNode) Compile(c *CompileContext) {
	f.Function.Compile(c)
	if f.Function.GetDataType().Kind.Kind != common.Func {
		panic(common.NewCompileErr(f.ErrorWithSource("not a function")))
	}

	meta := f.Function.GetDataType().LambdaMeta
	f.Meta = meta

	if len(meta.Out) == 1 {
		f.DataType = meta.Out[0]
	} else {
		f.DataType = common.BasicTypeMap["nil"]
	}

	num := len(f.Params)
	if num > len(meta.In) && meta.TailArray == false || num < len(meta.In) {
		panic(common.NewCompileErr(f.ErrorWithSource("parameter given mismatch with needed")))
	}

	if len(meta.Out) > 0 {
		f.AppendInstruction(instruction.GetAppendReturnPlace(len(meta.Out)))
	}
	f.StackIncrement = len(meta.Out)

	if f.Meta.TailArray {
		for i, param := range f.Params {
			idx := i
			if idx >= len(meta.In)-1 {
				idx = len(meta.In) - 1
			}
			param.SetRequiredType(meta.In[idx])
			param.Compile(c)
			f.AppendInstruction(param.GetInstructions()...)
		}
		f.AppendInstruction(instruction.GetTailArray(num - len(meta.In) + 1))
	} else {
		for i, param := range f.Params {
			param.SetRequiredType(meta.In[i])
			param.Compile(c)
			f.AppendInstruction(param.GetInstructions()...)
		}
	}

	if f.Function.IsVariadic() {
		f.AppendInstruction(f.Function.GetInstructions()...)
		f.AppendInstruction(instruction.GetDynamicCallFunc(len(meta.In)))
	} else {
		funcPtr := f.Function.GetConstantValue().(*common.Instruction)
		f.AppendInstruction(instruction.GetStaticCallFunc(funcPtr, len(meta.In)))
	}

	f.PostProcess()
}

type BuiltinFunctionNode struct {
	Node
	Params      []ASTNode
	BuiltinName string
	Aux         interface{}
}

func (b *BuiltinFunctionNode) Compile(c *CompileContext) {
	switch b.BuiltinName {
	case "pushBack", "pushFront":
		b.Variadic = true
		b.DataType = common.BasicTypeMap["illegal"]
		if len(b.Params) != 2 {
			panic(common.NewCompileErr(b.ErrorWithSource("pushBack/pushFront has two parameter")))
		}
		b.Params[0].SetLhs()
		b.Params[0].Compile(c)
		sliceType := b.Params[0].GetDataType()
		b.AppendInstruction(b.Params[0].GetInstructions()...)
		if sliceType.Kind.Kind != common.Slice {
			panic(common.NewCompileErr(b.ErrorWithSource("slice operation should have slice value as its first parameter")))
		}
		b.AppendInstruction()
		b.Params[1].SetRequiredType(sliceType.ItemType)
		b.Params[1].Compile(c)
		if b.Params[0].IsStackPtr() {
			b.AppendInstruction(instruction.StackOffsetToStackPtr(1))
		}
		b.AppendInstruction(b.Params[1].GetInstructions()...)
		if b.BuiltinName == "pushBack" {
			b.AppendInstruction(instruction.PushBack())
		} else {
			b.AppendInstruction(instruction.PushFront())
		}
	case "delete":
		b.Variadic = true
		b.DataType = common.BasicTypeMap["illegal"]
		if len(b.Params) != 2 {
			panic(common.NewCompileErr(b.ErrorWithSource("delete has two parameter")))
		}
		m := b.Params[0]
		key := b.Params[1]
		m.Compile(c)
		mType := m.GetDataType()
		b.AppendInstruction(m.GetInstructions()...)

		switch m.GetDataType().Kind.Kind {
		case common.Message:
			key.SetDataType(common.BasicTypeMap[common.StringType])
			key.Compile(c)
			b.AppendInstruction(key.GetInstructions()...)
			b.AppendInstruction(instruction.DeleteMessageField())
		case common.Map:
			key.SetDataType(mType.KeyType)
			key.Compile(c)
			b.AppendInstruction(key.GetInstructions()...)
			mapDelFunc := instruction.MapDelete(key.GetDataType())
			b.AppendInstruction(instruction.DeleteMapKey(mapDelFunc))
		default:
			panic(common.NewCompileErr(b.ErrorWithSource("delete not allowed on " + m.GetDataType().Type)))
		}
	case "len":
		b.Variadic = true
		b.DataType = common.BasicTypeMap[common.Int64Type]
		if len(b.Params) != 1 {
			panic(common.NewCompileErr(b.ErrorWithSource("len has one parameter")))
		}
		val := b.Params[0]
		val.Compile(c)
		b.AppendInstruction(val.GetInstructions()...)
		if val.GetDataType().Kind.Kind != common.Map && val.GetDataType().Kind.Kind != common.Slice && val.GetDataType().Kind.Kind != common.String && val.GetDataType().Kind.Kind != common.Bytes {
			panic(common.NewTypeErr(b.ErrorWithSource("len parameter could be map, slice, string or map")))
		}
		lengthFunc := instruction.GetLengthFunc(val.GetDataType())
		b.AppendInstruction(instruction.Len(lengthFunc))
	case "enumString":
		b.Variadic = true
		b.DataType = common.BasicTypeMap[common.StringType]
		if len(b.Params) != 1 {
			panic(common.NewCompileErr(b.ErrorWithSource("enumString has one parameter")))
		}
		val := b.Params[0]
		val.Compile(c)
		b.AppendInstruction(val.GetInstructions()...)
		if val.GetDataType().Kind.Kind != common.Int32 && val.GetDataType().Type != "int32" {
			panic(common.NewTypeErr(b.ErrorWithSource("enumName can only be applied on an enum value")))
		}
		rEnum := c.TypeRegistry.GetREnums(val.GetDataType().Type)
		b.AppendInstruction(instruction.EnumString(rEnum))
	case "typeof":
		b.Variadic = false
		b.DataType = common.BasicTypeMap["reflect"]
		if len(b.Params) != 1 {
			panic(common.NewCompileErr(b.ErrorWithSource("typeof has one parameter")))
		}
		val := b.Params[0]
		val.Compile(c)
		b.AppendInstruction(instruction.PushConstantToStack(val.GetDataType()))
	default:
		panic(common.NewCompileErr(b.ErrorWithSource("unknown builtin function")))
	}

	b.PostProcess()
	b.StackIncrement = 1
}
