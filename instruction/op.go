package instruction

import (
	"github.com/lwpyr/goscript/common"
	"math"
	"unsafe"
)

func FunctionAssign(convert []common.Instruction) []common.Instruction {
	num := len(convert)
	var ret []common.Instruction
	for i := 0; i < num; i++ {
		if convert[i] != nil {
			ret = append(ret, convert[i])
		}
		ret = append(ret, func(m *common.Memory, stk *common.Stack) {
			*stk.Top().(*interface{}) = stk.TopIndex(num)
			stk.Pop()
			stk.Pc++
		})
	}
	return ret
}

func Assign(op string, lhs *common.DataType) common.Instruction {
	switch op {
	case "=":
		return func(m *common.Memory, stk *common.Stack) {
			*stk.Top().(*interface{}) = stk.TopIndex(1)
			stk.Pop()
			stk.Pc++
		}
	case "+=":
		switch lhs.Kind.Kind {
		case common.UInt32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*uint32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) + stk.TopIndex(1).(uint32)
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*uint64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) + stk.TopIndex(1).(uint64)
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*int32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) + stk.TopIndex(1).(int32)
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*int64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) + stk.TopIndex(1).(int64)
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*float32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) + stk.TopIndex(1).(float32)
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*float64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) + stk.TopIndex(1).(float64)
				stk.Pop()
				stk.Pc++
			}
		case common.String:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*string)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) + stk.TopIndex(1).(string)
				stk.Pop()
				stk.Pc++
			}
		}
	case "-=":
		switch lhs.Kind.Kind {
		case common.UInt32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*uint32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) - stk.TopIndex(1).(uint32)
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*uint64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) - stk.TopIndex(1).(uint64)
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*int32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) - stk.TopIndex(1).(int32)
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*int64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) - stk.TopIndex(1).(int64)
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*float32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) - stk.TopIndex(1).(float32)
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*float64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) - stk.TopIndex(1).(float64)
				stk.Pop()
				stk.Pc++
			}
		}
	case "*=":
		switch lhs.Kind.Kind {
		case common.UInt32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*uint32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) * stk.TopIndex(1).(uint32)
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*uint64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) * stk.TopIndex(1).(uint64)
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*int32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) * stk.TopIndex(1).(int32)
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*int64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) * stk.TopIndex(1).(int64)
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*float32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) * stk.TopIndex(1).(float32)
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*float64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) * stk.TopIndex(1).(float64)
				stk.Pop()
				stk.Pc++
			}
		}
	case "/=":
		switch lhs.Kind.Kind {
		case common.UInt32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*uint32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) / stk.TopIndex(1).(uint32)
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*uint64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) / stk.TopIndex(1).(uint64)
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*int32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) / stk.TopIndex(1).(int32)
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*int64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) / stk.TopIndex(1).(int64)
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*float32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) / stk.TopIndex(1).(float32)
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			return func(_ *common.Memory, stk *common.Stack) {
				*(stk.TopIndex(0).(*interface{})) = *(*float64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) / stk.TopIndex(1).(float64)
				stk.Pop()
				stk.Pc++
			}
		}
	}
	return common.ErrorInstruction
}

func And(rhsLen int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if !(stk.Top()).(bool) {
			stk.Pc += rhsLen + 1
			return
		}
		stk.Pop()
		stk.Pc++
	}
}

func Or(rhsLen int) common.Instruction {
	return func(m *common.Memory, stk *common.Stack) {
		if (stk.Top()).(bool) {
			stk.Pc += rhsLen + 1
			return
		}
		stk.Pop()
		stk.Pc++
	}
}

