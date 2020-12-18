package instruction

import (
	"fmt"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/hack"
	"unsafe"
)

func GetConvertInstruction(from *common.DataType, to *common.DataType) common.Instruction {
	if to.Kind.Kind == common.Any {
		return nil
	}
	switch from.Kind.Kind {
	case common.Any:
		return nil
	case common.Nil:
		switch to.Kind.Kind {
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilUInt32
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilUInt64
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilInt32
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilInt64
				stk.Pc++
			}
		case common.Float32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilFloat32
				stk.Pc++
			}
		case common.Float64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilFloat64
				stk.Pc++
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilBool
				stk.Pc++
			}
		case common.String:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilString
				stk.Pc++
			}
		case common.Slice:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilSlice
				stk.Pc++
			}
		case common.Bytes:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = NilBytes
				stk.Pc++
			}
		case common.Any:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.Message:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = make(map[string]interface{})
				stk.Pc++
			}
		case common.Map:
			switch to.KeyType.Kind.Kind {
			case common.UInt32:
				return func(m *common.Memory, stk *common.Stack) {
					stk.Data[stk.Sp] = make(map[uint32]interface{})
					stk.Pc++
				}
			case common.UInt64:
				return func(m *common.Memory, stk *common.Stack) {
					stk.Data[stk.Sp] = make(map[uint64]interface{})
					stk.Pc++
				}
			case common.Int32:
				return func(m *common.Memory, stk *common.Stack) {
					stk.Data[stk.Sp] = make(map[int32]interface{})
					stk.Pc++
				}
			case common.Int64:
				return func(m *common.Memory, stk *common.Stack) {
					stk.Data[stk.Sp] = make(map[int64]interface{})
					stk.Pc++
				}
			case common.Float32:
				return func(m *common.Memory, stk *common.Stack) {
					stk.Data[stk.Sp] = make(map[float32]interface{})
					stk.Pc++
				}
			case common.Float64:
				return func(m *common.Memory, stk *common.Stack) {
					stk.Data[stk.Sp] = make(map[float32]interface{})
					stk.Pc++
				}
			case common.Bool:
				return func(m *common.Memory, stk *common.Stack) {
					stk.Data[stk.Sp] = make(map[bool]interface{})
					stk.Pc++
				}
			case common.String:
				return func(m *common.Memory, stk *common.Stack) {
					stk.Data[stk.Sp] = make(map[string]interface{})
					stk.Pc++
				}
			default:
				return common.ErrorInstruction
			}
		case common.Nil:
			return nil
		default:
			return common.ErrorInstruction
		}
	case common.UInt8:
		switch to.Kind.Kind {
		case common.Nil:
			return nil
		case common.UInt8:
			return nil
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = *(*uint8)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data) != 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.UInt32:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.UInt32:
			return nil
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = *(*uint32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data) != 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.UInt64:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.UInt64:
			return nil
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = *(*uint64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data) != 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Int32:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint32(*(*int32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint64(*(*int32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int32:
			return nil
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int64(*(*int32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float32(*(*int32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float64(*(*int32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = *(*int32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data) != 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Int64:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint32(*(*int64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint64(*(*int64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int32(*(*int64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int64:
			return nil
		case common.Float32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float32(*(*int64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float64(*(*int64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = *(*int64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data) != 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Float32:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint32(*(*float32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint64(*(*float32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int32(*(*float32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int64(*(*float32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float32:
			return nil
		case common.Float64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float64(*(*float32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = *(*float32)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data) != 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Float64:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.UInt32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint32(*(*float64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.UInt64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = uint64(*(*float64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int32(*(*float64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Int64:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = int64(*(*float64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float32:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = float32(*(*float64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Float64:
			return nil
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = *(*float64)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data) != 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.String:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.String:
			return nil
		case common.Bytes:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = []byte(*(*string)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = len(*(*string)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data)) != 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Bytes:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.String:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = string(*(*[]byte)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data))
				stk.Pc++
			}
		case common.Bytes:
			return nil
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = len(*(*[]byte)((*EmptyFace)(unsafe.Pointer(&stk.Data[stk.Sp])).Data)) != 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Bool:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.Bool:
			return nil
		default:
			return common.ErrorInstruction
		}
	case common.Slice:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.Slice:
			return nil
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = stk.Data[stk.Sp] != nil && hack.SliceLen(stk.Data[stk.Sp]) == 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Map:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.Map:
			if to.Type == from.Type {
				return nil
			} else {
				return common.ErrorInstruction
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = stk.Data[stk.Sp] != nil && hack.MLen(stk.Data[stk.Sp]) == 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Message:
		switch to.Kind.Kind {
		case common.Nil:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = nil
				stk.Pc++
			}
		case common.Message:
			if to.Type == from.Type {
				return nil
			} else {
				return common.ErrorInstruction
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = stk.Data[stk.Sp] != nil && hack.MLen(stk.Data[stk.Sp]) == 0
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.ReflectType:
		switch to.Kind.Kind {
		case common.ReflectType:
			return nil
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = true
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Func:
		switch to.Kind.Kind {
		case common.Func:
			if to.Type == from.Type {
				return nil
			} else {
				return common.ErrorInstruction
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = stk.Data[stk.Sp] != nil
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	case common.Channel:
		switch to.Kind.Kind {
		case common.Channel:
			if to.Type == from.Type {
				return nil
			} else {
				return common.ErrorInstruction
			}
		case common.Bool:
			return func(m *common.Memory, stk *common.Stack) {
				stk.Data[stk.Sp] = stk.Data[stk.Sp] != nil
				stk.Pc++
			}
		default:
			return common.ErrorInstruction
		}
	default:
		return common.ErrorInstruction
	}
}

type EmptyFace struct {
	_type unsafe.Pointer
	Data  unsafe.Pointer
}

type TypeConvertFunc32u func(from interface{}) uint32
type TypeConvertFunc64u func(from interface{}) uint64
type TypeConvertFunc32i func(from interface{}) int32
type TypeConvertFunc64i func(from interface{}) int64
type TypeConvertFunc32f func(from interface{}) float32
type TypeConvertFunc64f func(from interface{}) float64
type TypeConvertFuncStr func(from interface{}) string
type TypeConvertFuncBytes func(from interface{}) []byte
type TypeConvertFuncBool func(from interface{}) bool
type TypeConvertFuncMap func(from interface{}) interface{}
type TypeConvertFuncMessage func(from interface{}) interface{}
type TypeConvertFuncSlice func(from interface{}) []interface{}
type TypeConvertFuncObject func(from interface{}) interface{}

func GetConvertFunc32u(from *common.DataType) (TypeConvertFunc32u, error) {
	switch from.Kind.Kind {
	case common.Any:
		return func(from interface{}) uint32 {
			return *(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Nil:
		return func(from interface{}) uint32 {
			return 0
		}, nil
	case common.UInt8:
		return func(from interface{}) uint32 {
			return uint32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt32:
		return func(from interface{}) uint32 {
			return uint32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt64:
		return func(from interface{}) uint32 {
			return uint32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int32:
		return func(from interface{}) uint32 {
			return uint32(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int64:
		return func(from interface{}) uint32 {
			return uint32(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float32:
		return func(from interface{}) uint32 {
			return uint32(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float64:
		return func(from interface{}) uint32 {
			return uint32(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.String:
		return nil, fmt.Errorf("string to uint32 value cannot be auto-converted")
	case common.Bytes:
		return nil, fmt.Errorf("bytes to uint32 value cannot be auto-converted")
	case common.Bool:
		return nil, fmt.Errorf("bool to uint32 value cannot be auto-converted")
	case common.Slice:
		return nil, fmt.Errorf("slice to uint32 value cannot be auto-converted")
	case common.Map:
		return nil, fmt.Errorf("map to uint32 value cannot be auto-converted")
	case common.Message:
		return nil, fmt.Errorf("message to uint32 value cannot be auto-converted")
	case common.ReflectType:
		return nil, fmt.Errorf("reflect to uint32 value cannot be auto-converted")
	default:
		return nil, fmt.Errorf("unknown type convert")
	}
}

func GetConvertFunc32i(from *common.DataType) (TypeConvertFunc32i, error) {
	switch from.Kind.Kind {
	case common.Any:
		return func(from interface{}) int32 {
			return *(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Nil:
		return func(from interface{}) int32 {
			return 0
		}, nil
	case common.UInt8:
		return func(from interface{}) int32 {
			return int32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt32:
		return func(from interface{}) int32 {
			return int32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt64:
		return func(from interface{}) int32 {
			return int32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int32:
		return func(from interface{}) int32 {
			return int32(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int64:
		return func(from interface{}) int32 {
			return int32(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float32:
		return func(from interface{}) int32 {
			return int32(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float64:
		return func(from interface{}) int32 {
			return int32(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.String:
		return nil, fmt.Errorf("string to uint32 value cannot be auto-converted")
	case common.Bytes:
		return nil, fmt.Errorf("bytes to uint32 value cannot be auto-converted")
	case common.Bool:
		return nil, fmt.Errorf("bool to uint32 value cannot be auto-converted")
	case common.Slice:
		return nil, fmt.Errorf("slice to uint32 value cannot be auto-converted")
	case common.Map:
		return nil, fmt.Errorf("map to uint32 value cannot be auto-converted")
	case common.Message:
		return nil, fmt.Errorf("message to uint32 value cannot be auto-converted")
	case common.ReflectType:
		return nil, fmt.Errorf("reflect to uint32 value cannot be auto-converted")
	default:
		return nil, fmt.Errorf("unknown type convert")
	}
}

func GetConvertFunc32f(from *common.DataType) (TypeConvertFunc32f, error) {
	switch from.Kind.Kind {
	case common.Any:
		return func(from interface{}) float32 {
			return *(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Nil:
		return func(from interface{}) float32 {
			return 0
		}, nil
	case common.UInt8:
		return func(from interface{}) float32 {
			return float32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt32:
		return func(from interface{}) float32 {
			return float32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt64:
		return func(from interface{}) float32 {
			return float32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int32:
		return func(from interface{}) float32 {
			return float32(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int64:
		return func(from interface{}) float32 {
			return float32(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float32:
		return func(from interface{}) float32 {
			return float32(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float64:
		return func(from interface{}) float32 {
			return float32(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.String:
		return nil, fmt.Errorf("string to uint32 value cannot be auto-converted")
	case common.Bytes:
		return nil, fmt.Errorf("bytes to uint32 value cannot be auto-converted")
	case common.Bool:
		return nil, fmt.Errorf("bool to uint32 value cannot be auto-converted")
	case common.Slice:
		return nil, fmt.Errorf("slice to uint32 value cannot be auto-converted")
	case common.Map:
		return nil, fmt.Errorf("map to uint32 value cannot be auto-converted")
	case common.Message:
		return nil, fmt.Errorf("message to uint32 value cannot be auto-converted")
	case common.ReflectType:
		return nil, fmt.Errorf("reflect to uint32 value cannot be auto-converted")
	default:
		return nil, fmt.Errorf("unknown type convert")
	}
}

func GetConvertFunc64u(from *common.DataType) (TypeConvertFunc64u, error) {
	switch from.Kind.Kind {
	case common.Any:
		return func(from interface{}) uint64 {
			return *(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Nil:
		return func(from interface{}) uint64 {
			return 0
		}, nil
	case common.UInt8:
		return func(from interface{}) uint64 {
			return uint64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt32:
		return func(from interface{}) uint64 {
			return uint64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt64:
		return func(from interface{}) uint64 {
			return uint64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int32:
		return func(from interface{}) uint64 {
			return uint64(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int64:
		return func(from interface{}) uint64 {
			return uint64(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float32:
		return func(from interface{}) uint64 {
			return uint64(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float64:
		return func(from interface{}) uint64 {
			return uint64(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.String:
		return nil, fmt.Errorf("string to 64bit value cannot be auto-converted")
	case common.Bytes:
		return nil, fmt.Errorf("bytes to 64bit value cannot be auto-converted")
	case common.Bool:
		return nil, fmt.Errorf("bool to 64bit value cannot be auto-converted")
	case common.Slice:
		return nil, fmt.Errorf("slice to 64bit value cannot be auto-converted")
	case common.Map:
		return nil, fmt.Errorf("map to 64bit value cannot be auto-converted")
	case common.Message:
		return nil, fmt.Errorf("message to 64bit value cannot be auto-converted")
	case common.ReflectType:
		return nil, fmt.Errorf("reflect to 32bit value cannot be auto-converted")
	default:
		return nil, fmt.Errorf("unknown type convert")
	}
}

func GetConvertFunc64i(from *common.DataType) (TypeConvertFunc64i, error) {
	switch from.Kind.Kind {
	case common.Any:
		return func(from interface{}) int64 {
			return *(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Nil:
		return func(from interface{}) int64 {
			return 0
		}, nil
	case common.UInt8:
		return func(from interface{}) int64 {
			return int64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt32:
		return func(from interface{}) int64 {
			return int64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt64:
		return func(from interface{}) int64 {
			return int64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int32:
		return func(from interface{}) int64 {
			return int64(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int64:
		return func(from interface{}) int64 {
			return int64(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float32:
		return func(from interface{}) int64 {
			return int64(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float64:
		return func(from interface{}) int64 {
			return int64(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.String:
		return nil, fmt.Errorf("string to 64bit value cannot be auto-converted")
	case common.Bytes:
		return nil, fmt.Errorf("bytes to 64bit value cannot be auto-converted")
	case common.Bool:
		return nil, fmt.Errorf("bool to 64bit value cannot be auto-converted")
	case common.Slice:
		return nil, fmt.Errorf("slice to 64bit value cannot be auto-converted")
	case common.Map:
		return nil, fmt.Errorf("map to 64bit value cannot be auto-converted")
	case common.Message:
		return nil, fmt.Errorf("message to 64bit value cannot be auto-converted")
	case common.ReflectType:
		return nil, fmt.Errorf("reflect to 32bit value cannot be auto-converted")
	default:
		return nil, fmt.Errorf("unknown type convert")
	}
}

func GetConvertFunc64f(from *common.DataType) (TypeConvertFunc64f, error) {
	switch from.Kind.Kind {
	case common.Any:
		return func(from interface{}) float64 {
			return *(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Nil:
		return func(from interface{}) float64 {
			return 0
		}, nil
	case common.UInt8:
		return func(from interface{}) float64 {
			return float64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt32:
		return func(from interface{}) float64 {
			return float64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt64:
		return func(from interface{}) float64 {
			return float64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int32:
		return func(from interface{}) float64 {
			return float64(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Int64:
		return func(from interface{}) float64 {
			return float64(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float32:
		return func(from interface{}) float64 {
			return float64(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Float64:
		return func(from interface{}) float64 {
			return float64(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.String:
		return nil, fmt.Errorf("string to 64bit value cannot be auto-converted")
	case common.Bytes:
		return nil, fmt.Errorf("bytes to 64bit value cannot be auto-converted")
	case common.Bool:
		return nil, fmt.Errorf("bool to 64bit value cannot be auto-converted")
	case common.Slice:
		return nil, fmt.Errorf("slice to 64bit value cannot be auto-converted")
	case common.Map:
		return nil, fmt.Errorf("map to 64bit value cannot be auto-converted")
	case common.Message:
		return nil, fmt.Errorf("message to 64bit value cannot be auto-converted")
	case common.ReflectType:
		return nil, fmt.Errorf("reflect to 32bit value cannot be auto-converted")
	default:
		return nil, fmt.Errorf("unknown type convert")
	}
}

func GetConvertFuncBool(from *common.DataType) (TypeConvertFuncBool, error) {
	switch from.Kind.Kind {
	case common.Any:
		return func(from interface{}) bool {
			return *(*bool)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Nil:
		return func(from interface{}) bool {
			return false
		}, nil
	case common.UInt8:
		return func(from interface{}) bool {
			return *(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}, nil
	case common.UInt32:
		return func(from interface{}) bool {
			return *(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}, nil
	case common.UInt64:
		return func(from interface{}) bool {
			return *(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}, nil
	case common.Int32:
		return func(from interface{}) bool {
			return *(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}, nil
	case common.Int64:
		return func(from interface{}) bool {
			return *(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}, nil
	case common.Float32:
		return func(from interface{}) bool {
			return *(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}, nil
	case common.Float64:
		return func(from interface{}) bool {
			return *(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}, nil
	case common.String:
		return nil, fmt.Errorf("string to bool value cannot be auto-converted")
	case common.Bytes:
		return nil, fmt.Errorf("bytes to bool value cannot be auto-converted")
	case common.Bool:
		return func(from interface{}) bool {
			return *(*bool)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Slice:
		return nil, fmt.Errorf("slice to bool value cannot be auto-converted")
	case common.Map:
		return nil, fmt.Errorf("map to bool value cannot be auto-converted")
	case common.Message:
		return nil, fmt.Errorf("message to bool value cannot be auto-converted")
	case common.ReflectType:
		return nil, fmt.Errorf("reflect to bool value cannot be auto-converted")
	default:
		return nil, fmt.Errorf("unknown type convert")
	}
}

func GetConvertFuncStr(from *common.DataType) (TypeConvertFuncStr, error) {
	switch from.Kind.Kind {
	case common.Any:
		return func(from interface{}) string {
			return *(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Nil:
		return func(from interface{}) string {
			return ""
		}, nil
	case common.UInt8:
		return func(from interface{}) string {
			return string(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.UInt32:
		return nil, fmt.Errorf("uint32 to string value cannot be auto-converted")
	case common.UInt64:
		return nil, fmt.Errorf("uint64 to string value cannot be auto-converted")
	case common.Int32:
		return nil, fmt.Errorf("int32 to string value cannot be auto-converted")
	case common.Int64:
		return nil, fmt.Errorf("int64 to string value cannot be auto-converted")
	case common.Float32:
		return nil, fmt.Errorf("float32 to string value cannot be auto-converted")
	case common.Float64:
		return nil, fmt.Errorf("float64 to string value cannot be auto-converted")
	case common.String:
		return func(from interface{}) string {
			return *(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Bytes:
		return func(from interface{}) string {
			return string(*(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Bool:
		return nil, fmt.Errorf("bool to string value cannot be auto-converted")
	case common.Slice:
		return nil, fmt.Errorf("slice to string value cannot be auto-converted")
	case common.Map:
		return nil, fmt.Errorf("map to string value cannot be auto-converted")
	case common.Message:
		return nil, fmt.Errorf("message to string value cannot be auto-converted")
	case common.ReflectType:
		return nil, fmt.Errorf("reflect to string value cannot be auto-converted")
	default:
		return nil, fmt.Errorf("unknown type convert")
	}
}

func GetConvertFuncBytes(from *common.DataType) (TypeConvertFuncBytes, error) {
	switch from.Kind.Kind {
	case common.Any:
		return func(from interface{}) []byte {
			return *(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Nil:
		return func(from interface{}) []byte {
			return []byte{}
		}, nil
	case common.UInt8:
		return nil, fmt.Errorf("uint8 to bytes value cannot be auto-converted")
	case common.UInt32:
		return nil, fmt.Errorf("uint32 to bytes value cannot be auto-converted")
	case common.UInt64:
		return nil, fmt.Errorf("uint64 to bytes value cannot be auto-converted")
	case common.Int32:
		return nil, fmt.Errorf("int32 to bytes value cannot be auto-converted")
	case common.Int64:
		return nil, fmt.Errorf("int64 to bytes value cannot be auto-converted")
	case common.Float32:
		return nil, fmt.Errorf("float32 to bytes value cannot be auto-converted")
	case common.Float64:
		return nil, fmt.Errorf("float64 to bytes value cannot be auto-converted")
	case common.String:
		return func(from interface{}) []byte {
			return []byte(*(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}, nil
	case common.Bytes:
		return func(from interface{}) []byte {
			return *(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}, nil
	case common.Bool:
		return nil, fmt.Errorf("bool to string value cannot be auto-converted")
	case common.Slice:
		return nil, fmt.Errorf("slice to string value cannot be auto-converted")
	case common.Map:
		return nil, fmt.Errorf("map to string value cannot be auto-converted")
	case common.Message:
		return nil, fmt.Errorf("message to string value cannot be auto-converted")
	case common.ReflectType:
		return nil, fmt.Errorf("reflect to string value cannot be auto-converted")
	default:
		return nil, fmt.Errorf("unknown type convert")
	}
}
