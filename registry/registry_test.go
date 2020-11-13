package registry

import (
	"github.com/ToolPackage/pipe/functions"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestRegistry(t *testing.T) {
	RegisterFunction(functions.DefFunc("a.b.c.d",
		func(args functions.Parameters, in io.Reader, out io.Writer) error { return nil },
		functions.DefParams(),
	))

	node := commandHandlerTree
	assert.True(t, len(node.children) == 1)
	assert.Equal(t, "a", node.children[0].value)
	node = node.children[0]
	assert.True(t, len(node.children) == 1)
	assert.Equal(t, "b", node.children[0].value)
	node = node.children[0]
	assert.True(t, len(node.children) == 1)
	assert.Equal(t, "c", node.children[0].value)
	node = node.children[0]
	assert.True(t, len(node.children) == 1)
	assert.Equal(t, "d", node.children[0].value)
	node = node.children[0]

	assert.NotNil(t, node.funcList[0].Handler)

	handler, err := GetFunction("a.b.c.d")
	assert.Nil(t, err)
	assert.NotNil(t, handler)
}
