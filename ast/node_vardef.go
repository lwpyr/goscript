package ast

import "github.com/lwpyr/goscript/common"

type GlobalNode struct {
	Node
	Assign   ASTNode
	Variable *common.Variable
}

func (t *GlobalNode) Compile(c *Compiler) {
	if t.Assign != nil {
		t.Assign.Compile(c)
		assignInst := c.InstructionPop()
		variable := t.Variable
		t.Instructions = []common.Instruction{
			func(m *common.Memory, stk *common.Stack) {
				assignInst(m, stk)
				*m.MustGet(variable) = stk.Top()
				stk.Pop()
				stk.Pc++
			},
		}
	}
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
