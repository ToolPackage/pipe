package parser

import "io"

type PipeScript struct {
	Funcs []CompactFunction
}

type CompactFunction struct {
	Name     string
	Params   []Parameter
	Callable CompactFunctionCallable
}

type Parameter struct {
	Name     string
	Optional bool
	Type     ParameterType
}

type ParameterType int

const (
	IntegerValue = iota
	DecimalValue
	StringValue
	BoolValue
	DictValue
)

type CompactFunctionCallable struct {
	Pipes []Pipe
}

type Pipe struct {
	variables map[string]int
	nodes     []PipeNode
}

type PipeNode interface {
	Exec(in io.Reader, out io.Writer) error
}
