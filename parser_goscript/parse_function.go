package parser_goscript

import (
	"github.com/lwpyr/goscript/ast"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
	"strings"
)

// EnterFunctiondef is called when production functiondef is entered.
func (s *ASTBuilder) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	s.FunctionDepth++
}

// ExitFunctiondef is called when production functiondef is exited.
func (s *ASTBuilder) ExitFunctionDef(ctx *parser.FunctionDefContext) {
	s.FunctionDepth--
	fNode := &ast.FunctionDefNode{
		Node: ast.Node{
			Variadic: true,
		},
		FuncName:  ctx.Name().GetText(),
		TailArray: ctx.TAILARRAY() != nil,
	}
	fNode.Block = s.NodePop()
	lenOut := len(ctx.AllOutparam()) + len(ctx.AllReturntypename())
	fNode.OutParam = make([]ast.ASTNode, lenOut)
	for i := 0; i < lenOut; i++ {
		fNode.OutParam[lenOut-1-i] = s.NodePop()
	}
	lenIn := len(ctx.AllInparam())
	fNode.InParam = make([]ast.ASTNode, lenIn)
	for i := 0; i < lenIn; i++ {
		fNode.InParam[lenIn-1-i] = s.NodePop()
	}
	fNode.SetSource(ctx)
	s.NodePush(fNode)
}

// EnterLambdaDef is called when production LambdaDef is entered.
func (s *ASTBuilder) EnterLambdaDef(ctx *parser.LambdaDefContext) {
	s.FunctionDepth++
}

// ExitLambdaDef is called when production LambdaDef is exited.
func (s *ASTBuilder) ExitLambdaDef(ctx *parser.LambdaDefContext) {
	s.FunctionDepth--
	fNode := &ast.FunctionDefNode{
		Node: ast.Node{
			Variadic: true,
		},
		Block:     s.NodePop(),
		TailArray: ctx.TAILARRAY() != nil,
	}
	lenOut := len(ctx.AllOutparam()) + len(ctx.AllReturntypename())
	fNode.OutParam = make([]ast.ASTNode, lenOut)
	for i := 0; i < lenOut; i++ {
		fNode.OutParam[lenOut-1-i] = s.NodePop()
	}
	lenIn := len(ctx.AllInparam())
	fNode.InParam = make([]ast.ASTNode, lenIn)
	for i := 0; i < lenIn; i++ {
		fNode.InParam[lenIn-1-i] = s.NodePop()
	}
	s.NodePush(fNode)
}

// EnterParam is called when production param is entered.
func (s *ASTBuilder) ExitParam(ctx *parser.ParamContext) {
	node := &ast.ParamNode{
		Symbol:   ctx.Name().GetText(),
		TypeNode: s.NodePop(),
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitSimpleTypeName is called when production SimpleTypeName is exited.
func (s *ASTBuilder) ExitSimpleTypeName(ctx *parser.SimpleTypeNameContext) {
	node := &ast.TypeNode{
		TypeNodeType:   ast.SimpleType,
		SimpleTypeName: strings.Replace(ctx.GetText(), " ", "", -1),
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitChanTypeName is called when production ChanTypeName is exited.
func (s *ASTBuilder) ExitChanTypeName(ctx *parser.ChanTypeNameContext) {
	node := &ast.TypeNode{
		TypeNodeType: ast.ChanType,
		Value:        s.NodePop(),
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitMapTypeName is called when production MapTypeName is exited.
func (s *ASTBuilder) ExitMapTypeName(ctx *parser.MapTypeNameContext) {
	keyType := ctx.BasicTypeName().GetText()
	valNode := s.NodePop()
	s.NodePush(&ast.TypeNode{
		TypeNodeType: ast.MapType,
		Key:          common.BasicTypeMap[keyType],
		Value:        valNode,
	})
}

// ExitSliceTypeName is called when production SliceTypeName is exited.
func (s *ASTBuilder) ExitSliceTypeName(ctx *parser.SliceTypeNameContext) {
	s.NodePush(&ast.TypeNode{
		TypeNodeType: ast.SliceType,
		Value:        s.NodePop(),
	})
}

// ExitFunctionTypeName is called when production functionTypeName is exited.
func (s *ASTBuilder) ExitFunctionTypeName(ctx *parser.FunctionTypeNameContext) {
	node := &ast.TypeNode{
		TypeNodeType: ast.FuncType,
	}
	lenOut := len(ctx.AllReturntypename())
	node.OutParam = make([]ast.ASTNode, lenOut)
	for i := 0; i < lenOut; i++ {
		node.OutParam[lenOut-1-i] = s.NodePop()
	}
	lenIn := len(ctx.AllIntypename())
	node.InParam = make([]ast.ASTNode, lenIn)
	for i := 0; i < lenIn; i++ {
		node.InParam[lenIn-1-i] = s.NodePop()
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitFunctions is called when production functions is exited.
func (s *ASTBuilder) ExitCall(ctx *parser.CallContext) {
	numParam := len(ctx.AllExpr()) - 1
	params := make([]ast.ASTNode, numParam)
	for i := 0; i < numParam; i++ {
		params[numParam-1-i] = s.NodePop()
	}
	funcNode := s.NodePop()
	if typeNode, ok := funcNode.(*ast.TypeNode); ok {
		if len(params) != 1 {
			panic(common.NewMismatchErr(common.FormatError(ctx, "type convert only support one input")))
		}
		node := &ast.ConvertNode{
			Node: ast.Node{
				Variadic: true,
			},
			Type:  typeNode,
			Value: params[0],
		}
		node.SetSource(ctx)
		s.NodePush(node)
	} else if builtinNode, ok := funcNode.(*ast.BuiltinFunctionNode); ok {
		builtinNode.Params = params
		s.NodePush(builtinNode)
	} else {
		node := &ast.FunctionCallNode{
			Node: ast.Node{
				Variadic: true,
			},
			Function: funcNode,
			Params:   params,
		}
		node.SetSource(ctx)
		s.NodePush(node)
	}
}

func (s *ASTBuilder) ExitTypeConvert(ctx *parser.TypeConvertContext) {
	val := s.NodePop()
	typeNode := s.NodePop()
	node := &ast.ConvertNode{
		Node: ast.Node{
			Variadic: true,
		},
		Type:  typeNode,
		Value: val,
	}
	node.SetSource(ctx)
	s.NodePush(node)
}

// ExitConstruct is called when production Construct is exited.
func (s *ASTBuilder) ExitConstructor(ctx *parser.ConstructorContext) {
	num := len(ctx.AllExpr())
	node := &ast.ConstructorNode{}
	node.Params = make([]ast.ASTNode, num)
	for i := 0; i < num; i++ {
		node.Params[i] = s.NodePop()
	}
	node.Type = s.NodePop()
	node.SetSource(ctx)
	s.NodePush(node)
}
