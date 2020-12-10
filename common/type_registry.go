package common

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

type TypeRegistry struct {
	Types        map[string]*DataType
	Enums        map[string]map[string]int32
	REnums       map[string]map[int32]string
	TypesInBuild map[string]*DataType
}

func (t *TypeRegistry) AddTypePlaceHolderInBuild(name string) *DataType {
	if _, ok := BasicTypeMap[name]; ok {
		panic("try to overwritten a basic type: " + name)
	}
	if t.Types[name] != nil {
		panic(NewCompileErr("try to create duplicate type name " + name))
	}
	placeHolder := &DataType{
		Type: name,
	}
	t.Types[name] = placeHolder
	return placeHolder
}

func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		Types:  map[string]*DataType{},
		Enums:  map[string]map[string]int32{},
		REnums: map[string]map[int32]string{},
	}
}

func (t *TypeRegistry) AddEnum(name string, e map[string]int32) {
	if t.FindType(name) != nil {
		return
	}
	t.Types[name] = &DataType{
		Type: name,
		Kind: KindMap[Int32],
		Unmarshal: func(iter *jsoniter.Iterator) interface{} {
			return iter.ReadInt32()
		},
		Constructor: ConstructorMap[Int32],
	}
	t.Enums[name] = e
	re := map[int32]string{}
	for k, v := range e {
		re[v] = k
	}
	t.REnums[name] = re
}

func (t *TypeRegistry) GetEnums(name string) interface{} {
	if ret, ok := t.Enums[name]; ok {
		return ret
	}
	return nil
}

func (t *TypeRegistry) GetREnums(name string) map[int32]string {
	if ret, ok := t.REnums[name]; ok {
		return ret
	}
	return nil
}
func (t *TypeRegistry) FindFuncType(meta *FunctionMeta) *DataType {
	typeName := meta.GenerateTypeName()
	if _, ok := t.Types[typeName]; !ok {
		dType := &DataType{
			Type:       typeName,
			Kind:       KindMap[Func],
			LambdaMeta: meta,
		}
		t.Types[typeName] = dType
	}
	return t.Types[typeName]
}

func (t *TypeRegistry) FindSliceType(name string) *DataType {
	if t.FindType(name) == nil {
		return nil
	}
	sliceTypeName := "slice<" + name + ">"
	if _, ok := t.Types[sliceTypeName]; !ok {
		itemType := t.FindType(name)
		dtype := &DataType{
			Type:        sliceTypeName,
			Kind:        KindMap[Slice],
			ItemType:    itemType,
			Constructor: ConstructorMap[Slice],
		}
		dtype.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
			var ret []interface{}
			for iter.ReadArray() {
				ret = append(ret, itemType.Unmarshal(iter))
			}
			return ret
		}
		t.Types[sliceTypeName] = dtype
	}
	return t.Types[sliceTypeName]
}

func (t *TypeRegistry) FindMapType(key, val string) *DataType {
	mapTypeName := "map<" + key + "," + val + ">"
	if t.FindType(key) == nil || t.FindType(val) == nil {
		return nil
	}
	if _, ok := t.Types[mapTypeName]; !ok {
		t.Types[mapTypeName] = t.MakeMapType(key, val)

	}
	return t.Types[mapTypeName]
}

func (t *TypeRegistry) MakeMapType(key, val string) *DataType {
	mapTypeName := "map<" + key + "," + val + ">"
	keyType := t.FindType(key)
	valType := t.FindType(val)
	var constructor Constructor
	var unmarshal func(iter *jsoniter.Iterator) interface{}
	switch keyType.Kind.Kind {
	case Int32:
		constructor = NewMapInt32Value
		unmarshal = func(iter *jsoniter.Iterator) interface{} {
			ret := map[int32]interface{}{}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				k, _ := strconv.ParseInt(field, 10, 32)
				ret[int32(k)] = valType.Unmarshal(iter)
			}
			return ret
		}
	case Int64:
		constructor = NewMapInt64Value
		unmarshal = func(iter *jsoniter.Iterator) interface{} {
			ret := map[int64]interface{}{}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				k, _ := strconv.ParseInt(field, 10, 32)
				ret[k] = valType.Unmarshal(iter)
			}
			return ret
		}
	case UInt32:
		constructor = NewMapUInt32Value
		unmarshal = func(iter *jsoniter.Iterator) interface{} {
			ret := map[uint32]interface{}{}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				k, _ := strconv.ParseUint(field, 10, 32)
				ret[uint32(k)] = valType.Unmarshal(iter)
			}
			return ret
		}
	case UInt64:
		constructor = NewMapUInt64Value
		unmarshal = func(iter *jsoniter.Iterator) interface{} {
			ret := map[uint64]interface{}{}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				k, _ := strconv.ParseUint(field, 10, 32)
				ret[uint64(k)] = valType.Unmarshal(iter)
			}
			return ret
		}
	case Float32:
		constructor = NewMapFloat32Value
		unmarshal = func(iter *jsoniter.Iterator) interface{} {
			ret := map[float32]interface{}{}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				k, _ := strconv.ParseFloat(field, 32)
				ret[float32(k)] = valType.Unmarshal(iter)
			}
			return ret
		}
	case Float64:
		constructor = NewMapFloat64Value
		unmarshal = func(iter *jsoniter.Iterator) interface{} {
			ret := map[float64]interface{}{}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				k, _ := strconv.ParseFloat(field, 64)
				ret[k] = valType.Unmarshal(iter)
			}
			return ret
		}
	case String:
		constructor = NewMapStringValue
		unmarshal = func(iter *jsoniter.Iterator) interface{} {
			ret := map[string]interface{}{}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				ret[field] = valType.Unmarshal(iter)
			}
			return ret
		}
	case Bool:
		constructor = NewMapBoolValue
		unmarshal = func(iter *jsoniter.Iterator) interface{} {
			ret := map[bool]interface{}{}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				k, _ := strconv.ParseBool(field)
				ret[k] = valType.Unmarshal(iter)
			}
			return ret
		}
	}
	return &DataType{
		Type:        mapTypeName,
		Kind:        KindMap[Map],
		KeyType:     keyType,
		ValueType:   valType,
		Constructor: constructor,
		Unmarshal:   unmarshal,
	}
}

