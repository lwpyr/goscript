package instruction

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/hack"
	"unsafe"
)

type LengthFunc func(from interface{}) int64

func GetLengthFunc(from *common.DataType) LengthFunc {
	switch from.Kind.Kind {
	case common.String:
		return GetStringLenFunc()
	case common.Bytes:
		return GetBytesLenFunc()
	case common.Slice:
		return GetSliceLenFunc()
	case common.Map:
		return GetMapLenFunc()
	default:
		panic("cannot take length " + from.Type)
	}
}

func GetStringLenFunc() LengthFunc {
	return func(from interface{}) int64 {
		return int64(len(*(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data)))
	}
}

func GetBytesLenFunc() LengthFunc {
	return func(from interface{}) int64 {
		return int64(len(*(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data)))
	}
}

func GetSliceLenFunc() LengthFunc {
	return func(from interface{}) int64 {
		return int64(len(*(*[]interface{})((*EmptyFace)(unsafe.Pointer(&from)).Data)))
	}
}

func GetMapLenFunc() LengthFunc {
	return func(from interface{}) int64 {
		return hack.MLen(from)
	}
}
