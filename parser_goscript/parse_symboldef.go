package parser_goscript

import (
	"github.com/lwpyr/goscript/ast"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
)

func (s *ASTBuilder) ExitSymbolDef(ctx *parser.SymbolDefContext) {
	node := &ast.SymbolNode{}
	varType := common.Const
	if ctx.VAR() != nil {
		varType = common.Global
	} else if ctx.LOCAL() != nil || ctx.LOCALASSIGN() != nil {
		varType = common.Local
	}
	if ctx.Expr() != nil {
		node.ValueToAssign = s.NodePop()
	}
	if ctx.Typename() != nil {
		node.DataTypeNode = s.NodePop()
	}

	node.SymbolName = ctx.Name().GetText()
	node.SymbolType = varType
	node.SetSource(ctx)
	s.NodePush(node)
}
