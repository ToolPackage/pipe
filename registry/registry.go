package registry

import (
	"fmt"
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/ToolPackage/pipe/util"
	"os"
	"strings"
)

const commandPathSeparator = "."

var commandHandlerTree = &TreeNode{children: make([]*TreeNode, 0)}

func RegisterFunctions(funcList []*FunctionDefinition) {
	for _, function := range funcList {
		RegisterFunction(function)
	}
}

func RegisterFunction(funcDef *FunctionDefinition) {
	patterns := strings.Split(funcDef.Name, commandPathSeparator)
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
					children: make([]*TreeNode, 0),
				}
				root.children = append(root.children, node)
				root = node
			}
			root.funcDef = funcDef
			return
		}
	}

	fmt.Println("duplicate function name:", funcDef.Name)
	os.Exit(1)
}

func LookupFunctionHandler(funcNode *FunctionNode) {
	funcDef, err := GetFunction(funcNode.Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := funcDef.ParamConstraint.Validate(funcNode.Params); err != nil {
		fmt.Printf("validation error %s: %s\n", funcNode.Name, err)
		fmt.Printf("enter command \"pipe usage [funcName]\" to check function usage")
		os.Exit(1)
	}
	funcNode.Handler = funcDef.Handler
}

func GetFunction(funcName string) (*FunctionDefinition, error) {
	patterns := strings.Split(funcName, commandPathSeparator)
	patternIdx := 0

	root := commandHandlerTree
	for patternIdx < len(patterns) {
		nodes := root.matchChildren(patterns[patternIdx])
		if len(nodes) > 1 {
			result := make([]string, len(nodes))
			for idx, node := range nodes {
				result[idx] = node.value
			}
			return nil, fmt.Errorf("lookup function error: multiple results found for [%s]\n"+
				"target subname: %s(#%d), candidate: %v",
				funcName, patterns[patternIdx], patternIdx+1, strings.Join(result, ", "))
		}
		if len(nodes) == 0 {
			return nil, fmt.Errorf("function not found: %s, failed to match subname: %s",
				funcName, patterns[patternIdx])
		}
		root = nodes[0]
		patternIdx++
	}

	if root.funcDef == nil {
		return nil, fmt.Errorf("function not found: %s, incomplete function name", funcName)
	}
	return root.funcDef, nil
}

func PrintFunctionUsage(funcName string) error {
	funcDef, err := GetFunction(funcName)
	if err != nil {
		return err
	}

	desc := util.FuncDescription(funcDef.Handler)
	lines := strings.Split(desc, "\n")
	fmt.Println("->", lines[0])
	for i := 1; i < len(lines); i++ {
		fmt.Println("  ", lines[i])
	}

	return nil
}

func PrintFunctionUsages() {
	fmt.Println("example: pipe run in.text('Hello, world!')=out")
	fmt.Println("or like: pipe run i.t('Hello, world!')=o")
	fmt.Println("usages:")
	for _, child := range commandHandlerTree.children {
		printFuncUsages(1, child)
	}
}

func printFuncUsages(indent int, node *TreeNode) {
	fmt.Print(strings.Repeat(" ", indent*2), node.value)
	if node.funcDef != nil {
		prompt := " ->"
		fmt.Println(prompt, util.SimpleFuncDescription(node.funcDef.Handler))
	} else {
		fmt.Println()
	}
	for _, child := range node.children {
		printFuncUsages(indent+1, child)
	}
}

type TreeNode struct {
	value    string
	funcDef  *FunctionDefinition
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
