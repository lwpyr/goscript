package libcommon

import (
	"github.com/lwpyr/goscript/common"
)

func Register() {
	common.RegisterLib("common", Lib)
}

var Lib *commonLib

type commonLib struct{}

func (b *commonLib) Init(tr *common.TypeRegistry) map[string]*common.Function {
	return map[string]*common.Function{
		"itoa": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["int64"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["string"],
				},
			}),
			F: ItoA,
		},
		"atoi": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["string"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["int64"],
				},
			}),
			F: AtoI,
		},
		"ftoa": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["float64"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["string"],
				},
			}),
			F: FtoA,
		},
		"atof": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["string"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["float64"],
				},
			}),
			F: AtoF,
		},
		"random": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{},
				Out: []*common.DataType{
					common.BasicTypeMap["float64"],
				},
			}),
			F: Random,
		},
		"print": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["object"],
				},
				TailArray: true,
			}),
			F: Print,
		},
		"sleep": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["int64"],
				},
			}),
			F: Sleep,
		},
	}
}
