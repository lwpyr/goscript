package hack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceHack(t *testing.T) {
	s := []interface{}{"a", "b"}

	b := SliceIndex(s, 1)
	assert.Equal(t, "b", *b)

	assert.Equal(t, int64(2), SliceLen(s))
}
