package hack

import "unsafe"

type Eface struct {
	_type unsafe.Pointer
	Data  unsafe.Pointer
}

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func SliceLen(s interface{}) int64 {
	return int64((*slice)((*Eface)(unsafe.Pointer(&s)).Data).len)
}

func SliceIndex(s interface{}, i int64) *interface{} {
	return &(*(*[]interface{})((*Eface)(unsafe.Pointer(&s)).Data))[i]
}
