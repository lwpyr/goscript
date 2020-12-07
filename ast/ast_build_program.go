package ast

import (
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
	"strconv"
	"strings"
)

// EnterProgram is called when production program is entered.
func (s *ASTBuilder) EnterProgram(ctx *parser.ProgramContext) {
	cur := &ProgramRoot{
		Node: Node{
			Parent:   nil,
			Variadic: true,
		},
	}
	s.VisitPush(cur)
}

// ExitProgram is called when production program is exited.
func (s *ASTBuilder) ExitProgram(ctx *parser.ProgramContext) {
	node := s.VisitPop().(*ProgramRoot)
	num := len(ctx.AllStatement())
	for i := 0; i < num; i++ {
		temp := s.NodePop()
		switch temp.(type) {
		case *FunctionDefNode:
			node.FunctionDefNode = append(node.FunctionDefNode, temp)
		case *TypeDef, *TypeAlias:
			node.TypeDefNode = append(node.TypeDefNode, temp)
		case *EnumNode:
			node.EnumNode = append(node.EnumNode, temp)
		default:
			node.RunnableNode = append([]ASTNode{temp}, node.RunnableNode...)
		}
	}
	s.NodePush(node)
}

// EnterEnumdef is called when production enumdef is entered.
func (s *ASTBuilder) EnterEnumdef(ctx *parser.EnumdefContext) {
	cur := &EnumNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		EnumName: ctx.NAME(0).GetText(),
	}
	num := len(ctx.AllINT())
	for i := 0; i < num; i++ {
		v, _ := strconv.ParseInt(ctx.INT(i).GetText(), 10, 32)
		cur.Enum[ctx.NAME(i+1).GetText()] = int32(v)
		cur.REnum[int32(v)] = ctx.NAME(i + 1).GetText()
	}
	s.NodePush(cur)
}

// ExitSimpleTypeName is called when production SimpleTypeName is exited.
func (s *ASTBuilder) ExitSimpleTypeName(ctx *parser.SimpleTypeNameContext) {
	// todo
}

// EnterTypeDefAlias is called when production TypeDefAlias is entered.
func (s *ASTBuilder) EnterTypeDefAlias(ctx *parser.TypeDefAliasContext) {
	// todo
}

// EnterTypeDefComplex is called when production TypeDefComplex is entered.
func (s *ASTBuilder) EnterTypeDefComplex(ctx *parser.TypeDefComplexContext) {
	// todo
}

// EnterFunctiondef is called when production functiondef is entered.
func (s *ASTBuilder) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	s.InFunction = true
	s.Compiler.Scope = common.NewScope(s.Compiler.Scope)
	cur := &FunctionDefNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		Meta: &common.FunctionMeta{
			Name:      ctx.NAME().GetText(),
			ConstExpr: false,
		},
	}
	s.Compiler.FunctionLib[cur.Meta.Name] = cur.Meta
	s.VisitPush(cur)
}

// ExitFunctiondef is called when production functiondef is exited.
func (s *ASTBuilder) ExitFunctionDef(ctx *parser.FunctionDefContext) {
	s.InFunction = false
	fNode := s.VisitPop().(*FunctionDefNode)
	fNode.Block = s.NodePop()
	s.NodePush(fNode)
}

// ExitReturntypename is called when production returntypename is exited.
func (s *ASTBuilder) ExitReturntypename(ctx *parser.ReturntypenameContext) {
	fNode := s.VisitTop().(*FunctionDefNode)
	typeName := strings.Replace(ctx.Typename().GetText(), " ", "", -1)
	t := s.Compiler.TypeRegistry.FindType(typeName)
	if t == nil {
		panic(common.NewCompileErr("unknown param type " + typeName))
	}
	fNode.Meta.Out = append(fNode.Meta.Out, t)
}

// ExitInparam is called when production inparam is exited.
func (s *ASTBuilder) ExitInparam(ctx *parser.InparamContext) {
	fNode := s.VisitTop().(*FunctionDefNode)
	paramNode := s.NodePop().(*ParamNode)
	t := s.Compiler.TypeRegistry.FindType(paramNode.TypeName)
	if t == nil {
		panic(common.NewCompileErr("unknown param type " + paramNode.TypeName + " of parameter " + paramNode.Symbol))
	}
	s.Compiler.Scope.AddParameterVariable(&common.Variable{
		Symbol: paramNode.Symbol,
		Type:   t,
	})
	fNode.Meta.In = append(fNode.Meta.In, t)
}

// ExitOutparam is called when production outparam is exited.
func (s *ASTBuilder) ExitOutparam(ctx *parser.OutparamContext) {
	fNode := s.VisitTop().(*FunctionDefNode)
	paramNode := s.NodePop().(*ParamNode)
	t := s.Compiler.TypeRegistry.FindType(paramNode.TypeName)
	if t == nil {
		panic(common.NewCompileErr("unknown param type " + paramNode.TypeName + " of parameter " + paramNode.Symbol))
	}
	s.Compiler.Scope.AddReturnVariable(&common.Variable{
		Symbol: paramNode.Symbol,
		Type:   t,
	})
	fNode.Meta.Out = append(fNode.Meta.Out, t)
}

// EnterParam is called when production param is entered.
func (s *ASTBuilder) ExitParam(ctx *parser.ParamContext) {
	name := ctx.NAME().GetText()
	typeName := strings.Replace(ctx.Typename().GetText(), " ", "", -1)
	s.NodePush(&ParamNode{
		Symbol:   name,
		TypeName: typeName,
	})
}

// ExitMapTypeName is called when production MapTypeName is exited.
func (s *ASTBuilder) ExitMapTypeName(ctx *parser.MapTypeNameContext) {
	keyType := ctx.NAME().GetText()
	valType := strings.Replace(ctx.Typename().GetText(), " ", "", -1)
	s.Compiler.TypeRegistry.FindMapType(keyType, valType)
}

// ExitSliceTypeName is called when production SliceTypeName is exited.
func (s *ASTBuilder) ExitSliceTypeName(ctx *parser.SliceTypeNameContext) {
	itemType := strings.Replace(ctx.Typename().GetText(), " ", "", -1)
	s.Compiler.TypeRegistry.FindSliceType(itemType)
}
