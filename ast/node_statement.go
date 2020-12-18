package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/instruction"
)

type FetchSymbolNode struct {
	Node
	SymbolName string
	Symbol     *common.Symbol
}

type AssertNode struct {
	Node
	Type  ASTNode
	Value ASTNode
}

type SelectorNode struct {
	Node
	Data      ASTNode // selector or var
	FieldName string
}

type SliceFilterNode struct {
	Node
	Slice ASTNode
	Expr  ASTNode
}

type SliceMultiIndexNode struct {
	Node
	Slice   ASTNode
	Indices []ASTNode
}

type MapMultiIndexNode struct {
	Node
	Map    ASTNode
	Fields []ASTNode
}

type IndexNode struct {
	Node
	ToIndex ASTNode
	Index   ASTNode
}

type IndicesNode struct {
	Node
	From ASTNode
	To   ASTNode
	Step ASTNode
}

type BinaryNode struct {
	Node
	Lhs ASTNode
	Rhs ASTNode
	Op  string
}

type LeftUnaryNode struct {
	Node
	Operand ASTNode
	Op      string
}

type RightUnaryNode struct {
	Node
	Operand ASTNode
	Op      string
}

type ConstructorNode struct {
	Node
	Params []ASTNode
	Type   ASTNode
}

type ValueNode struct {
	Node
}

type ChanSend struct {
	Node
	ToSend   ASTNode
	Chan     ASTNode
	NonBlock bool
}

type ChanRecv struct {
	Node
	Chan     ASTNode
	NonBlock bool
}

type InitializationSliceNode struct {
	Node
	Items []ASTNode
	Type  ASTNode
}

type InitializationKVNode struct {
	Node
	Keys   []ASTNode
	Values []ASTNode
	Type   ASTNode
}
type FunctionAssignNode struct {
	Node
	Lhs      []ASTNode
	Function ASTNode
}

type AssignNode struct {
	Node
	Lhs ASTNode
	Rhs ASTNode
	Op  string
}

func (v *FetchSymbolNode) CheckIsConstant() {
	v.Variadic = true
	if v.Symbol.SymbolType == common.Const {
		v.Variadic = false
	}
}

func (v *FetchSymbolNode) Compile(c *Compiler) {
	symbol := c.Scope.GetSymbol(v.SymbolName)
	if symbol == nil {
		panic(common.NewSymbolErr(v.ErrorWithSource("unknown symbol")))
	}
	v.Symbol = symbol
	var inst common.Instruction
	if v.LhsFlag {
		inst = instruction.GetFetchSymbolFuncLhs(symbol, c.Scope)
	} else {
		inst = instruction.GetFetchSymbolFunc(symbol, c.Scope)
	}
	if common.IsError(inst) {
		panic(common.NewCompileErr(v.ErrorWithSource("unknown symbol type")))
	}
	v.AppendInstruction(inst)
	v.DataType = symbol.DataType
	if v.Symbol.SymbolType == common.Local && v.LhsFlag {
		v.StackPtr = true
	}

	// do optimize
	v.CheckIsConstant()
	v.PostProcess()
	v.StackIncrement = 1
}

func (a *AssertNode) CheckIsConstant() {
	a.Variadic = true
	if !a.Value.IsVariadic() {
		a.Variadic = false
	}
}

func (a *AssertNode) Compile(c *Compiler) {
	a.Type.Compile(c)
	a.DataType = a.Type.GetDataType()

	a.Value.Compile(c)
	if a.Value.GetDataType().Kind.Kind != common.Any {
		panic(common.NewTypeErr(a.ErrorWithSource("assert should have 'any' type value")))
	}

	a.AppendInstruction(a.Value.GetInstructions()...)
	convertInstruction := instruction.GetConvertInstruction(a.Value.GetDataType(), a.DataType)
	if common.IsError(convertInstruction) {
		panic(common.NewTypeErr(a.ErrorWithSource("cannot convert")))
	}
	if convertInstruction != nil {
		a.AppendInstruction(convertInstruction)
	}

	a.CheckIsConstant()
	a.PostProcess()
	a.StackIncrement = 1
}

func (s *SelectorNode) CheckIsConstant() {
	s.Variadic = true
	if !s.Data.IsVariadic() {
		s.Variadic = false
	}
}

