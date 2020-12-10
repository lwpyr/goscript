package lambda_chains

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/hack"
	"unsafe"
)

type EmptyFace struct {
	_type unsafe.Pointer
	Data  unsafe.Pointer
}

type TypeConvertFunc func(from interface{}) (to interface{})

func GetConvertFunc(from *common.DataType, to *common.DataType) TypeConvertFunc {
	if to.Kind.Kind == common.Object {
		return func(from interface{}) (to interface{}) {
			return from
		}
	}
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) (to interface{}) {
			return from
		}
	case common.Nil:
		switch to.Kind.Kind {
		case common.UInt32:
			return func(from interface{}) (to interface{}) {
				return NilUInt32
			}
		case common.UInt64:
			return func(from interface{}) (to interface{}) {
				return NilUInt64
			}
		case common.Int32:
			return func(from interface{}) (to interface{}) {
				return NilInt32
			}
		case common.Int64:
			return func(from interface{}) (to interface{}) {
				return NilInt64
			}
		case common.Float32:
			return func(from interface{}) (to interface{}) {
				return NilFloat32
			}
		case common.Float64:
			return func(from interface{}) (to interface{}) {
				return NilFloat64
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				return NilBool
			}
		case common.String:
			return func(from interface{}) (to interface{}) {
				return NilString
			}
		case common.Slice:
			return func(from interface{}) (to interface{}) {
				return NilSlice
			}
		case common.Bytes:
			return func(from interface{}) (to interface{}) {
				return NilBytes
			}
		case common.Object:
			return func(from interface{}) (to interface{}) {
				return NilArbitrary
			}
		case common.Message:
			return func(from interface{}) (to interface{}) {
				return make(map[string]interface{})
			}
		case common.Map:
			switch to.KeyType.Kind.Kind {
			case common.UInt32:
				return func(from interface{}) (to interface{}) {
					return make(map[uint32]interface{})
				}
			case common.UInt64:
				return func(from interface{}) (to interface{}) {
					return make(map[uint64]interface{})
				}
			case common.Int32:
				return func(from interface{}) (to interface{}) {
					return make(map[int32]interface{})
				}
			case common.Int64:
				return func(from interface{}) (to interface{}) {
					return make(map[int64]interface{})
				}
			case common.Float32:
				return func(from interface{}) (to interface{}) {
					return make(map[float32]interface{})
				}
			case common.Float64:
				return func(from interface{}) (to interface{}) {
					return make(map[float64]interface{})
				}
			case common.Bool:
				return func(from interface{}) (to interface{}) {
					return make(map[bool]interface{})
				}
			case common.String:
				return func(from interface{}) (to interface{}) {
					return make(map[string]interface{})
				}
			default:
				panic("unsupported map key type")
			}
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		default:
			panic("unknown type converter")
		}
	case common.UInt8:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.UInt8:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt32
				}
				return from
			}
		case common.UInt32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt32
				}
				return uint64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.UInt64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt64
				}
				return uint64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt32
				}
				return int32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt64
				}
				return int64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat32
				}
				return float32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat64
				}
				return float64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBool
				}
				return *(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
			}
		default:
			panic("unknown type converter")
		}
	case common.UInt32:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.UInt32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt32
				}
				return from
			}
		case common.UInt64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt64
				}
				var i interface{} = uint64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
				return i
			}
		case common.Int32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt32
				}
				return int32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt64
				}
				return int64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat32
				}
				return float32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat64
				}
				return float64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBool
				}
				return *(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
			}
		default:
			panic("unknown type converter")
		}
	case common.UInt64:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.UInt32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt32
				}
				return uint32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.UInt64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt64
				}
				return from
			}
		case common.Int32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt32
				}
				return int32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt64
				}
				return int64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat32
				}
				return float32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat64
				}
				return float64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBool
				}
				return *(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
			}
		default:
			panic("unknown type converter")
		}
	case common.Int32:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.UInt32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt32
				}
				return uint32(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.UInt64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt64
				}
				return uint64(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt32
				}
				return from
			}
		case common.Int64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt64
				}
				return int64(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat32
				}
				return float32(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat64
				}
				return float64(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBool
				}
				return *(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
			}
		default:
			panic("unknown type converter")
		}
	case common.Int64:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.UInt32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt32
				}
				return uint32(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.UInt64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt64
				}
				return uint64(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt32
				}
				return int32(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt64
				}
				return from
			}
		case common.Float32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat32
				}
				return float32(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat64
				}
				return float64(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBool
				}
				return *(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
			}
		default:
			panic("unknown type converter")
		}
	case common.Float32:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.UInt32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt32
				}
				return uint32(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.UInt64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt64
				}
				return uint64(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt32
				}
				return int32(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt64
				}
				return int64(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat32
				}
				return from
			}
		case common.Float64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat64
				}
				return float64(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBool
				}
				return *(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
			}
		default:
			panic("unknown type converter")
		}
	case common.Float64:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.UInt32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt32
				}
				return uint32(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.UInt64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilUInt64
				}
				return uint64(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt32
				}
				return int32(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Int64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilInt64
				}
				return int64(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float32:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat32
				}
				return float32(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Float64:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilFloat64
				}
				return from
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBool
				}
				return *(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
			}
		default:
			panic("unknown type converter")
		}
	case common.String:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.String:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilString
				}
				return from
			}
		case common.Bytes:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBytes
				}
				return []byte(*(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBool
				}
				return len(*(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data)) != 0
			}
		default:
			panic("unknown type convert")
		}
	case common.Bytes:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.String:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilString
				}
				return string(*(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data))
			}
		case common.Bytes:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBytes
				}
				return from
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilBool
				}
				return len(*(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data)) != 0
			}
		default:
			panic("unknown type converter")
		}
	case common.Bool:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				return from
			}
		default:
			panic("unknown type converter")
		}
	case common.Slice:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.Slice:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return NilSlice
				}
				return from
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				return from != nil && hack.SliceLen(from) == 0
			}
		default:
			panic("unknown type converter")
		}
	case common.Map:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.Map:
			switch to.KeyType.Kind.Kind {
			case common.UInt32:
				return func(from interface{}) (to interface{}) {
					if from == nil {
						return make(map[uint32]interface{})
					}
					return from
				}
			case common.UInt64:
				return func(from interface{}) (to interface{}) {
					if from == nil {
						return make(map[uint64]interface{})
					}
					return from
				}
			case common.Int32:
				return func(from interface{}) (to interface{}) {
					if from == nil {
						return make(map[int32]interface{})
					}
					return from
				}
			case common.Int64:
				return func(from interface{}) (to interface{}) {
					if from == nil {
						return make(map[int64]interface{})
					}
					return from
				}
			case common.Float32:
				return func(from interface{}) (to interface{}) {
					if from == nil {
						return make(map[float32]interface{})
					}
					return from
				}
			case common.Float64:
				return func(from interface{}) (to interface{}) {
					if from == nil {
						return make(map[float64]interface{})
					}
					return from
				}
			case common.Bool:
				return func(from interface{}) (to interface{}) {
					if from == nil {
						return make(map[bool]interface{})
					}
					return from
				}
			case common.String:
				return func(from interface{}) (to interface{}) {
					if from == nil {
						return make(map[string]interface{})
					}
					return from
				}
			default:
				panic("unsupported map key type")
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				return from != nil && hack.MLen(from) == 0
			}
		default:
			panic("unknown type converter")
		}
	case common.Message:
		switch to.Kind.Kind {
		case common.Nil:
			return func(from interface{}) (to interface{}) {
				return nil
			}
		case common.Message:
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return make(map[string]interface{})
				}
				return from
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				return from != nil && hack.MLen(from) == 0
			}
		default:
			panic("unknown type converter")
		}
	case common.ReflectType:
		switch to.Kind.Kind {
		case common.ReflectType:
			return func(from interface{}) (to interface{}) {
				return from
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				return true
			}
		default:
			panic("unknown type converter")
		}
	case common.Func:
		switch to.Kind.Kind {
		case common.Func:
			return func(from interface{}) (to interface{}) {
				return from
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				return true
			}
		default:
			panic("unknown type converter")
		}
	case common.Channel:
		switch to.Kind.Kind {
		case common.Channel:
			return func(from interface{}) (to interface{}) {
				return from
			}
		case common.Bool:
			return func(from interface{}) (to interface{}) {
				return true
			}
		default:
			panic("unknown type converter")
		}
	default:
		panic("unknown type converter")
	}
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

