package ast

import (
	"github.com/lw396285v/goscript/common"
	"github.com/lw396285v/goscript/lambda_chains"
	"math"
)

func MergeLeftUnary(op string, operand IConstantNode) IConstantNode {
	switch op {
	case "!":
		if operand.GetDataType().Kind.Kind == common.Bool {
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: !operand.GetConstantValue().(bool),
			}
		}
	case "-":
		switch operand.GetDataType().Kind.Kind {
		case common.UInt32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint32"],
					Variadic: false,
				},
				Val: -int32(operand.GetConstantValue().(uint32)),
			}
		case common.Int32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int32"],
					Variadic: false,
				},
				Val: -operand.GetConstantValue().(int32),
			}
		case common.Float32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float32"],
					Variadic: false,
				},
				Val: -operand.GetConstantValue().(float32),
			}
		case common.UInt64:
			ret := &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint64"],
					Variadic: false,
				},
				Val: -int64(operand.GetConstantValue().(uint64)),
			}
			return ret
		case common.Int64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int64"],
					Variadic: false,
				},
				Val: -operand.GetConstantValue().(int64),
			}
		case common.Float64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float64"],
					Variadic: false,
				},
				Val: -operand.GetConstantValue().(float64),
			}
		}
	}
	panic("unknown left unary")
}

