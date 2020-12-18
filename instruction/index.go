package instruction

import "github.com/lwpyr/goscript/common"

func SliceFilter(filter common.Instruction, boolConvertFunc TypeConvertFuncBool) common.Instruction {
	lengthFunc := SliceLen()
	return func(m *common.Memory, stk *common.Stack) {
		lenSlice := lengthFunc(stk.Top())
		val := make([]interface{}, 0, lenSlice)
		for i := int64(0); i < lenSlice; i++ {
			item := *GetSliceField(stk.Top(), i)
			stk.Push(nil)
			stk.Push(item)

			// make a call
			// do you know why this is not wrapper in a function?
			bp := stk.Bp
			pc := stk.Pc
			stk.Bp = stk.Sp
			stk.Pc = 0
			filter(m, stk)
			stk.Sp = stk.Bp - 1
			stk.Bp = bp
			stk.Data = stk.Data[:stk.Sp+1]
			stk.Pc = pc
			// answer is we cannot force inline function

			if boolConvertFunc(stk.Top()) == true {
				val = append(val, item)
			}
			stk.Pop()
		}
		stk.Set(0, val)
		stk.Pc++
	}
}

func StringArrIndex() common.Instruction {
	lengthFunc := StringLen()
	return func(m *common.Memory, stk *common.Stack) {
		slice := stk.Top()
		stk.Pop()
		index := (stk.Top()).([]int64)
		if index[0] == -1 && index[1] == -1 {
			index[0] = index[2]
			index[1] = index[2] + 1
			index[2] = 1
		} else if index[0] == -1 {
			index[0] = 0
		} else if index[1] == -1 {
			index[1] = lengthFunc(slice)
		}
		stk.Set(0, (slice.(string))[index[0]:index[1]])
		stk.Pc++
	}
}

func SliceArrIndex(numIndex int) common.Instruction {
	lengthFunc := SliceLen()
	return func(m *common.Memory, stk *common.Stack) {
		slice := stk.Top()
		stk.Pop()
		lenSlice := lengthFunc(slice)
		lenVal := int64(0)
		indices := make([][]int64, 0, numIndex)
		for i := 0; i < numIndex; i++ {
			index := (stk.Top()).([]int64)
			if index[0] == -1 && index[1] == -1 {
				index[0] = index[2]
				index[1] = index[2] + 1
				index[2] = 1
			} else if index[0] == -1 {
				index[0] = 0
			} else if index[1] == -1 {
				index[1] = lenSlice
			}
			lenIndex := (index[1]-index[0]-1)/index[2] + 1
			if lenIndex < 0 {
				lenIndex = 0
			}
			lenVal += lenIndex
			indices = append(indices, index)
			stk.Pop()
		}
		val := make([]interface{}, 0, lenVal)
		for _, index := range indices {
			for i := index[0]; i < index[1]; i += index[2] {
				val = append(val, *GetSliceField(slice, i))
			}
		}
		stk.Push(val)
		stk.Pc++
	}
}

func SliceIndexLhs() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		stk.Set(1, GetSliceField(*stk.Top().(*interface{}), stk.TopIndex(1).(int64)))
		stk.Pop()
		stk.Pc++
	}
}

func SliceIndex() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if stk.Top() != nil {
			stk.Set(1, *GetSliceField(stk.Top(), stk.TopIndex(1).(int64)))
			stk.Pop()
		} else {
			stk.Pop()
			stk.Set(0, nil)
		}
		stk.Pc++
	}
}

func StringIndex() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if stk.Top() != nil {
			stk.Set(1, string((stk.Top().(string))[stk.TopIndex(1).(int64)]))
			stk.Pop()
		}
		stk.Pc++
	}
}

func MapIndexLhs(constructor common.Constructor, mapMustGet func(m interface{}, k interface{}) (v *interface{})) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		ptr := stk.Top().(*interface{})
		stk.Pop()
		if *ptr == nil {
			*ptr = constructor()
		}
		stk.Set(0, mapMustGet(*ptr, stk.Top()))
		stk.Pc++
	}
}

func MapIndex(mapGet func(m interface{}, k interface{}) (v interface{})) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if stk.Top() != nil {
			stk.Set(1, mapGet(stk.Top(), stk.TopIndex(1)))
			stk.Pop()
		}
		stk.Pc++
	}
}

func MapArrIndex(mapGet func(m interface{}, k interface{}) (v interface{}), numFields int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		mapValue := stk.Top()
		stk.Pop()
		if mapValue != nil {
			val := make([]interface{}, numFields)
			for i := 0; i < numFields; i++ {
				if stk.Top() != nil {
					val[i] = mapGet(mapValue, stk.Top())
				} else {
					val[i] = nil
				}
				stk.Pop()
			}
			stk.Push(val)
		} else {
			stk.PopN(numFields)
			stk.Push(nil)
		}
		stk.Pc++
	}
}

func ArrIndex() common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		step := stk.Top()
		stk.Pop()
		to := stk.Top()
		stk.Pop()
		from := stk.Top()
		stk.Set(0, []int64{from.(int64), to.(int64), step.(int64)})
		stk.Pc++
	}
}
