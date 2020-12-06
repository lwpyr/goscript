package ast

import (
	"github.com/lw396285v/goscript/common"
	"github.com/lw396285v/goscript/lambda_chains"
	"github.com/lw396285v/goscript/parser"
	"strconv"
	"strings"
)

// EnterVardef is called when production vardef is entered.
func (s *ASTBuilder) EnterVardef(ctx *parser.VardefContext) {
	t := s.Compiler.TypeRegistry.FindType(ctx.NAME(1).GetText())
	if t == nil {
		panic(common.NewTypeNotFoundErr(ctx.NAME(1).GetText()))
	}
	v := &common.Variable{
		Symbol:      ctx.NAME(0).GetText(),
		IsParameter: false,
		Type:        t,
	}
	s.Compiler.Scope.AddVariable(v)
	s.NodePush(&GlobalNode{
		Node: Node{
			Parent: s.VisitTop(),
		},
		Assign:   nil,
		Variable: v,
	})
}

// EnterLocalVariable is called when production LocalVariable is entered.
func (s *ASTBuilder) EnterLocal(ctx *parser.LocalContext) {
	if s.Compiler.Scope.Outer == nil {
		panic(common.NewSymbolErr("you cannot define local variable in the root scope " + ctx.GetText()))
	}
	t := s.Compiler.TypeRegistry.FindType(ctx.NAME(1).GetText())
	if t == nil {
		panic(common.NewTypeNotFoundErr(ctx.NAME(1).GetText()))
	}
	v := &common.Variable{
		Symbol:      ctx.NAME(0).GetText(),
		IsParameter: false,
		Type:        t,
	}
	s.Compiler.Scope.AddParameterVariable(v)
	s.NodePush(&LocalNode{
		Node: Node{
			Parent:   s.VisitTop(),
			DataType: t,
			Variadic: true,
		},
		Variable: v,
	})
}

// EnterLocalVariableAssign is called when production LocalVariableAssign is entered.
func (s *ASTBuilder) EnterLocalAssign(ctx *parser.LocalAssignContext) {
	if s.Compiler.Scope.Outer == nil {
		panic(common.NewSymbolErr("you cannot define local variable in the root scope " + ctx.GetText()))
	}
	t := s.Compiler.TypeRegistry.FindType(ctx.NAME(1).GetText())
	if t == nil {
		panic(common.NewTypeNotFoundErr(ctx.NAME(1).GetText()))
	}
	v := &common.Variable{
		Symbol:      ctx.NAME(0).GetText(),
		IsParameter: true,
		Type:        t,
	}
	s.Compiler.Scope.AddParameterVariable(v)
	cur := &LocalNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: v.Type,
		},
		Variable: v,
	}
	cur.Assign = &BinaryNode{
		Node: Node{
			Parent:   cur,
			Variadic: true,
		},
		Op: "=",
		Lhs: &VarNode{
			Node: Node{
				Variadic: true,
				DataType: v.Type,
			},
			Variable: v,
		},
	}
	s.VisitPush(cur)
}

// ExitLocalVariableAssign is called when production LocalVariableAssign is exited.
func (s *ASTBuilder) ExitLocalAssign(ctx *parser.LocalAssignContext) {
	node := s.VisitPop().(*LocalNode)
	node.Assign.Rhs = s.NodePop()
	s.NodePush(node)
}

// EnterStartFunctionAssign is called when production StartFunctionAssign is entered.
func (s *ASTBuilder) EnterFunctionAssign(_ *parser.FunctionAssignContext) {
	cur := &FunctionAssignNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitStartFunctionAssign is called when production StartFunctionAssign is exited.
func (s *ASTBuilder) ExitFunctionAssign(ctx *parser.FunctionAssignContext) {
	node := s.VisitPop().(*FunctionAssignNode)
	node.Function = s.NodePop().(IFunctionNode)
	node.Lhs = []ASTNode{}
	num := len(ctx.AllLhs())
	if num > len(node.Function.(*FunctionNode).Meta.Out) {
		panic(common.NewMismatchErr("number of placeholders is larger than the function output " + ctx.GetText()))
	}
	for i := 0; i < num; i++ {
		temp := s.NodePop()
		if temp.GetDataType().Type != node.Function.(*FunctionNode).Meta.Out[num-1-i].Type {
			panic(common.NewMismatchErr("cannot assign different type to an variable " + ctx.GetText()))
		}
		node.Lhs = append(node.Lhs, temp)
	}
	s.NodePush(node)
}

// EnterStartAssign is called when production StartAssign is entered.
func (s *ASTBuilder) EnterMultiAssign(_ *parser.MultiAssignContext) {
	cur := &AssignNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitStartAssign is called when production StartAssign is exited.
func (s *ASTBuilder) ExitMultiAssign(ctx *parser.MultiAssignContext) {
	node := s.VisitPop().(*AssignNode)
	if len(ctx.AllLhs()) != len(ctx.AllExpr()) {
		panic(common.NewMismatchErr("number mismatch " + ctx.GetText()))
	}
	lenAssign := len(ctx.AllLhs())
	for i := 0; i < lenAssign; i++ {
		node.Rhs = append(node.Rhs, s.NodePop())
	}
	for i := 0; i < lenAssign; i++ {
		node.Lhs = append(node.Lhs, s.NodePop())
	}
	for i := 0; i < lenAssign; i++ {
		if !node.Lhs[i].GetDataType().CanConvertTo(node.Rhs[i].GetDataType()) {
			panic(common.NewMismatchErr("cannot assign different type to an variable " + ctx.GetText()))
		}
		if !node.Lhs[i].IsVariadic() {
			panic(common.NewMismatchErr("cannot assign to constant value " + ctx.GetText()))
		}
	}
	s.NodePush(node)
}

// EnterVariableSliceFilter is called when production VariableSliceFilter is entered.
func (s *ASTBuilder) EnterSliceFilter(_ *parser.SliceFilterContext) {
	cur := &SliceFilterNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			NodeType: "SliceFilterNode",
		},
	}
	s.VisitPush(cur)
}

