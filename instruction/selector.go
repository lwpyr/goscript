package instruction

import "github.com/lwpyr/goscript/common"

func GetSelectFuncLhs(constructor common.Constructor, fieldName string, oneOfs []string) common.Instruction {
	if len(oneOfs) > 0 {
		return func(m *common.Memory, stk *common.Stack) {
			ptr := stk.Top().(*interface{})
			if *ptr == nil {
				*ptr = constructor()
			}
			for _, name := range oneOfs {
				MessageResetField(*ptr, name)
			}
			ptr = MessageMustGetField(*ptr, fieldName)
			stk.Set(0, ptr)
			stk.Pc++
		}
	} else {
		return func(m *common.Memory, stk *common.Stack) {
			ptr := stk.Top().(*interface{})
			if *ptr == nil {
				*ptr = constructor()
			}
			ptr = MessageMustGetField(*ptr, fieldName)
			stk.Set(0, ptr)
			stk.Pc++
		}
	}
}

func GetSelectFunc(fieldName string) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if stk.Top() != nil {
			stk.Set(0, MessageGetField(stk.Top(), fieldName))
		}
		stk.Pc++
	}
}

func GetSelectEnumFunc(fieldName string) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if stk.Top() != nil {
			stk.Set(0, stk.Top().(map[string]int32)[fieldName])
		}
		stk.Pc++
	}
}