func (s *SelectorNode) Compile(c *Compiler) {
	if s.LhsFlag {
		s.Data.SetLhs()
	}
	s.Data.Compile(c)
	s.AppendInstruction(s.Data.GetInstructions()...)
	if s.Data.GetDataType().Kind.Kind == common.Message {
		fieldName := s.FieldName
		constructor := s.Data.GetDataType().Constructor
		fieldType := s.Data.GetDataType().FieldType[fieldName]
		if fieldType == nil {
			panic(common.NewTypeErr(s.ErrorWithSource("type does not have the field ")))
		}
		var oneOfs []string
		if OneOfGroupName, ok := s.Data.GetDataType().OneOf[fieldName]; ok {
			tempOneOfs := s.Data.GetDataType().OneOfGroup[OneOfGroupName]
			for _, fn := range tempOneOfs {
				if fn != fieldName {
					oneOfs = append(oneOfs, fn)
				}
			}
		}
		var selectInstruction common.Instruction
		if s.LhsFlag {
			selectInstruction = instruction.GetSelectFuncLhs(constructor, fieldName, oneOfs)
		} else {
			selectInstruction = instruction.GetSelectFunc(fieldName)
		}
		if s.Data.IsStackPtr() {
			s.AppendInstruction(instruction.GetStackOffsetToStackPtr(0))
		}
		s.AppendInstruction(selectInstruction)

		s.DataType = fieldType
	} else if s.Data.GetDataType().Kind.Kind == common.Enum {
		s.AppendInstruction(instruction.GetSelectEnumFunc(s.FieldName))
		s.DataType = c.FindType(s.Data.GetDataType().Type)
	} else {
		panic(common.NewTypeErr(s.ErrorWithSource("selector follow a non message type")))
	}

	s.CheckIsConstant()
	s.PostProcess()
	s.StackIncrement = 1
}

func (s *SliceFilterNode) CheckIsConstant() {
	s.Variadic = true
	// todo: function constexpr check
}

func (s *SliceFilterNode) Compile(c *Compiler) {
	if s.LhsFlag {
		panic(common.NewCompileErr(s.ErrorWithSource("cannot use slice-filter an array on the left hand side")))
	}

	s.Slice.Compile(c)
	if s.Slice.GetDataType().Kind.Kind != common.Slice {
		panic(common.NewTypeErr(s.ErrorWithSource("try to filter a non slice data")))
	}
	s.AppendInstruction(s.Slice.GetInstructions()...)

	c.MakeChildScope()
	defer c.ReturnParentScope()

	c.Scope.AddLocalVariable(&common.Symbol{
		Symbol:   "@",
		DataType: s.Slice.GetDataType().ItemType,
	})

	s.Expr.Compile(c)
	exprInstruction := s.Expr.GetInstructions()
	exprInstruction = append(exprInstruction, instruction.StackReturn())
	filter := func(m *common.Memory, stk *common.Stack) {
		for stk.Pc != -1 {
			exprInstruction[stk.Pc](m, stk)
		}
	}

	boolConvertFunc, err := instruction.GetConvertFuncBool(s.Expr.GetDataType())
	if err != nil {
		panic(common.NewTypeErr(s.ErrorWithSource("filter cannot return a bool-convertible value")))
	}
	s.AppendInstruction(instruction.GetSliceFilterFunc(filter, boolConvertFunc))

	s.DataType = s.Slice.GetDataType()
	s.CheckIsConstant()
	s.PostProcess()
	s.StackIncrement = 1
}

func (s *SliceMultiIndexNode) CheckIsConstant() {
	s.Variadic = true
	if !s.Slice.IsVariadic() {
		flag := true
		for _, index := range s.Indices {
			if index.IsVariadic() {
				flag = false
				break
			}
		}
		if flag {
			s.Variadic = false
		}
	}
}

func (s *SliceMultiIndexNode) Compile(c *Compiler) {
	if s.LhsFlag {
		panic(common.NewTypeErr(s.ErrorWithSource("cannot use slice-index an array on the left hand side")))
	}

	numIndex := len(s.Indices)
	for _, index := range s.Indices {
		index.Compile(c)
		s.AppendInstruction(index.GetInstructions()...)
	}

	s.Slice.Compile(c)
	s.AppendInstruction(s.Slice.GetInstructions()...)

	if s.Slice.GetDataType().Kind.Kind == common.Slice {
		s.AppendInstruction(instruction.GetSliceMultiIndexFunc(numIndex))
	} else if s.Slice.GetDataType().Kind.Kind == common.String {
		if numIndex > 1 {
			panic(common.NewIndexErr(s.ErrorWithSource("illegal string index")))
		}
		s.AppendInstruction(instruction.GetStringMultiIndexFunc())
	} else {
		panic(common.NewTypeErr(s.ErrorWithSource("try to index a value which cannot be indexed")))
	}

	s.DataType = s.Slice.GetDataType()
	s.CheckIsConstant()
	s.PostProcess()
	s.StackIncrement = 1
}

