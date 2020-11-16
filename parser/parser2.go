package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

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
	listener := NewMultiPipeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return nil
}

type PipeScriptListener struct {
	*BasePipeListener
}

func (s *PipeScriptListener) EnterFuncDef(ctx *FuncDefContext) {

}
