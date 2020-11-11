package main

import (
	"fmt"
	"github.com/ToolPackage/pipe/base64"
	"github.com/ToolPackage/pipe/gzip"
	"github.com/ToolPackage/pipe/input"
	"github.com/ToolPackage/pipe/json"
	"github.com/ToolPackage/pipe/output"
	"io"
	"strings"
)

func init() {
	_ = registerCommand("base64.encode", base64.Encode)
	_ = registerCommand("base64.decode", base64.Decode)
	_ = registerCommand("gzip.compress", gzip.Compress)
	_ = registerCommand("gzip.decompress", gzip.Decompress)
	_ = registerCommand("json.pretty", json.Pretty)
	_ = registerCommand("input", input.Input)
	_ = registerCommand("output", output.Output)
}

type Command struct {
	name    string
	params  []*CommandParameter
	handler CommandHandler
}

type CommandParameter struct {
	label     string
	valueType CommandParameterType
	value     string
}

const (
	NumberValue = iota
	StringValue
	BoolValue
)

type CommandParameterType int

var commandHandlers = &TreeNode{children: make([]*TreeNode, 0)}

type CommandHandler func(args []string, in io.Reader, out io.Writer)

const commandPathSeparator = "."

func registerCommand(commandName string, handler CommandHandler) error {
	patterns := strings.Split(commandName, commandPathSeparator)
	patternIdx := 0

	root := commandHandlers
	for patternIdx < len(patterns) {
		if node, ok := root.matchChildren(patterns[patternIdx]); ok {
			root = node
			patternIdx++
		} else {
			for ; patternIdx < len(patterns); patternIdx++ {
				node := &TreeNode{value: patterns[patternIdx], children: make([]*TreeNode, 0)}
				root.children = append(root.children, node)
				root = node
			}
			root.handler = handler
			return nil
		}
	}
	return fmt.Errorf("duplicate command [%s]", commandName)
}

func getCommandHandler(commandName string) (CommandHandler, error) {
	patterns := strings.Split(commandName, commandPathSeparator)
	patternIdx := 0

	root := commandHandlers
	for patternIdx < len(patterns) {
		if node, ok := root.matchChildren(patterns[patternIdx]); ok {
			root = node
			patternIdx++
		} else {
			return nil, fmt.Errorf("invalid command pattern [%s]", patterns[patternIdx])
		}
	}

	return root.handler, nil
}

type TreeNode struct {
	value    string
	handler  CommandHandler
	children []*TreeNode
}

func (n *TreeNode) matchChildren(pattern string) (*TreeNode, bool) {
	for _, node := range n.children {
		if node.value == pattern {
			return node, true
		}
	}
	return nil, false
}
