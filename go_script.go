package go_script

import (
	"fmt"
	"github.com/lw396285v/goscript/ast"
	"github.com/lw396285v/goscript/common"
	"github.com/lw396285v/goscript/functions/libbase64"
	"github.com/lw396285v/goscript/functions/libcommon"
	"github.com/lw396285v/goscript/functions/libdatetime"
	"github.com/lw396285v/goscript/functions/libjson"
	"github.com/lw396285v/goscript/functions/libstring"
	"github.com/lw396285v/goscript/program"
)

func CompileSingleLineProgram(expr string, c *ast.Compiler, s *common.Scope) (p *program.SingleLineProgram, err error) {
	if s != nil {
		c.Scope = s
	}
	inst, err := c.CompileExpression(expr)
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

func NewVariable(name string, t *common.DataType) *common.Variable {
	return &common.Variable{
		Symbol: name,
		Type:   t,
	}
}

func NewParameter(name string, t *common.DataType, offset int) *common.Variable {
	return &common.Variable{
		Offset:      offset,
		Symbol:      name,
		Type:        t,
		IsParameter: true,
	}
}

func NewCompiler() *ast.Compiler {
	return &ast.Compiler{
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
