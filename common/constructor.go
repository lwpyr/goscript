package common

import "reflect"

type Constructor func(...interface{}) interface{}

// constructor
func NewUInt8Value(params ...interface{}) interface{} {
	return uint8(0)
}

func NewInt32Value(params ...interface{}) interface{} {
	return int32(0)
}

func NewInt64Value(params ...interface{}) interface{} {
	return int64(0)
}

func NewUInt32Value(params ...interface{}) interface{} {
	return uint32(0)
}

func NewUInt64Value(params ...interface{}) interface{} {
	return uint64(0)
}

func NewFloat32Value(params ...interface{}) interface{} {
	return float32(0)
}

func NewFloat64Value(params ...interface{}) interface{} {
	return float64(0)
}

func NewBoolValue(params ...interface{}) interface{} {
	return false
}

func NewStringValue(params ...interface{}) interface{} {
	return ""
}

func NewBytesValue(params ...interface{}) interface{} {
	return make([]byte, 0)
}

func NewNilValue(params ...interface{}) interface{} {
	return nil
}

func NewSliceValue(params ...interface{}) interface{} {
	if len(params) == 0 {
		return make([]interface{}, 0)
	} else if len(params) == 1 {
		return make([]interface{}, reflect.ValueOf(params[0]).Int())
	} else {
		return make([]interface{}, reflect.ValueOf(params[0]).Int(), reflect.ValueOf(params[1]).Int())
	}
}

func NewChannelValue(params ...interface{}) interface{} {
	if len(params) == 0 {
		return make(chan interface{})
	}
	return make(chan interface{}, reflect.ValueOf(params[0]).Int())
}

func NewMapValue(params ...interface{}) interface{} {
	return make(map[interface{}]interface{})
}

func NewMapInt32Value(params ...interface{}) interface{} {
	return make(map[int32]interface{})
}

func NewMapInt64Value(params ...interface{}) interface{} {
	return make(map[int64]interface{})
}

func NewMapUInt32Value(params ...interface{}) interface{} {
	return make(map[uint32]interface{})
}

func NewMapUInt64Value(params ...interface{}) interface{} {
	return make(map[uint64]interface{})
}

func NewMapFloat32Value(params ...interface{}) interface{} {
	return make(map[float32]interface{})
}

func NewMapFloat64Value(params ...interface{}) interface{} {
	return make(map[float64]interface{})
}

func NewMapStringValue(params ...interface{}) interface{} {
	return make(map[string]interface{})
}

func NewMapBoolValue(params ...interface{}) interface{} {
	return make(map[bool]interface{})
}

func NewMessageValue(params ...interface{}) interface{} {
	return make(map[string]interface{})
}

func NewAnyValue(params ...interface{}) interface{} {
	return nil
}
