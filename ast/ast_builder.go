package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
)

type ASTBuilder struct {
	*parser.BasegoscriptListener

	NodeStack     []ASTNode
	VisitStack    []ASTNode
	InitTypeStack []*common.DataType

	InFunction bool
	LoopDepth  int

	Compiler *Compiler
}

// EnterEveryRule is called when any rule is entered.
//func (s *ASTBuilder) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	fmt.Println(ctx.GetText())
//}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{}
}

func (s *ASTBuilder) VisitErrorNode(node antlr.ErrorNode) {
	panic(common.NewCompileErr(node.GetText()))
}

func (s *ASTBuilder) VisitPush(node ASTNode) {
	s.VisitStack = append(s.VisitStack, node)
}

func (s *ASTBuilder) VisitTop() ASTNode {
	if len(s.VisitStack) == 0 {
		return nil
	}
	return s.VisitStack[len(s.VisitStack)-1]
}

func (s *ASTBuilder) VisitPop() ASTNode {
	ret := s.VisitTop()
	s.VisitStack = s.VisitStack[0 : len(s.VisitStack)-1]
	return ret
}

func (s *ASTBuilder) NodePush(node ASTNode) {
	s.NodeStack = append(s.NodeStack, node)
}

func (s *ASTBuilder) NodeTop() ASTNode {
	if len(s.NodeStack) == 0 {
		return nil
	}
	return s.NodeStack[len(s.NodeStack)-1]
}

func (s *ASTBuilder) NodePop() ASTNode {
	ret := s.NodeTop()
	s.NodeStack = s.NodeStack[0 : len(s.NodeStack)-1]
	return ret
}

func (s *ASTBuilder) TypePush(t *common.DataType) {
	s.InitTypeStack = append(s.InitTypeStack, t)
}

func (s *ASTBuilder) TypeTop() *common.DataType {
	if len(s.InitTypeStack) == 0 {
		return nil
	}
	return s.InitTypeStack[len(s.InitTypeStack)-1]
}

func (s *ASTBuilder) TypePop() *common.DataType {
	ret := s.TypeTop()
	s.InitTypeStack = s.InitTypeStack[0 : len(s.InitTypeStack)-1]
	return ret
}
