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
		return StringLen()
	case common.Bytes:
		return BytesLen()
	case common.Slice:
		return SliceLen()
	case common.Map:
		return MapLen()
	default:
		panic("cannot take length " + from.Type)
	}
}

func StringLen() LengthFunc {
	return func(from interface{}) int64 {
		return int64(len(*(*string)((*EmptyFace)(unsafe.Pointer(&from)).Data)))
	}
}

func BytesLen() LengthFunc {
	return func(from interface{}) int64 {
		return int64(len(*(*[]byte)((*EmptyFace)(unsafe.Pointer(&from)).Data)))
	}
}

func SliceLen() LengthFunc {
	return func(from interface{}) int64 {
		return int64(len(*(*[]interface{})((*EmptyFace)(unsafe.Pointer(&from)).Data)))
	}
}

func MapLen() LengthFunc {
	return func(from interface{}) int64 {
		return hack.MLen(from)
	}
}
