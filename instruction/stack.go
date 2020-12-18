package instruction

import "github.com/lwpyr/goscript/common"

func StackPopN(n int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.PopN(n)
		stk.Pc++
	}
}

func StackReturnN(n int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		for i := 0; i < n; i++ {
			stk.Data[stk.Bp-i-1] = stk.Data[stk.Sp-i]
		}
		stk.Pc = -1
	}
}

func StackReturn() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Data[stk.Bp-1] = stk.Data[stk.Sp]
		stk.Pc = -1
	}
}

func StackReturnVoid() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Pc = -1
	}
}
