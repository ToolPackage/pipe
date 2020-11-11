package core

import (
	"fmt"
	"github.com/ToolPackage/pipe/command"
	"strings"
)

const commandPathSeparator = "."

var commandHandlers = &TreeNode{children: make([]*TreeNode, 0)}

func RegisterCommand(commandName string, handler command.CommandHandler) error {
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

func getCommandHandler(commandName string) (command.CommandHandler, error) {
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
	handler  command.CommandHandler
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
