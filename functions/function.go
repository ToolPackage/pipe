package functions

import (
	"fmt"
	"io"
	"strconv"
)

const (
	DecimalValue = iota
	FloatValue
	StringValue
	BoolValue
)

type Function struct {
	Name    string
	Params  FunctionParameters
	Handler FunctionHandler
}

func (c *Function) Exec(in io.Reader, out io.Writer) error {
	return c.Handler(c.Params, in, out)
}

type FunctionParameterType int

type FunctionHandler func(params FunctionParameters, in io.Reader, out io.Writer) error

type FunctionParameters []FunctionParameter

func (p FunctionParameters) Size() int {
	return len(p)
}

func (p FunctionParameters) GetParameter(label string, idx int) (*FunctionParameter, bool) {
	var param *FunctionParameter

	for i := range p {
		if p[i].Label == label {
			return &p[i], true
		}
		if i == idx {
			param = &p[i]
			break
		}
	}

	if param != nil {
		return param, true
	}
	return nil, false
}

type FunctionParameter struct {
	Label     string
	ValueType FunctionParameterType
	Value     string
}

func (p *FunctionParameter) GetValue() interface{} {
	var v interface{}
	var err error

	switch p.ValueType {
	case DecimalValue:
		v, err = strconv.Atoi(p.Value)
	case FloatValue:
		v, err = strconv.ParseFloat(p.Value, 32)
	case StringValue:
		v = p.Value[1 : len(p.Value)-1]
	case BoolValue:
		v, err = strconv.ParseBool(p.Value)
	default:
		panic(fmt.Errorf("invalid function parameter type %d", p.ValueType))
	}

	if err != nil {
		panic(err)
	}
	return v
}
