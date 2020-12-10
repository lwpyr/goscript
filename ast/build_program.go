package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lwpyr/goscript/parser"
)

// EnterProgram is called when production program is entered.
func (s *ASTBuilder) EnterProgram(ctx *parser.ProgramContext) {
	cur := &ProgramRoot{
		Node: Node{
			Parent:   nil,
			Variadic: true,
		},
	}
	tr := NewTypeRegister()
	tr.Compiler = s.Compiler
	tb := NewTypeParser()
	tb.Compiler = s.Compiler
	// register all types declared
	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *parser.TypeDefEnumContext, *parser.TypeDefFunctionContext, *parser.TypeDefMapContext, *parser.TypeDefMessageContext, *parser.TypeDefSliceContext:
			antlr.ParseTreeWalkerDefault.Walk(tr, child)
		}
	}
	// parse all types declared
	var typeDefNodes []ITypeDefNode
	for _, child := range ctx.GetChildren() {
		switch child.(type) {
		case *parser.TypeDefEnumContext, *parser.TypeDefFunctionContext, *parser.TypeDefMapContext, *parser.TypeDefMessageContext, *parser.TypeDefSliceContext:
			antlr.ParseTreeWalkerDefault.Walk(tb, child)
			typeDefNodes = append(typeDefNodes, tb.NodePop().(ITypeDefNode))
		}
	}
	for _, node := range typeDefNodes {
		node.Compile(s.Compiler)
	}
	s.VisitPush(cur)
}

// ExitProgram is called when production program is exited.
func (s *ASTBuilder) ExitProgram(ctx *parser.ProgramContext) {
	node := s.VisitPop().(*ProgramRoot)
	num := len(ctx.AllExecution()) + len(ctx.AllFunctiondef())
	for i := 0; i < num; i++ {
		temp := s.NodePop()
		if temp == nil {
			// this is already done by TypeParser
			continue
		}
		switch temp.(type) {
		case *FunctionDefNode:
			node.FunctionDefNode = append(node.FunctionDefNode, temp)
		default:
			node.StatementNode = append([]ASTNode{temp}, node.StatementNode...)
		}
	}
	s.NodePush(node)
}
