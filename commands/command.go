package commands

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

type Command struct {
	Name    string
	Params  CommandParameters
	Handler CommandHandler
}

func (c *Command) Exec(in io.Reader, out io.Writer) error {
	return c.Handler(c.Params, in, out)
}

type CommandParameterType int

type CommandHandler func(params CommandParameters, in io.Reader, out io.Writer) error

type CommandParameters []CommandParameter

func (p CommandParameters) Size() int {
	return len(p)
}

func (p CommandParameters) GetParameter(label string, idx int) (interface{}, bool) {
	var param *CommandParameter

	for i, parameter := range p {
		if parameter.Label == label {
			return parameter.getValue(), true
		}
		if i == idx {
			param = &p[i]
			break
		}
	}

	if param != nil {
		return param.getValue(), true
	}
	return nil, false
}

type CommandParameter struct {
	Label     string
	ValueType CommandParameterType
	Value     string
}

func (p *CommandParameter) getValue() interface{} {
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
		panic(fmt.Errorf("invalid commands parameter type %d", p.ValueType))
	}

	if err != nil {
		panic(err)
	}
	return v
}
