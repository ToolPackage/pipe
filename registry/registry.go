package registry

import (
	"fmt"
	"github.com/ToolPackage/pipe/functions"
	"github.com/ToolPackage/pipe/util"
	"os"
	"strings"
)

const commandPathSeparator = "."

var commandHandlerTree = &TreeNode{children: make([]*TreeNode, 0)}

func RegisterFunction(functionName string, handler functions.FunctionHandler) {
	patterns := strings.Split(functionName, commandPathSeparator)
	patternIdx := 0

	root := commandHandlerTree
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

	fmt.Printf("unable to register duplicate command: %s\n", functionName)
	os.Exit(1)
}

func GetFunctionHandler(commandName string) (functions.FunctionHandler, error) {
	patterns := strings.Split(commandName, commandPathSeparator)
	patternIdx := 0

	root := commandHandlerTree
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

func PrintFunctionUsages() {
	fmt.Println("usages:")
	for _, child := range commandHandlerTree.children {
		printFuncUsages(1, child)
	}
}

func printFuncUsages(indent int, node *TreeNode) {
	fmt.Print(strings.Repeat("  ", indent), node.value)
	if node.handler != nil {
		usage := strings.Trim(util.FuncDescription(node.handler), " \n")
		fmt.Println(" ->", usage)
	} else {
		fmt.Println()
	}
	for _, child := range node.children {
		printFuncUsages(indent+1, child)
	}
}

type TreeNode struct {
	value    string
	handler  functions.FunctionHandler
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
