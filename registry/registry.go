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
		if node, ok := root.matchChild(patterns[patternIdx]); ok {
			root = node
			patternIdx++
		} else {
			// create empty nodes
			for ; patternIdx < len(patterns); patternIdx++ {
				node := &TreeNode{
					value:    patterns[patternIdx],
					funcList: make([]*FunctionDefinition, 0),
					children: make([]*TreeNode, 0),
				}
				root.children = append(root.children, node)
				root = node
			}
			break
		}
	}
	root.funcList = append(root.funcList, function)
}

func GetFunction(funcName string) ([]*FunctionDefinition, error) {
	patterns := strings.Split(funcName, commandPathSeparator)
	patternIdx := 0

	root := commandHandlerTree
	for patternIdx < len(patterns) {
		nodes := root.matchChildren(patterns[patternIdx])
		if len(nodes) > 1 {
			return nil, fmt.Errorf("lookup function error: %s, multiple results found while matching subname: %s",
				funcName, patterns[patternIdx])
		}
		if len(nodes) == 0 {
			return nil, fmt.Errorf("function not found: %s, failed to match subname: %s",
				funcName, patterns[patternIdx])
		}
		root = nodes[0]
		patternIdx++
	}

	if len(root.funcList) == 0 {
		return nil, fmt.Errorf("function not found: %s, incomplete function name", funcName)
	}
	return root.funcList, nil
}

func PrintFunctionUsage(funcName string) error {
	funcDefs, err := GetFunction(funcName)
	if err != nil {
		return err
	}

	for _, funcDef := range funcDefs {
		desc := util.FuncDescription(funcDef.Handler)
		lines := strings.Split(desc, "\n")
		fmt.Println("->", lines[0])
		for i := 1; i < len(lines); i++ {
			fmt.Println("  ", lines[i])
		}
	}

	return nil
}

func PrintFunctionUsages() {
	fmt.Println("usages:")
	for _, child := range commandHandlerTree.children {
		printFuncUsages(1, child)
	}
	fmt.Println("example: pipe run in.text('Hello, world!')=out")
	fmt.Println("or like: pipe run i.t('Hello, world!')=o")
}

func printFuncUsages(indent int, node *TreeNode) {
	fmt.Print(strings.Repeat(" ", indent*2), node.value)
	if len(node.funcList) > 0 {
		prompt := " ->"
		spacing := strings.Repeat(" ", indent*2+len(node.value)) + "  +"
		for _, function := range node.funcList {
			fmt.Println(prompt, util.SimpleFuncDescription(function.Handler))
			prompt = spacing
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

func (n *TreeNode) matchChild(pattern string) (*TreeNode, bool) {
	for _, node := range n.children {
		if node.value == pattern {
			return node, true
		}
	}
	return nil, false
}

func (n *TreeNode) matchChildren(pattern string) []*TreeNode {
	result := make([]*TreeNode, 0)
	for _, node := range n.children {
		if node.value == pattern {
			result = append(result, node)
		} else if strings.HasPrefix(node.value, pattern) {
			result = append(result, node)
		}
	}
	return result
}
