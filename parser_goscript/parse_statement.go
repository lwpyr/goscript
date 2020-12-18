package parser_goscript

import (
	"github.com/lwpyr/goscript/ast"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
	"strconv"
	"strings"
)

// ExitAsserted is called when production asserted is exited.
func (s *ASTBuilder) ExitTypeAssert(ctx *parser.TypeAssertContext) {
	typeNode := s.NodePop()
	val := s.NodePop()
	node := &ast.AssertNode{
		Node: ast.Node{
			Variadic: true,
		},
		Type:  typeNode,
		Value: val,
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitStartFunctionAssign is called when production StartFunctionAssign is exited.
func (s *ASTBuilder) ExitFunctionAssign(ctx *parser.FunctionAssignContext) {
	node := &ast.FunctionAssignNode{
		Node: ast.Node{
			Variadic: true,
		},
	}
	node.Function = s.NodePop()
	num := len(ctx.AllExpr()) - 1
	node.Lhs = make([]ast.ASTNode, num)
	for i := 0; i < num; i++ {
		temp := s.NodePop()
		node.Lhs[i] = temp
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitVariableSliceFilter is called when production VariableSliceFilter is exited.
func (s *ASTBuilder) ExitSliceFilter(ctx *parser.SliceFilterContext) {
	node := &ast.SliceFilterNode{
		Node: ast.Node{
			Variadic: true,
		},
	}
	node.Expr = s.NodePop()
	node.Slice = s.NodePop()
	node.DataType = node.Slice.GetDataType()
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitVariableSelect is called when production VariableSelect is exited.
func (s *ASTBuilder) ExitSelect(ctx *parser.SelectContext) {
	node := &ast.SelectorNode{
		Node: ast.Node{
			Variadic: true,
		},
		FieldName: ctx.Fieldname().GetText(),
	}
	node.Data = s.NodePop()
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitVariableIndex is called when production VariableIndex is exited.
func (s *ASTBuilder) ExitIndex(ctx *parser.IndexContext) {
	node := &ast.IndexNode{
		Node: ast.Node{
			Variadic: true,
		},
	}
	node.Index = s.NodePop()
	node.ToIndex = s.NodePop()
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitVariableName is called when production VariableName is exited.
func (s *ASTBuilder) ExitSymbol(ctx *parser.SymbolContext) {
	symbolName := ctx.GetText()
	var node ast.ASTNode
	if enum := s.Compiler.TypeRegistry.Enums[symbolName]; enum != nil {
		node = &ast.ValueNode{
			ast.Node{
				DataType: &common.DataType{
					Type: symbolName,
					Kind: common.KindMap[common.Enum],
				},
				Variadic:      false,
				ConstantValue: enum,
			},
		}
	} else if t := s.Compiler.TypeRegistry.FindType(symbolName); t != nil {
		node = &ast.TypeNode{
			SimpleTypeName: symbolName,
			TypeNodeType:   ast.SimpleType,
		}
	} else if common.IsBuiltIn(symbolName) {
		node = &ast.BuiltinFunctionNode{
			Node: ast.Node{
				Variadic: true,
			},
			BuiltinName: symbolName,
		}
	} else {
		node = &ast.FetchSymbolNode{
			Node: ast.Node{
				Variadic: true,
			},
			SymbolName: symbolName,
		}
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitVariableSliceIndex is called when production VariableSliceIndex is exited.
func (s *ASTBuilder) ExitSliceMultiIndex(ctx *parser.SliceMultiIndexContext) {
	node := &ast.SliceMultiIndexNode{
		Node: ast.Node{
			Variadic: true,
		},
	}
	allConstant := true
	for i := 0; i < len(ctx.AllIndexs()); i++ {
		temp := s.NodePop().(ast.ASTNode)
		allConstant = allConstant && !temp.IsVariadic()
		node.Indices = append(node.Indices, temp)
	}
	node.Slice = s.NodePop()
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitVariableMapSliceIndex is called when production VariableMapSliceIndex is exited.
func (s *ASTBuilder) ExitMapMultiIndex(ctx *parser.MapMultiIndexContext) {
	node := &ast.MapMultiIndexNode{
		Node: ast.Node{
			Variadic: true,
		},
	}
	lenKeys := len(ctx.AllExpr()) - 1
	node.Fields = make([]ast.ASTNode, lenKeys)
	for i := 0; i < lenKeys; i++ {
		node.Fields[i] = s.NodePop()
	}
	node.Map = s.NodePop()
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitIndexType1 is called when production IndexType1 is exited.
func (s *ASTBuilder) ExitIndexType1(ctx *parser.IndexType1Context) {
	node := &ast.IndicesNode{
		Node: ast.Node{
			Variadic: true,
		},
	}

	node.Step = s.NodePop().(ast.ASTNode)
	node.To = s.NodePop().(ast.ASTNode)
	node.From = s.NodePop().(ast.ASTNode)
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitIndexType2 is called when production IndexType2 is exited.
func (s *ASTBuilder) ExitIndexType2(ctx *parser.IndexType2Context) {
	node := &ast.IndicesNode{
		Node: ast.Node{
			Variadic: true,
		},
	}

	node.To = s.NodePop().(ast.ASTNode)
	node.From = s.NodePop().(ast.ASTNode)
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitIndexType3 is called when production IndexType3 is exited.
func (s *ASTBuilder) ExitIndexType3(ctx *parser.IndexType3Context) {
	node := &ast.IndicesNode{
		Node: ast.Node{
			Variadic: true,
		},
	}

	node.From = s.NodePop().(ast.ASTNode)
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitIndexType4 is called when production IndexType4 is exited.
func (s *ASTBuilder) ExitIndexType4(ctx *parser.IndexType4Context) {
	node := &ast.IndicesNode{
		Node: ast.Node{
			Variadic: true,
		},
	}

	node.Step = s.NodePop().(ast.ASTNode)
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitIndexType5 is called when production IndexType5 is exited.
func (s *ASTBuilder) ExitIndexType5(ctx *parser.IndexType5Context) {
	node := &ast.IndicesNode{
		Node: ast.Node{
			Variadic: true,
		},
	}

	node.To = s.NodePop().(ast.ASTNode)
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitExprBinary is called when production ExprBinary is exited.
func (s *ASTBuilder) ExitBinary(ctx *parser.BinaryContext) {
	node := &ast.BinaryNode{
		Node: ast.Node{
			Variadic: true,
		},
		Op: ctx.GetOp().GetText(),
	}

	node.Rhs = s.NodePop()
	node.Lhs = s.NodePop()
	node.SetSource(ctx)

	s.NodePush(node)
}

func (s *ASTBuilder) ExitLeftUnary(ctx *parser.LeftUnaryContext) {
	node := &ast.LeftUnaryNode{
		Node: ast.Node{
			Variadic: true,
		},
		Op: ctx.GetOp().GetText(),
	}
	operand := s.NodePop()
	node.DataType = operand.GetDataType()
	node.Operand = operand
	node.SetSource(ctx)
	s.NodePush(node)
}

func (s *ASTBuilder) ExitRightUnary(ctx *parser.RightUnaryContext) {
	node := &ast.RightUnaryNode{
		Node: ast.Node{
			Variadic: true,
		},
		Op: ctx.GetOp().GetText(),
	}
	operand := s.NodePop()
	node.DataType = operand.GetDataType()
	node.Operand = operand
	node.SetSource(ctx)
	s.NodePush(node)
}

// EnterConstantInt is called when production ConstantInt is entered.
func (s *ASTBuilder) EnterConstantInt(ctx *parser.ConstantIntContext) {
	num, _ := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
	node := &ast.ValueNode{
		Node: ast.Node{
			Variadic:      false,
			DataType:      common.BasicTypeMap[common.Int64Type],
			ConstantValue: num,
		},
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// EnterConstantFloat is called when production ConstantFloat is entered.
func (s *ASTBuilder) EnterConstantFloat(ctx *parser.ConstantFloatContext) {
	num, _ := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
	cur := &ast.ValueNode{
		Node: ast.Node{
			Variadic:      false,
			DataType:      common.BasicTypeMap[common.Float64Type],
			ConstantValue: num,
		},
	}
	s.NodePush(cur)
}

// EnterConstantBool is called when production ConstantBool is entered.
func (s *ASTBuilder) EnterConstantBool(ctx *parser.ConstantBoolContext) {
	b, _ := strconv.ParseBool(ctx.BOOLLITERAL().GetText())
	cur := &ast.ValueNode{
		Node: ast.Node{
			Variadic:      false,
			DataType:      common.BasicTypeMap[common.BoolType],
			ConstantValue: b,
		},
	}
	s.NodePush(cur)
}

// EnterConstantNil is called when production ConstantNil is entered.
func (s *ASTBuilder) EnterConstantNil(_ *parser.ConstantNilContext) {
	cur := &ast.ValueNode{
		Node: ast.Node{
			Variadic:      false,
			DataType:      common.BasicTypeMap["nil"],
			ConstantValue: nil,
		},
	}
	s.NodePush(cur)
}

// EnterConstantString is called when production ConstantString is entered.
func (s *ASTBuilder) EnterConstantString(ctx *parser.ConstantStringContext) {
	str := ctx.STRINGLITERAL().GetText()[1 : len(ctx.STRINGLITERAL().GetText())-1]
	str = strings.Replace(str, "\\'", "'", -1)
	cur := &ast.ValueNode{
		Node: ast.Node{
			Variadic:      false,
			DataType:      common.BasicTypeMap[common.StringType],
			ConstantValue: str,
		},
	}
	s.NodePush(cur)
}

// ExitInitSlice is called when production InitSlice is exited.
func (s *ASTBuilder) ExitInitSlice(ctx *parser.InitSliceContext) {
	cur := &ast.InitializationSliceNode{
		Node: ast.Node{
			Variadic: false,
		},
	}
	lenElem := len(ctx.AllExpr())
	cur.Items = make([]ast.ASTNode, lenElem)
	for i := 0; i < lenElem; i++ {
		cur.Items[i] = s.NodePop()
	}
	if ctx.Typename() != nil {
		cur.Type = s.NodePop()
	}
	s.NodePush(cur)
}

// ExitInitMap is called when production InitMap is exited.
func (s *ASTBuilder) ExitInitKV(ctx *parser.InitKVContext) {
	cur := &ast.InitializationKVNode{
		Node: ast.Node{
			Variadic: true,
		},
	}
	num := len(ctx.AllExpr()) / 2
	cur.Values = make([]ast.ASTNode, num)
	cur.Keys = make([]ast.ASTNode, num)
	for i := 0; i < num; i++ {
		cur.Values[i] = s.NodePop()
		cur.Keys[i] = s.NodePop()
	}
	if ctx.Typename() != nil {
		cur.Type = s.NodePop()
	}
	s.NodePush(cur)
}

// ExitSendOrRecv is called when production SendOrRecv is exited.
func (s *ASTBuilder) ExitSend(ctx *parser.SendContext) {
	rhs := s.NodePop()
	lhs := s.NodePop()
	s.NodePush(&ast.ChanSend{
		Node: ast.Node{
			Variadic: true,
		},
		ToSend:   rhs,
		Chan:     lhs,
		NonBlock: ctx.CHANOPNONBLOCK() != nil,
	})
}

// ExitRecv is called when production Recv is entered.
func (s *ASTBuilder) ExitRecv(ctx *parser.RecvContext) {
	rhs := s.NodePop()
	s.NodePush(&ast.ChanRecv{
		Node: ast.Node{
			Variadic: true,
		},
		Chan:     rhs,
		NonBlock: ctx.CHANOPNONBLOCK() != nil,
	})
}

// ExitAssign is called when production Assign is exited.
func (s *ASTBuilder) ExitAssign(ctx *parser.AssignContext) {
	node := &ast.AssignNode{
		Node: ast.Node{
			Variadic: true,
		},
		Op: ctx.GetOp().GetText(),
	}

	node.Rhs = s.NodePop()
	node.Lhs = s.NodePop()
	node.SetSource(ctx)

	s.NodePush(node)
}
