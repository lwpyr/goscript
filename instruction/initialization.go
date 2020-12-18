package instruction

import "github.com/lwpyr/goscript/common"

func InitializeMap(constructor common.Constructor, mapSet func(m interface{}, k interface{}, v interface{}), num int) common.Instruction {
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

func InitializeMessage(keys []string, num int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		ret := common.NewMessageValue()
		for j := 0; j < num; j++ {
			MessageSet(ret, keys[j], stk.Top())
			stk.Pop()
		}
		stk.Push(ret)
		stk.Pc++
	}
}

func InitializeSlice(num int) common.Instruction {
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
