package hack

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestMapBool(t *testing.T) {
	a := map[bool]interface{}{}

	// Get
	a[true] = "A"
	key1 := true
	aptr := MapGet(a, unsafe.Pointer(&key1))
	assert.Equal(t, "A", *(*interface{})(aptr))

	// Set
	key2 := false
	bptr := MapGetOrCreate(a, unsafe.Pointer(&key2))
	*(*interface{})(bptr) = 123
	assert.Equal(t, 123, a[false])

	// Len
	assert.Equal(t, int64(2), MLen(a))

	// Del
	MapDelete(a, unsafe.Pointer(&key1))
	assert.Equal(t, nil, a[true])

	// Len
	assert.Equal(t, int64(1), MLen(a))
}

func TestMap32(t *testing.T) {
	a := map[int32]interface{}{}

	// Get
	a[1] = "A"
	aptr := Map32Get(a, 1)
	assert.Equal(t, "A", *(*interface{})(aptr))

	// Set
	var key int32 = -1
	bptr := Map32GetOrCreate(a, *(*uint32)(unsafe.Pointer(&key)))
	*(*interface{})(bptr) = 123
	assert.Equal(t, 123, a[-1])

	// Len
	assert.Equal(t, int64(2), MLen(a))

	// Del
	Map32Delete(a, 1)
	assert.Equal(t, nil, a[1])

	// Len
	assert.Equal(t, int64(1), MLen(a))
}

func TestMap64(t *testing.T) {
	a := map[int64]interface{}{}

	// Get
	a[1] = "A"
	aptr := Map64Get(a, 1)
	assert.Equal(t, "A", *(*interface{})(aptr))

	// Set
	var key int64 = -1
	bptr := Map64GetOrCreate(a, *(*uint64)(unsafe.Pointer(&key)))
	*(*interface{})(bptr) = 123
	assert.Equal(t, 123, a[-1])

	// Len
	assert.Equal(t, int64(2), MLen(a))

	// Del
	Map64Delete(a, 1)
	assert.Equal(t, nil, a[1])

	// Len
	assert.Equal(t, int64(1), MLen(a))
}

func TestMapStr(t *testing.T) {
	a := map[string]interface{}{}

	// Get
	a["a"] = "A"
	aptr := MapStrGet(a, "a")
	assert.Equal(t, "A", *(*interface{})(aptr))

	// Set
	bptr := MapStrGetOrCreate(a, "b")
	*(*interface{})(bptr) = 123
	assert.Equal(t, 123, a["b"])

	// Len
	assert.Equal(t, int64(2), MLen(a))

	// Del
	MapStrDelete(a, "a")
	assert.Equal(t, nil, a["a"])

	// Len
	assert.Equal(t, int64(1), MLen(a))
}

func TestMap(t *testing.T) {

	// string key map
	s := map[interface{}]interface{}{}
	// Get
	s["a"] = "A"
	var keystr interface{} = "a"
	aptr := MapGet(s, unsafe.Pointer(&keystr))
	assert.Equal(t, "A", *(*interface{})(aptr))

	// Set
	keystr = "b"
	bptr := MapGetOrCreate(s, unsafe.Pointer(&keystr))
	*(*interface{})(bptr) = "123"
	assert.Equal(t, "123", s["b"])

	// Len
	assert.Equal(t, int64(2), MLen(s))

	// Del
	MapDelete(s, unsafe.Pointer(&keystr))
	assert.Equal(t, nil, s["b"])

	// Len
	assert.Equal(t, int64(1), MLen(s))

	// number key map
	n := map[interface{}]interface{}{}
	// Get
	n[int64(1)] = "A"
	var keyint interface{} = int64(1)
	var keyint2 interface{} = int64(2)
	aptr = MapGet(n, unsafe.Pointer(&keyint))
	assert.Equal(t, "A", *(*interface{})(aptr))

	// Set
	bptr = MapGetOrCreate(n, unsafe.Pointer(&keyint))
	assert.Equal(t, "A", *(*interface{})(bptr))
	bptr = MapGetOrCreate(n, unsafe.Pointer(&keyint2))
	*(*interface{})(bptr) = "123"
	assert.Equal(t, "123", n[int64(2)])

	// Len
	assert.Equal(t, int64(2), MLen(n))

	// Del
	MapDelete(n, unsafe.Pointer(&keyint))
	assert.Equal(t, nil, n[int64(1)])

	// Len
	assert.Equal(t, int64(1), MLen(n))
}
