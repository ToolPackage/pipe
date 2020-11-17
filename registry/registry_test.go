package registry

import (
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestRegistry(t *testing.T) {
	RegisterFunction(DefFunc("a1.b2.c3.d4",
		func(args Parameters, in io.Reader, out io.Writer) error { return nil },
		DefParams(),
	))

	node := commandHandlerTree
	assert.True(t, len(node.children) == 1)
	assert.Equal(t, "a1", node.children[0].value)
	node = node.children[0]
	assert.True(t, len(node.children) == 1)
	assert.Equal(t, "b2", node.children[0].value)
	node = node.children[0]
	assert.True(t, len(node.children) == 1)
	assert.Equal(t, "c3", node.children[0].value)
	node = node.children[0]
	assert.True(t, len(node.children) == 1)
	assert.Equal(t, "d4", node.children[0].value)
	node = node.children[0]

	assert.NotNil(t, node.funcList[0].Handler)

	handler, err := GetFunction("a.b.c.d")
	assert.Nil(t, err)
	assert.NotNil(t, handler)
}
