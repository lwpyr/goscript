package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
	"strconv"
)

// EnterTypeDefMap is called when production TypeDefMap is entered.
func (s *TypeRegister) EnterTypeDefMap(ctx *parser.TypeDefMapContext) {
	typeName := ctx.NAME().GetText()
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(typeName)
}

// EnterTypeDefSlice is called when production TypeDefSlice is entered.
func (s *TypeRegister) EnterTypeDefSlice(ctx *parser.TypeDefSliceContext) {
	typeName := ctx.NAME().GetText()
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(typeName)
}

// EnterTypeDefMessage is called when production TypeDefMessage is entered.
func (s *TypeRegister) EnterTypeDefMessage(ctx *parser.TypeDefMessageContext) {
	typeName := ctx.NAME().GetText()
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(typeName)
}

// EnterTypeDefEnum is called when production TypeDefEnum is entered.
func (s *TypeRegister) EnterTypeDefEnum(ctx *parser.TypeDefEnumContext) {
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(ctx.NAME(0).GetText())
}

// EnterTypeDefFunction is called when production TypeDefFunction is entered.
func (s *TypeRegister) EnterTypeDefFunction(ctx *parser.TypeDefFunctionContext) {
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(ctx.NAME().GetText())
}

// EnterTypeDefMap is called when production TypeDefMap is entered.
func (s *TypeParser) EnterTypeDefMap(ctx *parser.TypeDefMapContext) {
	typeName := ctx.NAME().GetText()
	s.VisitPush(&MapTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "MapTypeDef",
			DataType: s.Compiler.FindType(typeName),
			Variadic: false,
		},
		TypeDefName: TypeDefName{Name: typeName},
		KeyType:     common.BasicTypeMap[ctx.BasicTypeName().GetText()],
	})
}

// ExitTypeDefMap is called when production TypeDefMap is exited.
func (s *TypeParser) ExitTypeDefMap(ctx *parser.TypeDefMapContext) {
	node := s.VisitPop().(*MapTypeDef)
	node.Value = s.NodePop().(ITypeDefNode)
	s.NodePush(node)
}

// EnterTypeDefSlice is called when production TypeDefSlice is entered.
func (s *TypeParser) EnterTypeDefSlice(ctx *parser.TypeDefSliceContext) {
	typeName := ctx.NAME().GetText()
	s.VisitPush(&SliceTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "SliceTypeDef",
			DataType: s.Compiler.FindType(typeName),
			Variadic: false,
		},
		TypeDefName: TypeDefName{Name: typeName},
	})
}

// ExitTypeDefSlice is called when production TypeDefSlice is exited.
func (s *TypeParser) ExitTypeDefSlice(ctx *parser.TypeDefSliceContext) {
	node := s.VisitPop().(*SliceTypeDef)
	node.Item = s.NodePop().(ITypeDefNode)
	s.NodePush(node)
}

// EnterTypeDefMessage is called when production TypeDefMessage is entered.
func (s *TypeParser) EnterTypeDefMessage(ctx *parser.TypeDefMessageContext) {
	typeName := ctx.NAME().GetText()
	s.VisitPush(&MessageTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "MessageTypeDef",
			DataType: s.Compiler.FindType(typeName),
			Variadic: false,
		},
		TypeDefName: TypeDefName{Name: ctx.NAME().GetText()},
		Field:       map[string]ITypeDefNode{},
	})
}

// ExitTypeDefMessage is called when production TypeDefMessage is exited.
func (s *TypeParser) ExitTypeDefMessage(ctx *parser.TypeDefMessageContext) {
	node := s.VisitPop().(*MessageTypeDef)
	for i := 0; i < len(ctx.AllMessagefield()); i++ {
		fieldNode := s.NodePop().(*MessageFieldDef)
		if node.Field[fieldNode.Name] != nil {
			panic(common.NewCompileErr("duplicate field name in message definition " + node.GetTypeDefName()))
		}
		node.Field[fieldNode.Name] = fieldNode.Type
	}
	s.NodePush(node)
}

// EnterTypeDefEnum is called when production TypeDefEnum is entered.
func (s *TypeParser) EnterTypeDefEnum(ctx *parser.TypeDefEnumContext) {
	enumName := ctx.NAME(0).GetText()
	enum := map[string]int32{}
	num := len(ctx.AllINT())
	for i := 0; i < num; i++ {
		v, _ := strconv.ParseInt(ctx.INT(i).GetText(), 10, 32)
		enum[ctx.NAME(i+1).GetText()] = int32(v)
	}

	s.NodePush(&EnumNode{
		TypeDefName: TypeDefName{Name: enumName},
		Enum:        enum,
	})
}

// EnterTypeDefFunction is called when production TypeDefFunction is entered.
func (s *TypeParser) EnterTypeDefFunction(ctx *parser.TypeDefFunctionContext) {
	s.VisitPush(&FunctionTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "FunctionTypeDef",
			DataType: s.Compiler.FindType(ctx.NAME().GetText()),
			Variadic: false,
		},
		TypeDefName: TypeDefName{Name: ctx.NAME().GetText()},
	})
}

