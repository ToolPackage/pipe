package parser

import (
	"fmt"
	"github.com/ToolPackage/pipe/functions"
	"github.com/ToolPackage/pipe/registry"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
)

func ParseMultiPipe(script string) []functions.FunctionNode {
	input := antlr.NewInputStream(script)
	lexer := NewPipeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewPipeParser(stream)
	p.AddErrorListener(NewErrorListener()) // default is console error listener
	p.BuildParseTrees = true
	tree := p.MultiPipe()
	listener := NewMultiPipeListener()
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

func lookupFunction(f *functions.FunctionNode, funcDefs []*functions.FunctionDefinition) functions.FunctionHandler {
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

type MultiPipeListener struct {
	*BasePipeListener

	functions []functions.FunctionNode

	function           *functions.FunctionNode
	param              *functions.Parameter
	mapEntryLabel      string
	enterMapEntryValue bool
}

func NewMultiPipeListener() *MultiPipeListener {
	return &MultiPipeListener{functions: make([]functions.FunctionNode, 0)}
}

func (mpl *MultiPipeListener) EnterFunction(ctx *FunctionContext) {
	mpl.function = new(functions.FunctionNode)
}

func (mpl *MultiPipeListener) ExitFunction(ctx *FunctionContext) {
	mpl.functions = append(mpl.functions, *mpl.function)
	mpl.function = nil
}

func (mpl *MultiPipeListener) EnterFunctionName(ctx *FunctionNameContext) {
	mpl.function.Name = ctx.GetText()
}

func (mpl *MultiPipeListener) EnterFunctionParameters(ctx *FunctionParametersContext) {
	mpl.function.Params = make([]functions.Parameter, 0)
}

func (mpl *MultiPipeListener) EnterFunctionParameter(ctx *FunctionParameterContext) {
	mpl.param = new(functions.Parameter)
}

func (mpl *MultiPipeListener) ExitFunctionParameter(ctx *FunctionParameterContext) {
	mpl.function.Params = append(mpl.function.Params, *mpl.param)
	mpl.param = nil
}

func (mpl *MultiPipeListener) EnterFunctionParameterLabel(ctx *FunctionParameterLabelContext) {
	labelWithColon := ctx.GetText()
	mpl.param.Label = labelWithColon[:len(labelWithColon)-1]
}

// handle basic type

func (mpl *MultiPipeListener) EnterIntegerValue(ctx *IntegerValueContext) {
	mpl.updateParameterValue(functions.NewBaseParameterValue(functions.IntegerValue, ctx.GetText()))
}

func (mpl *MultiPipeListener) EnterDecimalValue(ctx *DecimalValueContext) {
	mpl.updateParameterValue(functions.NewBaseParameterValue(functions.DecimalValue, ctx.GetText()))
}

func (mpl *MultiPipeListener) EnterStringValue(ctx *StringValueContext) {
	mpl.updateParameterValue(functions.NewBaseParameterValue(functions.StringValue, ctx.GetText()))
}

func (mpl *MultiPipeListener) EnterBooleanValue(ctx *BooleanValueContext) {
	mpl.updateParameterValue(functions.NewBaseParameterValue(functions.BoolValue, ctx.GetText()))
}

func (mpl *MultiPipeListener) updateParameterValue(newItem functions.ParameterValue) {
	if mpl.enterMapEntryValue {
		v := mpl.param.Value.(*functions.DictParameterValue)
		v.AddEntry(mpl.mapEntryLabel, newItem)
		mpl.mapEntryLabel = ""
	} else {
		mpl.param.Value = newItem
	}
}

// handle dict

func (mpl *MultiPipeListener) EnterDictValue(ctx *DictValueContext) {
	mpl.param.Value = functions.NewDictParameterValue()
}

func (mpl *MultiPipeListener) EnterDictEntryLabel(c *DictEntryLabelContext) {
	mpl.mapEntryLabel = c.GetText()
}

func (mpl *MultiPipeListener) EnterDictEntryValue(c *DictEntryValueContext) {
	mpl.enterMapEntryValue = true
}

func (mpl *MultiPipeListener) ExitDictEntryValue(c *DictEntryValueContext) {
	mpl.enterMapEntryValue = false
}
