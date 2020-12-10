package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
)

type TypeParser struct {
	*parser.BasegoscriptListener

	NodeStack  []ASTNode
	VisitStack []ASTNode

	Compiler *Compiler
}

// EnterEveryRule is called when any rule is entered.
//func (s *TypeParser) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	fmt.Println(ctx.GetText(), ctx.GetRuleIndex())
//}

func NewTypeParser() *TypeParser {
	return &TypeParser{}
}

func (s *TypeParser) VisitErrorNode(node antlr.ErrorNode) {
	panic(common.NewCompileErr(node.GetText()))
}

func (s *TypeParser) VisitPush(node ASTNode) {
	s.VisitStack = append(s.VisitStack, node)
}

func (s *TypeParser) VisitTop() ASTNode {
	if len(s.VisitStack) == 0 {
		return nil
	}
	return s.VisitStack[len(s.VisitStack)-1]
}

func (s *TypeParser) VisitPop() ASTNode {
	ret := s.VisitTop()
	s.VisitStack = s.VisitStack[0 : len(s.VisitStack)-1]
	return ret
}

func (s *TypeParser) NodePush(node ASTNode) {
	s.NodeStack = append(s.NodeStack, node)
}

func (s *TypeParser) NodeTop() ASTNode {
	if len(s.NodeStack) == 0 {
		return nil
	}
	return s.NodeStack[len(s.NodeStack)-1]
}

func (s *TypeParser) NodePop() ASTNode {
	ret := s.NodeTop()
	s.NodeStack = s.NodeStack[0 : len(s.NodeStack)-1]
	return ret
}

type TypeRegister struct {
	*parser.BasegoscriptListener
	Compiler *Compiler
}

func NewTypeRegister() *TypeRegister {
	return &TypeRegister{}
}
