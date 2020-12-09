package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
)

// EnterVardef is called when production vardef is entered.
func (s *ASTBuilder) EnterVardef(ctx *parser.VardefContext) {
	s.VisitPush(&GlobalNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	})
}

func (s *ASTBuilder) ExitVardef(ctx *parser.VardefContext) {
	node := s.VisitPop().(*GlobalNode)
	var valNode ASTNode
	if ctx.Expr() != nil {
		valNode = s.NodePop()
	}
	typeNode := s.NodePop()
	node.Variable = &common.Variable{
		Symbol:       ctx.NAME().GetText(),
		VariableType: common.Global,
		Scope:        s.Compiler.Scope,
		DataType:     typeNode.GetDataType(),
	}
	s.Compiler.Scope.AddVariable(node.Variable)
	node.Assign = valNode
	s.NodePush(node)
}

// EnterLocalVariable is called when production LocalVariable is entered.
func (s *ASTBuilder) EnterLocaldef(ctx *parser.LocaldefContext) {
	s.VisitPush(&LocalNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	})
}

func (s *ASTBuilder) ExitLocaldef(ctx *parser.LocaldefContext) {
	node := s.VisitPop().(*LocalNode)
	var valNode ASTNode
	if ctx.Expr() != nil {
		valNode = s.NodePop()
	}
	typeNode := s.NodePop()
	node.Variable = &common.Variable{
		Symbol:       ctx.NAME().GetText(),
		VariableType: common.Global,
		Scope:        s.Compiler.Scope,
		DataType:     typeNode.GetDataType(),
	}
	s.Compiler.Scope.AddParameterVariable(node.Variable)
	node.Assign = valNode
	s.NodePush(node)
}
