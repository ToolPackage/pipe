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
	defer func() {
		if err, ok := recover().(error); ok {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	input := antlr.NewInputStream(script)
	lexer := NewPipeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewPipeParser(stream)
	p.AddErrorListener(NewErrorListener()) // default is console error listener
	p.BuildParseTrees = true
	tree := p.Cmd()
	listener := newMultiPipeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.multiPipe
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

	multiPipe  *MultiPipe
	scopeStack ScopeStack

	mapEntryLabel string
}

// NewMultiPipeListener
func newMultiPipeListener() *multiPipeListener {
	return &multiPipeListener{
		multiPipe: &MultiPipe{Variables: make(map[string]*ImmutableValue), PipeList: make(Pipes, 0)},
	}
}

// help function

func (m *multiPipeListener) lastPipe() *Pipe {
	return m.multiPipe.PipeList[len(m.multiPipe.PipeList)-1]
}

func (m *multiPipeListener) lastPipeNode() PipeNode {
	lastPipe := m.lastPipe()
	return lastPipe.Nodes[len(lastPipe.Nodes)-1]
}

func (m *multiPipeListener) lastPipeNodeArray() *PipeNodeArray {
	return m.lastPipeNode().(*PipeNodeArray)
}

func (m *multiPipeListener) lastFunctionNode() *FunctionNode {
	switch m.currentScope() {
	case ScopeStreamSplitter:
		node := m.lastPipeNode().(*StreamNode)
		return node.Splitter
	case ScopeStreamCollector:
		node := m.lastPipeNode().(*StreamNode)
		return node.Collector
	case ScopePipeNodeArray:
		node := m.lastPipeNodeArray()
		return node.Nodes[len(node.Nodes)-1].(*FunctionNode)
	}
	return nil
}

func (m *multiPipeListener) lastFunctionNodeParameter() *Parameter {
	param := m.lastFunctionNode().Params
	return param[len(param)-1]
}

func (m *multiPipeListener) enterScope(scope Scope) {
	m.scopeStack.Push(scope)
}

func (m *multiPipeListener) currentScope() Scope {
	return m.scopeStack.Peek()
}

func (m *multiPipeListener) exitScope() {
	m.scopeStack.Pop()
}

func (m *multiPipeListener) EnterPipe(c *PipeContext) {
	m.enterScope(ScopePipe)
	m.multiPipe.PipeList = append(m.multiPipe.PipeList, &Pipe{Nodes: make([]PipeNode, 0)})
}

func (m *multiPipeListener) ExitPipe(c *PipeContext) {
	m.exitScope()
}

func (m *multiPipeListener) EnterStreamNode(c *StreamNodeContext) {
	m.enterScope(ScopeStreamNode)
	pipe := m.lastPipe()
	pipe.Nodes = append(pipe.Nodes, &StreamNode{})
}

func (m *multiPipeListener) ExitStreamNode(c *StreamNodeContext) {
	m.exitScope()
}

func (m *multiPipeListener) EnterStreamSplitter(c *StreamSplitterContext) {
	m.enterScope(ScopeStreamSplitter)
}

func (m *multiPipeListener) ExitStreamSplitter(c *StreamSplitterContext) {
	m.exitScope()
}

func (m *multiPipeListener) EnterStreamCollector(c *StreamCollectorContext) {
	m.enterScope(ScopeStreamCollector)
}

func (m *multiPipeListener) ExitStreamCollector(c *StreamCollectorContext) {
	m.exitScope()
}

func (m *multiPipeListener) EnterPipeNodeArray(c *PipeNodeArrayContext) {
	m.enterScope(ScopePipeNodeArray)
	pipe := m.lastPipe()
	pipe.Nodes = append(pipe.Nodes, &PipeNodeArray{Nodes: make([]PipeNode, 0)})
}

func (m *multiPipeListener) ExitPipeNodeArray(c *PipeNodeArrayContext) {
	m.exitScope()
}

// EnterVariableNode
func (m *multiPipeListener) EnterVariableNode(c *VariableNodeContext) {
	m.enterScope(ScopeVariableNode)

	variableName := c.GetText()[1:]

	pipe := m.lastPipe()
	// check if current pipe has only 1 node (pipeNodeArray)
	if len(pipe.Nodes) == 1 {
		// read from variable, check if variable is defined
		if _, ok := m.multiPipe.Variables[variableName]; !ok {
			raise(c, UndefinedVariableError(variableName))
		}
	} else {
		// update and read value, if variable has been defined, raise error
		if _, ok := m.multiPipe.Variables[variableName]; ok {
			raise(c, UpdateImmutableVariableError(variableName))
		}
	}

	v := NewImmutableValue()
	m.multiPipe.Variables[variableName] = v
	newNode := &VariableNode{Name: variableName, Value: v}

	arr := m.lastPipeNodeArray()
	arr.Nodes = append(arr.Nodes, newNode)
}

func (m *multiPipeListener) ExitVariableNode(c *VariableNodeContext) {
	m.exitScope()
}

// EnterFunctionNode
func (m *multiPipeListener) EnterFunctionNode(c *FunctionNodeContext) {
	newNode := &FunctionNode{Params: make(Parameters, 0)}
	switch m.currentScope() {
	case ScopeStreamSplitter:
		node := m.lastPipeNode().(*StreamNode)
		node.Splitter = newNode
	case ScopeStreamCollector:
		node := m.lastPipeNode().(*StreamNode)
		node.Collector = newNode
	case ScopePipeNodeArray:
		node := m.lastPipeNode().(*PipeNodeArray)
		node.Nodes = append(node.Nodes, newNode)
	default:
		panic(fmt.Sprintf("unexpected scope: %d", m.currentScope()))
	}
}

func (m *multiPipeListener) ExitFunctionNode(c *FunctionNodeContext) {
	registry.LookupFunctionHandler(m.lastFunctionNode())
}

// EnterFunctionName
func (m *multiPipeListener) EnterFunctionName(c *FunctionNameContext) {
	m.lastFunctionNode().Name = c.GetText()
}

func (m *multiPipeListener) EnterFunctionParameter(c *FunctionParameterContext) {
	node := m.lastFunctionNode()
	node.Params = append(node.Params, &Parameter{})
}

func (m *multiPipeListener) EnterFunctionParameterLabel(c *FunctionParameterLabelContext) {
	labelWithColon := c.GetText()
	m.lastFunctionNodeParameter().Name = labelWithColon[:len(labelWithColon)-1]
}

// handle basic type

func (m *multiPipeListener) EnterIntegerValue(ctx *IntegerValueContext) {
	m.updateParameterValue(NewBaseParameterValue(IntegerValue, ctx.GetText()))
}

func (m *multiPipeListener) EnterDecimalValue(ctx *DecimalValueContext) {
	m.updateParameterValue(NewBaseParameterValue(FloatValue, ctx.GetText()))
}

func (m *multiPipeListener) EnterStringValue(ctx *StringValueContext) {
	// remove single quote
	v := ctx.GetText()
	m.updateParameterValue(NewBaseParameterValue(StringValue, v[1:len(v)-1]))
}

func (m *multiPipeListener) EnterBooleanValue(ctx *BooleanValueContext) {
	m.updateParameterValue(NewBaseParameterValue(BoolValue, ctx.GetText()))
}

// handle reference value

func (m *multiPipeListener) EnterVariableValue(c *VariableValueContext) {
	variableName := c.GetText()[1:]
	// if variable is undefined, raise error
	if v, ok := m.multiPipe.Variables[variableName]; !ok {
		panic(UndefinedVariableError(variableName))
	} else {
		m.updateParameterValue(NewReferenceParameterValue(variableName, v))
	}
}

func (m *multiPipeListener) updateParameterValue(newItem Value) {
	param := m.lastFunctionNodeParameter()
	// if param.Value is not nil and we are processing other type value like integer,
	// that means we are building a map
	if param.Value != nil {
		v := param.Value.(*DictParameterValue)
		v.AddEntry(m.mapEntryLabel, newItem)
		m.mapEntryLabel = ""
	} else {
		param.Value = newItem
	}
}

// handle dict

func (m *multiPipeListener) EnterDictValue(ctx *DictValueContext) {
	m.lastFunctionNodeParameter().Value = NewDictParameterValue()
}

func (m *multiPipeListener) EnterDictEntryLabel(c *DictEntryLabelContext) {
	m.mapEntryLabel = c.GetText()
}

type AstContext interface {
	GetStart() antlr.Token
}

func raise(ctx AstContext, err error) {
	start := ctx.GetStart()
	panic(fmt.Errorf("[%d:%d] %s", start.GetLine(), start.GetColumn(), err))
}
