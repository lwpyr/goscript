package instruction

import "github.com/lwpyr/goscript/common"

func GetInitializeMapFunc(constructor common.Constructor, mapSet func(m interface{}, k interface{}, v interface{}), num int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		ret := constructor()
		for i := 0; i < num; i++ {
			mapSet(ret, stk.TopIndex(i+num), stk.TopIndex(i))
		}
		stk.PopN(num * 2)
		stk.Push(ret)
		stk.Pc++
	}
}

func GetInitializeMessageFunc(keys []string, num int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		ret := common.NewMessageValue()
		for j := 0; j < num; j++ {
			MessageSetField(ret, keys[j], stk.Top())
			stk.Pop()
		}
		stk.Push(ret)
		stk.Pc++
	}
}

func GetInitializeSliceFunc(num int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		items := make([]interface{}, num)
		for j := 0; j < num; j++ {
			items[j] = stk.Top()
			stk.Pop()
		}
		stk.Push(items)
		stk.Pc++
	}
}
