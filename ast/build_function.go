package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
	"strings"
)

// EnterFunctiondef is called when production functiondef is entered.
func (s *ASTBuilder) EnterFunctionDef(ctx *parser.FunctionDefContext) {
	s.Compiler.Scope = common.NewScope(s.Compiler.Scope)
	s.Compiler.Scope.LocalIndex = 0
	funcName := ctx.NAME().GetText()
	funcPlaceHolder := common.Instruction(nil)
	cur := &FunctionDefNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: &common.DataType{
				Kind: common.KindMap[common.Func],
				LambdaMeta: &common.FunctionMeta{
					ConstExpr: false,
					TailArray: ctx.TAILARRAY() != nil,
				},
			},
		},
		FuncName: funcName,
		Function: &funcPlaceHolder,
	}
	s.Closure = append(s.Closure, cur.DataType.LambdaMeta)
	s.VisitPush(cur)
}

// ExitFunctiondef is called when production functiondef is exited.
func (s *ASTBuilder) ExitFunctionDef(ctx *parser.FunctionDefContext) {
	s.Closure = s.Closure[:len(s.Closure)-1]
	s.Compiler.Scope = s.Compiler.Scope.Outer
	fNode := s.VisitPop().(*FunctionDefNode)
	fNode.Block = s.NodePop()
	s.NodePush(fNode)
}

// EnterLambdaDef is called when production LambdaDef is entered.
func (s *ASTBuilder) EnterLambdaDef(ctx *parser.LambdaDefContext) {
	s.Compiler.Scope = common.NewScope(s.Compiler.Scope)
	s.Compiler.Scope.LocalIndex = 0
	s.Compiler.Scope.SetCaptureMode()
	cur := &LambdaDefNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: &common.DataType{
				Kind: common.KindMap[common.Func],
				LambdaMeta: &common.FunctionMeta{
					ConstExpr: false,
					TailArray: ctx.TAILARRAY() != nil,
				},
			},
		},
		Scope: s.Compiler.Scope,
	}
	s.Closure = append(s.Closure, cur.DataType.LambdaMeta)
	s.VisitPush(cur)
}

// ExitLambdaDef is called when production LambdaDef is exited.
func (s *ASTBuilder) ExitLambdaDef(ctx *parser.LambdaDefContext) {
	s.Closure = s.Closure[:len(s.Closure)-1]
	s.Compiler.Scope = s.Compiler.Scope.Outer
	fNode := s.VisitPop().(*LambdaDefNode)
	fNode.Block = s.NodePop()
	s.NodePush(fNode)
}

// ExitIntypename is called when production intypename is exited.
func (s *ASTBuilder) ExitIntypename(ctx *parser.IntypenameContext) {
	fNodeMeta := s.VisitTop().GetDataType().LambdaMeta
	fNodeMeta.In = append(fNodeMeta.In, s.NodePop().GetDataType())
}

// ExitReturntypename is called when production returntypename is exited.
func (s *ASTBuilder) ExitReturntypename(ctx *parser.ReturntypenameContext) {
	fNodeMeta := s.VisitTop().GetDataType().LambdaMeta
	fNodeMeta.Out = append(fNodeMeta.Out, s.NodePop().GetDataType())
}

// ExitInparam is called when production inparam is exited.
func (s *ASTBuilder) ExitInparam(ctx *parser.InparamContext) {
	fNodeMeta := s.VisitTop().GetDataType().LambdaMeta
	paramNode := s.NodePop().(*ParamNode)
	s.Compiler.Scope.AddParameterVariable(&common.Variable{
		Symbol:   paramNode.Symbol,
		DataType: paramNode.DataType,
		Scope:    s.Compiler.Scope,
	})
	fNodeMeta.In = append(fNodeMeta.In, paramNode.DataType)
}

// ExitOutparam is called when production outparam is exited.
func (s *ASTBuilder) ExitOutparam(ctx *parser.OutparamContext) {
	fNodeMeta := s.VisitTop().GetDataType().LambdaMeta
	paramNode := s.NodePop().(*ParamNode)
	s.Compiler.Scope.AddReturnVariable(&common.Variable{
		Symbol:   paramNode.Symbol,
		DataType: paramNode.DataType,
		Scope:    s.Compiler.Scope,
	})
	fNodeMeta.Out = append(fNodeMeta.Out, paramNode.DataType)
}

