package test

import (
	"github.com/lwpyr/goscript"
	"github.com/lwpyr/goscript/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func compileConstExpr(expr string) interface{} {
	node, err := goscript.BuildSingleLineAST(constC, expr)
	if err != nil {
		panic(err)
	}
	node.Compile(constC)
	return node.GetConstantValue()
}

func TestConstantExpression(t *testing.T) {
	expr := "5+3*2.3"
	assert.Equal(t, float64(5)+float64(3)*2.3, compileConstExpr(expr))
	expr = "5+3*23"
	assert.Equal(t, int64(5)+int64(3)*int64(23), compileConstExpr(expr))
	expr = "(5+3)*23"
	assert.Equal(t, (int64(5)+int64(3))*int64(23), compileConstExpr(expr))
	expr = "(5+3)/23"
	assert.Equal(t, int64(0), compileConstExpr(expr))
	expr = "3**2"
	assert.Equal(t, int64(9), compileConstExpr(expr))
	expr = "3**3**3"
	assert.Equal(t, int64(7625597484987), compileConstExpr(expr))
	expr = "(3**3)**3"
	assert.Equal(t, int64(19683), compileConstExpr(expr))
	expr = "10-5%2"
	assert.Equal(t, int64(9), compileConstExpr(expr))
	expr = "'Tom' + ' ' + 'and' + ' ' + 'Jerry'"
	assert.Equal(t, "Tom and Jerry", compileConstExpr(expr))
}

func TestConstantExpressionWithEnum(t *testing.T) {
	expr := "fruit.apple"
	assert.Equal(t, int32(0), compileConstExpr(expr))
	expr = "fruit.apple + sports.soccer"
	assert.Equal(t, int32(3), compileConstExpr(expr))
	expr = "fruit.orange * sports.soccer"
	assert.Equal(t, int32(6), compileConstExpr(expr))
	expr = "(fruit.orange * sports.soccer)**(fruit.banana + sports.basketball)"
	assert.Equal(t, int32(36), compileConstExpr(expr))
}

var constC *ast.Compiler

func init() {
	tr := goscript.NewTypeRegistry()
	tr.AddEnum("fruit", map[string]int32{
		"apple":      int32(0),
		"banana":     int32(1),
		"orange":     int32(2),
		"strawberry": int32(3),
	})
	tr.AddEnum("sports", map[string]int32{
		"football":   int32(0),
		"basketball": int32(1),
		"volleyball": int32(2),
		"soccer":     int32(3),
	})
	constC = &ast.Compiler{
		TypeRegistry: tr,
		Scope:        nil,
	}
}
