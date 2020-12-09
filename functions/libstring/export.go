package libstring

import (
	"github.com/lwpyr/goscript/common"
)

func Register() {
	common.RegisterLib("string", Lib)
}

var Lib *StringLib

type StringLib struct{}

func (s *StringLib) Init(tr *common.TypeRegistry) map[string]*common.Function {
	var test1Func common.Instruction = func(m *common.Memory, stk *common.Stack) {
		i1 := stk.TopIndex(1).(int64)
		i2 := stk.Top().(int64)
		stk.ReturnValues(i1/i2, i1%i2)
	}
	var StartWithFunc common.Instruction = StartWith
	var EndWithFunc common.Instruction = EndWith
	var ContainFunc common.Instruction = Contain
	var StrCatFunc common.Instruction = StrCat
	var StrSplitFunc common.Instruction = StrSplit
	return map[string]*common.Function{
		"startWith": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["string"],
					common.BasicTypeMap["string"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["bool"],
				},
			}),
			F: &StartWithFunc,
		},
		"endWith": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["string"],
					common.BasicTypeMap["string"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["bool"],
				},
			}),
			F: &EndWithFunc,
		},
		"contain": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["string"],
					common.BasicTypeMap["string"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["bool"],
				},
			}),
			F: &ContainFunc,
		},
		"strCat": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["string"],
					common.BasicTypeMap["string"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["string"],
				},
				TailArray: true,
			}),
			F: &StrCatFunc,
		},
		"strSplit": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["string"],
					common.BasicTypeMap["string"],
				},
				Out: []*common.DataType{
					tr.FindSliceType("string"),
				},
			}),
			F: &StrSplitFunc,
		},
		"test1": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["int64"],
					common.BasicTypeMap["int64"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["int64"],
					common.BasicTypeMap["int64"],
				},
			}),
			F: &test1Func,
		},
	}
}
