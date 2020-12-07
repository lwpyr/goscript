package lambda_chains

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/hack"
	"unsafe"
)

func GetMapDelFunc(keyType *common.DataType) (mapSet func(m interface{}, k interface{})) {
	switch keyType.Kind.Kind {
	case common.Int32, common.UInt32, common.Float32:
		mapSet = DelMap32Field
	case common.Int64, common.UInt64, common.Float64:
		mapSet = DelMap64Field
	case common.String:
		mapSet = DelMapStrField
	case common.Bool:
		mapSet = DelMapBoolField
	default:
		panic("unknown key type")
	}
	return
}

func GetMapMustGetFunc(keyType *common.DataType) (mapSet func(m interface{}, k interface{}) (v *interface{})) {
	switch keyType.Kind.Kind {
	case common.Int32, common.UInt32, common.Float32:
		mapSet = MustGetMap32Field
	case common.Int64, common.UInt64, common.Float64:
		mapSet = MustGetMap64Field
	case common.String:
		mapSet = MustGetMapStrField
	case common.Bool:
		mapSet = MustGetMapBoolField
	default:
		panic("unknown key type")
	}
	return
}

func GetMapSetFunc(keyType *common.DataType) (mapSet func(m interface{}, k interface{}, v interface{})) {
	switch keyType.Kind.Kind {
	case common.Int32, common.UInt32, common.Float32:
		mapSet = SetMap32Field
	case common.Int64, common.UInt64, common.Float64:
		mapSet = SetMap64Field
	case common.String:
		mapSet = SetMapStrField
	case common.Bool:
		mapSet = SetMapBoolField
	default:
		panic("unknown key type")
	}
	return
}

func GetMapGetFunc(keyType *common.DataType) (mapGet func(m interface{}, k interface{}) (v interface{})) {
	switch keyType.Kind.Kind {
	case common.Int32, common.UInt32, common.Float32:
		mapGet = GetMap32Field
	case common.Int64, common.UInt64, common.Float64:
		mapGet = GetMap64Field
	case common.String:
		mapGet = GetMapStrField
	case common.Bool:
		mapGet = GetMapBoolField
	default:
		panic("unknown key type")
	}
	return
}

func MessageGetField(m interface{}, k string) interface{} {
	return *(*interface{})(hack.MapStrGet(m, k))
}

func MessageMustGetField(m interface{}, k string) *interface{} {
	return (*interface{})(hack.MapStrGetOrCreate(m, k))
}

func MessageSetField(m interface{}, k string, v interface{}) {
	*(*interface{})(hack.MapStrGetOrCreate(m, k)) = v
}

func MessageResetField(m interface{}, k string) {
	hack.MapStrDelete(m, k)
}

func GetMapStrField(m interface{}, k interface{}) interface{} {
	return *(*interface{})(hack.MapStrGet(m, *(*string)((*EmptyFace)(unsafe.Pointer(&k)).Data)))
}

func MustGetMapStrField(m interface{}, k interface{}) *interface{} {
	return (*interface{})(hack.MapStrGetOrCreate(m, *(*string)((*EmptyFace)(unsafe.Pointer(&k)).Data)))
}

func SetMapStrField(m interface{}, k interface{}, v interface{}) {
	*(*interface{})(hack.MapStrGetOrCreate(m, *(*string)((*EmptyFace)(unsafe.Pointer(&k)).Data))) = v
}

func DelMapStrField(m interface{}, k interface{}) {
	hack.MapStrDelete(m, *(*string)((*EmptyFace)(unsafe.Pointer(&k)).Data))
}

func GetMap32Field(m interface{}, k interface{}) interface{} {
	return *(*interface{})(hack.Map32Get(m, *(*uint32)((*EmptyFace)(unsafe.Pointer(&k)).Data)))
}

func MustGetMap32Field(m interface{}, k interface{}) *interface{} {
	return (*interface{})(hack.Map32GetOrCreate(m, *(*uint32)((*EmptyFace)(unsafe.Pointer(&k)).Data)))
}

func SetMap32Field(m interface{}, k interface{}, v interface{}) {
	*(*interface{})(hack.Map32GetOrCreate(m, *(*uint32)((*EmptyFace)(unsafe.Pointer(&k)).Data))) = v
}

func DelMap32Field(m interface{}, k interface{}) {
	hack.Map32Delete(m, *(*uint32)((*EmptyFace)(unsafe.Pointer(&k)).Data))
}

func GetMap64Field(m interface{}, k interface{}) interface{} {
	return *(*interface{})(hack.Map64Get(m, *(*uint64)((*EmptyFace)(unsafe.Pointer(&k)).Data)))
}

func MustGetMap64Field(m interface{}, k interface{}) *interface{} {
	return (*interface{})(hack.Map64GetOrCreate(m, *(*uint64)((*EmptyFace)(unsafe.Pointer(&k)).Data)))
}

func SetMap64Field(m interface{}, k interface{}, v interface{}) {
	*(*interface{})(hack.Map64GetOrCreate(m, *(*uint64)((*EmptyFace)(unsafe.Pointer(&k)).Data))) = v
}

func DelMap64Field(m interface{}, k interface{}) {
	hack.Map64Delete(m, *(*uint64)((*EmptyFace)(unsafe.Pointer(&k)).Data))
}

func GetMapBoolField(m interface{}, k interface{}) interface{} {
	return *(*interface{})(hack.MapGet(m, (*EmptyFace)(unsafe.Pointer(&k)).Data))
}

func MustGetMapBoolField(m interface{}, k interface{}) *interface{} {
	return (*interface{})(hack.MapGetOrCreate(m, (*EmptyFace)(unsafe.Pointer(&k)).Data))
}

func SetMapBoolField(m interface{}, k interface{}, v interface{}) {
	*(*interface{})(hack.MapGetOrCreate(m, (*EmptyFace)(unsafe.Pointer(&k)).Data)) = v
}

func DelMapBoolField(m interface{}, k interface{}) {
	hack.MapDelete(m, (*EmptyFace)(unsafe.Pointer(&k)).Data)
}

func GetMapIFaceField(m interface{}, k interface{}) interface{} {
	return *(*interface{})(hack.MapGet(m, unsafe.Pointer(&k)))
}

func MustGetMapIFaceField(m interface{}, k interface{}) *interface{} {
	return (*interface{})(hack.MapGetOrCreate(m, unsafe.Pointer(&k)))
}

func SetMapIFaceField(m interface{}, k interface{}, v interface{}) {
	*(*interface{})(hack.MapGetOrCreate(m, unsafe.Pointer(&k))) = v
}

func DelMapIFaceField(m interface{}, k interface{}) {
	hack.MapDelete(m, unsafe.Pointer(&k))
}

func GetSliceField(s interface{}, i int64) *interface{} {
	return hack.SliceIndex(s, i)
}
