package libstring

import (
	"github.com/lwpyr/goscript/common"
)

func Register() {
	common.RegisterLib("string", Lib)
}

var Lib *StringLib

type StringLib struct{}

func (s *StringLib) Init(tr *common.TypeRegistry) map[string]*common.FunctionMeta {
	return map[string]*common.FunctionMeta{
		"startWith": {
			Name: "startWith",
			In: []*common.DataType{
				common.BasicTypeMap["string"],
				common.BasicTypeMap["string"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["bool"],
			},
			F:         StartWith,
			ConstExpr: true,
		},
		"endWith": {
			Name: "endWith",
			In: []*common.DataType{
				common.BasicTypeMap["string"],
				common.BasicTypeMap["string"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["bool"],
			},
			F:         EndWith,
			ConstExpr: true,
		},
		"contain": {
			Name: "contain",
			In: []*common.DataType{
				common.BasicTypeMap["string"],
				common.BasicTypeMap["string"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["bool"],
			},
			F:         Contain,
			ConstExpr: true,
		},
		"strCat": {
			Name: "strCat",
			In: []*common.DataType{
				common.BasicTypeMap["string"],
				common.BasicTypeMap["string"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			F:         StrCat,
			ConstExpr: true,
			TailArray: true,
		},
		"strSplit": {
			Name: "strSplit",
			In: []*common.DataType{
				common.BasicTypeMap["string"],
				common.BasicTypeMap["string"],
			},
			Out: []*common.DataType{
				tr.FindSliceType("string"),
			},
			F:         StrSplit,
			ConstExpr: true,
		},
		"test1": {
			Name: "test1",
			In: []*common.DataType{
				common.BasicTypeMap["int64"],
				common.BasicTypeMap["int64"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["int64"],
				common.BasicTypeMap["int64"],
			},
			F: func(m *common.Memory, stk *common.Stack) {
				i1 := stk.TopIndex(1).(int64)
				i2 := stk.Top().(int64)
				r1 := i1 / i2
				r2 := i1 % i2
				stk.Pop()
				stk.ReturnVal([]interface{}{r1, r2})
			},
			ConstExpr: true,
		},
	}
}
