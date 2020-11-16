package parser

import "io"

// PipeScript
type PipeScript struct {
	Funcs []CompactFunction
}

// CompactFunction
type CompactFunction struct {
	Name     string
	Params   []Parameter
	Callable CompactFunctionCallable
}

// Parameter
type Parameter struct {
	Name     string
	Optional bool
	Type     ParameterType
}

// ParameterType
type ParameterType int

const (
	IntegerValue = iota
	DecimalValue
	StringValue
	BoolValue
	DictValue
)

// CompactFunctionCallable
type CompactFunctionCallable struct {
	Pipes []Pipe
}

// Pipe
type Pipe struct {
	variables map[string]int
	nodes     []PipeNode
}

// PipeNode
type PipeNode interface {
	Exec(in io.Reader, out io.Writer) error
}
