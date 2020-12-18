package parser_goscript

import (
	"github.com/lwpyr/goscript/ast"
	"github.com/lwpyr/goscript/parser"
)

// ExitProgram is called when production program is exited.
func (s *ASTBuilder) ExitProgram(ctx *parser.ProgramContext) {
	root := &ast.ProgramRoot{}

	for i := 0; i < len(ctx.GetChildren()); i++ {
		node := s.NodePop()
		switch node.(type) {
		case *ast.TypeDefNode:
			root.TypeDefNode = append(root.TypeDefNode, node)
		default:
			root.StatementNode = append([]ast.ASTNode{node}, root.StatementNode...)
		}
	}
	s.NodePush(root)
}
