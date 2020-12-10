package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/lambda_chains"
	"github.com/lwpyr/goscript/parser"
)

// EnterVardef is called when production vardef is entered.
func (s *ASTBuilder) EnterVardef(ctx *parser.VardefContext) {
	s.VisitPush(&VariableDefNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		IsGlobal: ctx.VAR() != nil,
	})
}

func (s *ASTBuilder) ExitVardef(ctx *parser.VardefContext) {
	node := s.VisitPop().(*VariableDefNode)
	varType := common.Global
	if ctx.LOCAL() != nil {
		varType = common.Local
	}
	if ctx.InitializationListBegin() != nil {
		node.Assign = s.NodePop()
		typeNode := node.Assign.(*BinaryNode).Lhs
		node.Variable = &common.Variable{
			Symbol:       ctx.NAME().GetText(),
			VariableType: varType,
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
			VariableType: varType,
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

// ExitConstdef is called when production constdef is exited.
func (s *ASTBuilder) ExitConstdef(ctx *parser.ConstdefContext) {
	constName := ctx.NAME().GetText()
	constType := common.BasicTypeMap[ctx.BasicTypeName().GetText()]
	ValNode := s.NodePop().(IConstantNode)
	convert := lambda_chains.GetConvertFunc(ValNode.GetDataType(), constType)
	constValue := convert(ValNode.GetConstantValue())
	s.Compiler.Scope.AddConstant(constName, &common.Constant{
		Symbol: constName,
		Type:   constType,
		Data:   constValue,
	})
	s.NodePush(&VariableDefNode{IsGlobal: true, Assign: nil})
}
