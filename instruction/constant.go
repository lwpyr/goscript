package instruction

import "github.com/lwpyr/goscript/common"

func PushConstantToStack(data interface{}) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Push(data)
		stk.Pc++
	}
}
