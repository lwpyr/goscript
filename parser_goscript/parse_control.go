package parser_goscript

import (
	"github.com/lwpyr/goscript/ast"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
)

// ExitRunnableExec is called when production RunnableExec is exited.
func (s *ASTBuilder) ExitRestoreStack(ctx *parser.RestoreStackContext) {
	node := &ast.RestoreStackNode{
		Line: s.NodePop(),
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitBlock is called when production block is exited.
func (s *ASTBuilder) ExitBlock(ctx *parser.BlockContext) {
	node := &ast.BlockNode{}
	num := len(ctx.AllExecution())
	node.Executions = make([]ast.ASTNode, num)
	for i := num - 1; i >= 0; i-- {
		node.Executions[i] = s.NodePop()
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

func (s *ASTBuilder) ExitIf(ctx *parser.IfContext) {
	node := &ast.IfNode{}
	if ctx.Control() != nil || len(ctx.AllBlock()) == 2 {
		node.ElseBlock = s.NodePop()
	}
	node.Block = s.NodePop()
	node.Condition = s.NodePop()
	node.SetSource(ctx)
	s.NodePush(node)
}

// EnterControlReturnVoid is called when production ControlReturnVoid is entered.
func (s *ASTBuilder) EnterReturnVoid(ctx *parser.ReturnVoidContext) {
	if s.FunctionDepth == 0 {
		panic(common.NewCompileErr("cannot return outside a function"))
	}
}

// ExitControlReturnVoid is called when production ControlReturnVoid is exited.
func (s *ASTBuilder) ExitReturnVoid(ctx *parser.ReturnVoidContext) {
	node := &ast.ReturnNode{}
	node.SetSource(ctx)
	s.NodePush(node)
}

// EnterControlReturnVal is called when production ControlReturnVal is entered.
func (s *ASTBuilder) EnterReturnVal(ctx *parser.ReturnValContext) {
	if s.FunctionDepth == 0 {
		panic(common.NewCompileErr(common.FormatError(ctx, "cannot return outside a function ")))
	}
}

// ExitControlReturnVal is called when production ControlReturnVal is exited.
func (s *ASTBuilder) ExitReturnVal(ctx *parser.ReturnValContext) {
	node := &ast.ReturnNode{}
	for i := 0; i < len(ctx.AllExpr()); i++ {
		node.Expr = append(node.Expr, s.NodePop())
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// EnterControlBreak is called when production ControlBreak is entered.
func (s *ASTBuilder) EnterBreak(ctx *parser.BreakContext) {
	if s.LoopDepth == 0 {
		panic(common.NewCompileErr(common.FormatError(ctx, "cannot break outside a loop")))
	}
}

// ExitControlBreak is called when production ControlBreak is exited.
func (s *ASTBuilder) ExitBreak(ctx *parser.BreakContext) {
	node := &ast.BreakNode{}
	node.SetSource(ctx)
	s.NodePush(node)
}

// EnterControlContinue is called when production ControlContinue is entered.
func (s *ASTBuilder) EnterContinue(ctx *parser.ContinueContext) {
	if s.LoopDepth == 0 {
		panic(common.NewCompileErr(common.FormatError(ctx, "cannot break outside a loop")))
	}
}

// ExitControlContinue is called when production ControlContinue is exited.
func (s *ASTBuilder) ExitContinue(ctx *parser.ContinueContext) {
	node := &ast.ContinueNode{}
	node.SetSource(ctx)
	s.NodePush(node)
}

func (s *ASTBuilder) EnterForInSlice(ctx *parser.ForInSliceContext) {
	s.LoopDepth++
}

// ExitControlForInSlice is called when production ControlForInSlice is exited.
func (s *ASTBuilder) ExitForInSlice(ctx *parser.ForInSliceContext) {
	s.LoopDepth--
	node := &ast.ForSliceNode{
		ItemName: ctx.Name().GetText(),
	}
	node.Serial = s.NodePop()
	node.Slice = s.NodePop()
	node.SetSource(ctx)
	s.NodePush(node)
}

func (s *ASTBuilder) EnterForInMap(ctx *parser.ForInMapContext) {
	s.LoopDepth++
}

// ExitControlForInMap is called when production ControlForInMap is exited.
func (s *ASTBuilder) ExitForInMap(ctx *parser.ForInMapContext) {
	s.LoopDepth--
	node := &ast.ForMapNode{
		KeyName: ctx.Name(0).GetText(),
		ValName: ctx.Name(1).GetText(),
	}
	node.Serial = s.NodePop()
	node.Map = s.NodePop()
	node.SetSource(ctx)
	s.NodePush(node)
}

// EnterControlFor is called when production ControlFor is entered.
func (s *ASTBuilder) EnterFor(ctx *parser.ForContext) {
	s.LoopDepth++
}

// ExitControlFor is called when production ControlFor is exited.
func (s *ASTBuilder) ExitFor(ctx *parser.ForContext) {
	s.LoopDepth--
	node := &ast.ForNode{}
	node.Block = s.NodePop()
	node.Step = s.NodePop()
	node.Condition = s.NodePop()
	node.Init = s.NodePop()
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitControlSwitch is called when production ControlSwitch is exited.
func (s *ASTBuilder) ExitSwitch(ctx *parser.SwitchContext) {
	node := &ast.SwitchNode{}
	caseNum := len(ctx.AllConstant())
	node.Conditions = make([]interface{}, caseNum)
	node.Blocks = make([]ast.ASTNode, caseNum)
	for i := 0; i < caseNum; i++ {
		node.Blocks[caseNum-i-1] = s.NodePop()
		if caseNode := s.NodePop(); !caseNode.IsVariadic() {
			node.Conditions[caseNum-i-1] = caseNode.GetConstantValue()
		} else {
			panic(common.NewCompileErr("switch case should be a constant"))
			return
		}
	}
	node.SetSource(ctx)
	s.NodePush(node)
}
