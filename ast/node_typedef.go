package ast

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/lwpyr/goscript/common"
)

type ITypeDefNode interface {
	GetDataType() *common.DataType
	SetDataType(*common.DataType)
	Compile(c *Compiler)
	GetTypeDefName() string
}

type TypeDefName struct {
	Name string
}

func (b *TypeDefName) GetTypeDefName() string {
	return b.Name
}

type TerminalTypeDef struct {
	Node
	TypeDefName
}

func (t *TerminalTypeDef) Compile(c *Compiler) {
	if t.DataType == nil {
		t.DataType = c.TypeRegistry.FindType(t.GetTypeDefName())
	}
	if t.DataType == nil {
		panic(common.NewCompileErr("unknown type name " + t.GetTypeDefName()))
	}
}

type EnumNode struct {
	Node
	TypeDefName
	Enum map[string]int32
}

func (e *EnumNode) Compile(c *Compiler) {
	enumName := e.GetTypeDefName()
	dt := c.TypeRegistry.FindType(enumName)
	dt.Kind = common.KindMap[common.Int32]
	dt.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
		return iter.ReadInt32()
	}
	dt.Constructor = common.ConstructorMap[common.Int32]
	c.TypeRegistry.Enums[enumName] = e.Enum
	re := map[int32]string{}
	for k, v := range e.Enum {
		re[v] = k
	}
	c.TypeRegistry.REnums[enumName] = re
}

type ChanTypeDef struct {
	Node
	TypeDefName
	Item ITypeDefNode
}

func (t *ChanTypeDef) Compile(c *Compiler) {
	t.Item.Compile(c)
	ItemType := t.Item.GetDataType()
	dt := c.TypeRegistry.MakeChanType(ItemType.Type)
	if c.TypeRegistry.FindType(dt.Type) == nil {
		c.TypeRegistry.AddType(dt.Type, dt)
	}
	if t.DataType != nil {
		*t.DataType = *dt
	} else {
		t.DataType = dt
	}
}

type MapTypeDef struct {
	Node
	TypeDefName
	KeyType *common.DataType
	Value   ITypeDefNode
}

func (t *MapTypeDef) Compile(c *Compiler) {
	t.Value.Compile(c)
	ValType := t.Value.GetDataType()
	dt := c.TypeRegistry.MakeMapType(t.KeyType.Type, ValType.Type)
	if c.TypeRegistry.FindType(dt.Type) == nil {
		c.TypeRegistry.AddType(dt.Type, dt)
	}
	if t.DataType != nil {
		*t.DataType = *dt
	} else {
		t.DataType = dt
	}
}

type SliceTypeDef struct {
	Node
	TypeDefName
	Item ITypeDefNode
}

func (t *SliceTypeDef) Compile(c *Compiler) {
	t.Item.Compile(c)
	ItemType := t.Item.GetDataType()
	dt := c.TypeRegistry.MakeSliceType(ItemType.Type)
	if c.TypeRegistry.FindType(dt.Type) == nil {
		c.TypeRegistry.AddType(dt.Type, dt)
	}
	if t.DataType != nil {
		*t.DataType = *dt
	} else {
		t.DataType = dt
	}
}

type MessageTypeDef struct {
	Node
	TypeDefName
	Field map[string]ITypeDefNode
}

func (t *MessageTypeDef) Compile(c *Compiler) {
	dt := c.FindType(t.GetTypeDefName())
	dt.FieldType = map[string]*common.DataType{}
	for fieldName, field := range t.Field {
		field.Compile(c)
		dt.FieldType[fieldName] = field.GetDataType()
	}
	dt.Kind = common.KindMap[common.Message]
	dt.Constructor = common.ConstructorMap[common.Message]
	fieldInfos := dt.FieldType
	dt.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
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
}

type MessageFieldDef struct {
	Node
	Name string
	Type ITypeDefNode
}

func (t *MessageFieldDef) Compile(c *Compiler) {
	panic("this won't compile")
}

type FunctionTypeDef struct {
	Node
	TypeDefName
	InType    []ITypeDefNode
	OutType   []ITypeDefNode
	TailArray bool
}

func (t *FunctionTypeDef) Compile(c *Compiler) {
	meta := &common.FunctionMeta{
		TailArray: t.TailArray,
		ConstExpr: false,
	}
	for _, in := range t.InType {
		meta.In = append(meta.In, in.GetDataType())
	}
	for _, out := range t.OutType {
		meta.In = append(meta.In, out.GetDataType())
	}
	dt := c.FindFuncType(meta)
	if t.DataType != nil {
		*t.DataType = *dt
	} else {
		t.DataType = dt
	}
	if t.GetTypeDefName() != "" {
		c.AddType(t.GetTypeDefName(), dt)
	}
}
