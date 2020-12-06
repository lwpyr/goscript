package common

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type DataType struct {
	Type        string
	Kind        *DataKind
	FieldType   map[string]*DataType // for complex
	OneOf       map[string]string
	OneOfGroup  map[string][]string
	ItemType    *DataType // for sequence item type
	KeyType     *DataType // map key
	ValueType   *DataType // map val
	Unmarshal   func(iterator *jsoniter.Iterator) interface{}
	Constructor Constructor
}

type DataTypeBuilder struct {
	T *DataType
}

func NewDataTypeBuilder(name string) *DataTypeBuilder {
	if name == "" || len(name) >= 6 && name[:6] == "slice<" || len(name) >= 4 && name[:4] == "map<" {
		fmt.Printf("Forbiddened type name, empty string or name starts with 'slice' or 'map' is not " +
			"allowed, in the future this will panic")
		return nil
	}
	dtb := &DataTypeBuilder{
		T: &DataType{
			Type: name,
		},
	}
	return dtb
}

func (d *DataTypeBuilder) SetKind(k int) {
	if k > Message {
		panic("illegal kind in dataType")
	}
	d.T.Kind = KindMap[k]
	if k == Message {
		d.T.FieldType = map[string]*DataType{}
		d.T.OneOfGroup = map[string][]string{}
		d.T.OneOf = map[string]string{}
	}
}

func (d *DataTypeBuilder) SetField(fieldName string, fieldType *DataType) {
	if d.T.Kind.Kind != Message {
		panic("cannot set a field type of a none message dataType")
	}
	d.T.FieldType[fieldName] = fieldType
}

func (d *DataTypeBuilder) SetOneOf(oneOfName string, fieldName string) {
	if d.T.Kind.Kind != Message {
		panic("cannot set a field type of a none message dataType")
	}
	if _, ok := d.T.FieldType[fieldName]; !ok {
		panic("cannot set a none existing field type in SetOneOf")
	}
	if group, ok := d.T.OneOfGroup[oneOfName]; !ok {
		d.T.OneOfGroup[oneOfName] = []string{fieldName}
	} else {
		d.T.OneOfGroup[oneOfName] = append(group, fieldName)
	}
	d.T.OneOf[fieldName] = oneOfName
}

func (d *DataTypeBuilder) SetKey(keyType *DataType) {
	if d.T.Kind.Kind != Map {
		panic("cannot set key type of a none map dataType")
	}
	d.T.KeyType = keyType
}

func (d *DataTypeBuilder) SetValue(valType *DataType) {
	if d.T.Kind.Kind != Map {
		panic("cannot set value type of a none map dataType")
	}
	d.T.ValueType = valType
}

func (d *DataTypeBuilder) SetItem(itemType *DataType) {
	if d.T.Kind.Kind != Slice {
		panic("cannot set item type of a none slice dataType")
	}
	d.T.ItemType = itemType
}

func (d *DataType) CanConvertTo(b *DataType) bool {
	if d.Kind.Kind == Nil || b.Kind.Kind == Object || d.Kind.Kind == Object {
		return true
	}
	switch b.Kind.Kind {
	case Message:
		return d.Type == b.Type
	case Map:
		return d.KeyType == b.KeyType && d.ValueType == b.ValueType
	case Slice:
		return d.ItemType == b.ItemType
	case Bytes, String:
		return d.Kind.Kind == Bytes || d.Kind.Kind == String
	case Nil, UInt8: // now we don't open the uint8 type
		return false
	default:
		return true
	}
}

func (d *DataType) CanNegate() bool {
	return d.Kind.Kind == Bool
}

func CanCompare(d *DataType, b *DataType) *DataType {
	switch b.Kind.Kind {
	case Message, Map, Slice, Bool:
		return nil
	case String:
		if d.Kind.Kind == String {
			return BasicTypeMap["string"]
		}
		return nil
	case Int32, Int64, UInt32, UInt64, Float32, Float64:
		switch d.Kind.Kind {
		case Int32, Int64, UInt32, UInt64, Float32, Float64:
			return BasicTypeMap["bool"]
		default:
			return nil
		}
	default:
		return nil
	}
}