func GetConvertFunc32u(from *common.DataType, to *common.DataType) TypeConvertFunc32u {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) uint32 {
			return *(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) uint32 {
			return 0
		}
	case common.UInt8:
		return func(from interface{}) uint32 {
			return uint32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt32:
		return func(from interface{}) uint32 {
			return uint32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt64:
		return func(from interface{}) uint32 {
			return uint32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int32:
		return func(from interface{}) uint32 {
			return uint32(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int64:
		return func(from interface{}) uint32 {
			return uint32(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float32:
		return func(from interface{}) uint32 {
			return uint32(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float64:
		return func(from interface{}) uint32 {
			return uint32(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.String:
		panic("string to uint32 value cannot be auto-converted")
	case common.Bytes:
		panic("bytes to uint32 value cannot be auto-converted")
	case common.Bool:
		panic("bool to uint32 value cannot be auto-converted")
	case common.Slice:
		panic("slice to uint32 value cannot be auto-converted")
	case common.Map:
		panic("map to uint32 value cannot be auto-converted")
	case common.Message:
		panic("message to uint32 value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to uint32 value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFunc32i(from *common.DataType, to *common.DataType) TypeConvertFunc32i {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) int32 {
			return *(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) int32 {
			return 0
		}
	case common.UInt8:
		return func(from interface{}) int32 {
			return int32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt32:
		return func(from interface{}) int32 {
			return int32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt64:
		return func(from interface{}) int32 {
			return int32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int32:
		return func(from interface{}) int32 {
			return int32(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int64:
		return func(from interface{}) int32 {
			return int32(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float32:
		return func(from interface{}) int32 {
			return int32(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float64:
		return func(from interface{}) int32 {
			return int32(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.String:
		panic("string to uint32 value cannot be auto-converted")
	case common.Bytes:
		panic("bytes to uint32 value cannot be auto-converted")
	case common.Bool:
		panic("bool to uint32 value cannot be auto-converted")
	case common.Slice:
		panic("slice to uint32 value cannot be auto-converted")
	case common.Map:
		panic("map to uint32 value cannot be auto-converted")
	case common.Message:
		panic("message to uint32 value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to uint32 value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFunc32f(from *common.DataType, to *common.DataType) TypeConvertFunc32f {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) float32 {
			return *(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) float32 {
			return 0
		}
	case common.UInt8:
		return func(from interface{}) float32 {
			return float32(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt32:
		return func(from interface{}) float32 {
			return float32(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt64:
		return func(from interface{}) float32 {
			return float32(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int32:
		return func(from interface{}) float32 {
			return float32(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int64:
		return func(from interface{}) float32 {
			return float32(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float32:
		return func(from interface{}) float32 {
			return float32(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float64:
		return func(from interface{}) float32 {
			return float32(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.String:
		panic("string to uint32 value cannot be auto-converted")
	case common.Bytes:
		panic("bytes to uint32 value cannot be auto-converted")
	case common.Bool:
		panic("bool to uint32 value cannot be auto-converted")
	case common.Slice:
		panic("slice to uint32 value cannot be auto-converted")
	case common.Map:
		panic("map to uint32 value cannot be auto-converted")
	case common.Message:
		panic("message to uint32 value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to uint32 value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFunc64u(from *common.DataType, to *common.DataType) TypeConvertFunc64u {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) uint64 {
			return *(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) uint64 {
			return 0
		}
	case common.UInt8:
		return func(from interface{}) uint64 {
			return uint64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt32:
		return func(from interface{}) uint64 {
			return uint64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt64:
		return func(from interface{}) uint64 {
			return uint64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int32:
		return func(from interface{}) uint64 {
			return uint64(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int64:
		return func(from interface{}) uint64 {
			return uint64(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float32:
		return func(from interface{}) uint64 {
			return uint64(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float64:
		return func(from interface{}) uint64 {
			return uint64(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.String:
		panic("string to 64bit value cannot be auto-converted")
	case common.Bytes:
		panic("bytes to 64bit value cannot be auto-converted")
	case common.Bool:
		panic("bool to 64bit value cannot be auto-converted")
	case common.Slice:
		panic("slice to 64bit value cannot be auto-converted")
	case common.Map:
		panic("map to 64bit value cannot be auto-converted")
	case common.Message:
		panic("message to 64bit value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to 32bit value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFunc64i(from *common.DataType, to *common.DataType) TypeConvertFunc64i {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) int64 {
			return *(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) int64 {
			return 0
		}
	case common.UInt8:
		return func(from interface{}) int64 {
			return int64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt32:
		return func(from interface{}) int64 {
			return int64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt64:
		return func(from interface{}) int64 {
			return int64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int32:
		return func(from interface{}) int64 {
			return int64(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int64:
		return func(from interface{}) int64 {
			return int64(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float32:
		return func(from interface{}) int64 {
			return int64(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float64:
		return func(from interface{}) int64 {
			return int64(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.String:
		panic("string to 64bit value cannot be auto-converted")
	case common.Bytes:
		panic("bytes to 64bit value cannot be auto-converted")
	case common.Bool:
		panic("bool to 64bit value cannot be auto-converted")
	case common.Slice:
		panic("slice to 64bit value cannot be auto-converted")
	case common.Map:
		panic("map to 64bit value cannot be auto-converted")
	case common.Message:
		panic("message to 64bit value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to 32bit value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFunc64f(from *common.DataType, to *common.DataType) TypeConvertFunc64f {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) float64 {
			return *(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) float64 {
			return 0
		}
	case common.UInt8:
		return func(from interface{}) float64 {
			return float64(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt32:
		return func(from interface{}) float64 {
			return float64(*(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt64:
		return func(from interface{}) float64 {
			return float64(*(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int32:
		return func(from interface{}) float64 {
			return float64(*(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Int64:
		return func(from interface{}) float64 {
			return float64(*(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float32:
		return func(from interface{}) float64 {
			return float64(*(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Float64:
		return func(from interface{}) float64 {
			return float64(*(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.String:
		panic("string to 64bit value cannot be auto-converted")
	case common.Bytes:
		panic("bytes to 64bit value cannot be auto-converted")
	case common.Bool:
		panic("bool to 64bit value cannot be auto-converted")
	case common.Slice:
		panic("slice to 64bit value cannot be auto-converted")
	case common.Map:
		panic("map to 64bit value cannot be auto-converted")
	case common.Message:
		panic("message to 64bit value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to 32bit value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFuncBool(from *common.DataType, to *common.DataType) TypeConvertFuncBool {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) bool {
			return *(*bool)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) bool {
			return false
		}
	case common.UInt8:
		return func(from interface{}) bool {
			return *(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}
	case common.UInt32:
		return func(from interface{}) bool {
			return *(*uint32)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}
	case common.UInt64:
		return func(from interface{}) bool {
			return *(*uint64)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}
	case common.Int32:
		return func(from interface{}) bool {
			return *(*int32)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}
	case common.Int64:
		return func(from interface{}) bool {
			return *(*int64)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}
	case common.Float32:
		return func(from interface{}) bool {
			return *(*float32)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}
	case common.Float64:
		return func(from interface{}) bool {
			return *(*float64)((*EmptyFace)(unsafe.Pointer(&from)).Data) != 0
		}
	case common.String:
		panic("string to bool value cannot be auto-converted")
	case common.Bytes:
		panic("bytes to bool value cannot be auto-converted")
	case common.Bool:
		return func(from interface{}) bool {
			return *(*bool)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Slice:
		panic("slice to bool value cannot be auto-converted")
	case common.Map:
		panic("map to bool value cannot be auto-converted")
	case common.Message:
		panic("message to bool value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to bool value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFuncStr(from *common.DataType, to *common.DataType) TypeConvertFuncStr {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) string {
			return *(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) string {
			return ""
		}
	case common.UInt8:
		return func(from interface{}) string {
			return string(*(*uint8)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.UInt32:
		panic("uint32 to string value cannot be auto-converted")
	case common.UInt64:
		panic("uint64 to string value cannot be auto-converted")
	case common.Int32:
		panic("int32 to string value cannot be auto-converted")
	case common.Int64:
		panic("int64 to string value cannot be auto-converted")
	case common.Float32:
		panic("float32 to string value cannot be auto-converted")
	case common.Float64:
		panic("float64 to string value cannot be auto-converted")
	case common.String:
		return func(from interface{}) string {
			return *(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Bytes:
		return func(from interface{}) string {
			return string(*(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Bool:
		panic("bool to string value cannot be auto-converted")
	case common.Slice:
		panic("slice to string value cannot be auto-converted")
	case common.Map:
		panic("map to string value cannot be auto-converted")
	case common.Message:
		panic("message to string value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to string value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFuncBytes(from *common.DataType, to *common.DataType) TypeConvertFuncBytes {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) []byte {
			return *(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) []byte {
			return []byte{}
		}
	case common.UInt8:
		panic("uint8 to bytes value cannot be auto-converted")
	case common.UInt32:
		panic("uint32 to bytes value cannot be auto-converted")
	case common.UInt64:
		panic("uint64 to bytes value cannot be auto-converted")
	case common.Int32:
		panic("int32 to bytes value cannot be auto-converted")
	case common.Int64:
		panic("int64 to bytes value cannot be auto-converted")
	case common.Float32:
		panic("float32 to bytes value cannot be auto-converted")
	case common.Float64:
		panic("float64 to bytes value cannot be auto-converted")
	case common.String:
		return func(from interface{}) []byte {
			return []byte(*(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data))
		}
	case common.Bytes:
		return func(from interface{}) []byte {
			return *(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Bool:
		panic("bool to string value cannot be auto-converted")
	case common.Slice:
		panic("slice to string value cannot be auto-converted")
	case common.Map:
		panic("map to string value cannot be auto-converted")
	case common.Message:
		panic("message to string value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to string value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFuncMap(from *common.DataType, to *common.DataType) TypeConvertFuncMap {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) interface{} {
			return from
		}
	case common.Nil:
		switch to.KeyType.Kind.Kind {
		case common.UInt32:
			if from.KeyType.Kind.Kind != common.UInt32 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				return make(map[uint32]interface{})
			}
		case common.UInt64:
			if from.KeyType.Kind.Kind != common.UInt64 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				return make(map[uint64]interface{})
			}
		case common.Int32:
			if from.KeyType.Kind.Kind != common.Int32 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				return make(map[int32]interface{})
			}
		case common.Int64:
			if from.KeyType.Kind.Kind != common.Int64 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				return make(map[int64]interface{})
			}
		case common.Float32:
			if from.KeyType.Kind.Kind != common.Float32 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				return make(map[float32]interface{})
			}
		case common.Float64:
			if from.KeyType.Kind.Kind != common.Float64 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				return make(map[float64]interface{})
			}
		case common.Bool:
			if from.KeyType.Kind.Kind != common.Bool {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				return make(map[bool]interface{})
			}
		case common.String:
			if from.KeyType.Kind.Kind != common.String {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				return make(map[string]interface{})
			}
		default:
			panic("unsupported map key type")
		}
	case common.UInt32:
		panic("uint32 to map value cannot be auto-converted")
	case common.UInt64:
		panic("uint64 to map value cannot be auto-converted")
	case common.Int32:
		panic("int32 to map value cannot be auto-converted")
	case common.Int64:
		panic("int64 to map value cannot be auto-converted")
	case common.Float32:
		panic("float32 to map value cannot be auto-converted")
	case common.Float64:
		panic("float64 to map value cannot be auto-converted")
	case common.String:
		panic("string to map value cannot be auto-converted")
	case common.Bytes:
		panic("bytes to map value cannot be auto-converted")
	case common.Bool:
		panic("bool to map value cannot be auto-converted")
	case common.Slice:
		panic("slice to map value cannot be auto-converted")
	case common.Map:
		switch to.KeyType.Kind.Kind {
		case common.UInt32:
			if from.KeyType.Kind.Kind != common.UInt32 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return make(map[uint32]interface{})
				}
				return from
			}
		case common.UInt64:
			if from.KeyType.Kind.Kind != common.UInt64 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return make(map[uint64]interface{})
				}
				return from
			}
		case common.Int32:
			if from.KeyType.Kind.Kind != common.Int32 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return make(map[int32]interface{})
				}
				return from
			}
		case common.Int64:
			if from.KeyType.Kind.Kind != common.Int64 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return make(map[int64]interface{})
				}
				return from
			}
		case common.Float32:
			if from.KeyType.Kind.Kind != common.Float32 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return make(map[float32]interface{})
				}
				return from
			}
		case common.Float64:
			if from.KeyType.Kind.Kind != common.Float64 {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return make(map[float64]interface{})
				}
				return from
			}
		case common.Bool:
			if from.KeyType.Kind.Kind != common.Bool {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return make(map[bool]interface{})
				}
				return from
			}
		case common.String:
			if from.KeyType.Kind.Kind != common.String {
				panic("cannot covert to a different key type map")
			}
			return func(from interface{}) (to interface{}) {
				if from == nil {
					return make(map[string]interface{})
				}
				return from
			}
		default:
			panic("unsupported map key type")
		}
	case common.Message:
		panic("message to map value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to map value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFuncSlice(from *common.DataType, to *common.DataType) TypeConvertFuncSlice {
	switch from.Kind.Kind {
	case common.Object:
		return func(from interface{}) []interface{} {
			return *(*[]interface{})((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Nil:
		return func(from interface{}) []interface{} {
			return []interface{}{}
		}
	case common.UInt32:
		panic("uint32 to slice value cannot be auto-converted")
	case common.UInt64:
		panic("uint64 to slice value cannot be auto-converted")
	case common.Int32:
		panic("int32 to slice cannot be auto-converted")
	case common.Int64:
		panic("int64 to slice value cannot be auto-converted")
	case common.Float32:
		panic("float32 to slice value cannot be auto-converted")
	case common.Float64:
		panic("float64 to slice value cannot be auto-converted")
	case common.String:
		panic("string to slice value cannot be auto-converted")
	case common.Bytes:
		panic("bytes to slice value cannot be auto-converted")
	case common.Bool:
		panic("bool to slice value cannot be auto-converted")
	case common.Slice:
		return func(from interface{}) []interface{} {
			if from == nil {
				return []interface{}{}
			}
			return *(*[]interface{})((*EmptyFace)(unsafe.Pointer(&from)).Data)
		}
	case common.Map:
		panic("map to slice value cannot be auto-converted")
	case common.Message:
		panic("message to slice value cannot be auto-converted")
	case common.ReflectType:
		panic("reflect to slice value cannot be auto-converted")
	default:
		panic("unknown type convert")
	}
}

func GetConvertFuncObject(from *common.DataType, to *common.DataType) TypeConvertFuncObject {
	return func(from interface{}) interface{} {
		return from
	}
}

type LengthFunc func(from interface{}) int64

func GetLengthFunc(from *common.DataType) LengthFunc {
	switch from.Kind.Kind {
	case common.String:
		return func(from interface{}) int64 {
			return int64(len(*(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data)))
		}
	case common.Bytes:
		return func(from interface{}) int64 {
			return int64(len(*(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data)))
		}
	case common.Slice:
		return func(from interface{}) int64 {
			return int64(len(*(*[]interface{})((*EmptyFace)(unsafe.Pointer(&from)).Data)))
		}
	case common.Map:
		return func(from interface{}) int64 {
			return hack.MLen(from)
		}
	default:
		panic("cannot take length " + from.Type)
	}
}