// ExitVariableSliceFilter is called when production VariableSliceFilter is exited.
func (s *ASTBuilder) ExitSliceFilter(_ *parser.SliceFilterContext) {
	node := s.VisitPop().(*SliceFilterNode)
	node.Expr = s.NodePop()
	node.Slice = s.NodePop()
	node.DataType = node.Slice.GetDataType()
	s.NodePush(node)
}

// EnterVariableSelect is called when production VariableSelect is entered.
func (s *ASTBuilder) EnterSelect(ctx *parser.SelectContext) {
	cur := &SelectorNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		FieldName: ctx.NAME().GetText(),
	}
	s.VisitPush(cur)
}

// ExitVariableSelect is called when production VariableSelect is exited.
func (s *ASTBuilder) ExitSelect(ctx *parser.SelectContext) {
	node := s.VisitPop().(*SelectorNode)
	node.Data = s.NodePop()
	toSelectType := node.Data.GetDataType()
	if toSelectType.Kind.Kind == common.Message {
		if t, ok := toSelectType.FieldType[node.FieldName]; ok {
			node.DataType = t
		} else {
			panic(common.NewSymbolErr("type does not have the field " + ctx.GetText()))
		}
	} else if toSelectType.Kind.Kind == common.Enum {
		node.DataType = toSelectType.ValueType
	} else {
		panic(common.NewSymbolErr("Try to select on not a enum or message type " + ctx.GetText()))
	}

	if !node.Data.IsVariadic() {
		if toSelectType.Kind.Kind == common.Message {
			mergedNode := &ValueNode{
				Node: Node{
					Parent:   s.VisitTop(),
					NodeType: "ValueNode",
					DataType: node.DataType,
					Variadic: false,
				},
				Val: lambda_chains.MessageGetField(node.Data.(IConstantNode).GetConstantValue(), node.FieldName),
			}
			s.NodePush(mergedNode)
		} else {
			mergedNode := &ValueNode{
				Node: Node{
					Parent:   s.VisitTop(),
					NodeType: "ValueNode",
					DataType: node.DataType,
					Variadic: false,
				},
				Val: node.Data.(IConstantNode).GetConstantValue().(map[string]int32)[node.FieldName],
			}
			s.NodePush(mergedNode)
		}
	} else {
		s.NodePush(node)
	}
}

// EnterFilter is called when production filter is entered.
func (s *ASTBuilder) EnterFilter(_ *parser.FilterContext) {
	scope := common.NewScope(s.Compiler.Scope)
	scope.AddParameterVariable(&common.Variable{
		Offset:      0,
		Symbol:      "@",
		IsParameter: true,
		Type:        s.NodeTop().GetDataType().ItemType,
	})
	s.Compiler.Scope = scope
}

// ExitFilter is called when production filter is exited.
func (s *ASTBuilder) ExitFilter(_ *parser.FilterContext) {
	s.Compiler.Scope = s.Compiler.Scope.Outer
}

