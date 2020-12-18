package instruction

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/hack"
)

func GetStackOffsetToStackPtr(spOffset int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Set(0, &stk.Data[stk.TopIndex(spOffset).(int)+stk.Bp])
		stk.Pc++
	}
}

func GetIndirect() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Set(0, *stk.Top().(*interface{}))
		stk.Pc++
	}
}

func GetTakePtr() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		temp := stk.Top()
		stk.Set(0, &temp)
		stk.Pc++
	}
}

func GetLocalSymbolNil() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Push(nil)
		stk.Pc++
	}
}

func GetGlobalSymbolAssign(sym *common.Symbol) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		*m.MustGet(sym) = stk.Top()
		stk.Pop()
		stk.Pc++
	}
}

func GetFetchSymbolFuncLhs(symbol *common.Symbol, scope *common.Scope) common.Instruction {
	switch symbol.SymbolType {
	case common.Global:
		return func(m *common.Memory, stk *common.Stack) {
			stk.Push(m.MustGet(symbol))
			stk.Pc++
		}
	case common.Local:
		return func(m *common.Memory, stk *common.Stack) {
			stk.Push(symbol.Offset)
			stk.Pc++
		}
	case common.Captured:
		captures := scope.GetSymbol("#")
		return func(m *common.Memory, stk *common.Stack) {
			stk.Push(hack.SliceIndex(stk.Get(captures), int64(symbol.Offset)))
			stk.Pc++
		}
	default:
		return common.ErrorInstruction
	}
}

func GetFetchSymbolFunc(symbol *common.Symbol, scope *common.Scope) common.Instruction {
	switch symbol.SymbolType {
	case common.Global:
		return func(m *common.Memory, stk *common.Stack) {
			stk.Push(m.Get(symbol))
			stk.Pc++
		}
	case common.Local:
		return func(m *common.Memory, stk *common.Stack) {
			stk.Push(stk.Get(symbol))
			stk.Pc++
		}
	case common.Captured:
		captures := scope.GetSymbol("#")
		return func(m *common.Memory, stk *common.Stack) {
			stk.Push(*hack.SliceIndex(stk.Get(captures), int64(symbol.Offset)))
			stk.Pc++
		}
	case common.Const:
		val := symbol.Data
		return func(m *common.Memory, stk *common.Stack) {
			stk.Push(val)
			stk.Pc++
		}
	default:
		return common.ErrorInstruction
	}
}
