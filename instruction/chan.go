package instruction

import "github.com/lwpyr/goscript/common"

func ChanSend(nonBlock bool) common.Instruction {
	if nonBlock {
		return func(m *common.Memory, stk *common.Stack) {
			select {
			case stk.TopIndex(1).(chan interface{}) <- stk.Top():
				stk.PopN(2)
				stk.Push(true)
			default:
				stk.PopN(2)
				stk.Push(false)
			}
			stk.Pc++
		}
	} else {
		return func(m *common.Memory, stk *common.Stack) {
			stk.TopIndex(1).(chan interface{}) <- stk.Top()
			stk.PopN(2)
			stk.Push(true)
			stk.Pc++
		}
	}
}

func ChanRecv(nonBlock bool) common.Instruction {
	if nonBlock {
		return func(m *common.Memory, stk *common.Stack) {
			select {
			case temp := <-stk.Top().(chan interface{}):
				stk.Set(0, temp)
			default:
				stk.Set(0, nil)
			}
			stk.Pc++
		}
	} else {
		return func(m *common.Memory, stk *common.Stack) {
			stk.Set(0, <-stk.Top().(chan interface{}))
			stk.Pc++
		}
	}
}
