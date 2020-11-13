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
	tree := p.Script()
	listener := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	for idx := range listener.functions {
		f := &listener.functions[idx]
		funcDefs, err := registry.GetFunction(f.Name)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		f.Handler = lookupFunction(f, funcDefs)
	}

	return listener.functions
}

func lookupFunction(f *functions.Function, funcDefs []*functions.FunctionDefinition) functions.FunctionHandler {
	// TODO:
	funcDef := funcDefs[0]
	if err := funcDef.ParamConstraint.Validate(f.Params); err != nil {
		fmt.Printf("validation error %s: %s\n", f.Name, err.Error())
		os.Exit(1)
	}
	return funcDefs[0].Handler
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

	function           *functions.Function
	param              *functions.Parameter
	mapEntryLabel      string
	enterMapEntryValue bool
}

func NewTreeShapeListener() *TreeShapeListener {
	return &TreeShapeListener{functions: make([]functions.Function, 0)}
}

func (t *TreeShapeListener) EnterFunction(ctx *FunctionContext) {
	t.function = new(functions.Function)
}

func (t *TreeShapeListener) ExitFunction(ctx *FunctionContext) {
	t.functions = append(t.functions, *t.function)
	t.function = nil
}

func (t *TreeShapeListener) EnterFunctionName(ctx *FunctionNameContext) {
	t.function.Name = ctx.GetText()
}

func (t *TreeShapeListener) EnterFunctionParameters(ctx *FunctionParametersContext) {
	t.function.Params = make([]functions.Parameter, 0)
}

func (t *TreeShapeListener) EnterFunctionParameter(ctx *FunctionParameterContext) {
	t.param = new(functions.Parameter)
}

func (t *TreeShapeListener) ExitFunctionParameter(ctx *FunctionParameterContext) {
	t.function.Params = append(t.function.Params, *t.param)
	t.param = nil
}

func (t *TreeShapeListener) EnterFunctionParameterLabel(ctx *FunctionParameterLabelContext) {
	labelWithColon := ctx.GetText()
	t.param.Label = labelWithColon[:len(labelWithColon)-1]
}

// handle basic type

func (t *TreeShapeListener) EnterIntegerValue(ctx *IntegerValueContext) {
	t.updateParameterValue(functions.NewBaseParameterValue(functions.IntegerValue, ctx.GetText()))
}

func (t *TreeShapeListener) EnterDecimalValue(ctx *DecimalValueContext) {
	t.updateParameterValue(functions.NewBaseParameterValue(functions.DecimalValue, ctx.GetText()))
}

func (t *TreeShapeListener) EnterStringValue(ctx *StringValueContext) {
	t.updateParameterValue(functions.NewBaseParameterValue(functions.StringValue, ctx.GetText()))
}

func (t *TreeShapeListener) EnterBooleanValue(ctx *BooleanValueContext) {
	t.updateParameterValue(functions.NewBaseParameterValue(functions.BoolValue, ctx.GetText()))
}

func (t *TreeShapeListener) updateParameterValue(newItem functions.ParameterValue) {
	if t.enterMapEntryValue {
		v := t.param.Value.(*functions.DictParameterValue)
		v.AddEntry(t.mapEntryLabel, newItem)
		t.mapEntryLabel = ""
	} else {
		t.param.Value = newItem
	}
}

// handle dict

func (t *TreeShapeListener) EnterDictValue(ctx *DictValueContext) {
	t.param.Value = functions.NewDictParameterValue()
}

func (t *TreeShapeListener) EnterDictEntryLabel(c *DictEntryLabelContext) {
	t.mapEntryLabel = c.GetText()
}

func (t *TreeShapeListener) EnterDictEntryValue(c *DictEntryValueContext) {
	t.enterMapEntryValue = true
}

func (t *TreeShapeListener) ExitDictEntryValue(c *DictEntryValueContext) {
	t.enterMapEntryValue = false
}
