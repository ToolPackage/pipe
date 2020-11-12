package functions

import (
	"fmt"
	"io"
	"strconv"
)

const (
	IntegerValue = iota
	DecimalValue
	StringValue
	BoolValue
	DictValue
)

type Function struct {
	Name    string
	Params  Parameters
	Handler FunctionHandler
}

func (c *Function) Exec(in io.Reader, out io.Writer) error {
	return c.Handler(c.Params, in, out)
}

type ParameterType int

type FunctionHandler func(params Parameters, in io.Reader, out io.Writer) error

type Parameters []Parameter

func (p Parameters) Size() int {
	return len(p)
}

func (p Parameters) GetParameter(label string, idx int) (*Parameter, bool) {
	var param *Parameter

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

func (p Parameters) GetParameterByIndex(idx int) (*Parameter, bool) {
	if idx < 0 || idx >= len(p) {
		return nil, false
	}

	return &p[idx], true
}

type Parameter struct {
	Label string
	Value ParameterValue
}

type ParameterValue interface {
	Type() ParameterType
	Get() interface{}
}

type BaseParameterValue struct {
	ValueType ParameterType
	Value     string
}

func NewBaseParameterValue(valueType ParameterType, value string) *BaseParameterValue {
	return &BaseParameterValue{ValueType: valueType, Value: value}
}

func (bpv *BaseParameterValue) Type() ParameterType {
	return bpv.ValueType
}

func (bpv *BaseParameterValue) Get() interface{} {
	var v interface{}
	var err error

	switch bpv.ValueType {
	case IntegerValue:
		v, err = strconv.Atoi(bpv.Value)
	case DecimalValue:
		v, err = strconv.ParseFloat(bpv.Value, 32)
	case StringValue:
		v = bpv.Value[1 : len(bpv.Value)-1]
	case BoolValue:
		v, err = strconv.ParseBool(bpv.Value)
	default:
		panic(fmt.Errorf("invalid function parameter type %s", bpv.Value))
	}

	if err != nil {
		panic(err)
	}
	return v
}

type DictParameterValue struct {
	ValueType     ParameterType
	ValueOrderMap []string
	Value         map[string]ParameterValue
}

func NewDictParameterValue() *DictParameterValue {
	return &DictParameterValue{ValueType: DictValue, Value: make(map[string]ParameterValue)}
}

func (dpv *DictParameterValue) Size() int {
	return len(dpv.Value)
}

func (dpv *DictParameterValue) AddEntry(label string, value ParameterValue) {
	if len(label) == 0 {
		label = strconv.Itoa(len(dpv.Value))
	}
	dpv.Value[label] = value
	dpv.ValueOrderMap = append(dpv.ValueOrderMap, label)
}

func (dpv *DictParameterValue) GetValue(label string, idx int) (ParameterValue, bool) {
	v, ok := dpv.Value[label]
	if ok {
		return v, true
	}
	return dpv.GetValueByIndex(idx)
}

func (dpv *DictParameterValue) GetValueByLabel(label string) (ParameterValue, bool) {
	v, ok := dpv.Value[label]
	return v, ok
}

func (dpv *DictParameterValue) GetValueByIndex(idx int) (ParameterValue, bool) {
	if idx < 0 || idx >= dpv.Size() {
		return nil, false
	}
	label := dpv.ValueOrderMap[idx]
	return dpv.GetValueByLabel(label)
}

func (dpv *DictParameterValue) Type() ParameterType {
	return dpv.ValueType
}

func (dpv *DictParameterValue) Get() interface{} {
	return dpv.Value
}

type FunctionDefinition struct {
	Name            string
	Handler         FunctionHandler
	ParamConstraint FuncParamConstraint
}

type ParameterDefinition struct {
	Type       ParameterType
	Optional   bool
	ConstValue []string
}

type ParamDefSequence []ParameterDefinition
type FuncParamConstraint []ParamDefSequence

func (fpc FuncParamConstraint) Validate(params Parameters) error {
}