// EnterParam is called when production param is entered.
func (s *ASTBuilder) ExitParam(ctx *parser.ParamContext) {
	name := ctx.NAME().GetText()
	s.NodePush(&ParamNode{
		Node: Node{
			DataType: s.NodePop().GetDataType(),
		},
		Symbol: name,
	})
}

// ExitSimpleTypeName is called when production SimpleTypeName is exited.
func (s *ASTBuilder) ExitSimpleTypeName(ctx *parser.SimpleTypeNameContext) {
	typeName := strings.Replace(ctx.GetText(), " ", "", -1)
	t := s.Compiler.TypeRegistry.FindType(typeName)
	if t == nil {
		panic(common.NewTypeNotFoundErr(typeName))
	}
	s.NodePush(&TypeNode{
		Node: Node{
			DataType: t,
		},
	})
}

// ExitChanTypeName is called when production ChanTypeName is exited.
func (s *ASTBuilder) ExitChanTypeName(ctx *parser.ChanTypeNameContext) {
	itemType := s.NodePop().GetDataType().Type
	t := s.Compiler.TypeRegistry.FindChanType(itemType)
	if t == nil {
		panic(common.NewTypeNotFoundErr(fmt.Sprintf("chan<%s>", itemType)))
	}
	s.NodePush(&TypeNode{
		Node: Node{
			DataType: t,
		},
	})
}

// ExitMapTypeName is called when production MapTypeName is exited.
func (s *ASTBuilder) ExitMapTypeName(ctx *parser.MapTypeNameContext) {
	keyType := ctx.BasicTypeName().GetText()
	valType := s.NodePop().GetDataType().Type
	t := s.Compiler.TypeRegistry.FindMapType(keyType, valType)
	if t == nil {
		panic(common.NewTypeNotFoundErr(fmt.Sprintf("map<%s,%s>", keyType, valType)))
	}
	s.NodePush(&TypeNode{
		Node: Node{
			DataType: t,
		},
	})
}

// ExitSliceTypeName is called when production SliceTypeName is exited.
func (s *ASTBuilder) ExitSliceTypeName(ctx *parser.SliceTypeNameContext) {
	itemType := s.NodePop().GetDataType().Type
	t := s.Compiler.TypeRegistry.FindSliceType(itemType)
	if t == nil {
		panic(common.NewTypeNotFoundErr(fmt.Sprintf("slice<%s>", itemType)))
	}
	s.NodePush(&TypeNode{
		Node: Node{
			DataType: t,
		},
	})
}

func (s *ASTBuilder) EnterFunctionTypeName(ctx *parser.FunctionTypeNameContext) {
	s.VisitPush(&TypeNode{
		Node: Node{
			DataType: &common.DataType{
				LambdaMeta: &common.FunctionMeta{},
			},
		},
	})
}

// ExitFunctionTypeName is called when production functionTypeName is exited.
func (s *ASTBuilder) ExitFunctionTypeName(ctx *parser.FunctionTypeNameContext) {
	node := s.VisitPop().(*TypeNode)
	node.DataType = s.Compiler.FindFuncType(node.DataType.LambdaMeta)
	s.NodePush(node)
}

// EnterClosure is called when production closure is entered.
func (s *ASTBuilder) EnterClosure(ctx *parser.ClosureContext) {
	node := s.VisitTop()
	meta := node.GetDataType().LambdaMeta
	node.SetDataType(s.Compiler.FindFuncType(meta))
	if meta.TailArray {
		lenOut := len(meta.Out)
		v := s.Compiler.Scope.GetVariable(meta.Out[lenOut-1].Type)
		v.DataType = s.Compiler.FindSliceType(v.DataType.Type)
	}
	if _, ok := node.(*LambdaDefNode); ok {
		s.Compiler.Scope.AddParameterVariable(&common.Variable{
			Symbol:       "#",
			VariableType: common.Local,
			Scope:        s.Compiler.Scope,
		})
	}

	if fDef, ok := node.(*FunctionDefNode); ok {
		s.Compiler.Scope.Outer.AddConstant(fDef.FuncName, &common.Constant{
			Symbol: fDef.FuncName,
			Type:   fDef.DataType,
			Data:   fDef.Function,
		})
	}
}