func CanEqual(d *DataType, b *DataType) *DataType {
	if b.Kind.Kind == Nil || d.Kind.Kind == Nil {
		return BasicTypeMap["bool"]
	}
	switch b.Kind.Kind {
	case Message, Map, Slice:
		return nil
	case ReflectType:
		if d.Kind.Kind == ReflectType {
			return BasicTypeMap["bool"]
		}
		return nil
	case Bool:
		if d.Kind.Kind == Bool {
			return BasicTypeMap["bool"]
		}
		return nil
	case String:
		if d.Kind.Kind == String {
			return BasicTypeMap["bool"]
		}
		return nil
	case Int32, Int64, UInt32, UInt64, Float32, Float64:
		switch d.Kind.Kind {
		case Int32, Int64, UInt32, UInt64, Float32, Float64:
			return BasicTypeMap["bool"]
		default:
			return nil
		}
	default:
		return nil
	}
}

func CanAndOr(d *DataType, b *DataType) *DataType {
	if d.Kind.Kind == Bool && b.Kind.Kind == Bool {
		return BasicTypeMap["bool"]
	}
	return nil
}

func CanMod(d *DataType, b *DataType) *DataType {
	switch b.Kind.Kind {
	case Int32, Int64, UInt32, UInt64:
		switch d.Kind.Kind {
		case Int32, Int64, UInt32, UInt64:
			if b.Kind.Kind <= Int32 && d.Kind.Kind <= Int32 {
				return BasicTypeMap["int32"]
			} else {
				return BasicTypeMap["int64"]
			}
		default:
			return nil
		}
	default:
		return nil
	}
}

func CanAddSubMulPow(d *DataType, b *DataType) *DataType {
	switch b.Kind.Kind {
	case Int32, Int64, UInt32, UInt64:
		switch d.Kind.Kind {
		case Int32, Int64, UInt32, UInt64:
			if b.Kind.Kind <= Int32 && d.Kind.Kind <= Int32 {
				return BasicTypeMap["int32"]
			} else {
				return BasicTypeMap["int64"]
			}
		case Float32, Float64:
			if b.Kind.Kind <= Float32 && d.Kind.Kind <= Float32 {
				return BasicTypeMap["float32"]
			} else {
				return BasicTypeMap["float64"]
			}
		default:
			return nil
		}
	case Float32, Float64:
		switch d.Kind.Kind {
		case Int32, Int64, UInt32, UInt64, Float32, Float64:
			if b.Kind.Kind <= Float32 && d.Kind.Kind <= Float32 {
				return BasicTypeMap["float32"]
			} else {
				return BasicTypeMap["float64"]
			}
		default:
			return nil
		}
	default:
		return nil
	}
}

func CanDiv(d *DataType, b *DataType) *DataType {
	switch b.Kind.Kind {
	case Int32, Int64, UInt32, UInt64, Float32, Float64:
		switch d.Kind.Kind {
		case Int32, Int64, UInt32, UInt64, Float32, Float64:
			if b.Kind.Kind <= Float32 && d.Kind.Kind <= Float32 {
				return BasicTypeMap["float32"]
			} else {
				return BasicTypeMap["float64"]
			}
		default:
			return nil
		}
	default:
		return nil
	}
}

