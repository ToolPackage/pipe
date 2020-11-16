package definition

import (
	"fmt"
	"github.com/ToolPackage/pipe/util"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// PipeScript
type PipeScript struct {
	Funcs []CompactFunction
}

// CompactFunction
type CompactFunction struct {
	Name     string
	Params   []FuncParamDef
	Callable CompactFunctionCallable
}

// FuncParamDef
type FuncParamDef struct {
	Name     string
	Optional bool
	Type     ParameterType
}

// ParameterType
type ParameterType int

const (
	IntegerValue = iota
	FloatValue
	StringValue
	BoolValue
	DictValue
	Reference
)

var TypeMappings = map[string]ParameterType{
	"integer": IntegerValue,
	"float":   FloatValue,
	"string":  StringValue,
	"bool":    BoolValue,
	"dict":    DictValue,
}

// CompactFunctionCallable
type CompactFunctionCallable struct {
	Pipes []Pipe
}

// MultiPipe
type MultiPipe struct {
	Pipes []Pipe
}

// Pipe
type Pipe struct {
	Variables map[string]int
	Nodes     []PipeNode
}

// PipeNode
type PipeNode interface {
	Exec(in io.Reader, out io.Writer) error
}

type VariableNode struct {
	Name  string
	Value []byte
}

func (vn *VariableNode) Exec(in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	vn.Value = input
	_, err = out.Write(input)
	return err
}

// FunctionNode
type FunctionNode struct {
	Name    string
	Params  Parameters
	Handler FunctionHandler
}

type FunctionHandler func(params Parameters, in io.Reader, out io.Writer) error

func (fn *FunctionNode) Exec(in io.Reader, out io.Writer) error {
	return fn.Handler(fn.Params, in, out)
}

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

func (p *Parameter) Labeled() bool {
	return len(p.Label) > 0
}

type ParameterValue interface {
	Type() ParameterType
	Get() interface{}
}

// BaseParameterValue is the basic type set of parameters
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
	case FloatValue:
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

// DictParameterValue
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

// Reference
type ReferenceParameterValue struct {
	Name  string
	Value interface{}
	Sync  chan interface{}
}

func NewReferenceParameterValue(name string) *ReferenceParameterValue {
	return &ReferenceParameterValue{Name: name}
}

func (rpv *ReferenceParameterValue) Type() ParameterType {
	return Reference
}

func (rpv *ReferenceParameterValue) Get() interface{} {
	if rpv.Value == nil {
		rpv.Value = <-rpv.Sync
	}
	return rpv.Value
}

// FunctionDefinition

type FunctionDefinition struct {
	Name            string
	Handler         FunctionHandler
	ParamConstraint FuncParamConstraint
}

type FuncParamConstraint []ParameterDefinition

func (fpc FuncParamConstraint) Validate(params Parameters) error {
	// 1.validate labeled params
	idx := 0
	// skip all non-labeled parameters
	for ; idx < len(params) && !params[idx].Labeled(); idx++ {
	}
	// check all following parameters are labeled
	for ; idx < len(params); idx++ {
		if !params[idx].Labeled() {
			return InvalidNonLabeledParameterError
		}
	}

	paramSz := params.Size()
	for idx, paramDef := range fpc {
		// 2.validate parameters size
		if idx == paramSz {
			for ; idx < len(fpc); idx++ {
				if !paramDef.Optional {
					return NotEnoughParameterError
				}
			}
			return nil
		}
		param, ok := params.GetParameter(paramDef.Label, idx)
		if !ok {
			return nil
		}
		// 3.validate parameter type
		if param.Value.Type() != paramDef.Type {
			return InvalidTypeOfParameterError
		}
		// 4.validate parameter value
		if len(paramDef.ConstValue) > 0 && !util.SliceContains(paramDef.ConstValue, param.Value.Get()) {
			return InvalidConstValueError
		}
	}
	return nil
}

type ParameterDefinition struct {
	Label      string
	Type       ParameterType
	Optional   bool
	ConstValue []interface{}
}

// Define function utils

func DefFuncs(funcDefList ...*FunctionDefinition) []*FunctionDefinition {
	return funcDefList
}

func DefFunc(name string, handler FunctionHandler, paramConstraint FuncParamConstraint) *FunctionDefinition {
	return &FunctionDefinition{Name: name, Handler: handler, ParamConstraint: paramConstraint}
}

func DefParams(paramDefList ...ParameterDefinition) FuncParamConstraint {
	return paramDefList
}

func DefParam(paramType ParameterType, label string, optional bool, constValue ...interface{}) ParameterDefinition {
	label = strings.Trim(label, " \t\n\r")
	if len(label) == 0 {
		panic(InvalidFuncParamDefError)
	}
	return ParameterDefinition{Type: paramType, Label: label, Optional: optional, ConstValue: constValue}
}
