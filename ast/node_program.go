package ast

import "github.com/lwpyr/goscript/common"

func (p *ProgramRoot) Compile(c *Compiler) {
	p.BuildFunction(c)
	p.BuildStatement(c)
}

func (p *ProgramRoot) BuildFunction(c *Compiler) {
	for _, fDefNode := range p.FunctionDefNode {
		fDefNode.Compile(c)
	}
}

func (p *ProgramRoot) BuildStatement(c *Compiler) {
	num := len(p.StatementNode)
	var runnableInstructions []common.Instruction
	for i := 0; i < num; i++ {
		p.StatementNode[i].Compile(c)
		runnableInstructions = append(runnableInstructions, p.StatementNode[i].GetInstructions()...)
	}
	runnableInstructions = append(runnableInstructions, func(m *common.Memory, stk *common.Stack) {
		stk.Pc = -1
	}) // to make sure the program end
	c.InstructionPush(func(m *common.Memory, stk *common.Stack) {
		for stk.Pc != -1 {
			runnableInstructions[stk.Pc](m, stk)
		}
	})
}
