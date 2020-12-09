package lambda_chains

import (
	"github.com/lwpyr/goscript/common"
	"math"
	"unsafe"
)

func GetCalAssignOpFunc(op string, lhs *common.DataType, rhs *common.DataType) common.Instruction {
	switch lhs.Kind.Kind {
	case common.UInt32:
		rhsConvertFunc := GetConvertFunc32u(rhs, lhs)
		switch op {
		case "+=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*uint32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) += rhsConvertFunc(stk.TopIndex(1))
			}
		case "-=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*uint32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) -= rhsConvertFunc(stk.TopIndex(1))
			}
		case "*=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*uint32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) *= rhsConvertFunc(stk.TopIndex(1))
			}
		case "/=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*uint32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) /= rhsConvertFunc(stk.TopIndex(1))
			}
		}
	case common.UInt64:
		rhsConvertFunc := GetConvertFunc64u(rhs, lhs)
		switch op {
		case "+=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*uint64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) += rhsConvertFunc(stk.TopIndex(1))
			}
		case "-=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*uint64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) -= rhsConvertFunc(stk.TopIndex(1))
			}
		case "*=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*uint64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) *= rhsConvertFunc(stk.TopIndex(1))
			}
		case "/=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*uint64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) /= rhsConvertFunc(stk.TopIndex(1))
			}
		}
	case common.Int32:
		rhsConvertFunc := GetConvertFunc32i(rhs, lhs)
		switch op {
		case "+=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*int32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) += rhsConvertFunc(stk.TopIndex(1))
			}
		case "-=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*int32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) -= rhsConvertFunc(stk.TopIndex(1))
			}
		case "*=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*int32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) *= rhsConvertFunc(stk.TopIndex(1))
			}
		case "/=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*int32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) /= rhsConvertFunc(stk.TopIndex(1))
			}
		}
	case common.Int64:
		rhsConvertFunc := GetConvertFunc64i(rhs, lhs)
		switch op {
		case "+=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*int64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) += rhsConvertFunc(stk.TopIndex(1))
			}
		case "-=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*int64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) -= rhsConvertFunc(stk.TopIndex(1))
			}
		case "*=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*int64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) *= rhsConvertFunc(stk.TopIndex(1))
			}
		case "/=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*int64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) /= rhsConvertFunc(stk.TopIndex(1))
			}
		}
	case common.Float32:
		rhsConvertFunc := GetConvertFunc32f(rhs, lhs)
		switch op {
		case "+=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*float32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) += rhsConvertFunc(stk.TopIndex(1))
			}
		case "-=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*float32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) -= rhsConvertFunc(stk.TopIndex(1))
			}
		case "*=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*float32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) *= rhsConvertFunc(stk.TopIndex(1))
			}
		case "/=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*float32)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) /= rhsConvertFunc(stk.TopIndex(1))
			}
		}
	case common.Float64:
		rhsConvertFunc := GetConvertFunc64f(rhs, lhs)
		switch op {
		case "+=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*float64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) += rhsConvertFunc(stk.TopIndex(1))
			}
		case "-=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*float64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) -= rhsConvertFunc(stk.TopIndex(1))
			}
		case "*=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*float64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) *= rhsConvertFunc(stk.TopIndex(1))
			}
		case "/=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*float64)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) /= rhsConvertFunc(stk.TopIndex(1))
			}
		}
	case common.String:
		rhsConvertFunc := GetConvertFuncStr(rhs, lhs)
		switch op {
		case "+=":
			return func(_ *common.Memory, stk *common.Stack) {
				*(*string)((*EmptyFace)(unsafe.Pointer(stk.TopIndex(0).(*interface{}))).Data) += rhsConvertFunc(stk.TopIndex(1))
			}
		}
	}
	panic(common.NewTypeErr("not support " + op + " for Lhs " + lhs.Type + " and Rhs " + rhs.Type))
}

