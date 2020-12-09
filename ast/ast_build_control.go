package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
)

//// EnterLineProgram is called when production LineProgram is entered.
//func (s *ASTBuilder) EnterLineProgram(ctx *parser.LineProgramContext) {
//	s.VisitPush(&LineNode{
//		Node: Node{
//			Parent:   s.VisitTop(),
//			Variadic: true,
//		},
//	})
//}
//
//// ExitLineProgram is called when production LineProgram is exited.
//func (s *ASTBuilder) ExitLineProgram(ctx *parser.LineProgramContext) {
//	node := s.VisitPop().(*LineNode)
//	node.Line = s.NodePop()
//	s.NodePush(node)
//}

// EnterRunnableExec is called when production RunnableExec is entered.
func (s *ASTBuilder) EnterRestoreStack(ctx *parser.RestoreStackContext) {
	s.VisitPush(&RestoreStackNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	})
}

// ExitRunnableExec is called when production RunnableExec is exited.
func (s *ASTBuilder) ExitRestoreStack(ctx *parser.RestoreStackContext) {
	node := s.VisitPop().(*RestoreStackNode)
	node.Line = s.NodePop()
	s.NodePush(node)
}

// EnterBlock is called when production block is entered.
func (s *ASTBuilder) EnterBlock(ctx *parser.BlockContext) {
	s.Compiler.Scope = common.NewScope(s.Compiler.Scope)
	s.VisitPush(&BlockNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		Scope: s.Compiler.Scope,
	})
}

// ExitBlock is called when production block is exited.
func (s *ASTBuilder) ExitBlock(ctx *parser.BlockContext) {
	s.Compiler.Scope = s.Compiler.Scope.Outer
	node := s.VisitPop().(*BlockNode)
	num := len(ctx.AllExecution())
	node.Executions = make([]ASTNode, num)
	for i := num - 1; i >= 0; i-- {
		temp := s.NodePop()
		node.Executions[i] = temp
	}
	s.NodePush(node)
}

