package main

import (
	"fmt"
	"github.com/ToolPackage/pipe/base64"
	"github.com/ToolPackage/pipe/gzip"
	"github.com/ToolPackage/pipe/input"
	"github.com/ToolPackage/pipe/json"
	"github.com/ToolPackage/pipe/output"
	"io"
	"os"
	"regexp"
	"strings"
)

const commandPatternSeparator = "/"

const commandPathSeparator = "."

var commandPattern = regexp.MustCompile("(?P<CMD>[a-zA-Z0-9]+(\\.[a-zA-Z0-9]+)*)(?P<ARGS>\\(.*\\))?")

type Command struct {
	path    string
	args    []string
	handler CommandHandler
}

func parseCommands() []Command {
	cmdList := make([]Command, 0)
	pattern := strings.Split(strings.Join(os.Args[1:], ""), commandPatternSeparator)
	for _, command := range pattern {
		match := commandPattern.FindStringSubmatch(command)
		cmd := match[commandPattern.SubexpIndex("CMD")]
		args := match[commandPattern.SubexpIndex("ARGS")]
		if len(args) > 0 {
			args = args[1 : len(args)-1]
		}
		//fmt.Println(cmd + ":" + args)

		handler, err := getCommandHandler(cmd)
		if err != nil {
			panic(err)
		} else {
			cmdList = append(cmdList, Command{path: command, args: strings.Split(args, ","), handler: handler})
		}
	}
	return cmdList
}

var commandHandlers = &TreeNode{children: make([]*TreeNode, 0)}

type CommandHandler func(args []string, in io.Reader, out io.Writer)

func init() {
	_ = registerCommand("base64.encode", base64.Encode)
	_ = registerCommand("base64.decode", base64.Decode)
	_ = registerCommand("gzip.compress", gzip.Compress)
	_ = registerCommand("gzip.decompress", gzip.Decompress)
	_ = registerCommand("json.pretty", json.Pretty)
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
