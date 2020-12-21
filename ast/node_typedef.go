package ast

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/lwpyr/goscript/common"
)

type TypeDefNode struct {
	Node
	SimpleTypeName string
	TypeNodeType   TypeNodeType
	Key            *common.DataType // map
	Value          ASTNode          // map slice chan
	InParam        []ASTNode        // func
	OutParam       []ASTNode        // func
	TailArray      bool
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

func (t *TypeDefNode) Compile(c *CompileContext) {
	var dType *common.DataType
	switch t.TypeNodeType {
	case SimpleType:
		dType = c.FindType(t.SimpleTypeName)
	case SliceType:
		t.Value.Compile(c)
		ItemType := t.Value.GetDataType()
		dType = c.TypeRegistry.FindSliceType(ItemType.Type)
	case ChanType:
		t.Value.Compile(c)
		ItemType := t.Value.GetDataType()
		dType = c.TypeRegistry.FindChanType(ItemType.Type)
	case MapType:
		t.Value.Compile(c)
		ValType := t.Value.GetDataType()
		dType = c.TypeRegistry.FindMapType(t.Key.Type, ValType.Type)
	case FuncType:
		meta := &common.FunctionMeta{
			TailArray: t.TailArray,
			ConstExpr: false,
		}
		for _, in := range t.InParam {
			in.Compile(c)
			meta.In = append(meta.In, in.GetDataType())
		}
		for _, out := range t.OutParam {
			out.Compile(c)
			meta.Out = append(meta.Out, out.GetDataType())
		}
		dType = c.TypeRegistry.FindFuncType(meta)
	default:
		panic(common.NewCompileErr("bad node type, this should never happened because of the enum limitation"))
	}
	if t.DataType == nil {
		t.DataType = dType // not a definition, directly set
	} else {
		*t.DataType = *dType // make a alias
	}
}

func (e *EnumNode) Compile(c *CompileContext) {
	dt := e.DataType
	dt.Kind = common.KindMap[common.Int32]
	dt.Unmarshal = func(iter *jsoniter.Iterator) interface{} {
		return iter.ReadInt32()
	}
	dt.Constructor = common.ConstructorMap[common.Int32]
}

func (t *MessageTypeDef) Compile(c *CompileContext) {
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

type OneofFieldDef struct {
	Node
	Name    string
	Choices []*MessageFieldDef
}
