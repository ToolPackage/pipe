package parser

import (
	"fmt"
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/ToolPackage/pipe/registry"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
)

// ParseMultiPipe
func ParseMultiPipe(script string) *MultiPipe {
	input := antlr.NewInputStream(script)
	lexer := NewPipeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewPipeParser(stream)
	p.AddErrorListener(NewErrorListener()) // default is console error listener
	p.BuildParseTrees = true
	tree := p.MultiPipe()
	listener := newMultiPipeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	for _, pipe := range listener.multiPipe.Pipes {
		for _, node := range pipe.Nodes {
			funcNode, ok := node.(*FunctionNode)
			if ok {
				funcDefs, err := registry.GetFunction(funcNode.Name)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				funcNode.Handler = lookupFunction(funcNode, funcDefs)
			}
		}
	}

	return listener.multiPipe
}

func lookupFunction(f *FunctionNode, funcDefs []*FunctionDefinition) FunctionHandler {
	// TODO:
	funcDef := funcDefs[0]
	if err := funcDef.ParamConstraint.Validate(f.Params); err != nil {
		fmt.Printf("validation error %s: %s\n", f.Name, err.Error())
		os.Exit(1)
	}
	return funcDef.Handler
}

// ErrorListener
type ErrorListener struct {
	*antlr.DiagnosticErrorListener
}

// NewErrorListener
func NewErrorListener() *ErrorListener {
	return &ErrorListener{antlr.NewDiagnosticErrorListener(true)}
}

// SyntaxError
func (el *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	el.DiagnosticErrorListener.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
	os.Exit(1)
}

// multiPipeListener
type multiPipeListener struct {
	*BasePipeListener

	multiPipe *MultiPipe

	mapEntryLabel string
}

// NewMultiPipeListener
func newMultiPipeListener() *multiPipeListener {
	return &multiPipeListener{
		multiPipe: &MultiPipe{Pipes: make([]Pipe, 0)},
	}
}

func (mpl *multiPipeListener) lastPipe() *Pipe {
	return &mpl.multiPipe.Pipes[len(mpl.multiPipe.Pipes)-1]
}

func (mpl *multiPipeListener) lastPipeNode() PipeNode {
	lastPipe := mpl.lastPipe()
	return lastPipe.Nodes[len(lastPipe.Nodes)-1]
}

func (mpl *multiPipeListener) lastFunctionNode() *FunctionNode {
	return mpl.lastPipeNode().(*FunctionNode)
}

func (mpl *multiPipeListener) lastFunctionNodeParameter() *Parameter {
	param := mpl.lastFunctionNode().Params
	return &param[len(param)-1]
}

func (mpl *multiPipeListener) EnterPipe(c *PipeContext) {
	mpl.multiPipe.Pipes = append(mpl.multiPipe.Pipes,
		Pipe{Variables: make(map[string]int), Nodes: make([]PipeNode, 0)})
}

// EnterFunction
func (mpl *multiPipeListener) EnterFunctionNode(c *FunctionNodeContext) {
	pipe := mpl.lastPipe()
	pipe.Nodes = append(pipe.Nodes, &FunctionNode{Params: make([]Parameter, 0)})
}

// EnterFunctionName
func (mpl *multiPipeListener) EnterFunctionName(c *FunctionNameContext) {
	mpl.lastFunctionNode().Name = c.GetText()
}

func (mpl *multiPipeListener) EnterFunctionParameter(c *FunctionParameterContext) {
	node := mpl.lastFunctionNode()
	node.Params = append(node.Params, Parameter{})
}

func (mpl *multiPipeListener) EnterFunctionParameterLabel(c *FunctionParameterLabelContext) {
	labelWithColon := c.GetText()
	mpl.lastFunctionNodeParameter().Label = labelWithColon[:len(labelWithColon)-1]
}

// handle basic type

func (mpl *multiPipeListener) EnterIntegerValue(ctx *IntegerValueContext) {
	mpl.updateParameterValue(NewBaseParameterValue(IntegerValue, ctx.GetText()))
}

func (mpl *multiPipeListener) EnterDecimalValue(ctx *DecimalValueContext) {
	mpl.updateParameterValue(NewBaseParameterValue(FloatValue, ctx.GetText()))
}

func (mpl *multiPipeListener) EnterStringValue(ctx *StringValueContext) {
	mpl.updateParameterValue(NewBaseParameterValue(StringValue, ctx.GetText()))
}

func (mpl *multiPipeListener) EnterBooleanValue(ctx *BooleanValueContext) {
	mpl.updateParameterValue(NewBaseParameterValue(BoolValue, ctx.GetText()))
}

func (mpl *multiPipeListener) updateParameterValue(newItem ParameterValue) {
	param := mpl.lastFunctionNodeParameter()
	if param.Value != nil {
		v := param.Value.(*DictParameterValue)
		v.AddEntry(mpl.mapEntryLabel, newItem)
		mpl.mapEntryLabel = ""
	} else {
		param.Value = newItem
	}
}

// handle dict

func (mpl *multiPipeListener) EnterDictValue(ctx *DictValueContext) {
	mpl.lastFunctionNodeParameter().Value = NewDictParameterValue()
}

func (mpl *multiPipeListener) EnterDictEntryLabel(c *DictEntryLabelContext) {
	mpl.mapEntryLabel = c.GetText()
}
