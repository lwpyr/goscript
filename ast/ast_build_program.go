package ast

import (
	"github.com/lwpyr/goscript/parser"
	"strconv"
)

// EnterProgram is called when production program is entered.
func (s *ASTBuilder) EnterProgram(ctx *parser.ProgramContext) {
	cur := &ProgramRoot{
		Node: Node{
			Parent:   nil,
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitProgram is called when production program is exited.
func (s *ASTBuilder) ExitProgram(ctx *parser.ProgramContext) {
	node := s.VisitPop().(*ProgramRoot)
	num := len(ctx.AllStatement())
	for i := 0; i < num; i++ {
		temp := s.NodePop()
		switch temp.(type) {
		case *FunctionDefNode:
			node.FunctionDefNode = append(node.FunctionDefNode, temp)
		case *TypeDef, *TypeAlias:
			node.TypeDefNode = append(node.TypeDefNode, temp)
		case *EnumNode:
			node.EnumNode = append(node.EnumNode, temp)
		default:
			node.RunnableNode = append([]ASTNode{temp}, node.RunnableNode...)
		}
	}
	s.NodePush(node)
}

// EnterEnumdef is called when production enumdef is entered.
func (s *ASTBuilder) EnterEnumdef(ctx *parser.EnumdefContext) {
	cur := &EnumNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		EnumName: ctx.NAME(0).GetText(),
	}
	num := len(ctx.AllINT())
	for i := 0; i < num; i++ {
		v, _ := strconv.ParseInt(ctx.INT(i).GetText(), 10, 32)
		cur.Enum[ctx.NAME(i+1).GetText()] = int32(v)
		cur.REnum[int32(v)] = ctx.NAME(i + 1).GetText()
	}
	cur.Compile(s.Compiler)
	s.NodePush(cur)
}

// EnterTypeDefAlias is called when production TypeDefAlias is entered.
func (s *ASTBuilder) EnterTypeDefAlias(ctx *parser.TypeDefAliasContext) {
	// todo
}

// EnterTypeDefComplex is called when production TypeDefComplex is entered.
func (s *ASTBuilder) EnterTypeDefComplex(ctx *parser.TypeDefComplexContext) {
	// todo
}
