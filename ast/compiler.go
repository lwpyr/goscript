package ast

import (
	"github.com/lwpyr/goscript/common"
)

type CompileContext struct {
	TypeRegistry *common.TypeRegistry
	Scope        *common.Scope
	FunctionInfo []*common.FunctionMeta
}

func (c *CompileContext) PushFunctionMeta(meta *common.FunctionMeta) {
	c.FunctionInfo = append(c.FunctionInfo, meta)
}

func (c *CompileContext) CurrentFunctionMeta() *common.FunctionMeta {
	return c.FunctionInfo[len(c.FunctionInfo)-1]
}

func (c *CompileContext) PopFunctionMeta() {
	c.FunctionInfo = c.FunctionInfo[:len(c.FunctionInfo)-1]
}

func (c *CompileContext) MakeFunctionScope() {
	c.Scope = common.NewScope(c.Scope)
	c.Scope.LocalIndex = 0
}

func (c *CompileContext) MakeChildScope() {
	c.Scope = common.NewScope(c.Scope)
}

func (c *CompileContext) ReturnParentScope() {
	c.Scope = c.Scope.Outer
}

func (c *CompileContext) AddEnum(name string, e map[string]int32) {
	c.TypeRegistry.AddEnum(name, e)
}

func (c *CompileContext) GetEnums(name string) interface{} {
	return c.TypeRegistry.GetEnums(name)
}

func (c *CompileContext) GetREnums(name string) map[int32]string {
	return c.TypeRegistry.GetREnums(name)
}

func (c *CompileContext) FindSliceType(name string) *common.DataType {
	return c.TypeRegistry.FindSliceType(name)
}

func (c *CompileContext) FindMapType(key, value string) *common.DataType {
	return c.TypeRegistry.FindMapType(key, value)
}

func (c *CompileContext) AddType(name string, dType *common.DataType) {
	c.TypeRegistry.AddType(name, dType)
}

func (c *CompileContext) FindFuncType(meta *common.FunctionMeta) *common.DataType {
	return c.TypeRegistry.FindFuncType(meta)
}

func (c *CompileContext) FindType(name string) *common.DataType {
	return c.TypeRegistry.FindType(name)
}

func (c *CompileContext) AddBuiltType(dtb *common.DataTypeBuilder) {
	c.TypeRegistry.AddBuiltType(dtb)
}

func (c *CompileContext) Include(libName string) {
	lib := common.RegisteredLibs[libName]
	if lib == nil {
		panic("Lib name not recognized, skipped")
	}
	for name, f := range lib.Init(c.TypeRegistry) {
		c.Scope.AddConstant(name, &common.Symbol{
			Symbol:   name,
			DataType: f.Type,
			Data:     f.F,
		})
	}
}
