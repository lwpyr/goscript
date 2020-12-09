package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/lambda_chains"
)

type GlobalNode struct {
	Node
	Assign   ASTNode
	Variable *common.Variable
}

func (t *GlobalNode) Compile(c *Compiler) {
	if t.Assign != nil {
		if !t.Assign.GetDataType().CanConvertTo(t.Variable.DataType) {
			panic(common.NewTypeErr("try to use a wrong type value to initialize global variable " + t.Variable.Symbol))
		}
		t.Assign.Compile(c)
		assignInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFunc(t.Assign.GetDataType(), t.Variable.DataType)
		v := t.Variable
		t.Instructions = []common.Instruction{
			func(m *common.Memory, stk *common.Stack) {
				assignInst(m, stk)
				*m.MustGet(v) = convertFunc(stk.Top())
				stk.Pop()
				stk.Pc++
			},
		}
	}
}

type LocalNode struct {
	Node
	Assign   ASTNode
	Variable *common.Variable
}

func (t *LocalNode) Compile(c *Compiler) {
	if t.Assign == nil {
		c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
			stk.Push(nil)
			stk.Pc++
		})
	} else {
		if !t.Assign.GetDataType().CanConvertTo(t.Variable.DataType) {
			panic(common.NewTypeErr("try to use a wrong type value to initialize global variable " + t.Variable.Symbol))
		}
		t.Assign.Compile(c)
		assignInst := c.InstructionPop()
		convertFunc := lambda_chains.GetConvertFunc(t.Assign.GetDataType(), t.Variable.DataType)
		t.Instructions = []common.Instruction{
			func(m *common.Memory, stk *common.Stack) {
				assignInst(m, stk)
				stk.Set(0, convertFunc(stk.Top()))
				stk.Pc++
			},
		}
	}
}
