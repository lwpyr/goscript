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
	var ItoAFunc common.Instruction = ItoA
	var AtoIFunc common.Instruction = AtoI
	var FtoAFunc common.Instruction = FtoA
	var AtoFFunc common.Instruction = AtoF
	var RandomFunc common.Instruction = Random
	var PrintFunc common.Instruction = Print
	var SleepFunc common.Instruction = Sleep

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
			F: &ItoAFunc,
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
			F: &AtoIFunc,
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
			F: &FtoAFunc,
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
			F: &AtoFFunc,
		},
		"random": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{},
				Out: []*common.DataType{
					common.BasicTypeMap["float64"],
				},
			}),
			F: &RandomFunc,
		},
		"print": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["object"],
				},
				TailArray: true,
			}),
			F: &PrintFunc,
		},
		"sleep": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["int64"],
				},
			}),
			F: &SleepFunc,
		},
	}
}