// ExitTypeDefFunction is called when production TypeDefFunction is exited.
func (s *TypeParser) ExitTypeDefFunction(ctx *parser.TypeDefFunctionContext) {
	s.NodePush(s.VisitPop())
}

// EnterMessagefield is called when production messagefield is entered.
func (s *TypeParser) EnterMessagefield(ctx *parser.MessagefieldContext) {
	s.VisitPush(&MessageFieldDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "MessageFieldDef",
			Variadic: false,
		},
		Name: ctx.NAME().GetText(),
	})
}

// ExitMessagefield is called when production messagefield is exited.
func (s *TypeParser) ExitMessagefield(ctx *parser.MessagefieldContext) {
	node := s.VisitPop().(*MessageFieldDef)
	node.Type = s.NodePop().(ITypeDefNode)
	s.NodePush(node)
}

// EnterSimpleTypeNameInDef is called when production SimpleTypeNameInDef is entered.
func (s *TypeParser) EnterSimpleTypeNameInDef(ctx *parser.SimpleTypeNameInDefContext) {
	typeName := ctx.GetText()
	cur := &TerminalTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "TerminalTypeDef",
			Variadic: false,
		},
		TypeDefName: TypeDefName{Name: typeName},
	}
	cur.DataType = s.Compiler.FindType(typeName)
	s.NodePush(cur)
}

// EnterFunctionTypeInDef is called when production FunctionTypeInDef is entered.
func (s *TypeParser) EnterFunctionTypeNameindef(ctx *parser.FunctionTypeNameindefContext) {
	s.VisitPush(&FunctionTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "FunctionTypeDef",
			Variadic: false,
		},
		TailArray: ctx.TAILARRAY() != nil,
	})
}

// ExitFunctionTypeInDef is called when production FunctionTypeInDef is exited.
func (s *TypeParser) ExitFunctionTypeNameindef(ctx *parser.FunctionTypeNameindefContext) {
	s.NodePush(s.VisitPop())
}

// EnterMapTypeNameInDef is called when production MapTypeNameInDef is entered.
func (s *TypeParser) EnterMapTypeNameInDef(ctx *parser.MapTypeNameInDefContext) {
	s.VisitPush(&MapTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "MapTypeDef",
			Variadic: false,
		},
		KeyType: common.BasicTypeMap[ctx.BasicTypeName().GetText()],
	})
}

// ExitMapTypeNameInDef is called when production MapTypeNameInDef is exited.
func (s *TypeParser) ExitMapTypeNameInDef(ctx *parser.MapTypeNameInDefContext) {
	node := s.VisitPop().(*MapTypeDef)
	node.Value = s.NodePop().(ITypeDefNode)
	s.NodePush(node)
}

// EnterChanTypeNameInDef is called when production ChanTypeNameInDef is entered.
func (s *TypeParser) EnterChanTypeNameInDef(ctx *parser.ChanTypeNameInDefContext) {
	s.VisitPush(&ChanTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "ChanTypeDef",
			Variadic: false,
		},
	})
}

// ExitChanTypeNameInDef is called when production ChanTypeNameInDef is exited.
func (s *TypeParser) ExitChanTypeNameInDef(ctx *parser.ChanTypeNameInDefContext) {
	node := s.VisitPop().(*ChanTypeDef)
	node.Item = s.NodePop().(ITypeDefNode)
	s.NodePush(node)
}

// EnterSliceTypeNameInDef is called when production SliceTypeNameInDef is entered.
func (s *TypeParser) EnterSliceTypeNameInDef(ctx *parser.SliceTypeNameInDefContext) {
	s.VisitPush(&SliceTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "SliceTypeDef",
			Variadic: false,
		},
	})
}

// ExitSliceTypeNameInDef is called when production SliceTypeNameInDef is exited.
func (s *TypeParser) ExitSliceTypeNameInDef(ctx *parser.SliceTypeNameInDefContext) {
	node := s.VisitPop().(*SliceTypeDef)
	node.Item = s.NodePop().(ITypeDefNode)
	s.NodePush(node)
}

// ExitIntypenameindef is called when production intypenameindef is exited.
func (s *TypeParser) ExitIntypenameindef(ctx *parser.IntypenameindefContext) {
	funcNode := s.VisitTop().(*FunctionTypeDef)
	funcNode.InType = append(funcNode.InType, s.NodePop().(ITypeDefNode))
}

// ExitReturntypenameindef is called when production returntypenameindef is exited.
func (s *TypeParser) ExitReturntypenameindef(ctx *parser.ReturntypenameindefContext) {
	funcNode := s.VisitTop().(*FunctionTypeDef)
	funcNode.OutType = append(funcNode.OutType, s.NodePop().(ITypeDefNode))
}
