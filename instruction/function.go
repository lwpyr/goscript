package instruction

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/hack"
)

func PushLambdaWithCapture(rawFunc common.Instruction, capturedVariables []*common.Symbol, closureVariable *common.Symbol) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		captured := make([]interface{}, 0, len(capturedVariables))
		for _, v := range capturedVariables {
			switch v.SymbolType {
			case common.Local:
				captured = append(captured, stk.Get(v))
			case common.Captured:
				captured = append(captured, *hack.SliceIndex(*stk.MustGet(closureVariable), int64(v.Offset)))
			}
		}
		var lambda common.Instruction = func(m *common.Memory, stk *common.Stack) {
			stk.Push(captured)
			rawFunc(m, stk)
		}
		stk.Push(&lambda)
		stk.Pc++
	}
}

func PushLambda(lambda common.Instruction) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Push(&lambda)
		stk.Pc++
	}
}

func Construct(constructor common.Constructor, numParam int) common.Instruction {
	if numParam == 0 {
		return func(m *common.Memory, stk *common.Stack) {
			stk.Push(constructor())
			stk.Pc++
		}
	} else {
		return func(m *common.Memory, stk *common.Stack) {
			params := make([]interface{}, 0, numParam)
			for i := 0; i < numParam; i++ {
				params = append(params, stk.Top())
				stk.Pop()
			}
			stk.Push(constructor(params...))
			stk.Pc++
		}
	}
}

func GetDynamicCallFunc(lenIn int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		funcPtr := stk.Top().(*common.Instruction)
		stk.Pop()
		bp := stk.Bp
		pc := stk.Pc
		stk.Bp = stk.Sp - lenIn + 1
		stk.Pc = 0
		(*funcPtr)(m, stk)
		stk.Sp = stk.Bp - 1
		stk.Bp = bp
		stk.Data = stk.Data[:stk.Sp+1]
		stk.Pc = pc + 1
		// stk.Call(*funcPtr, m, stk, lenIn)
	}
}

func GetStaticCallFunc(funcPtr *common.Instruction, lenIn int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		bp := stk.Bp
		pc := stk.Pc
		stk.Bp = stk.Sp - lenIn + 1
		stk.Pc = 0
		(*funcPtr)(m, stk)
		stk.Sp = stk.Bp - 1
		stk.Bp = bp
		stk.Data = stk.Data[:stk.Sp+1]
		stk.Pc = pc + 1
		// stk.Call(*funcPtr, m, stk, lenIn)
	}
}

func GetTailArray(num int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		temp := make([]interface{}, num)
		copy(temp, stk.Data[stk.Sp-num+1:stk.Sp+1])
		stk.Set(num-1, temp)
		stk.PopN(num - 1)
		stk.Pc++
	}
}

func GetAppendReturnPlace(num int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		for i := 0; i < num; i++ {
			stk.Push(nil)
		}
		stk.Pc++
	}
}
