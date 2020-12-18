package ast

import (
	"github.com/lwpyr/goscript/common"
)

type Compiler struct {
	TypeRegistry *common.TypeRegistry
	Scope        *common.Scope
	FunctionInfo []*common.FunctionMeta
}

func (c *Compiler) PushFunctionMeta(meta *common.FunctionMeta) {
	c.FunctionInfo = append(c.FunctionInfo, meta)
}

func (c *Compiler) CurrentFunctionMeta() *common.FunctionMeta {
	return c.FunctionInfo[len(c.FunctionInfo)-1]
}

func (c *Compiler) PopFunctionMeta() {
	c.FunctionInfo = c.FunctionInfo[:len(c.FunctionInfo)-1]
}

func (c *Compiler) MakeFunctionScope() {
	c.Scope = common.NewScope(c.Scope)
	c.Scope.LocalIndex = 0
}

func (c *Compiler) MakeChildScope() {
	c.Scope = common.NewScope(c.Scope)
}

func (c *Compiler) ReturnParentScope() {
	c.Scope = c.Scope.Outer
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

func (c *Compiler) Include(libName string) {
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
