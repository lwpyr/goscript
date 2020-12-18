package parser_goscript

import (
	"github.com/lwpyr/goscript/ast"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
	"strconv"
)

func (s *ASTBuilder) ExitTypeDefAlias(ctx *parser.TypeDefAliasContext) {
	node := s.NodeTop()
	alias := ctx.Name().GetText()
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(alias)
	dt := s.Compiler.FindType(alias)
	node.SetDataType(dt)
}

func (s *ASTBuilder) EnterTypeDefEnum(ctx *parser.TypeDefEnumContext) {
	enumName := ctx.Name(0).GetText()
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(enumName)
	enum := map[string]int32{}
	num := len(ctx.AllINT())
	for i := 0; i < num; i++ {
		v, _ := strconv.ParseInt(ctx.INT(i).GetText(), 10, 32)
		enum[ctx.Name(i+1).GetText()] = int32(v)
	}
	re := map[int32]string{}
	for k, v := range enum {
		re[v] = k
	}
	node := &ast.EnumNode{
		Name: enumName,
		Enum: enum,
	}
	s.Compiler.TypeRegistry.Enums[enumName] = enum
	s.Compiler.TypeRegistry.REnums[enumName] = re
	node.DataType = s.Compiler.FindType(enumName)
	s.NodePush(node)
}

func (s *ASTBuilder) ExitTypeDefMessage(ctx *parser.TypeDefMessageContext) {
	typeName := ctx.Name().GetText()
	s.Compiler.TypeRegistry.AddTypePlaceHolderInBuild(typeName)
	node := &ast.MessageTypeDef{
		Field:      map[string]ast.ASTNode{},
		OneOfGroup: map[string][]string{},
		Name:       ctx.Name().GetText(),
	}
	for i := 0; i < len(ctx.AllMessagefield()); i++ {
		fieldNode := s.NodePop()
		switch fieldNode.(type) {
		case *ast.MessageFieldDef:
			messageFieldNode := fieldNode.(*ast.MessageFieldDef)
			if node.Field[messageFieldNode.Name] != nil {
				panic(common.NewCompileErr(common.FormatError(ctx, "duplicate field name in message definition ")))
			}
			node.Field[messageFieldNode.Name] = messageFieldNode.Type
		case *ast.OneofFieldDef:
			oneofFieldNode := fieldNode.(*ast.OneofFieldDef)
			oneofName := oneofFieldNode.Name
			if _, ok := node.OneOfGroup[oneofName]; ok {
				panic(common.NewCompileErr(common.FormatError(ctx, "duplicate oneof group name in message definition")))
			}
			node.OneOfGroup[oneofName] = []string{}
			for _, choice := range oneofFieldNode.Choices {
				if node.Field[choice.Name] != nil {
					panic(common.NewCompileErr(common.FormatError(ctx, "duplicate field name in message definition ")))
				}
				node.Field[choice.Name] = choice.Type
				node.OneOfGroup[oneofName] = append(node.OneOfGroup[oneofName], choice.Name)
			}
		}

	}
	node.DataType = s.Compiler.FindType(typeName)
	s.NodePush(node)
}

func (s *ASTBuilder) ExitFieldDef(ctx *parser.FieldDefContext) {
	node := &ast.MessageFieldDef{
		Name: ctx.Fieldname().GetText(),
	}
	node.Type = s.NodePop()
	s.NodePush(node)
}

func (s *ASTBuilder) ExitOneofDef(ctx *parser.OneofDefContext) {
	node := &ast.OneofFieldDef{
		Name: ctx.Fieldname().GetText(),
	}
	numChoice := len(ctx.AllOneoffield())
	for i := 0; i < numChoice; i++ {
		node.Choices = append(node.Choices, s.NodePop().(*ast.MessageFieldDef))
	}
	s.NodePush(node)
}

func (s *ASTBuilder) ExitOneoffield(ctx *parser.OneoffieldContext) {
	node := &ast.MessageFieldDef{
		Name: ctx.Fieldname().GetText(),
	}
	node.Type = s.NodePop()
	s.NodePush(node)
}

func (s *ASTBuilder) ExitMapTypeNameindef(ctx *parser.MapTypeNameindefContext) {
	node := &ast.TypeDefNode{
		TypeNodeType: ast.MapType,
		Key:          common.BasicTypeMap[ctx.BasicTypeName().GetText()],
		Value:        s.NodePop(),
	}
	s.NodePush(node)
}

func (s *ASTBuilder) ExitSliceTypeNameindef(_ *parser.SliceTypeNameindefContext) {
	s.NodePush(&ast.TypeDefNode{
		TypeNodeType: ast.SliceType,
		Value:        s.NodePop(),
	})
}

func (s *ASTBuilder) ExitSimpleTypeNameindef(ctx *parser.SimpleTypeNameindefContext) {
	typeName := ctx.GetText()
	cur := &ast.TypeDefNode{
		TypeNodeType:   ast.SimpleType,
		SimpleTypeName: typeName,
	}
	s.NodePush(cur)
}

func (s *ASTBuilder) ExitFunctionTypeNameindef(ctx *parser.FunctionTypeNameindefContext) {
	node := &ast.TypeDefNode{
		TypeNodeType: ast.FuncType,
	}
	lenIn := len(ctx.AllIntypenameindef())
	lenOut := len(ctx.AllReturntypenameindef())
	node.InParam = make([]ast.ASTNode, lenIn)
	node.OutParam = make([]ast.ASTNode, lenOut)
	for i := 0; i < lenOut; i++ {
		node.OutParam[lenOut-1-i] = s.NodePop()
	}
	for i := 0; i < lenIn; i++ {
		node.InParam[lenIn-1-i] = s.NodePop()
	}
	if ctx.TAILARRAY() != nil {
		node.TailArray = true
	}
	s.NodePush(node)
}

func (s *ASTBuilder) ExitChanTypeNameindef(_ *parser.ChanTypeNameindefContext) {
	node := &ast.TypeDefNode{
		TypeNodeType: ast.ChanType,
		Value:        s.NodePop(),
	}
	s.NodePush(node)
}
