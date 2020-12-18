package instruction

import "github.com/lwpyr/goscript/common"

func PushBack() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		ptr := stk.TopIndex(1).(*interface{})
		if *ptr == nil {
			*ptr = []interface{}{}
		}
		*ptr = append((*ptr).([]interface{}), stk.Top())
		stk.Pop()
		stk.Pc++
	}
}

func PushFront() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		ptr := stk.TopIndex(1).(*interface{})
		if *ptr == nil {
			*ptr = []interface{}{}
		}
		*ptr = append([]interface{}{stk.Top()}, (*ptr).([]interface{})...)
		stk.Pop()
		stk.Pc++
	}
}

func DeleteMessageField() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if stk.TopIndex(1) != nil {
			DelMapStrField(stk.TopIndex(1), stk.Top())
		}
		stk.Pop()
		stk.Pc++
	}
}

func DeleteMapKey(mapDelFunc func(m interface{}, k interface{})) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if stk.TopIndex(1) != nil {
			mapDelFunc(stk.TopIndex(1), stk.Top())
		}
		stk.Pop()
		stk.Pc++
	}
}

func Len(lengthFunc LengthFunc) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Set(0, lengthFunc(stk.Top()))
		stk.Pc++
	}
}

func EnumString(rEnum map[int32]string) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Set(0, rEnum[stk.Top().(int32)])
		stk.Pc++
	}
}
