package parser

import (
	"fmt"
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ParseScript(filename string) error {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return err
	}
	lexer := NewPipeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewPipeParser(stream)
	p.AddErrorListener(NewErrorListener()) // default is console error listener
	p.BuildParseTrees = true
	tree := p.Script()
	listener := newPipeScriptListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	fmt.Println(listener.pipeScript)

	return nil
}

type pipeScriptListener struct {
	*BasePipeListener

	pipeScript *PipeScript

	mapEntryLabel string
}

func newPipeScriptListener() *pipeScriptListener {
	return &pipeScriptListener{
		pipeScript: &PipeScript{Funcs: make([]CompactFunction, 0)},
	}
}

func (psl *pipeScriptListener) lastFuncDef() *CompactFunction {
	return &psl.pipeScript.Funcs[len(psl.pipeScript.Funcs)-1]
}

func (psl *pipeScriptListener) lastFuncParamDef() *FuncParamDef {
	return &psl.lastFuncDef().Params[len(psl.lastFuncDef().Params)-1]
}

func (psl *pipeScriptListener) lastPipe() *Pipe {
	return &psl.lastFuncDef().Callable.Pipes[len(psl.lastFuncDef().Callable.Pipes)-1]
}

func (psl *pipeScriptListener) lastPipeNode() PipeNode {
	return psl.lastPipe().Nodes[len(psl.lastPipe().Nodes)-1]
}

func (psl *pipeScriptListener) lastFunctionNode() *FunctionNode {
	return psl.lastPipeNode().(*FunctionNode)
}

func (psl *pipeScriptListener) lastFunctionNodeParameter() *Parameter {
	return &psl.lastFunctionNode().Params[len(psl.lastFunctionNode().Params)-1]
}

// create function signature

func (psl *pipeScriptListener) EnterFuncDef(ctx *FuncDefContext) {
	psl.pipeScript.Funcs = append(psl.pipeScript.Funcs, CompactFunction{
		Params:   make([]FuncParamDef, 0),
		Callable: CompactFunctionCallable{Pipes: make([]Pipe, 0)},
	})
}

func (psl *pipeScriptListener) EnterFuncName(c *FuncNameContext) {
	psl.lastFuncDef().Name = c.GetText()
}

func (psl *pipeScriptListener) EnterFuncParamDef(c *FuncParamDefContext) {
	psl.lastFuncDef().Params = append(psl.lastFuncDef().Params, FuncParamDef{})
}

func (psl *pipeScriptListener) EnterFuncParamName(c *FuncParamNameContext) {
	psl.lastFuncParamDef().Name = c.GetText()
}

func (psl *pipeScriptListener) EnterOptionalParamFlag(c *OptionalParamFlagContext) {
	psl.lastFuncParamDef().Optional = true
}

func (psl *pipeScriptListener) EnterFuncParamType(c *FuncParamTypeContext) {
	psl.lastFuncParamDef().Type = TypeMappings[c.GetText()]
}

// handle multiPipe

func (psl *pipeScriptListener) EnterPipe(c *PipeContext) {
	psl.lastFuncDef().Callable.Pipes = append(psl.lastFuncDef().Callable.Pipes,
		Pipe{Variables: make(map[string]int), Nodes: make([]PipeNode, 0)})
}

func (psl *pipeScriptListener) EnterVariableNode(c *VariableNodeContext) {
	dollarWithName := c.GetText()
	psl.lastPipe().Nodes = append(psl.lastPipe().Nodes, &VariableNode{Name: dollarWithName[1:]})
}

func (psl *pipeScriptListener) EnterFunctionNode(c *FunctionNodeContext) {
	pipe := psl.lastPipe()
	pipe.Nodes = append(pipe.Nodes, &FunctionNode{Params: make([]Parameter, 0)})
}

func (psl *pipeScriptListener) EnterFunctionName(c *FunctionNameContext) {
	psl.lastFunctionNode().Name = c.GetText()
}

func (psl *pipeScriptListener) EnterFunctionParameter(c *FunctionParameterContext) {
	node := psl.lastFunctionNode()
	node.Params = append(node.Params, Parameter{})
}

func (psl *pipeScriptListener) EnterFunctionParameterLabel(c *FunctionParameterLabelContext) {
	labelWithColon := c.GetText()
	psl.lastFunctionNodeParameter().Label = labelWithColon[:len(labelWithColon)-1]
}

// handle basic type

func (psl *pipeScriptListener) EnterIntegerValue(ctx *IntegerValueContext) {
	psl.updateParameterValue(NewBaseParameterValue(IntegerValue, ctx.GetText()))
}

func (psl *pipeScriptListener) EnterDecimalValue(ctx *DecimalValueContext) {
	psl.updateParameterValue(NewBaseParameterValue(FloatValue, ctx.GetText()))
}

func (psl *pipeScriptListener) EnterStringValue(ctx *StringValueContext) {
	psl.updateParameterValue(NewBaseParameterValue(StringValue, ctx.GetText()))
}

func (psl *pipeScriptListener) EnterBooleanValue(ctx *BooleanValueContext) {
	psl.updateParameterValue(NewBaseParameterValue(BoolValue, ctx.GetText()))
}

func (psl *pipeScriptListener) updateParameterValue(newItem ParameterValue) {
	param := psl.lastFunctionNodeParameter()
	if param.Value != nil {
		v := param.Value.(*DictParameterValue)
		v.AddEntry(psl.mapEntryLabel, newItem)
		psl.mapEntryLabel = ""
	} else {
		param.Value = newItem
	}
}

// handle dict

func (psl *pipeScriptListener) EnterDictValue(ctx *DictValueContext) {
	psl.lastFunctionNodeParameter().Value = NewDictParameterValue()
}

func (psl *pipeScriptListener) EnterDictEntryLabel(c *DictEntryLabelContext) {
	psl.mapEntryLabel = c.GetText()
}

// handle reference value

func (psl *pipeScriptListener) EnterVariableValue(c *VariableValueContext) {
	dollarWithName := c.GetText()
	psl.lastFunctionNodeParameter().Value = NewReferenceParameterValue(dollarWithName[1:])
}
