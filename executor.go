package main

import (
	"fmt"
	"github.com/ToolPackage/pipe/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strings"
)

func execute() {
	parseCommands()

	//buf := make([]byte, 0)
	//in := bytes.NewReader(buf)
	//out := new(strings.Builder)
	//for _, cmd := range commands {
	//	cmd.handler(cmd.args, in, out)
	//	buf = []byte(out.String())
	//	in = bytes.NewReader(buf)
	//	out = new(strings.Builder)
	//}
}

func parseCommands() []*Command {
	commands := strings.Join(os.Args[1:], "")
	input := antlr.NewInputStream(commands)
	lexer := parser.NewPipeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPipeParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Commands()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	fmt.Println(listener.commands[0])
	return listener.commands
	//for _, command := range pattern {
	//	match := commandPattern.FindStringSubmatch(command)
	//	cmd := match[groupCmd]
	//	argsRaw := match[groupArgs]
	//	var args []string
	//	if len(argsRaw) > 0 {
	//		// remove ()
	//		argsRaw = argsRaw[1 : len(argsRaw)-1]
	//		args = strings.Split(argsRaw, commandArgSeparator)
	//		// trim
	//		for idx, arg := range args {
	//			args[idx] = strings.Trim(arg, " ")
	//		}
	//	}
	//	fmt.Println(cmd, args)
	//
	//	handler, err := getCommandHandler(cmd)
	//	if err != nil {
	//		panic(err)
	//	} else {
	//		cmdList = append(cmdList, Command{name: command, args: args, handler: handler})
	//	}
	//}
}

type TreeShapeListener struct {
	*parser.BasePipeListener

	commands []*Command
	command  *Command
	param    *CommandParameter
}

func NewTreeShapeListener() *TreeShapeListener {
	return &TreeShapeListener{commands: make([]*Command, 0)}
}

func (t *TreeShapeListener) EnterPattern(ctx *parser.PatternContext) {
	t.command = new(Command)
}

func (t *TreeShapeListener) ExitPattern(ctx *parser.PatternContext) {
	t.commands = append(t.commands, t.command)
	t.command = nil
}

func (t *TreeShapeListener) EnterCommandName(ctx *parser.CommandNameContext) {
	t.command.name = ctx.GetText()
}

func (t *TreeShapeListener) EnterCommandArguments(ctx *parser.CommandArgumentsContext) {
	t.command.params = make([]*CommandParameter, 0)
}

func (t *TreeShapeListener) EnterCommandArgument(ctx *parser.CommandArgumentContext) {
	t.param = &CommandParameter{}
}

func (t *TreeShapeListener) ExitCommandArgument(ctx *parser.CommandArgumentContext) {
	t.command.params = append(t.command.params, t.param)
	t.command.params = nil
}

func (t *TreeShapeListener) EnterCommandArgumentLabel(ctx *parser.CommandArgumentLabelContext) {
	t.param.label = ctx.GetText()
}

func (t *TreeShapeListener) EnterNumberValue(ctx *parser.NumberValueContext) {
	t.param.valueType = NumberValue
	t.param.value = ctx.GetText()
}

func (t *TreeShapeListener) EnterStringValue(ctx *parser.StringValueContext) {
	t.param.valueType = StringValue
	t.param.value = ctx.GetText()
}

func (t *TreeShapeListener) EnterBooleanValue(ctx *parser.BooleanValueContext) {
	t.param.valueType = BoolValue
	t.param.value = ctx.GetText()
}
