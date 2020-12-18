package instruction

import "github.com/lwpyr/goscript/common"

func SelectorLhs(constructor common.Constructor, fieldName string, oneOfs []string) common.Instruction {
	if len(oneOfs) > 0 {
		return func(m *common.Memory, stk *common.Stack) {
			ptr := stk.Top().(*interface{})
			if *ptr == nil {
				*ptr = constructor()
			}
			for _, name := range oneOfs {
				MessageReset(*ptr, name)
			}
			ptr = MessageGetPtr(*ptr, fieldName)
			stk.Set(0, ptr)
			stk.Pc++
		}
	} else {
		return func(m *common.Memory, stk *common.Stack) {
			ptr := stk.Top().(*interface{})
			if *ptr == nil {
				*ptr = constructor()
			}
			ptr = MessageGetPtr(*ptr, fieldName)
			stk.Set(0, ptr)
			stk.Pc++
		}
	}
}

func Selector(fieldName string) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if stk.Top() != nil {
			stk.Set(0, MessageGet(stk.Top(), fieldName))
		}
		stk.Pc++
	}
}

func SelectorEnum(fieldName string) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if stk.Top() != nil {
			stk.Set(0, stk.Top().(map[string]int32)[fieldName])
		}
		stk.Pc++
	}
}
