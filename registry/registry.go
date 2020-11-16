package registry

import (
	"fmt"
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/ToolPackage/pipe/util"
	"strings"
)

const commandPathSeparator = "."

var commandHandlerTree = &TreeNode{children: make([]*TreeNode, 0)}

func RegisterFunctions(funcList []*FunctionDefinition) {
	for _, function := range funcList {
		RegisterFunction(function)
	}
}

func RegisterFunction(function *FunctionDefinition) {
	patterns := strings.Split(function.Name, commandPathSeparator)
	patternIdx := 0

	root := commandHandlerTree
	for patternIdx < len(patterns) {
		if node, ok := root.matchChildren(patterns[patternIdx]); ok {
			root = node
			patternIdx++
		} else {
			for ; patternIdx < len(patterns); patternIdx++ {
				node := &TreeNode{
					value:    patterns[patternIdx],
					funcList: make([]*FunctionDefinition, 0),
					children: make([]*TreeNode, 0),
				}
				root.children = append(root.children, node)
				root = node
			}
			root.funcList = append(root.funcList, function)
			return
		}
	}
}

func GetFunction(funcName string) ([]*FunctionDefinition, error) {
	patterns := strings.Split(funcName, commandPathSeparator)
	patternIdx := 0

	root := commandHandlerTree
	for patternIdx < len(patterns) {
		if node, ok := root.matchChildren(patterns[patternIdx]); ok {
			root = node
			patternIdx++
		} else {
			return nil, fmt.Errorf("function not found: %s, failed to match subname: %s",
				funcName, patterns[patternIdx])
		}
	}

	return root.funcList, nil
}

func PrintFunctionUsages() {
	fmt.Println("usages:")
	for _, child := range commandHandlerTree.children {
		printFuncUsages(1, child)
	}
	fmt.Println("example: pipe run in.text('Hello, world!')=out")
}

func printFuncUsages(indent int, node *TreeNode) {
	fmt.Print(strings.Repeat("  ", indent), node.value)
	if len(node.funcList) > 0 {
		for _, function := range node.funcList {
			usage := strings.Trim(util.FuncDescription(function.Handler), " \n")
			fmt.Println(" ->", usage)
		}
	} else {
		fmt.Println()
	}
	for _, child := range node.children {
		printFuncUsages(indent+1, child)
	}
}

type TreeNode struct {
	value    string
	funcList []*FunctionDefinition
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
