package parser

import (
	"fmt"
	"github.com/ToolPackage/pipe/functions"
	"github.com/ToolPackage/pipe/registry"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
)

func ParseScript(script string) []functions.Function {
	input := antlr.NewInputStream(script)
	lexer := NewPipeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewPipeParser(stream)
	p.AddErrorListener(NewErrorListener()) // default is console error listener
	p.BuildParseTrees = true
	tree := p.Commands()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	for idx := range listener.functions {
		handler, err := registry.GetFunctionHandler(listener.functions[idx].Name)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		listener.functions[idx].Handler = handler
	}

	return listener.functions
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

	functions []functions.Function
	function  *functions.Function
	param     *functions.FunctionParameter
}

func NewTreeShapeListener() *TreeShapeListener {
	return &TreeShapeListener{functions: make([]functions.Function, 0)}
}

func (t *TreeShapeListener) EnterPattern(ctx *PatternContext) {
	t.function = new(functions.Function)
}

func (t *TreeShapeListener) ExitPattern(ctx *PatternContext) {
	t.functions = append(t.functions, *t.function)
	t.function = nil
}

func (t *TreeShapeListener) EnterCommandName(ctx *CommandNameContext) {
	t.function.Name = ctx.GetText()
}

func (t *TreeShapeListener) EnterCommandArguments(ctx *CommandArgumentsContext) {
	t.function.Params = make([]functions.FunctionParameter, 0)
}

func (t *TreeShapeListener) EnterCommandArgument(ctx *CommandArgumentContext) {
	t.param = new(functions.FunctionParameter)
}

func (t *TreeShapeListener) ExitCommandArgument(ctx *CommandArgumentContext) {
	t.function.Params = append(t.function.Params, *t.param)
	t.param = nil
}

func (t *TreeShapeListener) EnterCommandArgumentLabel(ctx *CommandArgumentLabelContext) {
	labelWithColon := ctx.GetText()
	t.param.Label = labelWithColon[:len(labelWithColon)-1]
}

func (t *TreeShapeListener) EnterNumberValue(ctx *NumberValueContext) {
	t.param.ValueType = functions.DecimalValue
	t.param.Value = ctx.GetText()
}

func (t *TreeShapeListener) EnterStringValue(ctx *StringValueContext) {
	t.param.ValueType = functions.StringValue
	t.param.Value = ctx.GetText()
}

func (t *TreeShapeListener) EnterBooleanValue(ctx *BooleanValueContext) {
	t.param.ValueType = functions.BoolValue
	t.param.Value = ctx.GetText()
}