// EnterVariableIndex is called when production VarianleIndex is entered.
func (s *ASTBuilder) EnterIndex(_ *parser.IndexContext) {
	cur := &IndexNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitVariableIndex is called when production VariableIndex is exited.
func (s *ASTBuilder) ExitIndex(ctx *parser.IndexContext) {
	node := s.VisitPop().(*IndexNode)
	node.Index = s.NodePop()
	node.ToIndex = s.NodePop()
	switch node.ToIndex.GetDataType().Kind.Kind {
	case common.Slice:
		node.IndexType = "slice"
		if !common.IsInteger(node.Index.GetDataType()) {
			panic(common.NewTypeErr("type is not an integer " + ctx.GetText()))
		}
		if !node.Index.IsVariadic() && !node.ToIndex.IsVariadic() {
			mergedNode := MergeSliceIndex(node.ToIndex.(IConstantNode), node.Index.(IConstantNode))
			mergedNode.SetParent(node.Parent)
			s.NodePush(mergedNode)
		} else {
			node.DataType = node.ToIndex.GetDataType().ItemType
			s.NodePush(node)
		}
	case common.Map:
		node.IndexType = "map"
		if !node.Index.GetDataType().CanConvertTo(node.ToIndex.GetDataType().KeyType) {
			panic(common.NewTypeErr("map key type does not match " + ctx.GetText()))
		}
		if !node.Index.IsVariadic() && !node.ToIndex.IsVariadic() {
			mergedNode := MergeMapIndex(node.ToIndex.(IConstantNode), node.Index.(IConstantNode))
			mergedNode.SetParent(node.Parent)
			s.NodePush(mergedNode)
		} else {
			node.DataType = node.ToIndex.GetDataType().ValueType
			s.NodePush(node)
		}
	case common.String:
		node.IndexType = "string"
		if !node.Index.IsVariadic() && !node.ToIndex.IsVariadic() {
			mergedNode := MergeStringIndex(node.ToIndex.(IConstantNode), node.Index.(IConstantNode))
			mergedNode.SetParent(node.Parent)
			s.NodePush(mergedNode)
		} else {
			node.DataType = common.BasicTypeMap["string"]
			s.NodePush(node)
		}
	default:
		panic(common.NewTypeErr("type is not a slice " + ctx.GetText()))
	}
}

// EnterVariableName is called when production VariableName is entered.
func (s *ASTBuilder) EnterVariableName(ctx *parser.VariableNameContext) {
	if ctx.NAME() == nil {
		cur := &VarNode{
			Node: Node{
				Parent:   s.VisitTop(),
				Variadic: true,
				DataType: s.Compiler.Scope.GetVariable("@").Type,
			},
			Variable: s.Compiler.Scope.GetVariable("@"),
		}
		s.VisitPush(cur)
	} else {
		enum := s.Compiler.TypeRegistry.GetEnums(ctx.NAME().GetText())
		// todo: constant check
		if enum != nil {
			cur := &ValueNode{
				Node: Node{
					Parent:   s.VisitTop(),
					Variadic: false,
					DataType: common.BasicTypeMap["enum"],
				},
				Val: enum,
			}
			s.VisitPush(cur)
		} else if ctx.NAME() != nil {
			if constVal := s.Compiler.Scope.GetConstant(ctx.NAME().GetText()); constVal != nil {
				cur := &ValueNode{
					Node: Node{
						Parent:   s.VisitTop(),
						Variadic: true,
						DataType: constVal.Type,
					},
					Val: constVal.Data,
				}
				s.VisitPush(cur)
			} else if varVal := s.Compiler.Scope.GetVariable(ctx.NAME().GetText()); varVal != nil {
				cur := &VarNode{
					Node: Node{
						Parent:   s.VisitTop(),
						Variadic: true,
						DataType: varVal.Type,
					},
					Variable: s.Compiler.Scope.GetVariable(ctx.NAME().GetText()),
				}
				s.VisitPush(cur)
			} else {
				panic(common.NewSymbolErr("unknown symbol " + ctx.NAME().GetText()))
			}
		}
	}
}

// ExitVariableName is called when production VariableName is exited.
func (s *ASTBuilder) ExitVariableName(_ *parser.VariableNameContext) {
	s.NodePush(s.VisitPop())
}

// EnterVariableSliceIndex is called when production VariableSliceIndex is entered.
func (s *ASTBuilder) EnterSliceMultiIndex(_ *parser.SliceMultiIndexContext) {
	cur := &SliceMultiIndexNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitVariableSliceIndex is called when production VariableSliceIndex is exited.
func (s *ASTBuilder) ExitSliceMultiIndex(ctx *parser.SliceMultiIndexContext) {
	node := s.VisitPop().(*SliceMultiIndexNode)
	allConstant := true
	for i := 0; i < len(ctx.AllIndexs()); i++ {
		temp := s.NodePop().(ASTNode)
		allConstant = allConstant && !temp.IsVariadic()
		node.Indices = append(node.Indices, temp)
	}
	node.Slice = s.NodePop()
	if node.Slice.GetDataType().Kind.Kind == common.String {
		if len(node.Indices) > 1 || node.Indices[0].GetNodeType() == "IndexNodeType1" {
			panic(common.NewIndexErr("string index error " + ctx.GetText()))
		}
	}
	allConstant = allConstant && !node.Slice.IsVariadic()
	if allConstant {
		mergedNode := MergeSliceMultiIndex(node.Slice.(IConstantNode), node.Indices)
		mergedNode.SetParent(node.Parent)
		s.NodePush(mergedNode)
	} else {
		node.DataType = node.Slice.GetDataType()
		s.NodePush(node)
	}
}

// EnterVariableMapSliceIndex is called when production VariableMapSliceIndex is entered.
func (s *ASTBuilder) EnterMapMultiIndex(_ *parser.MapMultiIndexContext) {
	cur := &MapMultiIndexNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitVariableMapSliceIndex is called when production VariableMapSliceIndex is exited.
func (s *ASTBuilder) ExitMapMultiIndex(ctx *parser.MapMultiIndexContext) {
	node := s.VisitPop().(*MapMultiIndexNode)
	allConstant := true
	for i := 0; i < len(ctx.AllExpr()); i++ {
		temp := s.NodePop().(ASTNode)
		allConstant = allConstant && !temp.IsVariadic()
		node.Fields = append(node.Fields, temp)
	}
	node.Data = s.NodePop()
	if node.Data.GetDataType().Kind.Kind != common.Map {
		panic(common.NewTypeErr("access keys of a none map value " + ctx.GetText()))
	}
	for _, key := range node.Fields {
		if key.GetDataType().Type != node.Data.GetDataType().KeyType.Type {
			panic(common.NewTypeErr("map key type mismatch " + ctx.GetText()))
		}
	}
	allConstant = allConstant && !node.Data.IsVariadic()
	if allConstant {
		mergedNode := MergeMapMultiIndex(node.Data.(IConstantNode), node.Fields)
		mergedNode.SetParent(node.Parent)
		mergedNode.(*ValueNode).DataType = s.Compiler.TypeRegistry.FindSliceType(node.Data.GetDataType().ValueType.Type)
		s.NodePush(mergedNode)
	} else {
		node.DataType = s.Compiler.TypeRegistry.FindSliceType(node.Data.GetDataType().ValueType.Type)
		s.NodePush(node)
	}
}

// EnterIndexType1 is called when production IndexType1 is entered.
func (s *ASTBuilder) EnterIndexType1(_ *parser.IndexType1Context) {
	cur := &IndicesNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: s.Compiler.TypeRegistry.FindSliceType("int64"),
			NodeType: "IndexNodeType1",
		},
	}
	s.VisitPush(cur)
}

// ExitIndexType1 is called when production IndexType1 is exited.
func (s *ASTBuilder) ExitIndexType1(ctx *parser.IndexType1Context) {
	cur := s.VisitPop().(*IndicesNode)
	int64DataType := common.BasicTypeMap["int64"]

	cur.Step = s.NodePop().(ASTNode)
	cur.To = s.NodePop().(ASTNode)
	cur.From = s.NodePop().(ASTNode)
	if !cur.Step.GetDataType().CanConvertTo(int64DataType) || !cur.To.GetDataType().CanConvertTo(int64DataType) || !cur.From.GetDataType().CanConvertTo(int64DataType) {
		panic(common.NewIndexErr("index is not an integer " + ctx.GetText()))
	}
	s.NodePush(cur)
}

// EnterIndexType2 is called when production IndexType2 is entered.
func (s *ASTBuilder) EnterIndexType2(_ *parser.IndexType2Context) {
	cur := &IndicesNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: s.Compiler.TypeRegistry.FindSliceType("int64"),
			NodeType: "IndexNodeType2",
		},
	}
	s.VisitPush(cur)
}

