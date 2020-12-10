package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
)

// EnterVardef is called when production vardef is entered.
func (s *ASTBuilder) EnterVardef(ctx *parser.VardefContext) {
	s.VisitPush(&VariableDefNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		IsGlobal: true,
	})
}

func (s *ASTBuilder) ExitVardef(ctx *parser.VardefContext) {
	node := s.VisitPop().(*VariableDefNode)
	if ctx.InitializationListBegin() != nil {
		node.Assign = s.NodePop()
		typeNode := node.Assign.(*BinaryNode).Lhs
		node.Variable = &common.Variable{
			Symbol:       ctx.NAME().GetText(),
			VariableType: common.Global,
			Scope:        s.Compiler.Scope,
			DataType:     typeNode.GetDataType(),
		}
		node.Assign.(*BinaryNode).Lhs = &VarNode{
			Node: Node{
				DataType: node.Variable.DataType,
			},
			Variable: node.Variable,
		}
	} else if ctx.Expr() != nil {
		valNode := s.NodePop()
		typeNode := s.NodePop()
		node.Variable = &common.Variable{
			Symbol:       ctx.NAME().GetText(),
			VariableType: common.Global,
			Scope:        s.Compiler.Scope,
			DataType:     typeNode.GetDataType(),
		}
		node.Assign = &BinaryNode{
			Node: Node{},
			Lhs: &VarNode{
				Node: Node{
					DataType: node.Variable.DataType,
				},
				Variable: node.Variable,
			},
			Rhs: valNode,
			Op:  "=",
		}
	}

	s.Compiler.Scope.AddVariable(node.Variable)
	s.NodePush(node)
}

// EnterLocalVariable is called when production LocalVariable is entered.
func (s *ASTBuilder) EnterLocaldef(ctx *parser.LocaldefContext) {
	s.VisitPush(&VariableDefNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		IsGlobal: false,
	})
}

func (s *ASTBuilder) ExitLocaldef(ctx *parser.LocaldefContext) {
	node := s.VisitPop().(*VariableDefNode)
	if ctx.InitializationListBegin() != nil {
		node.Assign = s.NodePop()
		typeNode := node.Assign.(*BinaryNode).Lhs
		node.Variable = &common.Variable{
			Symbol:       ctx.NAME().GetText(),
			VariableType: common.Global,
			Scope:        s.Compiler.Scope,
			DataType:     typeNode.GetDataType(),
		}
		node.Assign.(*BinaryNode).Lhs = &VarNode{
			Node: Node{
				DataType: node.Variable.DataType,
			},
			Variable: node.Variable,
		}
	} else if ctx.Expr() != nil {
		valNode := s.NodePop()
		typeNode := s.NodePop()
		node.Variable = &common.Variable{
			Symbol:       ctx.NAME().GetText(),
			VariableType: common.Global,
			Scope:        s.Compiler.Scope,
			DataType:     typeNode.GetDataType(),
		}
		node.Assign = &BinaryNode{
			Node: Node{},
			Lhs: &VarNode{
				Node: Node{
					DataType: node.Variable.DataType,
				},
				Variable: node.Variable,
			},
			Rhs: valNode,
			Op:  "=",
		}
	}

	s.Compiler.Scope.AddVariable(node.Variable)
	s.NodePush(node)
}
