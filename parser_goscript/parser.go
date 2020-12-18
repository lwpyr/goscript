package parser_goscript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lwpyr/goscript/ast"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
)

type ASTBuilder struct {
	*parser.BasegoscriptListener

	NodeStack []ast.ASTNode

	FunctionDepth int
	LoopDepth     int

	Compiler *ast.Compiler
}

// EnterEveryRule is called when any rule is entered.
//func (s *ASTBuilder) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	fmt.Println(ctx.GetText(), ctx.GetRuleIndex())
//}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{}
}

func (s *ASTBuilder) VisitErrorNode(node antlr.ErrorNode) {
	panic(common.NewCompileErr(node.GetText()))
}

func (s *ASTBuilder) NodePush(node ast.ASTNode) {
	s.NodeStack = append(s.NodeStack, node)
}

func (s *ASTBuilder) NodeTop() ast.ASTNode {
	if len(s.NodeStack) == 0 {
		return nil
	}
	return s.NodeStack[len(s.NodeStack)-1]
}

func (s *ASTBuilder) NodePop() ast.ASTNode {
	ret := s.NodeTop()
	s.NodeStack = s.NodeStack[0 : len(s.NodeStack)-1]
	return ret
}