// ExitIndexType2 is called when production IndexType2 is exited.
func (s *ASTBuilder) ExitIndexType2(ctx *parser.IndexType2Context) {
	cur := s.VisitPop().(*IndicesNode)
	int64DataType := common.BasicTypeMap["int64"]

	cur.To = s.NodePop().(ASTNode)
	cur.From = s.NodePop().(ASTNode)
	if !cur.To.GetDataType().CanConvertTo(int64DataType) || !cur.From.GetDataType().CanConvertTo(int64DataType) {
		panic(common.NewIndexErr("index is not an integer " + ctx.GetText()))
	}
	s.NodePush(cur)
}

// EnterIndexType3 is called when production IndexType3 is entered.
func (s *ASTBuilder) EnterIndexType3(_ *parser.IndexType3Context) {
	cur := &IndicesNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: s.Compiler.TypeRegistry.FindSliceType("int64"),
			NodeType: "IndexNodeType3",
		},
	}
	s.VisitPush(cur)
}

// ExitIndexType3 is called when production IndexType3 is exited.
func (s *ASTBuilder) ExitIndexType3(ctx *parser.IndexType3Context) {
	cur := s.VisitPop().(*IndicesNode)
	int64DataType := common.BasicTypeMap["int64"]

	cur.From = s.NodePop().(ASTNode)
	if !cur.From.GetDataType().CanConvertTo(int64DataType) {
		panic(common.NewIndexErr("Index error: index is not an integer " + ctx.GetText()))
	}
	s.NodePush(cur)
}

// EnterIndexType4 is called when production IndexType4 is entered.
func (s *ASTBuilder) EnterIndexType4(_ *parser.IndexType4Context) {
	cur := &IndicesNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: s.Compiler.TypeRegistry.FindSliceType("int64"),
			NodeType: "IndexNodeType4",
		},
	}
	s.VisitPush(cur)
}