var BasicTypeMap = map[string]*DataType{
	"uint8": {
		Type: "uint8",
		Kind: KindMap[UInt8],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadInt32()
		},
		Constructor: ConstructorMap[UInt8],
	},
	"int32": {
		Type: "int32",
		Kind: KindMap[Int32],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadInt32()
		},
		Constructor: ConstructorMap[Int32],
	},
	"int64": {
		Type: "int64",
		Kind: KindMap[Int64],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadInt64()
		},
		Constructor: ConstructorMap[Int64],
	},
	"uint32": {
		Type: "uint32",
		Kind: KindMap[UInt32],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadUint32()
		},
		Constructor: ConstructorMap[UInt32],
	},
	"uint64": {
		Type: "uint64",
		Kind: KindMap[UInt64],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadUint32()
		},
		Constructor: ConstructorMap[UInt64],
	},
	"float32": {
		Type: "float32",
		Kind: KindMap[Float32],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadFloat32()
		},
		Constructor: ConstructorMap[Float32],
	},
	"float64": {
		Type: "float64",
		Kind: KindMap[Float64],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadFloat64()
		},
		Constructor: ConstructorMap[Float64],
	},
	"string": {
		Type: "string",
		Kind: KindMap[String],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadString()
		},
		Constructor: ConstructorMap[String],
	},
	"bytes": {
		Type: "bytes",
		Kind: KindMap[Bytes],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return []byte(iter.ReadString())
		},
		Constructor: ConstructorMap[Bytes],
	},
	"bool": {
		Type: "bool",
		Kind: KindMap[Bool],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadBool()
		},
		Constructor: ConstructorMap[Bool],
	},
	"nil": {
		Type: "nil",
		Kind: KindMap[Nil],
	},
	"reflect": { // only for reflect type
		Type: "reflect",
		Kind: KindMap[ReflectType],
	},
	"object": { // like Object
		Type:        "object",
		Kind:        KindMap[Object],
		Constructor: ConstructorMap[Object],
	},
	"tuple": { // for system use, function return more than one value
		Type: "tuple",
		Kind: KindMap[Object],
	},
}

func init() {
	BasicTypeMap["enumVal"] = &DataType{
		Type:        "enumVal",
		Kind:        KindMap[Int32],
		Constructor: ConstructorMap[Int32],
	}
	BasicTypeMap["enum"] = &DataType{
		Type:      "enum",
		ValueType: BasicTypeMap["enumVal"],
		Kind:      KindMap[Enum],
	}
}

func IsNumber(d *DataType) bool {
	return d.Kind.Kind == Int32 || d.Kind.Kind == Int64 ||
		d.Kind.Kind == UInt32 || d.Kind.Kind == UInt64 ||
		d.Kind.Kind == Float32 || d.Kind.Kind == Float64
}

func IsInteger(d *DataType) bool {
	return d.Kind.Kind == Int32 || d.Kind.Kind == Int64 ||
		d.Kind.Kind == UInt32 || d.Kind.Kind == UInt64
}

func IsFloat(d *DataType) bool {
	return d.Kind.Kind == Float32 || d.Kind.Kind == Float64
}

var ConstructorMap = map[int]Constructor{
	UInt8:   NewUInt8Value,
	Int32:   NewInt32Value,
	Int64:   NewInt64Value,
	UInt32:  NewUInt32Value,
	UInt64:  NewUInt64Value,
	Float32: NewFloat32Value,
	Float64: NewFloat64Value,
	Bool:    NewBoolValue,
	String:  NewStringValue,
	Bytes:   NewBytesValue,
	Nil:     NewNilValue,
	Message: NewMessageValue,
	Slice:   NewSliceValue,
	Object:  NewObjectValue,
}

var MapConstructorMap = map[int]Constructor{
	Int32:   NewMapInt32Value,
	Int64:   NewMapInt64Value,
	UInt32:  NewMapUInt32Value,
	UInt64:  NewMapUInt64Value,
	Float32: NewMapFloat32Value,
	Float64: NewMapFloat64Value,
	Bool:    NewMapBoolValue,
	String:  NewMapStringValue,
}

const ( // BasicTypeString
	UInt8Type   = "uint8"
	UInt32Type  = "uint32"
	UInt64Type  = "uint64"
	Int32Type   = "int32"
	Int64Type   = "int64"
	Float32Type = "float32"
	Float64Type = "float64"
	BoolType    = "bool"
	BytesType   = "bytes"
	StringType  = "string"
	SliceType   = "slice"
	MapType     = "map"
	MessageType = "message"
	ObjectType  = "object"
)
