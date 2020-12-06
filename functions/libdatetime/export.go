package libdatetime

import (
	"github.com/lw396285v/goscript/common"
)

func Register() {
	common.RegisterLib("datetime", Lib)
}

var Lib *DateTimeLib

type DateTimeLib struct{}

func (d *DateTimeLib) Init(tr *common.TypeRegistry) map[string]*common.FunctionMeta {
	return map[string]*common.FunctionMeta{
		"DateTimeAdd": {
			Name: "DateTimeAdd",
			In: []*common.DataType{
				common.BasicTypeMap["int64"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			F:         DateTimeAdd,
			ConstExpr: false,
		},
		"Now": {
			Name: "Now",
			Out: []*common.DataType{
				common.BasicTypeMap["int64"],
			},
			F:         Now,
			ConstExpr: false,
		},
		"NowStr": {
			Name: "NowStr",
			Out: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			F:         NowStr,
			ConstExpr: false,
		},
	}
}