// ExitIndexType4 is called when production IndexType4 is exited.
func (s *ASTBuilder) ExitIndexType4(ctx *parser.IndexType4Context) {
	cur := s.VisitPop().(*IndicesNode)
	int64DataType := common.BasicTypeMap["int64"]

	cur.Step = s.NodePop().(ASTNode)
	if !cur.Step.GetDataType().CanConvertTo(int64DataType) {
		panic(common.NewIndexErr("Index error: index is not an integer" + ctx.GetText()))
	}
	if !cur.Step.IsVariadic() {
		cur.Variadic = false
	}
	s.NodePush(cur)
}

// EnterIndexType5 is called when production IndexType5 is entered.
func (s *ASTBuilder) EnterIndexType5(_ *parser.IndexType5Context) {
	cur := &IndicesNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: s.Compiler.TypeRegistry.FindSliceType("int64"),
			NodeType: "IndexNodeType5",
		},
	}
	s.VisitPush(cur)
}

// ExitIndexType5 is called when production IndexType5 is exited.
func (s *ASTBuilder) ExitIndexType5(ctx *parser.IndexType5Context) {
	cur := s.VisitPop().(*IndicesNode)
	int64DataType := common.BasicTypeMap["int64"]

	cur.To = s.NodePop().(ASTNode)
	if !cur.To.GetDataType().CanConvertTo(int64DataType) {
		panic(common.NewIndexErr("Index error: index is not an integer" + ctx.GetText()))
	}
	s.NodePush(cur)
}

// EnterExprBinary is called when production ExprBinary is entered.
func (s *ASTBuilder) EnterBinary(ctx *parser.BinaryContext) {
	cur := &BinaryNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		Op: ctx.GetOp().GetText(),
	}
	s.VisitPush(cur)
}

// ExitExprBinary is called when production ExprBinary is exited.
func (s *ASTBuilder) ExitBinary(ctx *parser.BinaryContext) {
	node := s.VisitPop().(*BinaryNode)

	node.Rhs = s.NodePop()
	node.Lhs = s.NodePop()

	if !node.Lhs.IsVariadic() && !node.Rhs.IsVariadic() {
		lhs := node.Lhs.(IConstantNode)
		rhs := node.Rhs.(IConstantNode)
		mergedNode := MergeBinary(node.Op, lhs, rhs)
		mergedNode.SetParent(node.GetParent())
		s.NodePush(mergedNode)
		return
	} else {
		lhs := node.Lhs
		rhs := node.Rhs
		var ResType *common.DataType
		switch node.Op {
		case "/":
			if ResType = common.CanDiv(lhs.GetDataType(), rhs.GetDataType()); ResType == nil {
				panic(common.NewMathErr("type cannot '/' " + ctx.GetText()))
			}
		case "+", "-", "*", "**":
			if lhs.GetDataType().Kind.Kind == common.String && rhs.GetDataType().Kind.Kind == common.String {
				if node.Op != "+" {
					panic(common.NewMathErr("string can only '+' " + ctx.GetText()))
				} else {
					ResType = common.BasicTypeMap["string"]
				}
			} else if ResType = common.CanAddSubMulPow(lhs.GetDataType(), rhs.GetDataType()); ResType == nil {
				panic(common.NewMathErr("type cannot '+','-','*','**' " + ctx.GetText()))
			}
		case "%", "//":
			if ResType = common.CanMod(lhs.GetDataType(), rhs.GetDataType()); ResType == nil {
				panic(common.NewMathErr("type cannot '%','//'" + ctx.GetText()))
			}
		case "&&", "||":
			if ResType = common.CanAndOr(lhs.GetDataType(), rhs.GetDataType()); ResType == nil {
				panic(common.NewTypeErr("type is not boolean" + ctx.GetText()))
			}
		case ">", "<", ">=", "<=":
			if ResType = common.CanCompare(lhs.GetDataType(), rhs.GetDataType()); ResType == nil {
				panic(common.NewMathErr("type cannot compare" + ctx.GetText()))
			}
		case "==", "!=":
			if ResType = common.CanEqual(lhs.GetDataType(), rhs.GetDataType()); ResType == nil {
				panic(common.NewMathErr("type cannot compare" + ctx.GetText()))
			}
		case "=":
			if rhs.GetDataType().Kind.Kind == common.Nil {
				panic(common.NewTypeErr(" cannot explicit assign nil" + ctx.GetText()))
			}
			if !rhs.GetDataType().CanConvertTo(lhs.GetDataType()) {
				panic(common.NewTypeErr("cannot assign type " + rhs.GetDataType().Type + " to type " + lhs.GetDataType().Type + " " + ctx.GetText()))
			}
			ResType = lhs.GetDataType()
		default:
			panic("wtf?")
		}
		node.DataType = ResType
		node.NodeType = "BinaryNode"
		s.NodePush(node)
	}
}

