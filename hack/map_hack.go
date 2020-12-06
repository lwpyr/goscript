package hack

import (
	"unsafe"
)

func MapStrGetOrCreate(m interface{}, field string) unsafe.Pointer {
	mt, h := mapTypeAndValue(m)
	return mapassign_faststr(mt, h, field)
}

func MapStrGet(m interface{}, field string) unsafe.Pointer {
	mt, h := mapTypeAndValue(m)
	return mapaccess1_faststr(mt, h, field)
}

func MapStrDelete(m interface{}, field string) {
	mt, h := mapTypeAndValue(m)
	mapdelete_faststr(mt, h, field)
}

func MapGetOrCreate(m interface{}, field unsafe.Pointer) unsafe.Pointer {
	mt, h := mapTypeAndValue(m)
	return mapassign(mt, h, field)
}

func MapGet(m interface{}, field unsafe.Pointer) unsafe.Pointer {
	mt, h := mapTypeAndValue(m)
	return mapaccess1(mt, h, field)
}

func MapDelete(m interface{}, field unsafe.Pointer) {
	mt, h := mapTypeAndValue(m)
	mapdelete(mt, h, field)
}

func Map32GetOrCreate(m interface{}, field uint32) unsafe.Pointer {
	mt, h := mapTypeAndValue(m)
	return mapassign_fast32(mt, h, field)
}

func Map32Get(m interface{}, field uint32) unsafe.Pointer {
	mt, h := mapTypeAndValue(m)
	return mapaccess1_fast32(mt, h, field)
}

func Map32Delete(m interface{}, field uint32) {
	mt, h := mapTypeAndValue(m)
	mapdelete_fast32(mt, h, field)
}

func Map64GetOrCreate(m interface{}, field uint64) unsafe.Pointer {
	mt, h := mapTypeAndValue(m)
	return mapassign_fast64(mt, h, field)
}

func Map64Get(m interface{}, field uint64) unsafe.Pointer {
	mt, h := mapTypeAndValue(m)
	return mapaccess1_fast64(mt, h, field)
}

func Map64Delete(m interface{}, field uint64) {
	mt, h := mapTypeAndValue(m)
	mapdelete_fast64(mt, h, field)
}

func MLen(m interface{}) int64 {
	_, h := mapTypeAndValue(m)
	return int64(h.count)
}

// belows comes from go runtime map header, no need to read the codes
const (
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits
)

type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32

	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr

	extra *mapextra // optional fields
}

type mapextra struct {
	overflow     *[]*bmap
	oldoverflow  *[]*bmap
	nextOverflow *bmap
}

type bmap struct {
	tophash [bucketCnt]uint8
}

type Maptype struct {
	typ        _type
	key        *_type
	elem       *_type
	bucket     *_type // internal type representing a hash bucket
	keysize    uint8  // size of key slot
	elemsize   uint8  // size of elem slot
	bucketsize uint16 // size of bucket
	flags      uint32
}

type typeAlg struct {
	hash  func(unsafe.Pointer, uintptr) uintptr
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}

type tflag uint8
type nameOff int32
type typeOff int32

type _type struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32
	tflag      tflag
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *typeAlg
	gcdata     *byte
	str        nameOff
	ptrToThis  typeOff
}

type emptyInterface struct {
	_type unsafe.Pointer
	value unsafe.Pointer
}

func mapTypeAndValue(m interface{}) (*Maptype, *hmap) {
	ei := (*emptyInterface)(unsafe.Pointer(&m))
	return (*Maptype)(ei._type), (*hmap)(ei.value)
}

//go:linkname mapassign runtime.mapassign
func mapassign(t *Maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer

//go:linkname mapaccess1 runtime.mapaccess1
func mapaccess1(t *Maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer

//go:linkname mapdelete runtime.mapdelete
func mapdelete(t *Maptype, h *hmap, key unsafe.Pointer)

//go:linkname mapaccess1_fast32 runtime.mapaccess1_fast32
func mapaccess1_fast32(t *Maptype, h *hmap, key uint32) unsafe.Pointer

//go:linkname mapassign_fast32 runtime.mapassign_fast32
func mapassign_fast32(t *Maptype, h *hmap, key uint32) unsafe.Pointer

//go:linkname mapdelete_fast32 runtime.mapdelete_fast32
func mapdelete_fast32(t *Maptype, h *hmap, key uint32)

//go:linkname mapaccess1_fast64 runtime.mapaccess1_fast64
func mapaccess1_fast64(t *Maptype, h *hmap, key uint64) unsafe.Pointer

//go:linkname mapassign_fast64 runtime.mapassign_fast64
func mapassign_fast64(t *Maptype, h *hmap, key uint64) unsafe.Pointer

//go:linkname mapdelete_fast64 runtime.mapdelete_fast64
func mapdelete_fast64(t *Maptype, h *hmap, key uint64)

//go:linkname mapaccess1_faststr runtime.mapaccess1_faststr
func mapaccess1_faststr(t *Maptype, h *hmap, ky string) unsafe.Pointer

//go:linkname mapassign_faststr runtime.mapassign_faststr
func mapassign_faststr(t *Maptype, h *hmap, s string) unsafe.Pointer

//go:linkname mapdelete_faststr runtime.mapdelete_faststr
func mapdelete_faststr(t *Maptype, h *hmap, ky string)