func BinaryOp(op string, lhs *common.DataType, rhs *common.DataType) common.Instruction {
	switch op {
	case "+":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.String:
			lhsConvertFunc, err := GetConvertFuncStr(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFuncStr(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case "-":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case "*":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case "/":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case "%":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		if t.Kind.Kind > common.Int64 {
			return common.ErrorInstruction
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))%rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))%rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))%rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))%rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case "**":
		t := common.BasicTypeMap["int32"]
		if lhs.Kind.Kind > t.Kind.Kind {
			t = lhs
		}
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		if t.Kind.Kind > common.Float64 {
			return common.ErrorInstruction
		}
		switch t.Kind.Kind {
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, IPow32(lhsConvertFunc(stk.TopIndex(1)), rhsConvertFunc(stk.TopIndex(0))))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, IPow64(lhsConvertFunc(stk.TopIndex(1)), rhsConvertFunc(stk.TopIndex(0))))
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, float32(math.Pow(float64(lhsConvertFunc(stk.TopIndex(1))), float64(rhsConvertFunc(stk.TopIndex(0))))))
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, math.Pow(lhsConvertFunc(stk.TopIndex(1)), rhsConvertFunc(stk.TopIndex(0))))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case "==":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Bool:
			lhsConvertFunc, err := GetConvertFuncBool(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFuncBool(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.String:
			lhsConvertFunc, err := GetConvertFuncStr(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFuncStr(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Nil:
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, stk.TopIndex(1) == nil && stk.TopIndex(0) == nil)
				stk.Pop()
				stk.Pc++
			}
		default:
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
				stk.Pop()
				stk.Pc++
			}
		}
	case "!=":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Bool:
			lhsConvertFunc, err := GetConvertFuncBool(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFuncBool(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.String:
			lhsConvertFunc, err := GetConvertFuncStr(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFuncStr(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Pc++
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Nil:
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, stk.TopIndex(1) != nil || stk.TopIndex(0) != nil)
				stk.Pop()
				stk.Pc++
			}
		default:
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
				stk.Pop()
				stk.Pc++
			}
		}
	case ">":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.String:
			lhsConvertFunc, err := GetConvertFuncStr(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFuncStr(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case "<":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.String:
			lhsConvertFunc, err := GetConvertFuncStr(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFuncStr(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case ">=":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.String:
			lhsConvertFunc, err := GetConvertFuncStr(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFuncStr(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case "<=":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc, err := GetConvertFunc32u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.UInt64:
			lhsConvertFunc, err := GetConvertFunc64u(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64u(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int32:
			lhsConvertFunc, err := GetConvertFunc32i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Int64:
			lhsConvertFunc, err := GetConvertFunc64i(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64i(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float32:
			lhsConvertFunc, err := GetConvertFunc32f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc32f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.Float64:
			lhsConvertFunc, err := GetConvertFunc64f(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFunc64f(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		case common.String:
			lhsConvertFunc, err := GetConvertFuncStr(lhs)
			if err != nil {
				return common.ErrorInstruction
			}
			rhsConvertFunc, err := GetConvertFuncStr(rhs)
			if err != nil {
				return common.ErrorInstruction
			}
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	default:
		// (=, ~=, &&, ||) are not process here
		return common.ErrorInstruction
	}
}

func IPow32(a, b int32) int32 {
	var result int32 = 1

	for 0 != b {
		if 0 != (b & 1) {
			result *= a

		}
		b >>= 1
		a *= a
	}

	return result
}

func IPow64(a, b int64) int64 {
	var result int64 = 1

	for 0 != b {
		if 0 != (b & 1) {
			result *= a

		}
		b >>= 1
		a *= a
	}

	return result
}

func LeftUnaryOp(op string, opType *common.DataType) common.Instruction {
	switch op {
	case "!":
		if opType.Kind.Kind != common.Bool {
			return common.ErrorInstruction
		}
		return func(m *common.Memory, stk *common.Stack) {
			stk.Set(0, !stk.Top().(bool))
			stk.Pc++
		}
	case "-":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -int32(stk.Top().(uint32)))
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -stk.Top().(int32))
				stk.Pc++
			}
		case common.Float32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -stk.Top().(float32))
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -stk.Top().(uint64))
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -stk.Top().(int64))
				stk.Pc++
			}
		case common.Float64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -stk.Top().(float64))
				stk.Pc++
			}
		}
	case "++":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*uint32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Set(0, *ptr)
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*int32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Set(0, *ptr)
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*uint64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Set(0, *ptr)
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*int64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Set(0, *ptr)
				stk.Pc++
			}
		}
	case "--":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*uint32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Set(0, *ptr)
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*int32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Set(0, *ptr)
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*uint64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Set(0, *ptr)
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*int64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Set(0, *ptr)
				stk.Pc++
			}
		}
	}
	return common.ErrorInstruction
}

func RightUnaryOp(op string, opType *common.DataType) common.Instruction {
	switch op {
	case "++":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*uint32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*int32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*uint64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*int64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Pc++
			}
		}
	case "--":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*uint32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*int32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*uint64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*int64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Pc++
			}
		}
	}
	return common.ErrorInstruction
}
