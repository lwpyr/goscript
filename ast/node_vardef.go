package ast

import (
	"github.com/lwpyr/goscript/common"
)

type VariableDefNode struct {
	Node
	IsGlobal bool
	Assign   ASTNode
	Variable *common.Variable
}

func (t *VariableDefNode) Compile(c *Compiler) {
	if t.IsGlobal {
		if t.Assign != nil {
			t.Assign.Compile(c)
			assignInst := c.InstructionPop()
			t.Instructions = []common.Instruction{
				func(m *common.Memory, stk *common.Stack) {
					assignInst(m, stk)
					stk.Pop()
					stk.Pc++
				},
			}
		}
	} else {
		if t.Assign == nil {
			c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
				stk.Push(nil)
				stk.Pc++
			})
		} else {
			t.Assign.Compile(c)
			assignInst := c.InstructionPop()
			t.Instructions = []common.Instruction{
				func(m *common.Memory, stk *common.Stack) {
					stk.Push(nil)
					assignInst(m, stk)
					stk.Pop()
					stk.Pc++
				},
			}
		}
	}
}
