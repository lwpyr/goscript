package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
	"strconv"
)

// EnterTypeDefMap is called when production TypeDefMap is entered.
func (s *TypeRegister) EnterTypeDefAlias(ctx *parser.TypeDefAliasContext) {
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(ctx.Name().GetText())
}

// EnterTypeDefMessage is called when production TypeDefMessage is entered.
func (s *TypeRegister) EnterTypeDefMessage(ctx *parser.TypeDefMessageContext) {
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(ctx.Name().GetText())
}

// EnterTypeDefEnum is called when production TypeDefEnum is entered.
func (s *TypeRegister) EnterTypeDefEnum(ctx *parser.TypeDefEnumContext) {
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(ctx.Name(0).GetText())
}

// ExitTypeDefAlias is called when production TypeDefAlias is exited.
func (s *TypeParser) ExitTypeDefAlias(ctx *parser.TypeDefAliasContext) {
	node := s.NodeTop()
	dt := s.Compiler.FindType(ctx.Name().GetText())
	node.SetDataType(dt)
}

// EnterTypeDefMap is called when production TypeDefMap is entered.
func (s *TypeParser) EnterMapTypeNameindef(ctx *parser.MapTypeNameindefContext) {
	s.VisitPush(&MapTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "MapTypeDef",
			Variadic: false,
		},
		KeyType: common.BasicTypeMap[ctx.BasicTypeName().GetText()],
	})
}

// ExitTypeDefMap is called when production TypeDefMap is exited.
func (s *TypeParser) ExitMapTypeNameindef(ctx *parser.MapTypeNameindefContext) {
	node := s.VisitPop().(*MapTypeDef)
	node.Value = s.NodePop()
	s.NodePush(node)
}

// EnterTypeDefSlice is called when production TypeDefSlice is entered.
func (s *TypeParser) EnterSliceTypeNameindef(ctx *parser.SliceTypeNameindefContext) {
	s.VisitPush(&SliceTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "SliceTypeDef",
			Variadic: false,
		},
	})
}

// ExitTypeDefSlice is called when production TypeDefSlice is exited.
func (s *TypeParser) ExitSliceTypeNameindef(ctx *parser.SliceTypeNameindefContext) {
	node := s.VisitPop().(*SliceTypeDef)
	node.Item = s.NodePop()
	s.NodePush(node)
}

// EnterSimpleTypeNameInDef is called when production SimpleTypeNameInDef is entered.
func (s *TypeParser) EnterSimpleTypeNameindef(ctx *parser.SimpleTypeNameindefContext) {
	typeName := ctx.GetText()
	cur := &TerminalTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "TerminalTypeDef",
			Variadic: false,
			DataType: s.Compiler.FindType(typeName),
		},
		Name: typeName,
	}
	s.NodePush(cur)
}

// EnterTypeDefFunction is called when production TypeDefFunction is entered.
func (s *TypeParser) EnterFunctionTypeNameindef(ctx *parser.FunctionTypeNameindefContext) {
	s.VisitPush(&FunctionTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "FunctionTypeDef",
			Variadic: false,
		},
	})
}

// ExitIntypenameindef is called when production intypenameindef is exited.
func (s *TypeParser) ExitIntypenameindef(ctx *parser.IntypenameindefContext) {
	funcNode := s.VisitTop().(*FunctionTypeDef)
	funcNode.InType = append(funcNode.InType, s.NodePop())
}

// ExitReturntypenameindef is called when production returntypenameindef is exited.
func (s *TypeParser) ExitReturntypenameindef(ctx *parser.ReturntypenameindefContext) {
	funcNode := s.VisitTop().(*FunctionTypeDef)
	funcNode.OutType = append(funcNode.OutType, s.NodePop())
}

// ExitTypeDefFunction is called when production TypeDefFunction is exited.
func (s *TypeParser) ExitFunctionTypeNameindef(ctx *parser.FunctionTypeNameindefContext) {
	s.NodePush(s.VisitPop())
}

// EnterChanTypeNameInDef is called when production ChanTypeNameInDef is entered.
func (s *TypeParser) EnterChanTypeNameindef(ctx *parser.ChanTypeNameindefContext) {
	s.VisitPush(&ChanTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "ChanTypeDef",
			Variadic: false,
		},
	})
}

// ExitChanTypeNameInDef is called when production ChanTypeNameInDef is exited.
func (s *TypeParser) ExitChanTypeNameindef(ctx *parser.ChanTypeNameindefContext) {
	node := s.VisitPop().(*ChanTypeDef)
	node.Item = s.NodePop()
	s.NodePush(node)
}