func (m *MapMultiIndexNode) CheckIsConstant() {
	m.Variadic = true
	if !m.Map.IsVariadic() {
		flag := true
		for _, field := range m.Fields {
			if field.IsVariadic() {
				flag = false
				break
			}
		}
		if flag {
			m.Variadic = false
		}
	}
}

func (m *MapMultiIndexNode) Compile(c *Compiler) {
	if m.LhsFlag {
		panic(common.NewCompileErr(m.ErrorWithSource("cannot use slice-index a map on the left hand side")))
	}

	m.Map.Compile(c)
	if m.Map.GetDataType().Kind.Kind != common.Map {
		panic(common.NewTypeErr(m.ErrorWithSource("try to multi-index a none map value")))
	}
	numFields := len(m.Fields)
	for _, field := range m.Fields {
		field.SetRequiredType(m.Map.GetDataType().KeyType)
		field.Compile(c)
		m.AppendInstruction(field.GetInstructions()...)
	}

	m.AppendInstruction(m.Map.GetInstructions()...)

	mapGet := instruction.GetMapGetFunc(m.Map.GetDataType().KeyType)
	m.AppendInstruction(instruction.GetMapMultiIndexFunc(mapGet, numFields))

	m.DataType = c.FindSliceType(m.Map.GetDataType().ValueType.Type)
	m.CheckIsConstant()
	m.PostProcess()
	m.StackIncrement = 1
}

func (i *IndexNode) CheckIsConstant() {
	i.Variadic = true
	if !i.Index.IsVariadic() && !i.ToIndex.IsVariadic() {
		i.Variadic = false
	}
}

func (i *IndexNode) Compile(c *Compiler) {
	if i.LhsFlag {
		i.ToIndex.SetLhs()
	}
	i.ToIndex.Compile(c)

	var indexInstruction common.Instruction

	switch i.ToIndex.GetDataType().Kind.Kind {
	case common.Slice:
		i.Index.SetRequiredType(common.BasicTypeMap[common.Int64Type])
		i.Index.Compile(c)

		if i.LhsFlag {
			indexInstruction = instruction.GetSliceIndexFuncLhs()
		} else {
			indexInstruction = instruction.GetSliceIndexFunc()
		}
		i.DataType = i.ToIndex.GetDataType().ItemType
	case common.String:
		i.Index.SetRequiredType(common.BasicTypeMap[common.Int64Type])
		i.Index.Compile(c)
		if i.LhsFlag {
			panic(common.NewCompileErr(i.ErrorWithSource("string is immutable")))
		}
		indexInstruction = instruction.GetStringIndexFunc()
		i.DataType = i.ToIndex.GetDataType()
	case common.Map:
		i.Index.SetRequiredType(i.ToIndex.GetDataType().KeyType)
		i.Index.Compile(c)
		constructor := i.ToIndex.GetDataType().Constructor

		if i.LhsFlag {
			mapMustGet := instruction.GetMapMustGetFunc(i.ToIndex.GetDataType().KeyType)
			indexInstruction = instruction.GetMapIndexFuncLhs(constructor, mapMustGet)
		} else {
			mapGet := instruction.GetMapGetFunc(i.ToIndex.GetDataType().KeyType)
			indexInstruction = instruction.GetMapIndexFunc(mapGet)
		}
		i.DataType = i.ToIndex.GetDataType().ValueType
	default:
		panic(common.NewCompileErr(i.ErrorWithSource("try to index a value which cannot be indexed")))
	}

	i.AppendInstruction(i.Index.GetInstructions()...)
	i.AppendInstruction(i.ToIndex.GetInstructions()...)
	if i.ToIndex.IsStackPtr() {
		i.AppendInstruction(instruction.GetStackOffsetToStackPtr(0))
	}
	i.AppendInstruction(indexInstruction)

	i.CheckIsConstant()
	i.PostProcess()
	i.StackIncrement = 1
}

func (i *IndicesNode) CheckIsConstant() {
	i.Variadic = true
	if (i.From == nil || !i.From.IsVariadic()) && (i.To == nil || !i.To.IsVariadic()) && (i.Step == nil || !i.Step.IsVariadic()) {
		i.Variadic = false
	}
}