func (s *ASTBuilder) EnterLeftUnary(ctx *parser.LeftUnaryContext) {
	cur := &LeftUnaryNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		Op: ctx.GetOp().GetText(),
	}

	s.VisitPush(cur)
}

func (s *ASTBuilder) ExitLeftUnary(ctx *parser.LeftUnaryContext) {
	node := s.VisitPop().(*LeftUnaryNode)
	operand := s.NodePop()
	node.DataType = operand.GetDataType()
	node.Operand = operand
	if !operand.IsVariadic() {
		mergedNode := MergeLeftUnary(node.Op, operand.(IConstantNode))
		mergedNode.SetParent(node.Parent)
		s.NodePush(mergedNode)
	} else {
		s.NodePush(node)
	}
}

func (s *ASTBuilder) EnterRightUnary(ctx *parser.RightUnaryContext) {
	cur := &RightUnaryNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		Op: ctx.GetOp().GetText(),
	}
	s.VisitPush(cur)
}

func (s *ASTBuilder) ExitRightUnary(ctx *parser.RightUnaryContext) {
	node := s.VisitPop().(*RightUnaryNode)
	operand := s.NodePop()
	node.DataType = operand.GetDataType()
	node.Operand = operand
	s.NodePush(node)
}

// EnterConstantInt is called when production ConstantInt is entered.
func (s *ASTBuilder) EnterConstantInt(ctx *parser.ConstantIntContext) {
	num, _ := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
	cur := &ValueNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: false,
			DataType: common.BasicTypeMap["int64"],
			NodeType: "ValueNode",
		},
		Val: num,
	}
	s.VisitPush(cur)
}

// ExitConstantInt is called when production ConstantInt is exited.
func (s *ASTBuilder) ExitConstantInt(_ *parser.ConstantIntContext) {
	s.NodePush(s.VisitPop())
}

// EnterConstantFloat is called when production ConstantFloat is entered.
func (s *ASTBuilder) EnterConstantFloat(ctx *parser.ConstantFloatContext) {
	num, _ := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
	cur := &ValueNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: false,
			DataType: common.BasicTypeMap["float64"],
			NodeType: "ValueNode",
		},
		Val: num,
	}
	s.VisitPush(cur)
}

// ExitConstantFloat is called when production ConstantFloat is exited.
func (s *ASTBuilder) ExitConstantFloat(_ *parser.ConstantFloatContext) {
	s.NodePush(s.VisitPop())
}

// EnterConstantBool is called when production ConstantBool is entered.
func (s *ASTBuilder) EnterConstantBool(ctx *parser.ConstantBoolContext) {
	b, _ := strconv.ParseBool(ctx.BOOL().GetText())
	cur := &ValueNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: false,
			DataType: common.BasicTypeMap["bool"],
			NodeType: "ValueNode",
		},
		Val: b,
	}
	s.VisitPush(cur)
}

// ExitConstantBool is called when production ConstantBool is exited.
func (s *ASTBuilder) ExitConstantBool(_ *parser.ConstantBoolContext) {
	s.NodePush(s.VisitPop())
}

// EnterConstantNil is called when production ConstantNil is entered.
func (s *ASTBuilder) EnterConstantNil(_ *parser.ConstantNilContext) {
	cur := &ValueNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: false,
			DataType: common.BasicTypeMap["nil"],
			NodeType: "ValueNode",
		},
		Val: nil,
	}
	s.VisitPush(cur)
}

// ExitConstantNil is called when production ConstantNil is exited.
func (s *ASTBuilder) ExitConstantNil(_ *parser.ConstantNilContext) {
	s.NodePush(s.VisitPop())
}

// EnterConstantString is called when production ConstantString is entered.
func (s *ASTBuilder) EnterConstantString(ctx *parser.ConstantStringContext) {
	str := ctx.STRING().GetText()[1 : len(ctx.STRING().GetText())-1]
	str = strings.Replace(str, "\\'", "'", -1)
	cur := &ValueNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: false,
			DataType: common.BasicTypeMap["string"],
			NodeType: "ValueNode",
		},
		Val: str,
	}
	s.VisitPush(cur)
}

// ExitConstantString is called when production ConstantString is exited.
func (s *ASTBuilder) ExitConstantString(_ *parser.ConstantStringContext) {
	s.NodePush(s.VisitPop())
}

// EnterFunctions is called when production functions is entered.
func (s *ASTBuilder) EnterFunctions(ctx *parser.FunctionsContext) {
	cur := &FunctionNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: common.BasicTypeMap["tuple"],
			NodeType: "FunctionNode",
		},
	}
	s.VisitPush(cur)
}

