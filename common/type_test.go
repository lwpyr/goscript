package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestType(t *testing.T) {
	tr := NewTypeRegistry()

	tr.AddType("HttpHeader", tr.FindMapType("string", "string"))
	assert.Equal(t, "map<string,string>", tr.FindType("HttpHeader").Type)
	assert.Equal(t, BasicTypeMap["string"], tr.FindType("HttpHeader").KeyType)
	assert.Equal(t, BasicTypeMap["string"], tr.FindType("HttpHeader").ValueType)

	tr.AddType("CookieStrings", tr.FindSliceType("string"))
	assert.Equal(t, "slice<string>", tr.FindType("CookieStrings").Type)
	assert.Equal(t, BasicTypeMap["string"], tr.FindType("CookieStrings").ItemType)

	// build message type
	dtb := NewDataTypeBuilder("HttpRequest")
	dtb.SetKind(Message)
	dtb.SetField("header", tr.FindType("HttpHeader"))
	dtb.SetField("cookie", tr.FindType("CookieStrings"))
	tr.AddBuiltType(dtb)

	assert.Equal(t, "HttpRequest", tr.FindType("HttpRequest").Type)
	assert.Equal(t, Message, tr.FindType("HttpRequest").Kind.Kind)
	assert.Equal(t, tr.FindType("HttpHeader"), tr.FindType("HttpRequest").FieldType["header"])
	assert.Equal(t, tr.FindType("CookieStrings"), tr.FindType("HttpRequest").FieldType["cookie"])

	// build message type with one of
	dtb = NewDataTypeBuilder("oneOf")
	dtb.SetKind(Message)
	dtb.SetField("str", tr.FindType("string"))
	dtb.SetField("int", tr.FindType("int64"))
	dtb.SetField("float", tr.FindType("float64"))
	dtb.SetOneOf("number", "float")
	dtb.SetOneOf("number", "int")
	tr.AddBuiltType(dtb)

	assert.Equal(t, "oneOf", tr.FindType("oneOf").Type)
	assert.Equal(t, Message, tr.FindType("oneOf").Kind.Kind)
	assert.Equal(t, tr.FindType("string"), tr.FindType("oneOf").FieldType["str"])
	assert.Equal(t, tr.FindType("int64"), tr.FindType("oneOf").FieldType["int"])
	assert.Equal(t, tr.FindType("float64"), tr.FindType("oneOf").FieldType["float"])

	assert.Equal(t, []string{"float", "int"}, tr.FindType("oneOf").OneOfGroup["number"])
	assert.Equal(t, "number", tr.FindType("oneOf").OneOf["int"])
	assert.Equal(t, "number", tr.FindType("oneOf").OneOf["float"])

	// build map
	dtb = NewDataTypeBuilder("RequestMap")
	dtb.SetKind(Map)
	dtb.SetKey(tr.FindType("string"))
	dtb.SetValue(tr.FindType("HttpRequest"))
	tr.AddBuiltType(dtb)

	assert.Equal(t, "map<string,HttpRequest>", tr.FindType("RequestMap").Type)
	assert.Equal(t, Map, tr.FindType("RequestMap").Kind.Kind)

	assert.Equal(t, tr.FindType("string"), tr.FindType("RequestMap").KeyType)
	assert.Equal(t, tr.FindType("HttpRequest"), tr.FindType("RequestMap").ValueType)

	// build sliceType
	dtb = NewDataTypeBuilder("IntArray")
	dtb.SetKind(Slice)
	dtb.SetItem(tr.FindType("int64"))
	tr.AddBuiltType(dtb)

	assert.Equal(t, "slice<int64>", tr.FindType("IntArray").Type)
	assert.Equal(t, Slice, tr.FindType("IntArray").Kind.Kind)

	assert.Equal(t, tr.FindType("int64"), tr.FindType("IntArray").ItemType)

	// overwrite
	tr.AddType("IntArray", tr.FindSliceType("int32"))
	assert.Equal(t, "slice<int32>", tr.FindType("IntArray").Type)
	tr.AddBuiltType(dtb)
	assert.Equal(t, "slice<int64>", tr.FindType("IntArray").Type)
}

func TestPanic(t *testing.T) {
	tr := NewTypeRegistry()

	assert.Panics(t, func() { tr.AddType("int64", tr.FindMapType("string", "string")) })

	dtb := NewDataTypeBuilder("int64")
	assert.Panics(t, func() { dtb.SetKind(Object) })
	assert.Panics(t, func() { tr.AddBuiltType(dtb) })
	assert.Panics(t, func() { dtb.SetKind(Map); dtb.SetField("number", tr.FindType("int64")) })
	assert.Panics(t, func() { dtb.SetKind(Map); dtb.SetOneOf("number", "foo") })
	assert.Panics(t, func() { dtb.SetKind(Message); dtb.SetKey(tr.FindType("int64")) })
	assert.Panics(t, func() { dtb.SetKind(Message); dtb.SetValue(tr.FindType("int64")) })
	assert.Panics(t, func() { dtb.SetKind(Message); dtb.SetOneOf("foo", "bar") })
	assert.Panics(t, func() { dtb.SetKind(Message); dtb.SetItem(tr.FindType("int64")) })
}

func TestNil(t *testing.T) {
	tr := NewTypeRegistry()

	assert.Nil(t, tr.FindType("foo"))
	assert.Nil(t, tr.FindSliceType("foo"))
	assert.Nil(t, tr.FindMapType("foo", "bar"))
}

func TestUtil(t *testing.T) {
	tr := NewTypeRegistry()

	tr.AddType("HttpHeader", tr.FindMapType("string", "string"))
	tr.AddType("CookieStrings", tr.FindSliceType("string"))

	dtb := NewDataTypeBuilder("HttpRequest")
	dtb.SetKind(Message)
	dtb.SetField("header", tr.FindType("HttpHeader"))
	dtb.SetField("cookie", tr.FindType("CookieStrings"))
	tr.AddBuiltType(dtb)

	assert.True(t, tr.FindType("HttpRequest").CanConvertTo(tr.FindType("HttpRequest")))
	assert.True(t, tr.FindType("int64").CanConvertTo(tr.FindType("float64")))
	assert.True(t, tr.FindType("nil").CanConvertTo(tr.FindType("float64")))
	assert.True(t, tr.FindType("HttpHeader").CanConvertTo(tr.FindType("HttpHeader")))
	assert.True(t, tr.FindType("CookieStrings").CanConvertTo(tr.FindType("CookieStrings")))
	assert.True(t, tr.FindType("string").CanConvertTo(tr.FindType("bytes")))
	assert.False(t, tr.FindType("string").CanConvertTo(tr.FindType("nil")))
}