func MergeBinary(op string, lhs IConstantNode, rhs IConstantNode) IConstantNode {
	switch op {
	case "+":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint32) + rhsConvertFunc(rhs.GetConstantValue()).(uint32),
			}
		case common.UInt64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint64) + rhsConvertFunc(rhs.GetConstantValue()).(uint64),
			}
		case common.Int32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int32) + rhsConvertFunc(rhs.GetConstantValue()).(int32),
			}
		case common.Int64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int64) + rhsConvertFunc(rhs.GetConstantValue()).(int64),
			}
		case common.Float32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(float32) + rhsConvertFunc(rhs.GetConstantValue()).(float32),
			}
		case common.Float64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(float64) + rhsConvertFunc(rhs.GetConstantValue()).(float64),
			}
		case common.String:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["string"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(string) + rhsConvertFunc(rhs.GetConstantValue()).(string),
			}
		default:
			panic("not support + : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "-":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint32) - rhsConvertFunc(rhs.GetConstantValue()).(uint32),
			}
		case common.UInt64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint64) - rhsConvertFunc(rhs.GetConstantValue()).(uint64),
			}
		case common.Int32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int32) - rhsConvertFunc(rhs.GetConstantValue()).(int32),
			}
		case common.Int64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int64) - rhsConvertFunc(rhs.GetConstantValue()).(int64),
			}
		case common.Float32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(float32) - rhsConvertFunc(rhs.GetConstantValue()).(float32),
			}
		case common.Float64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(float64) - rhsConvertFunc(rhs.GetConstantValue()).(float64),
			}
		default:
			panic("not support - : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "*":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint32) * rhsConvertFunc(rhs.GetConstantValue()).(uint32),
			}
		case common.UInt64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint64) * rhsConvertFunc(rhs.GetConstantValue()).(uint64),
			}
		case common.Int32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int32) * rhsConvertFunc(rhs.GetConstantValue()).(int32),
			}
		case common.Int64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int64) * rhsConvertFunc(rhs.GetConstantValue()).(int64),
			}
		case common.Float32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(float32) * rhsConvertFunc(rhs.GetConstantValue()).(float32),
			}
		case common.Float64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(float64) * rhsConvertFunc(rhs.GetConstantValue()).(float64),
			}
		default:
			panic("not support * : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "/":
		t := common.BasicTypeMap["float32"]
		if lhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = lhs.GetDataType()
		}
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		if t.Kind.Kind > common.Float64 {
			panic("not support type for division: " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.Float32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(float32) / rhsConvertFunc(rhs.GetConstantValue()).(float32),
			}
		case common.Float64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(float64) / rhsConvertFunc(rhs.GetConstantValue()).(float64),
			}
		default:
			panic("not support / : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "//":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		if t.Kind.Kind > common.Int64 {
			panic("not support type for integer division: " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint32) / rhsConvertFunc(rhs.GetConstantValue()).(uint32),
			}
		case common.UInt64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint64) / rhsConvertFunc(rhs.GetConstantValue()).(uint64),
			}
		case common.Int32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int32) / rhsConvertFunc(rhs.GetConstantValue()).(int32),
			}
		case common.Int64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int64) / rhsConvertFunc(rhs.GetConstantValue()).(int64),
			}
		default:
			panic("not support // : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "%":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		if t.Kind.Kind > common.Int64 {
			panic("not support type for integer mod: " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint32) % rhsConvertFunc(rhs.GetConstantValue()).(uint32),
			}
		case common.UInt64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["uint64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(uint64) % rhsConvertFunc(rhs.GetConstantValue()).(uint64),
			}
		case common.Int32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int32"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int32) % rhsConvertFunc(rhs.GetConstantValue()).(int32),
			}
		case common.Int64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int64"],
					Variadic: false,
				},
				Val: lhsConvertFunc(lhs.GetConstantValue()).(int64) % rhsConvertFunc(rhs.GetConstantValue()).(int64),
			}
		default:
			panic("not support % : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "**":
		t := common.BasicTypeMap["int32"]
		if lhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		if t.Kind.Kind > common.Float64 {
			panic("not support type for division: " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.Int32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int32"],
					Variadic: false,
				},
				Val: lambda_chains.IPow32(lhsConvertFunc(lhs.GetConstantValue()).(int32), rhsConvertFunc(rhs.GetConstantValue()).(int32)),
			}
		case common.Int64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["int64"],
					Variadic: false,
				},
				Val: lambda_chains.IPow64(lhsConvertFunc(lhs.GetConstantValue()).(int64), rhsConvertFunc(rhs.GetConstantValue()).(int64)),
			}
		case common.Float32:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float32"],
					Variadic: false,
				},
				Val: float32(math.Pow(float64(lhsConvertFunc(lhs.GetConstantValue()).(float32)), float64(rhsConvertFunc(rhs.GetConstantValue()).(float32)))),
			}
		case common.Float64:
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["float64"],
					Variadic: false,
				},
				Val: math.Pow(lhsConvertFunc(lhs.GetConstantValue()).(float64), rhsConvertFunc(rhs.GetConstantValue()).(float64)),
			}
		default:
			panic("not support ** : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "==":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint32) == rhsConvertFunc(rhs.GetConstantValue()).(uint32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.UInt64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint64) == rhsConvertFunc(rhs.GetConstantValue()).(uint64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int32) == rhsConvertFunc(rhs.GetConstantValue()).(int32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int64) == rhsConvertFunc(rhs.GetConstantValue()).(int64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float32) == rhsConvertFunc(rhs.GetConstantValue()).(float32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float64) == rhsConvertFunc(rhs.GetConstantValue()).(float64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Bool:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(bool) == rhsConvertFunc(rhs.GetConstantValue()).(bool)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.String:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(string) == rhsConvertFunc(rhs.GetConstantValue()).(string)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.ReflectType:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(*common.DataType).Type == rhsConvertFunc(rhs.GetConstantValue()).(*common.DataType).Type
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		default:
			panic("not support == : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "!=":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint32) != rhsConvertFunc(rhs.GetConstantValue()).(uint32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.UInt64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint64) != rhsConvertFunc(rhs.GetConstantValue()).(uint64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int32) != rhsConvertFunc(rhs.GetConstantValue()).(int32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int64) != rhsConvertFunc(rhs.GetConstantValue()).(int64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float32) != rhsConvertFunc(rhs.GetConstantValue()).(float32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float64) != rhsConvertFunc(rhs.GetConstantValue()).(float64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Bool:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(bool) != rhsConvertFunc(rhs.GetConstantValue()).(bool)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.String:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(string) != rhsConvertFunc(rhs.GetConstantValue()).(string)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		default:
			panic("not support != : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case ">":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint32) > rhsConvertFunc(rhs.GetConstantValue()).(uint32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.UInt64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint64) > rhsConvertFunc(rhs.GetConstantValue()).(uint64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int32) > rhsConvertFunc(rhs.GetConstantValue()).(int32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int64) > rhsConvertFunc(rhs.GetConstantValue()).(int64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float32) > rhsConvertFunc(rhs.GetConstantValue()).(float32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float64) > rhsConvertFunc(rhs.GetConstantValue()).(float64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.String:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(string) > rhsConvertFunc(rhs.GetConstantValue()).(string)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		default:
			panic("not support > : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "<":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint32) < rhsConvertFunc(rhs.GetConstantValue()).(uint32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.UInt64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint64) < rhsConvertFunc(rhs.GetConstantValue()).(uint64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int32) < rhsConvertFunc(rhs.GetConstantValue()).(int32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int64) < rhsConvertFunc(rhs.GetConstantValue()).(int64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float32) < rhsConvertFunc(rhs.GetConstantValue()).(float32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float64) < rhsConvertFunc(rhs.GetConstantValue()).(float64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.String:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(string) < rhsConvertFunc(rhs.GetConstantValue()).(string)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		default:
			panic("not support < : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case ">=":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint32) >= rhsConvertFunc(rhs.GetConstantValue()).(uint32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.UInt64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint64) >= rhsConvertFunc(rhs.GetConstantValue()).(uint64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int32) >= rhsConvertFunc(rhs.GetConstantValue()).(int32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int64) >= rhsConvertFunc(rhs.GetConstantValue()).(int64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float32) >= rhsConvertFunc(rhs.GetConstantValue()).(float32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float64) >= rhsConvertFunc(rhs.GetConstantValue()).(float64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.String:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(string) >= rhsConvertFunc(rhs.GetConstantValue()).(string)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		default:
			panic("not support >= : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "<=":
		t := lhs.GetDataType()
		if rhs.GetDataType().Kind.Kind > t.Kind.Kind {
			t = rhs.GetDataType()
		}
		lhsConvertFunc := lambda_chains.GetConvertFunc(lhs.GetDataType(), t)
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), t)
		switch t.Kind.Kind {
		case common.UInt32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint32) <= rhsConvertFunc(rhs.GetConstantValue()).(uint32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.UInt64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(uint64) <= rhsConvertFunc(rhs.GetConstantValue()).(uint64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int32) <= rhsConvertFunc(rhs.GetConstantValue()).(int32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Int64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(int64) <= rhsConvertFunc(rhs.GetConstantValue()).(int64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float32:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float32) <= rhsConvertFunc(rhs.GetConstantValue()).(float32)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.Float64:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(float64) <= rhsConvertFunc(rhs.GetConstantValue()).(float64)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		case common.String:
			var res interface{} = lhsConvertFunc(lhs.GetConstantValue()).(string) <= rhsConvertFunc(rhs.GetConstantValue()).(string)
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["bool"],
					Variadic: false,
				},
				Val: res,
			}
		default:
			panic("not support * : " + lhs.GetDataType().Type + " " + rhs.GetDataType().Type)
		}
	case "=":
		rhsConvertFunc := lambda_chains.GetConvertFunc(rhs.GetDataType(), lhs.GetDataType())
		var res interface{} = rhsConvertFunc(rhs.GetConstantValue())
		return &ValueNode{
			Node: Node{
				NodeType: "ValueNode",
				DataType: common.BasicTypeMap["bool"],
				Variadic: false,
			},
			Val: res,
		}
	case "&&":
		var res interface{} = lhs.GetConstantValue().(bool) && rhs.GetConstantValue().(bool)
		return &ValueNode{
			Node: Node{
				NodeType: "ValueNode",
				DataType: common.BasicTypeMap["bool"],
				Variadic: false,
			},
			Val: res,
		}
	case "||":
		var res interface{} = lhs.GetConstantValue().(bool) || rhs.GetConstantValue().(bool)
		return &ValueNode{
			Node: Node{
				NodeType: "ValueNode",
				DataType: common.BasicTypeMap["bool"],
				Variadic: false,
			},
			Val: res,
		}
	case "~=":
		// todo
		panic("todo")
	default:
		panic("unknown op " + op)
	}
}

func MergeMapIndex(m IConstantNode, key IConstantNode) IConstantNode {
	keyConvertFunc := lambda_chains.GetConvertFunc(key.GetDataType(), m.GetDataType().KeyType)
	switch m.GetDataType().KeyType.Kind.Kind {
	case common.UInt32, common.Int32, common.Float32:
		mapVal := m.GetConstantValue()
		if mapVal == nil {
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["nil"],
					Variadic: false,
				},
				Val: nil,
			}
		} else {
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: m.GetDataType().ValueType,
					Variadic: false,
				},
				Val: lambda_chains.GetMap32Field(mapVal, keyConvertFunc(key.GetConstantValue())),
			}
		}
	case common.UInt64, common.Int64, common.Float64:
		mapVal := m.GetConstantValue()
		if mapVal == nil {
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["nil"],
					Variadic: false,
				},
				Val: nil,
			}
		} else {
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: m.GetDataType().ValueType,
					Variadic: false,
				},
				Val: lambda_chains.GetMap64Field(mapVal, keyConvertFunc(key.GetConstantValue())),
			}
		}
	case common.String:
		mapVal := m.GetConstantValue()
		if mapVal == nil {
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["nil"],
					Variadic: false,
				},
				Val: nil,
			}
		} else {
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: m.GetDataType().ValueType,
					Variadic: false,
				},
				Val: lambda_chains.GetMapStrField(mapVal, keyConvertFunc(key.GetConstantValue())),
			}
		}
	case common.Bool:
		mapVal := m.GetConstantValue()
		if mapVal == nil {
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: common.BasicTypeMap["nil"],
					Variadic: false,
				},
				Val: nil,
			}
		} else {
			return &ValueNode{
				Node: Node{
					NodeType: "ValueNode",
					DataType: m.GetDataType().ValueType,
					Variadic: false,
				},
				Val: lambda_chains.GetMapBoolField(mapVal, keyConvertFunc(key.GetConstantValue())),
			}
		}
	default:
		panic("unsupported key type")
	}
}

