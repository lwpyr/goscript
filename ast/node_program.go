package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/instruction"
)

func (p *ProgramRoot) Compile(c *CompileContext) {
	// first compile all the type definition
	for _, node := range p.TypeDefNode {
		node.Compile(c)
	}
	// then compile all the last statement
	num := len(p.StatementNode)
	var runnableInstructions []common.Instruction
	for i := 0; i < num; i++ {
		p.StatementNode[i].Compile(c)
		runnableInstructions = append(runnableInstructions, p.StatementNode[i].GetInstructions()...)
	}
	runnableInstructions = append(runnableInstructions, instruction.ProgramEnding())
	p.Instructions = []common.Instruction{
		// to make sure the program end
		func(m *common.Memory, stk *common.Stack) {
			for stk.Pc != -1 {
				runnableInstructions[stk.Pc](m, stk)
			}
		},
	}
}
