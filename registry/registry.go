package registry

import (
	"fmt"
	"github.com/ToolPackage/pipe/commands"
	"os"
	"strings"
)

const commandPathSeparator = "."

var commandHandlers = &TreeNode{children: make([]*TreeNode, 0)}

func RegisterCommand(commandName string, handler commands.CommandHandler) {
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
			return
		}
	}

	fmt.Printf("unable to register duplicate command: %s\n", commandName)
	os.Exit(1)
}

func GetCommandHandler(commandName string) (commands.CommandHandler, error) {
	patterns := strings.Split(commandName, commandPathSeparator)
	patternIdx := 0

	root := commandHandlers
	for patternIdx < len(patterns) {
		if node, ok := root.matchChildren(patterns[patternIdx]); ok {
			root = node
			patternIdx++
		} else {
			return nil, fmt.Errorf("command not found: %s, failed to match subname: %s", commandName, patterns[patternIdx])
		}
	}

	return root.handler, nil
}

type TreeNode struct {
	value    string
	handler  commands.CommandHandler
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
