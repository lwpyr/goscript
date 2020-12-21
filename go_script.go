package goscript

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lwpyr/goscript/ast"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/functions/libbase64"
	"github.com/lwpyr/goscript/functions/libcommon"
	"github.com/lwpyr/goscript/functions/libdatetime"
	"github.com/lwpyr/goscript/functions/libjson"
	"github.com/lwpyr/goscript/functions/libstring"
	"github.com/lwpyr/goscript/instruction"
	"github.com/lwpyr/goscript/parser"
	"github.com/lwpyr/goscript/parser_goscript"
	"github.com/lwpyr/goscript/program"
	"runtime/debug"
)

func BuildFunctionDefAST(c *ast.CompileContext, expr string) (root ast.ASTNode, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(common.ScriptError)
		}
	}()
	l := parser_goscript.NewASTBuilder()
	l.Compiler = c
	is := antlr.NewInputStream(expr)
	lexer := parser.NewgoscriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	d := parser.NewgoscriptParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(l, d.Functiondef())
	return l.NodeTop(), nil
}

func CompileFunctionDef(c *ast.CompileContext, expr string) error {
	node, err := BuildFunctionDefAST(c, expr)
	if err != nil {
		return err
	}
	node.Compile(c)
	return nil
}

func BuildSingleLineAST(c *ast.CompileContext, expr string) (root ast.ASTNode, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	l := parser_goscript.NewASTBuilder()
	l.Compiler = c
	is := antlr.NewInputStream(expr)
	lexer := parser.NewgoscriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	d := parser.NewgoscriptParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(l, d.KeepStack())
	return l.NodeTop(), nil
}

func CompileExpression(c *ast.CompileContext, expr string) (common.Instruction, error) {
	node, err := BuildSingleLineAST(c, expr)
	if err != nil {
		return nil, err
	}
	node.Compile(c)
	instructions := node.GetInstructions()
	instructions = append(instructions, instruction.ProgramEnding())
	ret := func(m *common.Memory, stk *common.Stack) {
		for stk.Pc != -1 {
			instructions[stk.Pc](m, stk)
		}
	}
	return ret, nil
}

func CompileSingleLineProgram(expr string, c *ast.CompileContext, s *common.Scope) (p *program.SingleLineProgram, err error) {
	if s != nil {
		c.Scope = s
	}
	inst, err := CompileExpression(c, expr)
	defer func() {
		r := recover()
		if r != nil {
			p = nil
			err = fmt.Errorf("compile failed %v", r)
		}
	}()
	if err != nil {
		return nil, err
	}
	return &program.SingleLineProgram{Root: inst}, nil
}

func BuildAST(c *ast.CompileContext, expr string) (root ast.ASTNode, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(common.ScriptError); ok {
				err = fmt.Errorf("%v\n%s", e, string(debug.Stack()))
			} else {
				err = fmt.Errorf("%v\n%s", r, string(debug.Stack()))
			}
		}
	}()
	l := parser_goscript.NewASTBuilder()
	l.Compiler = c
	is := antlr.NewInputStream(expr)
	lexer := parser.NewgoscriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	d := parser.NewgoscriptParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(l, d.Program())
	return l.NodeTop(), nil
}

func CompileScript(c *ast.CompileContext, script string) (common.Instruction, error) {
	node, err := BuildAST(c, script)
	if err != nil {
		return nil, err
	}
	node.Compile(c)
	return node.GetInstructions()[0], nil
}

func CompileScriptProgram(expr string, c *ast.CompileContext, s *common.Scope) (p *program.ScriptProgram, err error) {
	if s != nil {
		c.Scope = s
	}
	inst, err := CompileScript(c, expr)
	defer func() {
		r := recover()
		if r != nil {
			p = nil
			err = fmt.Errorf("compile failed %v", r)
		}
	}()
	if err != nil {
		return nil, err
	}
	return &program.ScriptProgram{Root: inst}, nil
}

func AcquireStack() *common.Stack {
	return common.Get()
}

func ReturnStack(stk *common.Stack) {
	common.Put(stk)
}

func NewScope(outer *common.Scope) *common.Scope {
	return common.NewScope(outer)
}

func NewTypeRegistry() *common.TypeRegistry {
	return common.NewTypeRegistry()
}

func NewVariable(name string, t *common.DataType) *common.Symbol {
	return &common.Symbol{
		Symbol:   name,
		DataType: t,
	}
}

func NewMemory(size int) *common.Memory {
	return &common.Memory{
		Data: make([]interface{}, 100),
	}
}

func NewCompiler() *ast.CompileContext {
	return &ast.CompileContext{
		TypeRegistry: NewTypeRegistry(),
		Scope:        nil,
		//FunctionLib:  nil,
	}
}

func init() {
	libjson.Register()
	libstring.Register()
	libbase64.Register()
	libdatetime.Register()
	libcommon.Register()
}
