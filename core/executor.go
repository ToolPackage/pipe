package core

import (
	"github.com/ToolPackage/pipe/command"
	"github.com/ToolPackage/pipe/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strings"
)

func Execute() {
	commands := parseCommands()

	in := NewDualChannel()
	out := NewDualChannel()

	for _, cmd := range commands {
		cmd.Exec(in, out)
		in, out = out, in
		out.Reset()
	}
}

func parseCommands() []command.Command {
	commands := strings.Join(os.Args[1:], "")
	input := antlr.NewInputStream(commands)
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
			panic(err)
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

	commands []command.Command
	command  *command.Command
	param    *command.CommandParameter
}

func NewTreeShapeListener() *TreeShapeListener {
	return &TreeShapeListener{commands: make([]command.Command, 0)}
}

func (t *TreeShapeListener) EnterPattern(ctx *parser.PatternContext) {
	t.command = new(command.Command)
}

func (t *TreeShapeListener) ExitPattern(ctx *parser.PatternContext) {
	t.commands = append(t.commands, *t.command)
	t.command = nil
}

func (t *TreeShapeListener) EnterCommandName(ctx *parser.CommandNameContext) {
	t.command.Name = ctx.GetText()
}

func (t *TreeShapeListener) EnterCommandArguments(ctx *parser.CommandArgumentsContext) {
	t.command.Params = make([]command.CommandParameter, 0)
}

func (t *TreeShapeListener) EnterCommandArgument(ctx *parser.CommandArgumentContext) {
	t.param = new(command.CommandParameter)
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
	t.param.ValueType = command.DecimalValue
	t.param.Value = ctx.GetText()
}

func (t *TreeShapeListener) EnterStringValue(ctx *parser.StringValueContext) {
	t.param.ValueType = command.StringValue
	t.param.Value = ctx.GetText()
}

func (t *TreeShapeListener) EnterBooleanValue(ctx *parser.BooleanValueContext) {
	t.param.ValueType = command.BoolValue
	t.param.Value = ctx.GetText()
}
