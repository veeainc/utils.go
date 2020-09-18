package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestIsSlice(t *testing.T) {
	assert.Equal(t, IsSlice([]byte{}), true)
	assert.Equal(t, IsSlice([]interface{}{42}), true)
	assert.Equal(t, IsSlice("45"), false)
	assert.Equal(t, IsSlice(12), false)
}

func TestIsMap(t *testing.T) {
	assert.Equal(t, IsMap(map[string]int{"foo": 12}), true)
	assert.Equal(t, IsMap(map[string]int{}), true)
	assert.Equal(t, IsMap(map[string]interface{}{}), true)
	assert.Equal(t, IsMap(map[int]interface{}{}), true)
	assert.Equal(t, IsMap([]byte{}), false)
	assert.Equal(t, IsMap("[]byte{}"), false)
	assert.Equal(t, IsMap(54), false)
}

func TestHasKey(t *testing.T) {
	assert.Equal(t, HasKey(42, "key"), false)
	assert.Equal(t, HasKey(map[string]interface{}{"value": 12}, "value"), true)
	assert.Equal(t, HasKey(map[string]interface{}{"value": 12}, "noValue"), false)
}


func TestGetKey(t *testing.T) {
	assert.Equal(t, GetKey(42, "key"), nil)
	assert.Equal(t, GetKey(map[string]interface{}{"value": 12}, "value"), 12)
	assert.Equal(t, GetKey(map[string]interface{}{"value": 12}, "noValue"), nil)
}

func TestToString(t *testing.T) {
	assert.Equal(t, ToString("foo"), "foo")
	assert.Equal(t, ToString(45), "45")
}