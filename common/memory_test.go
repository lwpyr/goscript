package common

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA(t *testing.T) {
	alphabet := map[interface{}]interface{}{
		"a": "A",
		"b": "B",
	}
	v, err := jsoniter.Marshal(alphabet)
	fmt.Println(v)
	fmt.Println(err)
}

func TestSliceMemory(t *testing.T) {
	varDef := &Variable{
		Offset: 0,
		Symbol: "test",
		Type:   BasicTypeMap["int64"],
	}
	sliceMem := &Memory{Data: []interface{}{nil}}

	// take value from memory
	val := sliceMem.MustGet(varDef)
	assert.Equal(t, int64(0), *val)

	// modify the value
	*val = int64(1)
	assert.Equal(t, int64(1), *val)

	// memory also modified
	assert.Equal(t, int64(1), sliceMem.Get(varDef))
}