// ExitFunctions is called when production functions is exited.
func (s *ASTBuilder) ExitFunctions(ctx *parser.FunctionsContext) {
	functionName := ctx.NAME().GetText()
	cur := s.VisitPop().(*FunctionNode)
	if newNode := s.BuiltinCheck(functionName, cur, ctx); newNode != nil {
		s.NodePush(newNode)
	} else {
		// lib
		meta := s.Compiler.FunctionLib[ctx.NAME().GetText()]
		if meta == nil {
			panic(common.NewSymbolErr("undefined function " + functionName))
		}
		cur.Meta = meta
		if len(meta.Out) == 1 {
			cur.DataType = meta.Out[0]
		}
		provided := len(ctx.AllExpr())
		var object ASTNode
		if ctx.Variable() != nil {
			object = s.NodePop()
			provided++
		}
		if provided > len(meta.In) && meta.TailArray == false || provided < len(meta.In) {
			panic(common.NewMismatchErr("more parameter given than needed " + functionName))
		}
		if object != nil {
			if object.GetDataType().Type != meta.In[0].Type {
				panic(common.NewTypeErr(object.GetDataType().Type + " type does not has the method " + functionName))
			}
			cur.Params = append(cur.Params, object)
		}
		for i := 0; i < len(ctx.AllExpr()); i++ {
			temp := s.NodePop()
			cur.Params = append(cur.Params, temp)
			var requireType *common.DataType
			if provided-1-i < len(meta.In) {
				requireType = meta.In[provided-1-i]
			} else {
				requireType = meta.In[len(meta.In)-1]
			}
			if i < len(meta.In) {
				if requireType.Kind.Kind != common.Object && !temp.GetDataType().CanConvertTo(requireType) {
					panic(common.NewTypeErr("function parameter type cannot be implicitly casted " + functionName))
				}
			}
		}
		s.NodePush(cur)
	}
}

// EnterInitSlice is called when production InitSlice is entered.
func (s *ASTBuilder) EnterInitSlice(ctx *parser.InitSliceContext) {
	curType := s.TypePop()
	if curType.Kind.Kind != common.Slice {
		panic(common.NewInitializationErr("type not match" + ctx.GetText()))
	}
	for i := 0; i < len(ctx.AllInitializationList()); i++ {
		s.TypePush(curType.ItemType)
	}
	cur := &InitializationSliceNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: false,
			DataType: curType,
			NodeType: "InitializationSliceNode",
		},
	}
	s.VisitPush(cur)
}

// ExitInitSlice is called when production InitSlice is exited.
func (s *ASTBuilder) ExitInitSlice(ctx *parser.InitSliceContext) {
	cur := s.VisitPop().(*InitializationSliceNode)
	for i := 0; i < len(ctx.AllInitializationList()); i++ {
		cur.Items = append(cur.Items, s.NodePop())
	}
	s.NodePush(cur)
}

// EnterInitMap is called when production InitMap is entered.
func (s *ASTBuilder) EnterInitMap(ctx *parser.InitMapContext) {
	curType := s.TypePop()
	if curType.Kind.Kind != common.Map {
		panic(common.NewInitializationErr("not a map type " + ctx.GetText()))
	}
	cur := &InitializationMapNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: curType,
			NodeType: "InitializationMapNode",
		},
	}
	for i := 0; i < len(ctx.AllInitializationList())/2; i++ {
		s.TypePush(curType.ValueType)
		s.TypePush(curType.KeyType)
	}
	s.VisitPush(cur)
}

// ExitInitMap is called when production InitMap is exited.
func (s *ASTBuilder) ExitInitMap(ctx *parser.InitMapContext) {
	cur := s.VisitPop().(*InitializationMapNode)
	num := len(ctx.AllInitializationList()) / 2
	for i := 0; i < num; i++ {
		cur.Values = append(cur.Values, s.NodePop())
		key := s.NodePop()
		cur.Keys = append(cur.Keys, key)
	}
	s.NodePush(cur)
}

// EnterInitMessage is called when production InitMessage is entered.
func (s *ASTBuilder) EnterInitMessage(ctx *parser.InitMessageContext) {
	curType := s.TypePop()
	if curType.Kind.Kind != common.Message {
		panic(common.NewInitializationErr("not a message type"))
	}
	cur := &InitializationMessageNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: curType,
			NodeType: "InitializationMessageNode",
		},
	}
	for i := 0; i < len(ctx.AllNAME()); i++ {
		if _, ok := curType.FieldType[ctx.NAME(i).GetText()]; !ok {
			panic(common.NewInitializationErr("message does not have the field"))
		}
		cur.Keys = append(cur.Keys, ctx.NAME(i).GetText())
		s.TypePush(curType.FieldType[ctx.NAME(len(ctx.AllNAME())-1-i).GetText()])
	}
	s.VisitPush(cur)
}

// ExitInitMessage is called when production InitMessage is exited.
func (s *ASTBuilder) ExitInitMessage(ctx *parser.InitMessageContext) {
	cur := s.VisitPop().(*InitializationMessageNode)
	num := len(ctx.AllNAME())
	for i := 0; i < num; i++ {
		cur.Values = append(cur.Values, s.NodePop())
	}
	s.NodePush(cur)
}

