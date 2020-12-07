package libjson

import (
	"github.com/lwpyr/goscript/common"
)

func Register() {
	common.RegisterLib("json", Lib)
}

var Lib *JsonLib

type JsonLib struct{}

func (j *JsonLib) Init(tr *common.TypeRegistry) map[string]*common.FunctionMeta {
	jsonObject := &common.DataType{
		Type:        "JSONObject",
		Kind:        common.KindMap[common.Object],
		Constructor: common.ConstructorMap[common.Object],
		Unmarshal:   nil,
	}

	if tr.FindType("JSONObject") == nil {
		tr.AddType("JSONObject", jsonObject)
	}

	jwtType := &common.DataType{
		Type: "JWT",
		Kind: common.KindMap[common.Message],
		FieldType: map[string]*common.DataType{
			"header":    tr.FindMapType("string", "string"),
			"payload":   tr.FindType("JSONObject"),
			"signature": tr.FindType("string"),
		},
	}

	if tr.FindType("JWT") == nil {
		tr.AddType("JWT", jwtType)
	}

	return map[string]*common.FunctionMeta{
		"JsonMarshal": {
			Name: "JsonMarshal",
			In: []*common.DataType{
				common.BasicTypeMap["object"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["bytes"],
			},
			F: JsonMarshal,
		},
		"JsonUnmarshal": {
			Name: "JsonUnmarshal",
			In: []*common.DataType{
				common.BasicTypeMap["bytes"],
				common.BasicTypeMap["reflect"],
			},
			Out: []*common.DataType{
				common.BasicTypeMap["object"],
			},
			F: JsonUnmarshal,
		},
		"JwtParse": {
			Name: "JwtParse",
			In: []*common.DataType{
				common.BasicTypeMap["string"],
			},
			Out: []*common.DataType{
				tr.FindType("JWT"),
			},
			F: JwtParse,
		},
		"JsonObject": {
			Name: "JsonObject",
			In: []*common.DataType{
				common.BasicTypeMap["bytes"],
			},
			Out: []*common.DataType{
				tr.FindType("JSONObject"),
			},
			F: JsonObject,
		},
		"JsonPath": {
			Name: "JsonPath",
			In: []*common.DataType{
				tr.FindType("JSONObject"),
				tr.FindType("string"),
			},
			Out: []*common.DataType{
				tr.FindType("JSONObject"),
			},
			F: JsonPath,
		},
	}
}
