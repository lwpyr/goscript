package libbase64

import (
	"github.com/lwpyr/goscript/common"
)

func Register() {
	common.RegisterLib("base64", Lib)
}

var Lib *Base64Lib

type Base64Lib struct{}

func (b *Base64Lib) Init(tr *common.TypeRegistry) map[string]*common.FunctionMeta {
	return map[string]*common.FunctionMeta{
		"EncodeBase64": {
			Name: "EncodeBase64",
			In: []*common.DataType{
				common.BasicTypeMap["bytes"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			F:         EncodeBase64,
			ConstExpr: false,
		},
		"DecodeBase64": {
			Name: "DecodeBase64",
			In: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["bytes"],
			},
			F:         DecodeBase64,
			ConstExpr: false,
		},
	}
}