// EnterFunctions is called when production functions is entered.
func (s *ASTBuilder) EnterDirectCall(ctx *parser.DirectCallContext) {
	cur := &FunctionCallNode{
		Node: Node{
			Parent:   s.VisitTop(),
			Variadic: true,
			DataType: common.BasicTypeMap["nil"],
			NodeType: "FunctionCallNode",
		},
	}
	s.VisitPush(cur)
}

// ExitFunctions is called when production functions is exited.
func (s *ASTBuilder) ExitDirectCall(ctx *parser.DirectCallContext) {
	params := make([]ASTNode, len(ctx.AllExpr()))
	for i := 0; i < len(ctx.AllExpr()); i++ {
		params[i] = s.NodePop()
	}
	functionNameNode := s.NodePop()
	funcName := ctx.Variable().GetText()
	cur := s.VisitPop().(*FunctionCallNode)
	var meta *common.FunctionMeta
	lambdaVarType := functionNameNode.GetDataType()
	if !functionNameNode.IsVariadic() {
		// named function
		functionConstant := functionNameNode.(*ValueNode)
		meta = functionConstant.DataType.LambdaMeta
		if meta == nil {
			panic(common.NewSymbolErr("undefined function " + funcName))
		}
		cur.Function = functionConstant
	} else {
		// lambda
		if lambdaVarType.Kind.Kind != common.Func {
			panic(common.NewTypeErr("symbol is not lambda: " + funcName))
		}
		meta = lambdaVarType.LambdaMeta
		cur.Function = functionNameNode
	}
	cur.Meta = meta
	if len(meta.Out) == 1 {
		cur.DataType = meta.Out[0]
	}
	provided := len(ctx.AllExpr())
	if provided > len(meta.In) && meta.TailArray == false || provided < len(meta.In) {
		panic(common.NewMismatchErr("parameter given mismatch with needed " + funcName))
	}
	for i := 0; i < len(ctx.AllExpr()); i++ {
		temp := params[i]
		var requireType *common.DataType
		if provided-1-i < len(meta.In) {
			requireType = meta.In[provided-1-i]
		} else {
			requireType = meta.In[len(meta.In)-1]
		}
		if i < len(meta.In) {
			if requireType.Kind.Kind != common.Object && !temp.GetDataType().CanConvertTo(requireType) {
				panic(common.NewTypeErr("function parameter type cannot be implicitly casted " + funcName))
			}
		}
	}
	cur.Params = params
	s.NodePush(cur)
}

func (s *ASTBuilder) ExitBuiltin(ctx *parser.BuiltinContext) {
	var params []ASTNode
	var dType *common.DataType
	builtinName := ctx.GetChild(0).GetPayload().(antlr.Token).GetText()
	if ctx.PUSHBACK() != nil || ctx.PUSHFRONT() != nil || ctx.DELETE() != nil {
		val := s.NodePop()
		data := s.NodePop()
		params = []ASTNode{val, data}
		dType = nil
	} else if ctx.LEN() != nil {
		params = []ASTNode{s.NodePop()}
		dType = common.BasicTypeMap["int64"]
	} else if ctx.TYPEOF() != nil {
		dType = s.NodePop().GetDataType()
		s.NodePush(&ValueNode{
			Node: Node{
				Parent:   s.VisitTop(),
				NodeType: "ValueNode",
				DataType: common.BasicTypeMap["reflect"],
				Variadic: false,
			},
			Val: dType,
		})
		return
	} else if ctx.ENUMSTRING() != nil {
		params = []ASTNode{s.NodePop()}
		val := params[0]
		if val.GetDataType().Kind.Kind != common.Int32 || val.GetDataType().Type == "int32" {
			panic(common.NewTypeErr("enumString() can only be applied on an enum value " + val.GetDataType().Type))
		}
		dType = common.BasicTypeMap["string"]
	} else {
		data := s.NodePop()
		params = []ASTNode{data}
		dType = common.BasicTypeMap[builtinName]
	}
	s.NodePush(&BuiltinFunctionNode{
		Node: Node{
			Parent:   s.VisitTop(),
			NodeType: "BuiltinFunctionNode",
			DataType: dType,
			LhsFlag:  false,
			Variadic: true,
		},
		Params:      params,
		BuiltinName: builtinName,
	})
}