// EnterExprConstant is called when production ExprConstant is entered.
func (s *ASTBuilder) EnterInitConstant(ctx *parser.InitConstantContext) {
	cur := &InitializationConstantNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: false,
			DataType: s.TypePop(),
			NodeType: "InitializationConstantNode",
		},
	}
	s.VisitPush(cur)
}

// ExitExprConstant is called when production ExprConstant is exited.
func (s *ASTBuilder) ExitInitConstant(ctx *parser.InitConstantContext) {
	cur := s.VisitPop().(*InitializationConstantNode)
	cur.Constant = s.NodePop()
	if !cur.Constant.GetDataType().CanConvertTo(cur.DataType) {
		panic(common.NewInitializationErr("type not match"))
	}
	s.NodePush(cur)
}

// EnterExprAssignInitlist is called when production ExprAssignInitlist is entered.
func (s *ASTBuilder) EnterInitializationListBegin(_ *parser.InitializationListBeginContext) {
	s.TypePush(s.NodeTop().GetDataType())
	cur := &BinaryNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		Op: "=",
	}
	s.VisitPush(cur)
}

// ExitExprAssignInitlist is called when production ExprAssignInitlist is exited.
func (s *ASTBuilder) ExitInitializationListBegin(_ *parser.InitializationListBeginContext) {
	cur := s.VisitPop().(*BinaryNode)
	cur.Rhs = s.NodePop()
	cur.Lhs = s.NodePop()
	s.NodePush(cur)
}

func (s *ASTBuilder) BuiltinCheck(builtinName string, cur *FunctionNode, ctx *parser.FunctionsContext) ASTNode {
	var params []ASTNode
	providedParamNum := len(ctx.AllExpr())
	if ctx.Variable() != nil {
		providedParamNum++
	}
	switch builtinName {
	case "pushBack", "pushFront", "delete", "enumString", "len", "typeof", "uint32", "uint64", "int32", "int64", "float32", "float64", "string", "bytes", "bool", "uint8":
		var dType *common.DataType
		switch builtinName {
		case "uint32", "uint64", "int32", "int64", "float32", "float64", "string", "bytes", "bool", "uint8":
			if providedParamNum != 1 {
				panic(common.NewMismatchErr("wrong arity for type conversion " + ctx.GetText()))
			}
			data := s.NodePop()
			params = []ASTNode{data}
			dType = common.BasicTypeMap[builtinName]
		case "pushBack", "pushFront", "delete":
			if providedParamNum != 2 {
				panic(common.NewMismatchErr("wrong arity for " + builtinName))
			}
			val := s.NodePop()
			data := s.NodePop()
			params = []ASTNode{val, data}
			dType = nil
		case "len":
			if providedParamNum != 1 {
				panic(common.NewMismatchErr("wrong arity for len " + ctx.GetText()))
			}
			params = []ASTNode{s.NodePop()}
			dType = common.BasicTypeMap["int64"]
		case "typeof":
			if len(ctx.AllExpr()) != 1 {
				panic(common.NewMismatchErr("wrong arity for typeof " + ctx.GetText()))
			}
			dType = s.NodePop().GetDataType()
			return &ValueNode{
				Node: Node{
					Parent:   cur.Parent,
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["reflect"],
					Variadic: false,
				},
				Val: dType,
			}
		case "enumString":
			if providedParamNum != 1 {
				panic(common.NewMismatchErr("wrong arity for enum " + ctx.GetText()))
			}
			params = []ASTNode{s.NodePop()}
			val := params[0]
			if val.GetDataType().Kind.Kind != common.Int32 || val.GetDataType().Type == "int32" {
				panic(common.NewTypeErr("enum can only be applied on an enum value " + ctx.GetText()))
			}
			dType = common.BasicTypeMap["string"]
		}
		return &BuiltinFunctionNode{
			Node: Node{
				Parent:   cur.Parent,
				NodeType: "BuiltinFunctionNode",
				DataType: dType,
				Lhs:      false,
				Variadic: true,
			},
			Params:      params,
			BuiltinName: builtinName,
		}
	default:
		return nil
	}
}

// EnterConstruct is called when production Construct is entered.
func (s *ASTBuilder) EnterConstructor(ctx *parser.ConstructorContext) {
	dt := s.Compiler.FindType(ctx.NAME().GetText())
	if dt == nil {
		panic(common.NewTypeNotFoundErr(ctx.GetText()))
	}
	cur := &ConstructorNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: false,
			DataType: dt,
			NodeType: "ValueNode",
		},
	}
	s.VisitPush(cur)
}

// ExitConstruct is called when production Construct is exited.
func (s *ASTBuilder) ExitConstructor(ctx *parser.ConstructorContext) {
	num := len(ctx.AllExpr())
	cur := s.VisitPop().(*ConstructorNode)
	for i := 0; i < num; i++ {
		cur.Params = append(cur.Params, s.NodePop())
	}
	s.NodePush(cur)
}
