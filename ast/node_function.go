package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/lambda_chains"
)

type LambdaDefNode struct {
	Node
	Block ASTNode
	Scope *common.Scope
}

type FunctionDefNode struct {
	Node
	FuncName string
	Block    ASTNode
	Function *common.Function
}

type ParamNode struct {
	Node
	Symbol string
}

func (p *ParamNode) Compile(_ *Compiler) {
	panic("should never reach here, this must be internal error")
}

type TypeNode struct {
	Node
	dType *common.DataType
}

func (t *TypeNode) Compile(c *Compiler) {
	panic("should never reach here, this must be internal error")
}

type TypeNameNode struct {
	Node
	TypeName string
}

func (t *TypeNameNode) Compile(c *Compiler) {
	panic("should never reach here, this must be internal error")
}

func (f *FunctionDefNode) Compile(c *Compiler) {
	f.Block.Compile(c)
	blockInst := f.Block.GetInstructions()
	blockInst = append(blockInst, func(m *common.Memory, stk *common.Stack) {
		stk.Pc = -1
	})
	f.Function.F = func(m *common.Memory, stk *common.Stack) {
		for stk.Pc != -1 {
			blockInst[stk.Pc](m, stk)
		}
	}
}

type IFunctionNode interface {
	ASTNode
	GetReturnType() []*common.DataType
}

type FunctionCallNode struct {
	Node
	Function ASTNode
	Params   []ASTNode
	Meta     *common.FunctionMeta
	Lambda   *common.Variable
}

func (f *FunctionCallNode) GetReturnType() []*common.DataType {
	return f.Meta.Out
}

func (f *LambdaDefNode) Compile(c *Compiler) {
	f.Block.Compile(c)
	capturedVariables := f.Scope.Capture
	if len(capturedVariables) > 0 {
		captureInstruction := func(m *common.Memory, stk *common.Stack) {
			captured := make([]interface{}, 0, len(capturedVariables))
			for _, v := range capturedVariables {
				captured = append(captured, stk.Get(v))
			}
			stk.Push(captured)
		}
		blockInst := f.Block.GetInstructions()
		blockInst = append(blockInst, func(m *common.Memory, stk *common.Stack) {
			stk.Pc = -1
		})
		rawFunc := func(m *common.Memory, stk *common.Stack) {
			for stk.Pc != -1 {
				blockInst[stk.Pc](m, stk)
			}
		}
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			captureInstruction(m, stk)
			captured := stk.Top().([]interface{})
			stk.Pop()
			lambda := &common.Function{
				Type: f.DataType,
				F: func(m *common.Memory, stk *common.Stack) {
					stk.Push(captured)
					rawFunc(m, stk)
				},
			}
			stk.Push(lambda)
		})
	} else {
		blockInst := f.Block.GetInstructions()
		blockInst = append(blockInst, func(m *common.Memory, stk *common.Stack) {
			stk.Pc = -1
		})
		lambda := &common.Function{
			Type: f.DataType,
			F: func(m *common.Memory, stk *common.Stack) {
				for stk.Pc != -1 {
					blockInst[stk.Pc](m, stk)
				}
			},
		}
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			stk.Push(lambda)
		})
	}
}

func (f *FunctionCallNode) Compile(c *Compiler) {
	num := len(f.Params)
	paramInstructions := make([]common.Instruction, 0, num)
	for _, param := range f.Params {
		param.Compile(c)
	}
	for i := 0; i < num; i++ {
		paramInstructions = append(paramInstructions, c.InstructionPop())
	}
	paramConvertFunc := make([]lambda_chains.TypeConvertFunc, 0, num)
	for i := num - 1; i >= 0; i-- {
		idx := num - i - 1
		if idx >= len(f.Meta.In) {
			idx = len(f.Meta.In) - 1
		}
		paramConvertFunc = append(paramConvertFunc, lambda_chains.GetConvertFunc(f.Params[i].GetDataType(), f.Meta.In[idx]))
	}
	lenOut := len(f.Meta.Out)
	if lenOut == 0 {
		lenOut = 1
	}
	if f.Function.IsVariadic() {
		f.Function.Compile(c)
		functionInstruction := c.InstructionPop()
		if f.Meta.TailArray {
			lenIn := len(f.Meta.In)
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				for i := 0; i < lenOut; i++ {
					stk.Push(nil)
				}
				for i := 0; i < lenIn-1; i++ {
					paramInstructions[i](m, stk)
					stk.Set(0, paramConvertFunc[i](stk.Top()))
				}
				tailArray := make([]interface{}, num-lenIn+1)
				for i := lenIn - 1; i < num; i++ {
					paramInstructions[i](m, stk)
					tailArray[i-lenIn+1] = paramConvertFunc[i](stk.Top())
					stk.Pop()
				}
				stk.Push(tailArray)
				functionInstruction(m, stk)
				function := stk.Top().(*common.Function)
				stk.Pop()
				stk.Call(function.F, m, stk, lenIn)
			})
		} else {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				for i := 0; i < lenOut; i++ {
					stk.Push(nil)
				}
				for i := 0; i < num; i++ {
					paramInstructions[i](m, stk)
					stk.Set(0, paramConvertFunc[i](stk.Top()))
				}
				functionInstruction(m, stk)
				function := stk.Top().(*common.Function)
				stk.Pop()
				stk.Call(function.F, m, stk, num)
			})
		}
	} else {
		function := f.Function.(IConstantNode).GetConstantValue().(*common.Function)
		if function.Type.LambdaMeta.TailArray {
			lenIn := len(f.Meta.In)
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				for i := 0; i < lenOut; i++ {
					stk.Push(nil)
				}
				for i := 0; i < lenIn-1; i++ {
					paramInstructions[i](m, stk)
					stk.Set(0, paramConvertFunc[i](stk.Top()))
				}
				tailArray := make([]interface{}, num-lenIn+1)
				for i := lenIn - 1; i < num; i++ {
					paramInstructions[i](m, stk)
					tailArray[i-lenIn+1] = paramConvertFunc[i](stk.Top())
					stk.Pop()
				}
				stk.Push(tailArray)
				stk.Call(function.F, m, stk, lenIn)
			})
		} else {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				for i := 0; i < lenOut; i++ {
					stk.Push(nil)
				}
				for i := 0; i < num; i++ {
					paramInstructions[i](m, stk)
					stk.Set(0, paramConvertFunc[i](stk.Top()))
				}
				stk.Call(function.F, m, stk, num)
			})
		}
	}
}