func (i *IndicesNode) Compile(c *Compiler) {
	var from, to, step []common.Instruction
	if i.From != nil {
		i.From.SetRequiredType(common.BasicTypeMap[common.Int64Type])
		i.From.Compile(c)
		from = i.From.GetInstructions()
	} else {
		from = []common.Instruction{instruction.GetPushConstantFunc(int64(-1))}
	}
	if i.To != nil {
		i.To.SetRequiredType(common.BasicTypeMap[common.Int64Type])
		i.To.Compile(c)
		to = i.To.GetInstructions()
	} else {
		to = []common.Instruction{instruction.GetPushConstantFunc(int64(-1))}
	}
	if i.Step != nil {
		i.Step.SetRequiredType(common.BasicTypeMap[common.Int64Type])
		i.Step.Compile(c)
		step = i.Step.GetInstructions()
	} else {
		step = []common.Instruction{instruction.GetPushConstantFunc(int64(1))}
	}
	i.AppendInstruction(from...)
	i.AppendInstruction(to...)
	i.AppendInstruction(step...)
	i.AppendInstruction(instruction.GetIndicesFunc())

	// no need to set data type
	i.CheckIsConstant()
	i.PostProcess()
}

func (n *BinaryNode) CheckIsConstant() {
	n.Variadic = true
	if !n.Lhs.IsVariadic() && !n.Rhs.IsVariadic() {
		n.Variadic = false
	}
}

func (n *BinaryNode) Compile(c *Compiler) {
	if n.LhsFlag {
		panic(common.NewCompileErr(n.ErrorWithSource("cannot use binary op on the left hand side")))
	}
	switch n.Op {
	case "&&", "||":
		n.Lhs.SetRequiredType(common.BasicTypeMap[common.BoolType])
		n.Lhs.Compile(c)
		n.Rhs.SetRequiredType(common.BasicTypeMap[common.BoolType])
		n.Rhs.Compile(c)
		n.DataType = n.Lhs.GetDataType()
	default:
		n.Lhs.Compile(c)
		n.Rhs.Compile(c)
		switch n.Op {
		case "==", "!=", ">", "<", ">=", "<=":
			n.DataType = common.BasicTypeMap[common.BoolType]
		default:
			n.DataType = n.Lhs.GetDataType()
			if n.Rhs.GetDataType().Kind.Kind > n.DataType.Kind.Kind {
				n.DataType = n.Rhs.GetDataType()
			}
		}
	}

	switch n.Op {
	case "+", "-", "*", "**", "/", "%", "//", "==", "!=", ">", "<", ">=", "<=":
		n.AppendInstruction(n.Lhs.GetInstructions()...)
		n.AppendInstruction(n.Rhs.GetInstructions()...)

		opFunc := instruction.GetBinaryOpFunc(n.Op, n.Lhs.GetDataType(), n.Rhs.GetDataType())
		if common.IsError(opFunc) {
			panic(common.NewCompileErr(n.ErrorWithSource("cannot do binary op")))
		}
		n.AppendInstruction(opFunc)
	case "&&":
		lhsInstructions := n.Lhs.GetInstructions()
		rhsInstructions := n.Rhs.GetInstructions()
		opFunc := instruction.GetAndFunc(len(rhsInstructions))
		n.AppendInstruction(lhsInstructions...)
		n.AppendInstruction(opFunc)
		n.AppendInstruction(rhsInstructions...)
	case "||":
		lhsInstructions := n.Lhs.GetInstructions()
		rhsInstructions := n.Rhs.GetInstructions()
		opFunc := instruction.GetOrFunc(len(rhsInstructions))
		n.AppendInstruction(lhsInstructions...)
		n.AppendInstruction(opFunc)
		n.AppendInstruction(rhsInstructions...)
	case "~=":
		// todo: regexp
	default:
		panic(common.NewCompileErr(n.ErrorWithSource("unknown binary op " + n.Op)))
	}

	n.CheckIsConstant()
	n.PostProcess()
	n.StackIncrement = 1
}

func (u *LeftUnaryNode) CheckIsConstant() {
	u.Variadic = true
	if !u.Operand.IsVariadic() {
		u.Variadic = false
	}
}

