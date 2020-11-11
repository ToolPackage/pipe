package parser

import (
	"fmt"
	"github.com/ToolPackage/pipe/commands"
	"github.com/ToolPackage/pipe/registry"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
)

func ParseCommands(commands string) []commands.Command {
	input := antlr.NewInputStream(commands)
	lexer := NewPipeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewPipeParser(stream)
	p.AddErrorListener(NewErrorListener()) // default is console error listener
	p.BuildParseTrees = true
	tree := p.Commands()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	for idx := range listener.commands {
		handler, err := registry.GetCommandHandler(listener.commands[idx].Name)
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
	*BasePipeListener

	commands []commands.Command
	command  *commands.Command
	param    *commands.CommandParameter
}

func NewTreeShapeListener() *TreeShapeListener {
	return &TreeShapeListener{commands: make([]commands.Command, 0)}
}

func (t *TreeShapeListener) EnterPattern(ctx *PatternContext) {
	t.command = new(commands.Command)
}

func (t *TreeShapeListener) ExitPattern(ctx *PatternContext) {
	t.commands = append(t.commands, *t.command)
	t.command = nil
}

func (t *TreeShapeListener) EnterCommandName(ctx *CommandNameContext) {
	t.command.Name = ctx.GetText()
}

func (t *TreeShapeListener) EnterCommandArguments(ctx *CommandArgumentsContext) {
	t.command.Params = make([]commands.CommandParameter, 0)
}

func (t *TreeShapeListener) EnterCommandArgument(ctx *CommandArgumentContext) {
	t.param = new(commands.CommandParameter)
}

func (t *TreeShapeListener) ExitCommandArgument(ctx *CommandArgumentContext) {
	t.command.Params = append(t.command.Params, *t.param)
	t.param = nil
}

func (t *TreeShapeListener) EnterCommandArgumentLabel(ctx *CommandArgumentLabelContext) {
	labelWithColon := ctx.GetText()
	t.param.Label = labelWithColon[:len(labelWithColon)-1]
}

func (t *TreeShapeListener) EnterNumberValue(ctx *NumberValueContext) {
	t.param.ValueType = commands.DecimalValue
	t.param.Value = ctx.GetText()
}

func (t *TreeShapeListener) EnterStringValue(ctx *StringValueContext) {
	t.param.ValueType = commands.StringValue
	t.param.Value = ctx.GetText()
}

func (t *TreeShapeListener) EnterBooleanValue(ctx *BooleanValueContext) {
	t.param.ValueType = commands.BoolValue
	t.param.Value = ctx.GetText()
}
