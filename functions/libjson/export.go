package libjson

import (
	"github.com/lwpyr/goscript/common"
)

func Register() {
	common.RegisterLib("json", Lib)
}

var Lib *JsonLib

type JsonLib struct{}

func (j *JsonLib) Init(tr *common.TypeRegistry) map[string]*common.Function {
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

	var JsonMarshalFunc common.Instruction = JsonMarshal
	var JsonUnmarshalFunc common.Instruction = JsonUnmarshal
	var JwtParseFunc common.Instruction = JwtParse
	var JsonObjectFunc common.Instruction = JsonObject
	var JsonPathFunc common.Instruction = JsonPath

	return map[string]*common.Function{
		"JsonMarshal": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["object"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["bytes"],
				},
			}),
			F: &JsonMarshalFunc,
		},
		"JsonUnmarshal": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["bytes"],
					common.BasicTypeMap["reflect"],
				},
				Out: []*common.DataType{
					common.BasicTypeMap["object"],
				},
			}),
			F: &JsonUnmarshalFunc,
		},
		"JwtParse": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["string"],
				},
				Out: []*common.DataType{
					tr.FindType("JWT"),
				},
			}),
			F: &JwtParseFunc,
		},
		"JsonObject": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					common.BasicTypeMap["bytes"],
				},
				Out: []*common.DataType{
					tr.FindType("JSONObject"),
				},
			}),
			F: &JsonObjectFunc,
		},
		"JsonPath": {
			Type: tr.FindFuncType(&common.FunctionMeta{
				In: []*common.DataType{
					tr.FindType("JSONObject"),
					tr.FindType("string"),
				},
				Out: []*common.DataType{
					tr.FindType("JSONObject"),
				},
			}),
			F: &JsonPathFunc,
		},
	}
}
