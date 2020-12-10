package ast

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/lwpyr/goscript/common"
	"github.com/lwpyr/goscript/parser"
	"runtime/debug"
)

type Compiler struct {
	stack        []common.Instruction
	TypeRegistry *common.TypeRegistry
	Scope        *common.Scope
}

func (c *Compiler) AddEnum(name string, e map[string]int32) {
	c.TypeRegistry.AddEnum(name, e)
}

func (c *Compiler) GetEnums(name string) interface{} {
	return c.TypeRegistry.GetEnums(name)
}

func (c *Compiler) GetREnums(name string) map[int32]string {
	return c.TypeRegistry.GetREnums(name)
}

func (c *Compiler) FindSliceType(name string) *common.DataType {
	return c.TypeRegistry.FindSliceType(name)
}

func (c *Compiler) FindMapType(key, value string) *common.DataType {
	return c.TypeRegistry.FindMapType(key, value)
}

func (c *Compiler) AddType(name string, dType *common.DataType) {
	c.TypeRegistry.AddType(name, dType)
}

func (c *Compiler) FindFuncType(meta *common.FunctionMeta) *common.DataType {
	return c.TypeRegistry.FindFuncType(meta)
}

func (c *Compiler) FindType(name string) *common.DataType {
	return c.TypeRegistry.FindType(name)
}

func (c *Compiler) AddBuiltType(dtb *common.DataTypeBuilder) {
	c.TypeRegistry.AddBuiltType(dtb)
}

func (c *Compiler) InstructionPush(i common.Instruction) {
	c.stack = append(c.stack, i)
}

func (c *Compiler) InstructionTop() common.Instruction {
	return c.stack[len(c.stack)-1]
}

func (c *Compiler) InstructionPop() common.Instruction {
	if len(c.stack) == 0 {
		return nil
	}
	ret := c.InstructionTop()
	c.stack = c.stack[:len(c.stack)-1]
	return ret
}

func (c *Compiler) Include(libName string) {
	lib := common.RegisteredLibs[libName]
	if lib == nil {
		panic("Lib name not recognized, skipped")
	}
	for name, f := range lib.Init(c.TypeRegistry) {
		c.Scope.AddConstant(name, &common.Constant{
			Symbol: name,
			Type:   f.Type,
			Data:   f.F,
		})
	}
}

func (c *Compiler) BuildSingleLineAST(expr string) (root ASTNode, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(common.ScriptError)
		}
	}()
	l := NewASTBuilder()
	l.Compiler = c
	is := antlr.NewInputStream(expr)
	lexer := parser.NewgoscriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	d := parser.NewgoscriptParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(l, d.KeepStack())
	return l.NodeTop(), nil
}

func (c *Compiler) CompileNode(node ASTNode) common.Instruction {
	node.Compile(c)
	return c.InstructionPop()
}

func (c *Compiler) CompileExpression(expr string) (common.Instruction, error) {
	node, err := c.BuildSingleLineAST(expr)
	if err != nil {
		return nil, err
	}
	return c.CompileNode(node), nil
}

func (c *Compiler) CompileSetter(expr string) (common.Instruction, error) {
	node, err := c.BuildSingleLineAST(expr)
	if err != nil {
		return nil, err
	}
	node.SetLhs()
	return c.CompileNode(node), nil
}

func (c *Compiler) BuildFunctionDefAST(expr string) (root ASTNode, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(common.ScriptError)
		}
	}()
	l := NewASTBuilder()
	l.Compiler = c
	is := antlr.NewInputStream(expr)
	lexer := parser.NewgoscriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	d := parser.NewgoscriptParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(l, d.Functiondef())
	return l.NodeTop(), nil
}

func (c *Compiler) CompileFunctionDef(expr string) error {
	node, err := c.BuildFunctionDefAST(expr)
	if err != nil {
		return err
	}
	node.Compile(c)
	return nil
}

func (c *Compiler) BuildAST(expr string) (root ASTNode, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(common.ScriptError); ok {
				err = errors.New(e.Error() + "\n" + string(debug.Stack()))
			} else {
				err = errors.New(r.(string) + "\n" + string(debug.Stack()))
			}
		}
	}()
	l := NewASTBuilder()
	l.Compiler = c
	is := antlr.NewInputStream(expr)
	lexer := parser.NewgoscriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	d := parser.NewgoscriptParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(l, d.Program())
	return l.NodeTop(), nil
}

func (c *Compiler) CompileScript(script string) (common.Instruction, error) {
	node, err := c.BuildAST(script)
	if err != nil {
		return nil, err
	}
	return c.CompileNode(node), nil
}
