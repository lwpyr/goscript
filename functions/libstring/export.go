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
			F: StartWith,
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
			F: EndWith,
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
			F: Contain,
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
			F: StrCat,
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
			F: StrSplit,
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
			F: func(m *common.Memory, stk *common.Stack) {
				i1 := stk.TopIndex(1).(int64)
				i2 := stk.Top().(int64)
				stk.ReturnValues(i1/i2, i1%i2)
			},
		},
	}
}