func MergeSliceIndex(slice IConstantNode, index IConstantNode) IConstantNode {
	sliceVal := slice.GetConstantValue()
	indexVal := index.GetConstantValue()
	i64ConvertFunc := lambda_chains.GetConvertFunc(index.GetDataType(), common.BasicTypeMap["int64"])
	return &ValueNode{
		Node: Node{
			NodeType: "ValueNode",
			DataType: slice.GetDataType().ItemType,
			Variadic: false,
		},
		Val: lambda_chains.GetSliceField(sliceVal, i64ConvertFunc(indexVal).(int64)),
	}
}

func MergeStringIndex(stringNode IConstantNode, index IConstantNode) IConstantNode {
	str := stringNode.GetConstantValue().(string)
	indexVal := index.GetConstantValue()

	i64ConvertFunc := lambda_chains.GetConvertFunc(index.GetDataType(), common.BasicTypeMap["int64"])
	return &ValueNode{
		Node: Node{
			NodeType: "ValueNode",
			DataType: common.BasicTypeMap["string"],
			Variadic: false,
		},
		Val: string(str[i64ConvertFunc(indexVal).(int64)]),
	}
}

func MergeMapMultiIndex(m IConstantNode, keys []ASTNode) IConstantNode {
	mapData := m.GetConstantValue()
	var val []interface{}
	switch m.GetDataType().KeyType.Kind.Kind {
	case common.UInt32, common.Int32, common.Float32:
		for _, key := range keys {
			keyConvertFunc := lambda_chains.GetConvertFunc(key.GetDataType(), m.GetDataType().KeyType)
			val = append(val, lambda_chains.GetMap32Field(mapData, keyConvertFunc(key.(IConstantNode).GetConstantValue())))
		}
	case common.UInt64, common.Int64, common.Float64:
		for _, key := range keys {
			keyConvertFunc := lambda_chains.GetConvertFunc(key.GetDataType(), m.GetDataType().KeyType)
			val = append(val, lambda_chains.GetMap64Field(mapData, keyConvertFunc(key.(IConstantNode).GetConstantValue())))
		}
	case common.String:
		for _, key := range keys {
			keyConvertFunc := lambda_chains.GetConvertFunc(key.GetDataType(), m.GetDataType().KeyType)
			val = append(val, lambda_chains.GetMapStrField(mapData, keyConvertFunc(key.(IConstantNode).GetConstantValue())))
		}
	case common.Bool:
		for _, key := range keys {
			keyConvertFunc := lambda_chains.GetConvertFunc(key.GetDataType(), m.GetDataType().KeyType)
			val = append(val, lambda_chains.GetMapBoolField(mapData, keyConvertFunc(key.(IConstantNode).GetConstantValue())))
		}
	}
	return &ValueNode{
		Node: Node{
			NodeType: "ValueNode",
			Variadic: false,
		},
		Val: val,
	}
}

func MergeSliceMultiIndex(slice IConstantNode, keys []ASTNode) IConstantNode {
	if slice.GetDataType().Kind.Kind == common.Slice {
		sliceData := slice.GetConstantValue()
		var val []interface{}
		for _, key := range keys {
			fromToStep := key.(IConstantNode).GetConstantValue().([]int64)
			for i := fromToStep[0]; i < fromToStep[1]; i += fromToStep[2] {
				val = append(val, lambda_chains.GetMapBoolField(sliceData, i))
			}
		}
		return &ValueNode{
			Node: Node{
				NodeType: "ValueNode",
				DataType: slice.GetDataType(),
				Variadic: false,
			},
			Val: val,
		}
	} else if slice.GetDataType().Kind.Kind == common.String {
		sliceData := slice.GetConstantValue().(string)
		key := keys[0]
		fromToStep := key.(IConstantNode).GetConstantValue().([]int64)
		return &ValueNode{
			Node: Node{
				NodeType: "ValueNode",
				DataType: common.BasicTypeMap["string"],
				Variadic: false,
			},
			Val: sliceData[fromToStep[0]:fromToStep[1]],
		}
	} else {
		panic("wtf")
	}
}
