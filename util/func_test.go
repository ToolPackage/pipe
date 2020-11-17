package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test123
// asd
func TestFuncDescription(t *testing.T) {
	v := FuncDescription(TestFuncDescription)
	assert.Equal(t, "test123\nasd", v)
	v = SimpleFuncDescription(TestFuncDescription)
	assert.Equal(t, "test123", v)
}
