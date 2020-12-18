package instruction

import "github.com/lwpyr/goscript/common"

func ConnectInstructions(instructions []common.Instruction) common.Instruction {
	instructions = append(instructions, GetProgramEnding())
	return func(m *common.Memory, stk *common.Stack) {
		for stk.Pc != -1 {
			instructions[stk.Pc](m, stk)
		}
	}
}

func GetProgramEnding() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Pc = -1
	}
}
