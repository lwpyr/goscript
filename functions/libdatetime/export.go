package libdatetime

import (
	"github.com/lwpyr/goscript/common"
)

func Register() {
	common.RegisterLib("datetime", Lib)
}

var Lib *DateTimeLib

type DateTimeLib struct{}

func (d *DateTimeLib) Init(tr *common.TypeRegistry) map[string]*common.Function {
	var DateTimeAddFunc common.Instruction = DateTimeAdd
	var NowFunc common.Instruction = Now
	var NowStrFunc common.Instruction = NowStr
	return map[string]*common.Function{
		"DateTimeAdd": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["int64"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["string"],
				},
			}),
			F: &DateTimeAddFunc,
		},
		"Now": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				Out: []*common.DataType{
					common.BasicTypeMap["int64"],
				},
			}),
			F: &NowStrFunc,
		},
		"NowStr": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				Out: []*common.DataType{
					common.BasicTypeMap["string"],
				},
			}),
			F: &NowFunc,
		},
	}
}