// EnterTypeDefMessage is called when production TypeDefMessage is entered.
func (s *TypeParser) EnterTypeDefMessage(ctx *parser.TypeDefMessageContext) {
	typeName := ctx.Name().GetText()
	s.VisitPush(&MessageTypeDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "MessageTypeDef",
			DataType: s.Compiler.FindType(typeName),
			Variadic: false,
		},
		Field:      map[string]ASTNode{},
		OneOfGroup: map[string][]string{},
		Name:       ctx.Name().GetText(),
	})
}

// ExitTypeDefMessage is called when production TypeDefMessage is exited.
func (s *TypeParser) ExitTypeDefMessage(ctx *parser.TypeDefMessageContext) {
	node := s.VisitPop().(*MessageTypeDef)
	for i := 0; i < len(ctx.AllMessagefield()); i++ {
		fieldNode := s.NodePop()
		switch fieldNode.(type) {
		case *MessageFieldDef:
			messageFieldNode := fieldNode.(*MessageFieldDef)
			if node.Field[messageFieldNode.Name] != nil {
				panic(common.NewCompileErr("duplicate field name in message definition " + node.Name))
			}
			node.Field[messageFieldNode.Name] = messageFieldNode.Type
		case *OneofFieldDef:
			oneofFieldNode := fieldNode.(*OneofFieldDef)
			oneofName := oneofFieldNode.Name
			if _, ok := node.OneOfGroup[oneofName]; ok {
				panic(common.NewCompileErr("duplicate oneof group name in message definition " + node.Name))
			}
			node.OneOfGroup[oneofName] = []string{}
			for _, choice := range oneofFieldNode.Choices {
				if node.Field[choice.Name] != nil {
					panic(common.NewCompileErr("duplicate field name in message definition " + node.Name))
				}
				node.Field[choice.Name] = choice.Type
				node.OneOfGroup[oneofName] = append(node.OneOfGroup[oneofName], choice.Name)
			}
		}

	}
	s.NodePush(node)
}

// EnterTypeDefEnum is called when production TypeDefEnum is entered.
func (s *TypeParser) EnterTypeDefEnum(ctx *parser.TypeDefEnumContext) {
	enumName := ctx.Name(0).GetText()
	enum := map[string]int32{}
	num := len(ctx.AllINT())
	for i := 0; i < num; i++ {
		v, _ := strconv.ParseInt(ctx.INT(i).GetText(), 10, 32)
		enum[ctx.Name(i+1).GetText()] = int32(v)
	}

	s.NodePush(&EnumNode{
		Name: enumName,
		Enum: enum,
	})
}

// EnterMessagefield is called when production messagefield is entered.
func (s *TypeParser) EnterFieldDef(ctx *parser.FieldDefContext) {
	s.VisitPush(&MessageFieldDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "MessageFieldDef",
			Variadic: false,
		},
		Name: ctx.Fieldname().GetText(),
	})
}

// ExitMessagefield is called when production messagefield is exited.
func (s *TypeParser) ExitFieldDef(ctx *parser.FieldDefContext) {
	node := s.VisitPop().(*MessageFieldDef)
	node.Type = s.NodePop()
	s.NodePush(node)
}

// EnterOneofDef is called when production OneofDef is entered.
func (s *TypeParser) EnterOneofDef(ctx *parser.OneofDefContext) {
	s.VisitPush(&OneofFieldDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "OneofFieldDef",
			Variadic: false,
		},
		Name: ctx.Fieldname().GetText(),
	})
}

// ExitOneofDef is called when production OneofDef is exited.
func (s *TypeParser) ExitOneofDef(ctx *parser.OneofDefContext) {
	node := s.VisitPop().(*OneofFieldDef)
	numChoice := len(ctx.AllOneoffield())
	for i := 0; i < numChoice; i++ {
		node.Choices = append(node.Choices, s.NodePop().(*MessageFieldDef))
	}
	s.NodePush(node)
}

// EnterOneoffield is called when production oneoffield is entered.
func (s *TypeParser) EnterOneoffield(ctx *parser.OneoffieldContext) {
	s.VisitPush(&MessageFieldDef{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "OneofField",
			Variadic: false,
		},
		Name: ctx.Fieldname().GetText(),
	})
}

// ExitOneoffield is called when production oneoffield is exited.
func (s *TypeParser) ExitOneoffield(ctx *parser.OneoffieldContext) {
	node := s.VisitPop().(*MessageFieldDef)
	node.Type = s.NodePop()
	s.NodePush(node)
}
