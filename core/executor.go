package core

import (
	"fmt"
	"github.com/ToolPackage/pipe/commands"
	"github.com/ToolPackage/pipe/parser"
	"github.com/ToolPackage/pipe/util"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strings"
)

func Execute() {
	cmds := parseCommands()

	in := NewDualChannel()
	out := NewDualChannel()
	for _, cmd := range cmds {
		if err := cmd.Exec(in, out); err != nil {
			if err == commands.NotEnoughParameterError {
				fmt.Print("not enough parameters, usage: ", util.FuncDescription(cmd.Handler))
			} else {
				fmt.Println(err.Error())
			}
			os.Exit(1)
		}

		in, out = out, in
		out.Reset()
	}
}

func parseCommands() []commands.Command {
	cmds := strings.Join(os.Args[1:], "")
	input := antlr.NewInputStream(cmds)
	lexer := parser.NewPipeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPipeParser(stream)
	p.AddErrorListener(NewErrorListener()) // default is console error listener
	p.BuildParseTrees = true
	tree := p.Commands()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	for idx := range listener.commands {
		handler, err := getCommandHandler(listener.commands[idx].Name)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		listener.commands[idx].Handler = handler
	}

	return listener.commands
}

type ErrorListener struct {
	*antlr.DiagnosticErrorListener
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{antlr.NewDiagnosticErrorListener(true)}
}

func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	el.DiagnosticErrorListener.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
	os.Exit(1)
}

type TreeShapeListener struct {
	*parser.BasePipeListener

	commands []commands.Command
	command  *commands.Command
	param    *commands.CommandParameter
}

func NewTreeShapeListener() *TreeShapeListener {
	return &TreeShapeListener{commands: make([]commands.Command, 0)}
}

func (t *TreeShapeListener) EnterPattern(ctx *parser.PatternContext) {
	t.command = new(commands.Command)
}

func (t *TreeShapeListener) ExitPattern(ctx *parser.PatternContext) {
	t.commands = append(t.commands, *t.command)
	t.command = nil
}

func (t *TreeShapeListener) EnterCommandName(ctx *parser.CommandNameContext) {
	t.command.Name = ctx.GetText()
}

func (t *TreeShapeListener) EnterCommandArguments(ctx *parser.CommandArgumentsContext) {
	t.command.Params = make([]commands.CommandParameter, 0)
}

func (t *TreeShapeListener) EnterCommandArgument(ctx *parser.CommandArgumentContext) {
	t.param = new(commands.CommandParameter)
}

func (t *TreeShapeListener) ExitCommandArgument(ctx *parser.CommandArgumentContext) {
	t.command.Params = append(t.command.Params, *t.param)
	t.param = nil
}

func (t *TreeShapeListener) EnterCommandArgumentLabel(ctx *parser.CommandArgumentLabelContext) {
	labelWithColon := ctx.GetText()
	t.param.Label = labelWithColon[:len(labelWithColon)-1]
}

func (t *TreeShapeListener) EnterNumberValue(ctx *parser.NumberValueContext) {
	t.param.ValueType = commands.DecimalValue
	t.param.Value = ctx.GetText()
}

func (t *TreeShapeListener) EnterStringValue(ctx *parser.StringValueContext) {
	t.param.ValueType = commands.StringValue
	t.param.Value = ctx.GetText()
}

func (t *TreeShapeListener) EnterBooleanValue(ctx *parser.BooleanValueContext) {
	t.param.ValueType = commands.BoolValue
	t.param.Value = ctx.GetText()
}
