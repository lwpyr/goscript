package ast

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/lwpyr/goscript/common"
)

type ChanTypeDef struct {
	Node
	Item ASTNode
}

type MapTypeDef struct {
	Node
	KeyType *common.DataType
	Value   ASTNode
}

type SliceTypeDef struct {
	Node
	Item ASTNode
}

type FunctionTypeDef struct {
	Node
	InType    []ASTNode
	OutType   []ASTNode
	TailArray bool
}

type MessageTypeDef struct {
	Node
	Name       string
	Field      map[string]ASTNode
	OneOfGroup map[string][]string
}

type EnumNode struct {
	Node
	Name string
	Enum map[string]int32
}

type TerminalTypeDef struct {
	Node
	Name string
}

func (t *TerminalTypeDef) Compile(c *Compiler) {
	if t.DataType == nil {
		t.DataType = c.TypeRegistry.FindType(t.Name)
	}
	if t.DataType == nil {
		panic(common.NewCompileErr("unknown type name " + t.Name))
	}
}

func (e *EnumNode) Compile(c *Compiler) {
	enumName := e.Name
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

func (t *ChanTypeDef) Compile(c *Compiler) {
	t.Item.Compile(c)
	ItemType := t.Item.GetDataType()
	dt := c.TypeRegistry.FindChanType(ItemType.Type)
	if t.DataType != nil {
		*t.DataType = *dt
	} else {
		t.DataType = dt
	}
}

func (t *MapTypeDef) Compile(c *Compiler) {
	t.Value.Compile(c)
	ValType := t.Value.GetDataType()
	dt := c.TypeRegistry.FindMapType(t.KeyType.Type, ValType.Type)
	if t.DataType != nil {
		*t.DataType = *dt
	} else {
		t.DataType = dt
	}
}

func (t *SliceTypeDef) Compile(c *Compiler) {
	t.Item.Compile(c)
	ItemType := t.Item.GetDataType()
	dt := c.TypeRegistry.FindSliceType(ItemType.Type)
	if c.TypeRegistry.FindType(dt.Type) == nil {
		c.TypeRegistry.AddType(dt.Type, dt)
	}
	if t.DataType != nil {
		*t.DataType = *dt
	} else {
		t.DataType = dt
	}
}

func (t *MessageTypeDef) Compile(c *Compiler) {
	dt := t.DataType
	dt.FieldType = map[string]*common.DataType{}
	for fieldName, field := range t.Field {
		field.Compile(c)
		dt.FieldType[fieldName] = field.GetDataType()
	}
	dt.OneOfGroup = t.OneOfGroup
	for oneOfName, group := range dt.OneOfGroup {
		if dt.OneOf == nil {
			dt.OneOf = map[string]string{}
		}
		for _, choice := range group {
			dt.OneOf[choice] = oneOfName
		}
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
	Type ASTNode
}

func (t *MessageFieldDef) Compile(c *Compiler) {
	panic("this won't compile")
}

type OneofFieldDef struct {
	Node
	Name    string
	Choices []*MessageFieldDef
}

func (t *OneofFieldDef) Compile(c *Compiler) {
	panic("this won't compile")
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
		meta.Out = append(meta.Out, out.GetDataType())
	}
	dt := c.TypeRegistry.FindFuncType(meta)
	if t.DataType != nil {
		*t.DataType = *dt
	} else {
		t.DataType = dt
	}
}
