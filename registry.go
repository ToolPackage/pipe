package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"util/base64"
	"util/gzip"
	"util/input"
	"util/json"
	"util/output"
)

const commandPatternSeparator = "/"

type Command struct {
	path    string
	handler CommandHandler
}

func parseCommands() []Command {
	cmdList := make([]Command, 0)
	pattern := strings.Split(strings.Join(os.Args[1:], ""), commandPatternSeparator)
	for _, command := range pattern {
		handler, err := getCommandHandler(command)
		if err != nil {
			panic(err)
		} else {
			cmdList = append(cmdList, Command{path: command, handler: handler})
		}
	}
	return cmdList
}

const commandPathSeparator = "."

var commandHandlers = &TreeNode{children: make([]*TreeNode, 0)}

type CommandHandler func(args []string, in io.Reader, out io.Writer)

func init() {
	_ = registerCommand("base64.encode", base64.Base64Encode)
	_ = registerCommand("base64.decode", base64.Base64Encode)
	_ = registerCommand("gzip.compress", gzip.GzipCompress)
	_ = registerCommand("gzip.decompress", gzip.GzipDecompress)
	_ = registerCommand("json.pretty", json.JsonPretty)
	_ = registerCommand("input", input.Input)
	_ = registerCommand("output", output.Output)
}

func registerCommand(command string, handler CommandHandler) error {
	patterns := strings.Split(command, commandPathSeparator)
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
	return fmt.Errorf("duplicate command [%s]", command)
}

func getCommandHandler(command string) (CommandHandler, error) {
	patterns := strings.Split(command, commandPathSeparator)
	patternIdx := 0

	root := commandHandlers
	for patternIdx < len(patterns) {
		if node, ok := root.matchChildren(patterns[patternIdx]); ok {
			root = node
			patternIdx++
		} else {
			return nil, fmt.Errorf("invalid command pattern [%s]" + patterns[patternIdx])
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
