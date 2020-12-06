package ast

import "github.com/lw396285v/goscript/common"

type GlobalNode struct {
	Node
	Assign   *BinaryNode
	Variable *common.Variable
}

func (t *GlobalNode) Compile(c *Compiler) {
	panic("todo pre-allocate")
}

type LocalNode struct {
	Node
	Assign   *BinaryNode
	Variable *common.Variable
}

func (t *LocalNode) Compile(c *Compiler) {
	if t.Assign == nil {
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			stk.Push(nil)
			stk.Push(nil)
		})
	} else {
		t.Assign.Compile(c)
		assignInst := c.InstructionPop()
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			stk.Push(nil)
			assignInst(m, stk)
		})
	}
}
