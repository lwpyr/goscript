package libcommon

import (
	"github.com/lwpyr/goscript/common"
)

func Register() {
	common.RegisterLib("common", Lib)
}

var Lib *commonLib

type commonLib struct{}

func (b *commonLib) Init(tr *common.TypeRegistry) map[string]*common.FunctionMeta {
	return map[string]*common.FunctionMeta{
		"itoa": {
			Name: "itoa",
			In: []*common.DataType{
				common.BasicTypeMap["int64"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			F:         ItoA,
			ConstExpr: false,
		},
		"atoi": {
			Name: "atoi",
			In: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["int64"],
			},
			F:         AtoI,
			ConstExpr: false,
		},
		"ftoa": {
			Name: "ftoa",
			In: []*common.DataType{
				common.BasicTypeMap["float64"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			F:         FtoA,
			ConstExpr: false,
		},
		"atof": {
			Name: "atof",
			In: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["float64"],
			},
			F:         AtoF,
			ConstExpr: false,
		},
		"random": {
			Name: "random",
			In:   []*common.DataType{},
			Out: []*common.DataType{
				common.BasicTypeMap["float64"],
			},
			F:         Random,
			ConstExpr: false,
		},
		"print": {
			Name: "print",
			In: []*common.DataType{
				common.BasicTypeMap["object"],
			},
			F:         Print,
			ConstExpr: false,
			TailArray: true,
		},
		"sleep": {
			Name: "sleep",
			In: []*common.DataType{
				common.BasicTypeMap["int64"],
			},
			F:         Sleep,
			ConstExpr: false,
			TailArray: false,
		},
	}
}