func (u *LeftUnaryNode) Compile(c *Compiler) {
	if u.LhsFlag {
		panic(common.NewCompileErr(u.ErrorWithSource("cannot use unary op on the left hand side")))
	}
	if u.Op == "++" || u.Op == "--" {
		u.Operand.SetLhs()
	}
	u.Operand.Compile(c)

	opFunc := instruction.GetLeftUnaryOpFunc(u.Op, u.Operand.GetDataType())
	if common.IsError(opFunc) {
		panic(common.NewCompileErr(u.ErrorWithSource("invalid op")))
	}

	u.AppendInstruction(u.Operand.GetInstructions()...)
	if u.Operand.IsStackPtr() {
		u.AppendInstruction(instruction.GetStackOffsetToStackPtr(0))
	}
	u.AppendInstruction(opFunc)

	u.DataType = u.Operand.GetDataType()
	u.CheckIsConstant()
	u.PostProcess()
	u.StackIncrement = 1
}

func (u *RightUnaryNode) Compile(c *Compiler) {
	if u.LhsFlag {
		panic(common.NewCompileErr(u.ErrorWithSource("cannot use unary op on the left hand side")))
	}
	if u.Op == "++" || u.Op == "--" {
		u.Operand.SetLhs()
	}
	u.Operand.Compile(c)

	opFunc := instruction.GetRightUnaryOpFunc(u.Op, u.Operand.GetDataType())
	if common.IsError(opFunc) {
		panic(common.NewCompileErr(u.ErrorWithSource("invalid op")))
	}

	u.AppendInstruction(u.Operand.GetInstructions()...)
	if u.Operand.IsStackPtr() {
		u.AppendInstruction(instruction.GetStackOffsetToStackPtr(0))
	}
	u.AppendInstruction(opFunc)

	u.DataType = u.Operand.GetDataType()
	u.PostProcess()
	u.StackIncrement = 1
}

func (n *ConstructorNode) Compile(c *Compiler) {
	if n.LhsFlag {
		panic(common.NewCompileErr(n.ErrorWithSource("constructor cannot be on the left hand side")))
	}

	n.Type.Compile(c)

	num := len(n.Params)
	for _, p := range n.Params {
		p.Compile(c)
		n.AppendInstruction(p.GetInstructions()...)
	}
	n.AppendInstruction(instruction.GetConstructFunc(n.Type.GetDataType().Constructor, num))

	n.DataType = n.Type.GetDataType()
	n.PostProcess()
	n.StackIncrement = 1
}

func (v *ValueNode) Compile(_ *Compiler) {
	if v.LhsFlag {
		panic(common.NewCompileErr(v.ErrorWithSource("constant value cannot be on the left hand side")))
	}
	v.AppendInstruction(instruction.GetPushConstantFunc(v.ConstantValue))
	v.PostProcess()
}

func (v *ChanSend) Compile(c *Compiler) {
	if v.LhsFlag {
		panic(common.NewCompileErr(v.ErrorWithSource("chan send cannot be on the left hand side")))
	}
	v.Variadic = true
	v.DataType = common.BasicTypeMap[common.BoolType]

	v.Chan.Compile(c)
	if v.Chan.GetDataType().Kind.Kind != common.Channel {
		panic(common.NewCompileErr(v.ErrorWithSource("chan op must be on a chan type value")))
	}
	v.AppendInstruction(v.Chan.GetInstructions()...)
	v.ToSend.SetRequiredType(v.Chan.GetDataType().ItemType)
	v.ToSend.Compile(c)
	v.AppendInstruction(v.ToSend.GetInstructions()...)
	v.AppendInstruction(instruction.GetChanSend(v.NonBlock))

	v.PostProcess()
	v.StackIncrement = 1
}

func (v *ChanRecv) Compile(c *Compiler) {
	if v.LhsFlag {
		panic(common.NewCompileErr(v.ErrorWithSource("chan recv cannot be on the left hand side")))
	}
	v.Chan.Compile(c)
	if v.Chan.GetDataType().Kind.Kind != common.Channel {
		panic(common.NewCompileErr(v.ErrorWithSource("chan op must be on a chan type value")))
	}
	v.AppendInstruction(v.Chan.GetInstructions()...)
	v.AppendInstruction(instruction.GetChanRecv(v.NonBlock))
	v.DataType = v.Chan.GetDataType().ItemType

	v.PostProcess()
	v.StackIncrement = 1
}

func (i *InitializationSliceNode) Compile(c *Compiler) {
	i.Variadic = true
	if i.Type != nil {
		i.Type.Compile(c)
		i.DataType = i.Type.GetDataType()
	} else {
		i.DataType = i.RequiredType
	}
	if i.DataType == nil {
		panic(common.NewTypeErr(i.ErrorWithSource("initialization type missing")))
	}

	num := len(i.Items)
	for _, item := range i.Items {
		item.SetRequiredType(i.DataType.ItemType)
		item.Compile(c)
		i.AppendInstruction(item.GetInstructions()...)
	}
	i.AppendInstruction(instruction.GetInitializeSliceFunc(num))
	i.StackIncrement = 1
}