func (s *ASTBuilder) EnterIf(ctx *parser.IfContext) {
	cur := &IfNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

func (s *ASTBuilder) ExitIf(ctx *parser.IfContext) {
	node := s.VisitPop().(*IfNode)
	if ctx.Control() != nil || len(ctx.AllBlock()) == 2 {
		node.ElseBlock = s.NodePop()
	}
	node.Block = s.NodePop()
	node.Condition = s.NodePop()
	s.NodePush(node)
}

// EnterControlReturnVoid is called when production ControlReturnVoid is entered.
func (s *ASTBuilder) EnterReturnVoid(ctx *parser.ReturnVoidContext) {
	if len(s.Closure) == 0 {
		panic(common.NewCompileErr("cannot return outside a function"))
	}
	cur := &ReturnNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitControlReturnVoid is called when production ControlReturnVoid is exited.
func (s *ASTBuilder) ExitReturnVoid(ctx *parser.ReturnVoidContext) {
	s.NodePush(s.VisitPop())
}

// EnterControlReturnVal is called when production ControlReturnVal is entered.
func (s *ASTBuilder) EnterReturnVal(ctx *parser.ReturnValContext) {
	if len(s.Closure) == 0 {
		panic(common.NewCompileErr("cannot return outside a function"))
	}
	cur := &ReturnNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitControlReturnVal is called when production ControlReturnVal is exited.
func (s *ASTBuilder) ExitReturnVal(ctx *parser.ReturnValContext) {
	node := s.VisitPop().(*ReturnNode)
	for i := 0; i < len(ctx.AllExpr()); i++ {
		node.Expr = append(node.Expr, s.NodePop())
	}
	s.NodePush(node)
}

// EnterControlBreak is called when production ControlBreak is entered.
func (s *ASTBuilder) EnterBreak(ctx *parser.BreakContext) {
	if s.LoopDepth == 0 {
		panic(common.NewCompileErr("cannot break outside a loop"))
	}
	cur := &BreakNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitControlBreak is called when production ControlBreak is exited.
func (s *ASTBuilder) ExitBreak(ctx *parser.BreakContext) {
	s.NodePush(s.VisitPop())
}

// EnterControlContinue is called when production ControlContinue is entered.
func (s *ASTBuilder) EnterContinue(ctx *parser.ContinueContext) {
	if s.LoopDepth == 0 {
		panic(common.NewCompileErr("cannot continue outside a loop"))
	}
	cur := &ContinueNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitControlContinue is called when production ControlContinue is exited.
func (s *ASTBuilder) ExitContinue(ctx *parser.ContinueContext) {
	s.NodePush(s.VisitPop())
}

// ExitCollection is called when production collection is exited.
func (s *ASTBuilder) ExitCollection(ctx *parser.CollectionContext) {
	collection := s.NodeTop()
	forNode := s.VisitTop()
	if collection.GetDataType().Kind.Kind == common.Slice {
		if node, ok := forNode.(*ForSliceNode); ok {
			s.Compiler.Scope.AddParameterVariable(&common.Variable{
				Symbol:   "$",
				DataType: common.BasicTypeMap[common.Int64Type],
				Scope:    s.Compiler.Scope,
			})
			s.Compiler.Scope.AddParameterVariable(&common.Variable{
				Symbol:   node.ItemName,
				DataType: collection.GetDataType().ItemType,
				Scope:    s.Compiler.Scope,
			})
			node.Item = s.Compiler.Scope.GetVariable(node.ItemName)
		} else {
			panic(common.NewCompileErr("for each in map get a slice"))
		}
	} else if collection.GetDataType().Kind.Kind == common.Map {
		if node, ok := forNode.(*ForMapNode); ok {
			s.Compiler.Scope.AddParameterVariable(&common.Variable{
				Symbol:   "$",
				DataType: common.BasicTypeMap[common.ObjectType],
				Scope:    s.Compiler.Scope,
			})
			node.Key = s.Compiler.Scope.GetVariable(node.KeyName)
			s.Compiler.Scope.AddParameterVariable(&common.Variable{
				Symbol:   node.KeyName,
				DataType: collection.GetDataType().KeyType,
				Scope:    s.Compiler.Scope,
			})
			node.Key = s.Compiler.Scope.GetVariable(node.KeyName)
			s.Compiler.Scope.AddParameterVariable(&common.Variable{
				Symbol:   node.ValName,
				DataType: collection.GetDataType().ValueType,
				Scope:    s.Compiler.Scope,
			})
			node.Val = s.Compiler.Scope.GetVariable(node.ValName)
		} else {
			panic(common.NewCompileErr("for each in map slice a map"))
		}
	} else {
		panic(common.NewCompileErr("for each collection type only support map and slice"))
	}
}

// EnterControlForInSlice is called when production ControlForInSlice is entered.
func (s *ASTBuilder) EnterForInSlice(ctx *parser.ForInSliceContext) {
	s.LoopDepth++
	s.Compiler.Scope = common.NewScope(s.Compiler.Scope)

	cur := &ForSliceNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		ItemName: ctx.NAME().GetText(),
	}
	s.VisitPush(cur)
}

// ExitControlForInSlice is called when production ControlForInSlice is exited.
func (s *ASTBuilder) ExitForInSlice(ctx *parser.ForInSliceContext) {
	s.LoopDepth--
	s.Compiler.Scope = s.Compiler.Scope.Outer
	node := s.VisitPop().(*ForSliceNode)
	node.Serial = s.NodePop()
	node.Slice = s.NodePop()
	s.NodePush(node)
}

// EnterControlForInMap is called when production ControlForInMap is entered.
func (s *ASTBuilder) EnterForInMap(ctx *parser.ForInMapContext) {
	s.LoopDepth++
	s.Compiler.Scope = common.NewScope(s.Compiler.Scope)

	cur := &ForMapNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		KeyName: ctx.NAME(0).GetText(),
		ValName: ctx.NAME(1).GetText(),
	}
	s.VisitPush(cur)
}

// ExitControlForInMap is called when production ControlForInMap is exited.
func (s *ASTBuilder) ExitForInMap(ctx *parser.ForInMapContext) {
	s.LoopDepth--
	s.Compiler.Scope = s.Compiler.Scope.Outer
	node := s.VisitPop().(*ForMapNode)
	node.Serial = s.NodePop()
	node.Map = s.NodePop()
	s.NodePush(node)
}

// EnterControlFor is called when production ControlFor is entered.
func (s *ASTBuilder) EnterFor(ctx *parser.ForContext) {
	s.LoopDepth++
	s.Compiler.Scope = common.NewScope(s.Compiler.Scope)

	cur := &ForNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		Scope: s.Compiler.Scope,
	}
	s.VisitPush(cur)
}

// ExitControlFor is called when production ControlFor is exited.
func (s *ASTBuilder) ExitFor(ctx *parser.ForContext) {
	s.LoopDepth--
	s.Compiler.Scope = s.Compiler.Scope.Outer

	node := s.VisitPop().(*ForNode)
	node.Block = s.NodePop()
	node.Step = s.NodePop()
	node.Condition = s.NodePop()
	node.Init = s.NodePop()
	s.NodePush(node)
}

// EnterControlSwitch is called when production ControlSwitch is entered.
func (s *ASTBuilder) EnterSwitch(ctx *parser.SwitchContext) {
	cur := &SwitchNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitControlSwitch is called when production ControlSwitch is exited.
func (s *ASTBuilder) ExitSwitch(ctx *parser.SwitchContext) {
	node := s.VisitPop().(*SwitchNode)
	caseNum := len(ctx.AllConstant())
	node.Conditions = make([]interface{}, caseNum)
	node.Blocks = make([]ASTNode, caseNum)
	for i := 0; i < caseNum; i++ {
		node.Blocks[caseNum-i-1] = s.NodePop()
		if caseNode, ok := s.NodePop().(IConstantNode); ok {
			node.Conditions[caseNum-i-1] = caseNode.GetConstantValue()
		} else {
			panic(common.NewCompileErr("switch case should be a constant"))
			return
		}
	}
	s.NodePush(node)
}
