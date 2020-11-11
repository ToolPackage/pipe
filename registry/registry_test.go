package registry

import (
	"github.com/ToolPackage/pipe/commands"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestRegistry(t *testing.T) {
	RegisterCommand("a.b.c.d", func(args commands.CommandParameters, in io.Reader, out io.Writer) error { return nil })

	node := commandHandlers
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

	assert.NotNil(t, node.handler)

	handler, err := GetCommandHandler("a.b.c.d")
	assert.Nil(t, err)
	assert.NotNil(t, handler)
}
