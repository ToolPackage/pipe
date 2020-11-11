package core

import (
	"github.com/ToolPackage/pipe/command"
	"io"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestRegistry(t *testing.T) {
	err := RegisterCommand("a.b.c.d", func(args command.CommandParameters, in io.Reader, out io.Writer) {})
	assert.Nil(t, err)

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

	handler, err := getCommandHandler("a.b.c.d")
	assert.Nil(t, err)
	assert.NotNil(t, handler)
}