type BuiltinFunctionNode struct {
	Node
	Params      []ASTNode
	BuiltinName string
	Aux         interface{}
}

func (b *BuiltinFunctionNode) Compile(c *Compiler) {
	switch b.BuiltinName {
	case "uint8":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFunc32u(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "uint32":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFunc32u(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "uint64":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFunc64u(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "int32":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFunc32i(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "int64":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFunc64i(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "float32":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFunc32f(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "float64":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFunc64f(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "string":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFuncStr(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "bytes":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFuncBytes(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "bool":
		b.Params[0].Compile(c)
		valInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFuncBool(b.Params[0].GetDataType(), b.DataType)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, convertFunc(stk.Top()))
		})
	case "pushBack", "pushFront":
		b.Params[0].Compile(c)
		b.Params[1].SetLhs()
		b.Params[1].Compile(c)
		itemConvertFunc := lambda_chains.GetConvertFunc(b.Params[0].GetDataType(), b.Params[1].GetDataType().ItemType)
		sliceInst := c.InstructionPop()
		itemInst := c.InstructionPop()
		if b.BuiltinName == "pushBack" {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				sliceInst(m, stk)
				itemInst(m, stk)
				ptr := stk.TopIndex(1).(*interface{})
				if *ptr == nil {
					*ptr = []interface{}{}
				}
				*ptr = append((*ptr).([]interface{}), itemConvertFunc(stk.Top()))
				stk.Pop()
			})
		} else {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				sliceInst(m, stk)
				itemInst(m, stk)
				ptr := stk.TopIndex(1).(*interface{})
				if *ptr == nil {
					*ptr = []interface{}{}
				}
				*ptr = append([]interface{}{itemConvertFunc(stk.Top())}, (*ptr).([]interface{})...)
				stk.Pop()
			})
		}
	case "delete":
		m := b.Params[1]
		key := b.Params[0]
		key.Compile(c)
		m.Compile(c)
		mapInst := c.InstructionPop()
		keyInst := c.InstructionPop()

		switch m.GetDataType().Kind.Kind {
		case common.Message:
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				mapInst(m, stk)
				if stk.Top() != nil {
					keyInst(m, stk)
					lambda_chains.DelMapStrField(stk.TopIndex(1), stk.Top())
					stk.Pop()
				}
			})
		case common.Map:
			keyConvertFunc := lambda_chains.GetConvertFunc(key.GetDataType(), m.GetDataType().KeyType)
			mapDelFunc := lambda_chains.GetMapDelFunc(key.GetDataType())
			// todo: optimization use type specific convert function
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				mapInst(m, stk)
				if stk.Top() != nil {
					keyInst(m, stk)
					mapDelFunc(stk.TopIndex(1), keyConvertFunc(stk.Top()))
					stk.Pop()
				}
			})
		default:
			panic("delete not allowed on " + m.GetDataType().Type)
		}
	case "len":
		val := b.Params[0]
		val.Compile(c)
		valInst := c.InstructionPop()
		if val.GetDataType().Kind.Kind != common.Map && val.GetDataType().Kind.Kind != common.Slice && val.GetDataType().Kind.Kind != common.String && val.GetDataType().Kind.Kind != common.Bytes {
			panic("map delete error, first parameter should be a map, second parameter should be the key type of the first parameter")
		}
		lengthFunc := lambda_chains.GetLengthFunc(val.GetDataType())
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, lengthFunc(stk.Top()))
		})
	case "enumString":
		val := b.Params[0]
		val.Compile(c)
		valInst := c.InstructionPop()
		if val.GetDataType().Kind.Kind != common.Int32 && val.GetDataType().Type != "int32" {
			panic("enumName can only be applied on an enum value")
		}
		rEnum := c.TypeRegistry.GetREnums(val.GetDataType().Type)
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			valInst(m, stk)
			stk.Set(0, rEnum[stk.Top().(int32)])
		})
	}
}
