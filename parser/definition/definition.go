package definition

import (
	"fmt"
	"github.com/ToolPackage/pipe/util"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

// PipeScript
type PipeScript struct {
	Funcs []CompactFunction
}

// CompactFunction
type CompactFunction struct {
	Name      string
	Variables map[string]*ImmutableValue
	Params    []FuncParamDef
	Callable  CompactFunctionCallable
}

// FuncParamDef
type FuncParamDef struct {
	Name     string
	Optional bool
	Type     ParameterType
}

// ParameterType
type ParameterType int
type ValueType int

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
	Variables map[string]*ImmutableValue
	Pipes     []Pipe
}

// Pipe
type Pipe struct {
	Nodes []PipeNode
}

// PipeNode
type PipeNode interface {
	Exec(in io.Reader, out io.Writer) error
}

type VariableNode struct {
	Name  string
	Value *ImmutableValue
}

func (vn *VariableNode) Exec(in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	vn.Value.Assign(string(input))
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
	Value Value
}

func (p *Parameter) Labeled() bool {
	return len(p.Label) > 0
}

type Value interface {
	Type() ValueType
	// get converted value
	Get() interface{}
}

// BaseParameterValue is the basic type set of parameters
type BaseParameterValue struct {
	ValueType ValueType
	Value     string
}

func NewBaseParameterValue(valueType ValueType, value string) Value {
	return &BaseParameterValue{ValueType: valueType, Value: value}
}

func (bpv *BaseParameterValue) Type() ValueType {
	return bpv.ValueType
}

func (bpv *BaseParameterValue) Get() interface{} {
	v, err := parseStringToValue(bpv.Value, bpv.ValueType)
	if err != nil {
		panic(err)
	}
	return v
}

// DictParameterValue
type DictParameterValue struct {
	ValueOrderMap []string
	Value         map[string]Value
}

func NewDictParameterValue() Value {
	return &DictParameterValue{Value: make(map[string]Value)}
}

func (dpv *DictParameterValue) Type() ValueType {
	return DictValue
}

func (dpv *DictParameterValue) Size() int {
	return len(dpv.Value)
}

func (dpv *DictParameterValue) AddEntry(label string, value Value) {
	if len(label) == 0 {
		label = strconv.Itoa(len(dpv.Value))
	}
	dpv.Value[label] = value
	dpv.ValueOrderMap = append(dpv.ValueOrderMap, label)
}

func (dpv *DictParameterValue) GetValue(label string, idx int) (Value, bool) {
	v, ok := dpv.Value[label]
	if ok {
		return v, true
	}
	return dpv.GetValueByIndex(idx)
}

func (dpv *DictParameterValue) GetValueByLabel(label string) (Value, bool) {
	v, ok := dpv.Value[label]
	return v, ok
}

func (dpv *DictParameterValue) GetValueByIndex(idx int) (Value, bool) {
	if idx < 0 || idx >= dpv.Size() {
		return nil, false
	}
	label := dpv.ValueOrderMap[idx]
	return dpv.GetValueByLabel(label)
}

func (dpv *DictParameterValue) Get() interface{} {
	return dpv.Value
}

// Reference
type ReferenceParameterValue struct {
	Name    string
	RefType ValueType
	Value   *ImmutableValue
}

func NewReferenceParameterValue(name string, value *ImmutableValue) Value {
	return &ReferenceParameterValue{Name: name, Value: value}
}

func (rpv *ReferenceParameterValue) Type() ValueType {
	return Reference
}

func (rpv *ReferenceParameterValue) Get() interface{} {
	v, err := parseStringToValue(rpv.Value.SyncAndGet().(string), rpv.RefType)
	if err != nil {
		panic(err)
	}
	return v
}

// parseStringToValue is used to parse text value to typed value, including int, float, string and bool
func parseStringToValue(text string, valueType ValueType) (v interface{}, err error) {
	switch valueType {
	case IntegerValue:
		v, err = strconv.Atoi(text)
	case FloatValue:
		v, err = strconv.ParseFloat(text, 32)
	case StringValue:
		v = text
	case BoolValue:
		v, err = strconv.ParseBool(text)
	default:
		err = fmt.Errorf("invalid function parameter type %s", text)
	}
	return
}

// ImmutableValue
type ImmutableValue struct {
	Value    interface{}
	mu       sync.Mutex
	assigned bool
	sync     chan bool
}

func NewImmutableValue() *ImmutableValue {
	return &ImmutableValue{sync: make(chan bool, 1)}
}

func (v *ImmutableValue) Assign(newValue interface{}) {
	v.mu.Lock()
	v.Value = newValue
	v.assigned = true
	v.mu.Unlock()
	// notify
	v.sync <- true
}

func (v *ImmutableValue) IsAssigned() bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.assigned
}

func (v *ImmutableValue) Get() interface{} {
	if !v.IsAssigned() {
		panic("cannot access undefined value")
	}
	return v.Value
}

func (v *ImmutableValue) SyncAndGet() interface{} {
	for !v.IsAssigned() {
		// wait
		<-v.sync
	}
	return v.Value
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
		if param.Value.Type() == Reference {
			// if parameter is a variable, check if function param type is dict
			if paramDef.Type == DictValue {
				return InvalidUsageOfVairableError
			}
			// set reference value's type which will be used to convert variable text value to typed value
			refValue := param.Value.(*ReferenceParameterValue)
			refValue.RefType = paramDef.Type
		} else if param.Value.Type() != paramDef.Type {
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
	Type       ValueType
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

func DefParam(paramType ValueType, label string, optional bool, constValue ...interface{}) ParameterDefinition {
	label = strings.Trim(label, " \t\n\r")
	if len(label) == 0 {
		panic(InvalidFuncParamDefError)
	}
	return ParameterDefinition{Type: paramType, Label: label, Optional: optional, ConstValue: constValue}
}
