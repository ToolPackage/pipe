package parser

import (
	"crypto/md5"
	"fmt"
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
)

func ParseScript(filename string) *PipeScript {
	defer func() {
		if err, ok := recover().(error); ok {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		panic(err)
	}
	lexer := NewPipeLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewPipeParser(stream)
	p.AddErrorListener(NewErrorListener()) // default is console error listener
	p.BuildParseTrees = true
	tree := p.Script()
	listener := newPipeScriptListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	// fmt.Println(listener.pipeScript.String())

	return listener.pipeScript
}

type pipeScriptListener struct {
	multiPipeListener

	pipeScript *PipeScript
}

func newPipeScriptListener() *pipeScriptListener {
	return &pipeScriptListener{
		pipeScript: &PipeScript{Funcs: make([]CompactFunction, 0)},
	}
}

func (p *pipeScriptListener) lastFuncDef() *CompactFunction {
	return &p.pipeScript.Funcs[len(p.pipeScript.Funcs)-1]
}

func (p *pipeScriptListener) lastFuncParamDef() *ParameterDefinition {
	funcDef := p.lastFuncDef()
	return &funcDef.Params[len(funcDef.Params)-1]
}

// create function signature

func (p *pipeScriptListener) EnterFuncDef(ctx *FuncDefContext) {
	p.pipeScript.Funcs = append(p.pipeScript.Funcs, CompactFunction{
		Params: make([]ParameterDefinition, 0),
		Md5:    fmt.Sprintf("%x", md5.Sum([]byte(ctx.GetText()))),
	})
}

func (p *pipeScriptListener) EnterFuncName(c *FuncNameContext) {
	p.lastFuncDef().Name = c.GetText()
}

func (p *pipeScriptListener) EnterFuncParamDef(c *FuncParamDefContext) {
	p.lastFuncDef().Params = append(p.lastFuncDef().Params, ParameterDefinition{})
}

func (p *pipeScriptListener) EnterFuncParamName(c *FuncParamNameContext) {
	p.lastFuncParamDef().Name = c.GetText()
}

func (p *pipeScriptListener) EnterOptionalParamFlag(c *OptionalParamFlagContext) {
	p.lastFuncParamDef().Optional = true
}

func (p *pipeScriptListener) EnterFuncParamType(c *FuncParamTypeContext) {
	p.lastFuncParamDef().Type = TypeMappings[c.GetText()]
}

func (p *pipeScriptListener) EnterFuncBody(c *FuncBodyContext) {
	// create local variables for function params
	funcDef := p.lastFuncDef()
	localVars := make(map[string]*ImmutableValue)
	for _, paramDef := range funcDef.Params {
		if _, ok := localVars[paramDef.Name]; ok {
			panic(DuplicateFuncParamNameError)
		}
		localVars[paramDef.Name] = NewImmutableValue()
	}

	p.multiPipeListener.multiPipe = &MultiPipe{Variables: localVars, Pipes: make([]Pipe, 0)}
}

func (p *pipeScriptListener) ExitFuncBody(c *FuncBodyContext) {
	p.lastFuncDef().Callable = &CompactFunctionCallable{Pipes: p.multiPipeListener.multiPipe}
	p.multiPipeListener.multiPipe = nil
}
