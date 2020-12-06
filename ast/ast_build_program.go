package ast

import (
	"github.com/lw396285v/goscript/common"
	"github.com/lw396285v/goscript/parser"
	"strconv"
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
		case *GlobalNode:
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

// ExitGenericTypeName is called when production GenericTypeName is exited.
func (s *ASTBuilder) ExitGenericTypeName(ctx *parser.GenericTypeNameContext) {
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
func (s *ASTBuilder) EnterFunctiondef(ctx *parser.FunctiondefContext) {
	s.InFunction = true
	s.Compiler.Scope = common.NewScope(s.Compiler.Scope)
	cur := &FunctionDefNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
		},
		Meta: &common.FunctionMeta{
			Name:      ctx.NAME(0).GetText(),
			ConstExpr: false,
		},
	}
	if len(ctx.AllNAME()) > 1 {
		t := s.Compiler.TypeRegistry.FindType(ctx.NAME(1).GetText())
		if t == nil {
			panic(common.NewCompileErr("unknown return type of function " + cur.Meta.Name))
		}
		cur.Meta.Out = append(cur.Meta.Out, t)
	}
	s.Compiler.FunctionLib[cur.Meta.Name] = cur.Meta
	s.VisitPush(cur)
}

// ExitFunctiondef is called when production functiondef is exited.
func (s *ASTBuilder) ExitFunctiondef(ctx *parser.FunctiondefContext) {
	s.InFunction = false
	fNode := s.VisitPop().(*FunctionDefNode)
	fNode.Block = s.NodePop()
	s.NodePush(fNode)
}

// EnterParam is called when production param is entered.
func (s *ASTBuilder) EnterParam(ctx *parser.ParamContext) {
	fNode := s.VisitTop().(*FunctionDefNode)
	name := ctx.NAME(0).GetText()
	t := s.Compiler.TypeRegistry.FindType(ctx.NAME(1).GetText())
	if t == nil {
		panic(common.NewCompileErr("unknown param type " + ctx.NAME(1).GetText() + " of function " + fNode.Meta.Name))
	}
	s.Compiler.Scope.AddParameterVariable(&common.Variable{
		Symbol: name,
		Type:   t,
	})
	fNode.Meta.In = append(fNode.Meta.In, t)
}