func GetBinaryOpFunc(op string, lhs *common.DataType, rhs *common.DataType) common.Instruction {
	switch op {
	case "+":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.String:
			lhsConvertFunc := GetConvertFuncStr(lhs, t)
			rhsConvertFunc := GetConvertFuncStr(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))+rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		default:
			panic("not support + : " + lhs.Type + " " + rhs.Type)
		}
	case "-":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))-rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		default:
			panic("not support - : " + lhs.Type + " " + rhs.Type)
		}
	case "*":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))*rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		default:
			panic("not support * : " + lhs.Type + " " + rhs.Type)
		}
	case "/":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))/rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		default:
			panic("not support / : " + lhs.Type + " " + rhs.Type)
		}
	case "%":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		if t.Kind.Kind > common.Int64 {
			panic("not support type for integer mod: " + lhs.Type + " " + rhs.Type)
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))%rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))%rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))%rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1))%rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		default:
			panic("not support % : " + lhs.Type + " " + rhs.Type)
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
			panic("not support type for division: " + lhs.Type + " " + rhs.Type)
		}
		switch t.Kind.Kind {
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, IPow32(lhsConvertFunc(stk.TopIndex(1)), rhsConvertFunc(stk.TopIndex(0))))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, IPow64(lhsConvertFunc(stk.TopIndex(1)), rhsConvertFunc(stk.TopIndex(0))))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, float32(math.Pow(float64(lhsConvertFunc(stk.TopIndex(1))), float64(rhsConvertFunc(stk.TopIndex(0))))))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, math.Pow(lhsConvertFunc(stk.TopIndex(1)), rhsConvertFunc(stk.TopIndex(0))))
				stk.Pop()
			}
		default:
			panic("not support ** : " + lhs.Type + " " + rhs.Type)
		}
	case "==":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Bool:
			lhsConvertFunc := GetConvertFuncBool(lhs, t)
			rhsConvertFunc := GetConvertFuncBool(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) == rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.String:
			lhsConvertFunc := GetConvertFuncStr(lhs, t)
			rhsConvertFunc := GetConvertFuncStr(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
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
			}
		default:
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, stk.TopIndex(1) == stk.TopIndex(0))
				stk.Pop()
			}
		}
	case "!=":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Bool:
			lhsConvertFunc := GetConvertFuncBool(lhs, t)
			rhsConvertFunc := GetConvertFuncBool(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				if stk.TopIndex(1) == nil || stk.TopIndex(0) == nil {
					stk.Set(1, stk.TopIndex(1) != stk.TopIndex(0))
					stk.Pop()
					return
				}
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) != rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.String:
			lhsConvertFunc := GetConvertFuncStr(lhs, t)
			rhsConvertFunc := GetConvertFuncStr(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
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
			}
		default:
			panic("not support != : " + lhs.Type + " " + rhs.Type)
		}
	case ">":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.String:
			lhsConvertFunc := GetConvertFuncStr(lhs, t)
			rhsConvertFunc := GetConvertFuncStr(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) > rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		default:
			panic("not support > : " + lhs.Type + " " + rhs.Type)
		}
	case "<":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.String:
			lhsConvertFunc := GetConvertFuncStr(lhs, t)
			rhsConvertFunc := GetConvertFuncStr(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) < rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		default:
			panic("not support < : " + lhs.Type + " " + rhs.Type)
		}
	case ">=":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.String:
			lhsConvertFunc := GetConvertFuncStr(lhs, t)
			rhsConvertFunc := GetConvertFuncStr(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) >= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		default:
			panic("not support >= : " + lhs.Type + " " + rhs.Type)
		}
	case "<=":
		t := lhs
		if rhs.Kind.Kind > t.Kind.Kind {
			t = rhs
		}
		switch t.Kind.Kind {
		case common.UInt32:
			lhsConvertFunc := GetConvertFunc32u(lhs, t)
			rhsConvertFunc := GetConvertFunc32u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.UInt64:
			lhsConvertFunc := GetConvertFunc64u(lhs, t)
			rhsConvertFunc := GetConvertFunc64u(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int32:
			lhsConvertFunc := GetConvertFunc32i(lhs, t)
			rhsConvertFunc := GetConvertFunc32i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Int64:
			lhsConvertFunc := GetConvertFunc64i(lhs, t)
			rhsConvertFunc := GetConvertFunc64i(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float32:
			lhsConvertFunc := GetConvertFunc32f(lhs, t)
			rhsConvertFunc := GetConvertFunc32f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.Float64:
			lhsConvertFunc := GetConvertFunc64f(lhs, t)
			rhsConvertFunc := GetConvertFunc64f(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		case common.String:
			lhsConvertFunc := GetConvertFuncStr(lhs, t)
			rhsConvertFunc := GetConvertFuncStr(rhs, t)
			return func(_ *common.Memory, stk *common.Stack) {
				stk.Set(1, lhsConvertFunc(stk.TopIndex(1)) <= rhsConvertFunc(stk.TopIndex(0)))
				stk.Pop()
			}
		default:
			panic("not support <= : " + lhs.Type + " " + rhs.Type)
		}
	default:
		// (=, ~=, &&, ||) are not process here
		panic("unknown op " + op)
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

func GetLeftUnaryOpFunc(op string, opType *common.DataType) common.Instruction {
	switch op {
	case "!":
		return func(m *common.Memory, stk *common.Stack) {
			stk.Set(0, !stk.Top().(bool))
		}
	case "-":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -int32(stk.Top().(uint32)))
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -stk.Top().(int32))
			}
		case common.Float32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -stk.Top().(float32))
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -int64(stk.Top().(uint64)))
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -stk.Top().(int64))
			}
		case common.Float64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Set(0, -stk.Top().(float64))
			}
		}
	case "++":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*uint32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Set(0, *ptr)
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*int32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Set(0, *ptr)
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*uint64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Set(0, *ptr)
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*int64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
				stk.Set(0, *ptr)
			}
		}
	case "--":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*uint32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Set(0, *ptr)
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*int32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Set(0, *ptr)
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*uint64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Set(0, *ptr)
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				*ptr = (*(*int64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
				stk.Set(0, *ptr)
			}
		}
	}
	panic("wtf")
}

func GetRightUnaryOpFunc(op string, opType *common.DataType) common.Instruction {
	switch op {
	case "++":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*uint32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*int32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*uint64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*int64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) + 1
			}
		}
	case "--":
		switch opType.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*uint32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*int32)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*uint64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				ptr := stk.Top().(*interface{})
				stk.Set(0, *ptr)
				*ptr = (*(*int64)((*EmptyFace)(unsafe.Pointer(ptr)).Data)) - 1
			}
		}
	}
	panic("wtf")
}