func (i *InitializationKVNode) Compile(c *Compiler) {
	i.Variadic = true
	if i.Type != nil {
		i.Type.Compile(c)
		i.DataType = i.Type.GetDataType()
	} else {
		i.DataType = i.RequiredType
	}
	if i.DataType == nil {
		panic(common.NewTypeErr(i.ErrorWithSource("initialization type missing")))
	}

	switch i.DataType.Kind.Kind {
	case common.Map:
		num := len(i.Keys)
		for _, key := range i.Keys {
			key.SetRequiredType(i.DataType.KeyType)
			key.Compile(c)
			i.AppendInstruction(key.GetInstructions()...)
		}
		for _, val := range i.Values {
			val.SetRequiredType(i.DataType.ValueType)
			val.Compile(c)
			i.AppendInstruction(val.GetInstructions()...)
		}
		constructor := common.MapConstructorMap[i.DataType.KeyType.Kind.Kind]
		mapSet := instruction.GetMapSetFunc(i.DataType.KeyType)
		i.AppendInstruction(instruction.GetInitializeMapFunc(constructor, mapSet, num))
	case common.Message:
		num := len(i.Keys)
		keyStrings := make([]string, num)

		for idx, key := range i.Keys {
			keyStrings[num-1-idx] = key.GetText()
		}
		for idx, val := range i.Values {
			val.SetRequiredType(i.DataType.FieldType[keyStrings[num-1-idx]])
			val.Compile(c)
			i.AppendInstruction(val.GetInstructions()...)
		}
		i.AppendInstruction(instruction.GetInitializeMessageFunc(keyStrings, num))
	default:
		panic(common.NewTypeErr(i.ErrorWithSource("initialization type should be message or map")))
	}
	i.StackIncrement = 1
}

func (f *FunctionAssignNode) Compile(c *Compiler) {
	num := len(f.Lhs)
	if funcCallNode, ok := f.Function.(*FunctionCallNode); ok {
		funcCallNode.Compile(c)
		if len(funcCallNode.Meta.Out) != num {
			panic(common.NewMismatchErr(f.ErrorWithSource("number of function return value and lhs mismatch")))
		}
		f.StackIncrement = funcCallNode.StackIncrement
		f.AppendInstruction(funcCallNode.Instructions...)

		for _, lhs := range f.Lhs {
			lhs.SetLhs()
			lhs.Compile(c)
			f.AppendInstruction(lhs.GetInstructions()...)
		}

		convertFunctions := make([]common.Instruction, 0, num)
		for i := 0; i < num; i++ {
			temp := instruction.GetConvertInstruction(f.Lhs[num-1-i].GetDataType(), funcCallNode.Meta.Out[i])
			if common.IsError(temp) {
				panic(common.NewCompileErr(f.ErrorWithSource("function return value cannot be implicit convert to lhs")))
			}
			convertFunctions = append(convertFunctions, temp)
		}
		for i, lhs := range f.Lhs {
			if lhs.IsStackPtr() {
				f.AppendInstruction(instruction.GetStackOffsetToStackPtr(num - 1 - i))
			}
		}
		f.AppendInstruction(instruction.GetMultiAssignFunc(convertFunctions)...)
	} else {
		panic(common.NewCompileErr(f.ErrorWithSource("multi-assign RHS is not a function call")))
	}
}

func (n *AssignNode) Compile(c *Compiler) {
	n.Lhs.SetLhs()
	n.Lhs.Compile(c)
	n.Rhs.SetRequiredType(n.Lhs.GetDataType())
	n.Rhs.Compile(c)
	n.DataType = n.Lhs.GetDataType()

	n.AppendInstruction(n.Rhs.GetInstructions()...)
	n.AppendInstruction(n.Lhs.GetInstructions()...)
	if n.Lhs.IsStackPtr() {
		n.AppendInstruction(instruction.GetStackOffsetToStackPtr(0))
	}
	opFunc := instruction.GetAssignFunc(n.Op, n.Lhs.GetDataType())
	if common.IsError(opFunc) {
		panic(common.NewCompileErr(n.ErrorWithSource("invalid op")))
	}
	n.AppendInstruction(opFunc)

	n.StackIncrement = 1
}