// not suggested to use, use DataTypeBuilder instead
func (t *TypeRegistry) AddType(name string, dType *DataType) {
	if _, ok := BasicTypeMap[name]; ok {
		panic("try to overwritten a basic type: " + name)
	}
	if _, ok := t.Types[name]; ok {
		fmt.Printf("Datatype name duplicated [%s], previous is overwrittend. old: %v, new %v\n", name, t.Types[name], dType)
	}
	t.Types[name] = dType
}

func (t *TypeRegistry) FindType(name string) *DataType {
	if ret, ok := BasicTypeMap[name]; ok {
		return ret
	}
	if ret, ok := t.Types[name]; ok {
		return ret
	}
	return nil
}

func (t *TypeRegistry) AddBuiltType(dtb *DataTypeBuilder) {
	if _, ok := BasicTypeMap[dtb.T.Type]; ok {
		panic("try to overwritten a basic type: " + dtb.T.Type)
	}
	typeName := dtb.T.Type // this may be alias
	if dtb.T.Kind.Kind == Slice {
		dtb.T.Type = "slice<" + dtb.T.ItemType.Type + ">"
		itemType := dtb.T.ItemType
		dtb.T.Constructor = ConstructorMap[Slice]
		dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
			var ret []interface{}
			for iter.ReadArray() {
				ret = append(ret, itemType.Unmarshal(iter))
			}
			return ret
		}
	} else if dtb.T.Kind.Kind == Map {
		dtb.T.Type = "map<" + dtb.T.KeyType.Type + "," + dtb.T.ValueType.Type + ">"
		valUnmarshal := dtb.T.ValueType.Unmarshal
		switch dtb.T.KeyType.Kind.Kind {
		case Int32:
			dtb.T.Constructor = NewMapInt32Value
			dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
				ret := map[int32]interface{}{}
				for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
					k, _ := strconv.ParseInt(field, 10, 32)
					ret[int32(k)] = valUnmarshal(iter)
				}
				return ret
			}
		case Int64:
			dtb.T.Constructor = NewMapInt64Value
			dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
				ret := map[int64]interface{}{}
				for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
					k, _ := strconv.ParseInt(field, 10, 32)
					ret[k] = valUnmarshal(iter)
				}
				return ret
			}
		case UInt32:
			dtb.T.Constructor = NewMapUInt32Value
			dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
				ret := map[uint32]interface{}{}
				for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
					k, _ := strconv.ParseUint(field, 10, 32)
					ret[uint32(k)] = valUnmarshal(iter)
				}
				return ret
			}
		case UInt64:
			dtb.T.Constructor = NewMapUInt64Value
			dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
				ret := map[uint64]interface{}{}
				for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
					k, _ := strconv.ParseUint(field, 10, 32)
					ret[k] = valUnmarshal(iter)
				}
				return ret
			}
		case Float32:
			dtb.T.Constructor = NewMapFloat32Value
			dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
				ret := map[float32]interface{}{}
				for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
					k, _ := strconv.ParseFloat(field, 32)
					ret[float32(k)] = valUnmarshal(iter)
				}
				return ret
			}
		case Float64:
			dtb.T.Constructor = NewMapFloat64Value
			dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
				ret := map[float64]interface{}{}
				for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
					k, _ := strconv.ParseFloat(field, 64)
					ret[k] = valUnmarshal(iter)
				}
				return ret
			}
		case String:
			dtb.T.Constructor = NewMapStringValue
			dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
				ret := map[string]interface{}{}
				for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
					ret[field] = valUnmarshal(iter)
				}
				return ret
			}
		case Bool:
			dtb.T.Constructor = NewMapBoolValue
			dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
				ret := map[bool]interface{}{}
				for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
					k, _ := strconv.ParseBool(field)
					ret[k] = valUnmarshal(iter)
				}
				return ret
			}
		}
	} else if dtb.T.Kind.Kind == Message {
		fieldInfos := dtb.T.FieldType
		dtb.T.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
			ret := map[string]interface{}{}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				if dType, ok := fieldInfos[field]; ok {
					ret[field] = dType.Unmarshal(iter)
				} else {
					iter.Skip()
				}
			}
			return ret
		}
		dtb.T.Constructor = ConstructorMap[Message]
	} else if dtb.T.Kind.Kind == Object {
		dtb.T.Unmarshal = nil
		dtb.T.Constructor = ConstructorMap[Object]
	}
	if ret, ok := t.Types[typeName]; ok {
		if !(len(typeName) >= 6 && typeName[:6] == "slice<" || len(typeName) >= 4 && typeName[:4] == "map<") {
			fmt.Printf("Datatype name duplicated [%s], previous is overwrittend. old: %v, new %v\n", typeName, ret, dtb.T)
			t.Types[typeName] = dtb.T
			t.Types[dtb.T.Type] = dtb.T
		}
	} else {
		t.Types[typeName] = dtb.T
		t.Types[dtb.T.Type] = dtb.T
	}
}
