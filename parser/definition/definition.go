package definition

import (
	"bytes"
	"fmt"
	"github.com/ToolPackage/pipe/util"
	"github.com/vipally/binary"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

const indentSpacing = "  "
const FuncNameLenLimit = 255
const FuncParamLenLimit = 255
const FuncParamConstValueLenLimit = 255
const FuncVariableLenLimit = 255
const VariableNameLenLimit = 255
const FuncMultiPipeLenLimit = 65535

// pipe node type
const (
	variableNode = iota
	streamNode
	pipeNodeArray
	functionNode
)

// PipeScript
type PipeScript struct {
	Funcs []CompactFunction
}

func (p *PipeScript) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("PipeScript [")
	builder.IncIndent()

	for _, funcDef := range p.Funcs {
		builder.WriteMultiLine(funcDef.String(), true)
	}

	builder.DecIndent()
	builder.WriteWithIndent("]")

	return builder.String()
}

// CompactFunction
type CompactFunction struct {
	Name     string
	Params   []ParameterDefinition
	Md5      string
	Callable *CompactFunctionCallable
}

func (c *CompactFunction) WriteTo(out *binary.Encoder) {
	// write func name length
	out.Uint8(uint8(len(c.Name)))
	// write func name
	out.String(c.Name)
	// write func param length
	out.Uint8(uint8(len(c.Name)))
	// write func params
	for idx := range c.Params {
		c.Params[idx].WriteTo(out)
	}
	// write func md5 (16)
	out.String(c.Md5)
	// write callable
	c.Callable.WriteTo(out)
}

func (c *CompactFunction) ReadFrom(in *binary.Decoder) {
	// TODO:
}

func (c *CompactFunction) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("CompactFunction {")

	builder.IncIndent()
	builder.WriteLine("Name: ", c.Name)

	builder.WriteLine("Params: [")
	builder.IncIndent()
	for _, paramDef := range c.Params {
		builder.WriteMultiLine(paramDef.String(), true)
	}
	builder.DecIndent()
	builder.WriteLine("]")

	builder.WriteIndent()
	builder.WriteString("Callable: ")
	builder.WriteMultiLine(c.Callable.String(), false)

	builder.DecIndent()
	builder.WriteWithIndent("}")

	return builder.String()
}

type ParameterDefinition struct {
	Name       string
	Type       ValueType
	Optional   bool
	ConstValue []interface{}
}

func (p *ParameterDefinition) WriteTo(out *binary.Encoder) {
	// write param name length
	out.Uint8(uint8(len(p.Name)))
	// write param name
	out.String(p.Name)
	// write param type
	out.Uint8(uint8(p.Type))
	// write optional flag
	if p.Optional {
		out.Uint8(1)
	} else {
		out.Uint8(0)
	}
	// write param const value length
	out.Uint8(uint8(len(p.ConstValue)))
	// write param const value
	for _, v := range p.ConstValue {
		// TODO: handle other type of const value
		// write value length
		out.Uint8(uint8(len(v.(string))))
		// write value
		out.String(v.(string))
	}
}

func (p *ParameterDefinition) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteWithIndent("FuncParamDef { Name: ", p.Name,
		", Optional: ", strconv.FormatBool(p.Optional),
		", Type: ", ValueType(p.Type).String(),
		" }")
	return builder.String()
}

type ValueType int

func (v ValueType) String() string {
	switch v {
	case IntegerValue:
		return "int"
	case FloatValue:
		return "float"
	case StringValue:
		return "string"
	case BoolValue:
		return "bool"
	case DictValue:
		return "dict"
	default:
		return "unknown"
	}
}

const (
	IntegerValue = iota
	FloatValue
	StringValue
	BoolValue
	DictValue
	ReferenceValue
	Unknown
)

var TypeMappings = map[string]ValueType{
	"int":    IntegerValue,
	"float":  FloatValue,
	"string": StringValue,
	"bool":   BoolValue,
	"dict":   DictValue,
}

// CompactFunctionCallable
type CompactFunctionCallable struct {
	Pipes *MultiPipe
}

func (c *CompactFunctionCallable) Exec(params Parameters, in io.Reader, out io.Writer) error {
	// TODO:
	return nil
}

func (c *CompactFunctionCallable) WriteTo(out *binary.Encoder) {
	c.Pipes.WriteTo(out)
}

func (c *CompactFunctionCallable) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("CompactFunctionCallable {")
	builder.IncIndent()

	builder.WriteMultiLine(c.Pipes.String(), true)

	builder.DecIndent()
	builder.WriteWithIndent("}")
	return builder.String()
}

// MultiPipe
type MultiPipe struct {
	Variables map[string]*ImmutableValue
	Pipes     []Pipe
}

func (m *MultiPipe) WriteTo(out *binary.Encoder) {
	// write variables length
	out.Uint8(uint8(len(m.Variables)))
	// write variables
	for name := range m.Variables {
		// write variable name len
		out.Uint8(uint8(len(name)))
		// write variable name
		out.String(name)
		// no need to serialize ImmutableValue
	}
	// write pipe length
	out.Uint16(uint16(len(m.Pipes)), false)
	// write pipes
	for idx := range m.Pipes {
		m.Pipes[idx].WriteTo(out)
	}
}

func (m *MultiPipe) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("MultiPipe [")
	builder.IncIndent()

	for _, pipe := range m.Pipes {
		builder.WriteMultiLine(pipe.String(), true)
	}

	builder.DecIndent()
	builder.WriteWithIndent("]")
	return builder.String()
}

// Pipe
type Pipe struct {
	Nodes []PipeNode
}

func (p *Pipe) Exec(in io.Reader, out io.Writer) error {
	i := util.NewDualChannel()
	if _, err := io.Copy(i, in); err != nil {
		return err
	}
	o := util.NewDualChannel()

	for _, node := range p.Nodes {
		if err := node.Exec(i, o); err != nil {
			return err
		}
		i, o = o, i
		o.Reset()
	}

	_, err := io.Copy(out, o)
	return err
}

func (p *Pipe) WriteTo(out *binary.Encoder) {
	// write nodes len
	out.Uint16(uint16(len(p.Nodes)), false)
	// write nodes
	for idx := range p.Nodes {
		p.Nodes[idx].WriteTo(out)
	}
}

func (p *Pipe) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("Pipe [")
	builder.IncIndent()

	for _, node := range p.Nodes {
		builder.WriteMultiLine(node.String(), true)
	}

	builder.DecIndent()
	builder.WriteWithIndent("]")
	return builder.String()
}

// PipeNode
type PipeNode interface {
	Exec(in io.Reader, out io.Writer) error
	WriteTo(out *binary.Encoder)
	String() string
}

type VariableNode struct {
	Name  string
	Value *ImmutableValue
}

func (v *VariableNode) Exec(in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	v.Value.Assign(string(input))
	_, err = out.Write(input)
	return err
}

func (v *VariableNode) WriteTo(out *binary.Encoder) {
	// write node type
	out.Uint8(uint8(variableNode))
	// write variable name len
	out.Uint8(uint8(len(v.Name)))
	// write variable name
	out.String(v.Name)
}

func (v *VariableNode) String() string {
	return fmt.Sprintf("VariableNode { Name: %s, Value: %s }",
		v.Name, v.Value.String())
}

// StreamNode
type StreamNode struct {
	Splitter  *FunctionNode
	Processor Pipe
	Collector *FunctionNode
}

func (s *StreamNode) Exec(in io.Reader, out io.Writer) error {
	var err error
	var frameSz uint32
	var frameBuf []byte
	frameSzBuf := make([]byte, 4)

	frameIn := util.NewSyncDualChannel()
	frameOut := util.NewSyncDualChannel()
	defer frameIn.Close()
	defer frameOut.Close()

	go func() {
		if err := s.Splitter.Exec(in, frameIn); err != nil {
			// TODO:
			fmt.Println(err)
		}
		frameIn.Close()
	}()
	go func() {
		if err := s.Collector.Exec(frameOut, out); err != nil {
			// TODO:
			fmt.Println(err)
		}
		frameOut.Close()
	}()

	for {
		// read frame
		if _, err = frameIn.Read(frameSzBuf); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		frameSz = util.ConvertByteToUint32(frameSzBuf, 0)
		frameBuf = make([]byte, frameSz)
		if _, err = frameIn.Read(frameBuf); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		// process frame
		if err = s.Processor.Exec(bytes.NewReader(frameBuf), frameOut); err != nil {
			return err
		}
	}
}

func (s *StreamNode) WriteTo(out *binary.Encoder) {
	// write node type
	out.Uint8(uint8(streamNode))

	s.Splitter.WriteTo(out)
	s.Processor.WriteTo(out)
	s.Collector.WriteTo(out)
}

func (s *StreamNode) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("StreamNode {")
	builder.IncIndent()

	builder.WriteWithIndent("Splitter: ")
	builder.WriteMultiLine(s.Splitter.String(), false)

	builder.WriteWithIndent("Processor: ")
	builder.WriteMultiLine(s.Processor.String(), false)

	builder.WriteWithIndent("Collector: ")
	builder.WriteMultiLine(s.Collector.String(), false)

	builder.DecIndent()
	builder.WriteWithIndent("}")
	return builder.String()
}

// PipeNodeArray
type PipeNodeArray struct {
	Nodes []PipeNode
}

// Exec: exec each function node sequentially
func (m *PipeNodeArray) Exec(in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	sep := ""
	for _, node := range m.Nodes {
		if _, err = out.Write([]byte(sep)); err != nil {
			return err
		}
		if err = node.Exec(bytes.NewReader(input), out); err != nil {
			return err
		}
		sep = "\n"
	}

	return nil
}

func (m *PipeNodeArray) WriteTo(out *binary.Encoder) {
	// write node type
	out.Uint8(uint8(pipeNodeArray))
	// write nodes len
	out.Uint8(uint8(len(m.Nodes)))
	// write nodes
	for idx := range m.Nodes {
		m.Nodes[idx].WriteTo(out)
	}
}

func (m *PipeNodeArray) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("PipeNodeArray [")
	builder.IncIndent()

	for _, node := range m.Nodes {
		builder.WriteMultiLine(node.String(), true)
	}

	builder.DecIndent()
	builder.WriteWithIndent("]")
	return builder.String()
}

// FunctionNode
type FunctionNode struct {
	Name    string
	Params  Parameters
	Handler FunctionHandler
}

type FunctionHandler func(params Parameters, in io.Reader, out io.Writer) error

func (f *FunctionNode) Exec(in io.Reader, out io.Writer) error {
	if err := f.Handler(f.Params, in, out); err != nil {
		switch err {
		case NotEnoughParameterError:
			return fmt.Errorf("not enough parameters, function usage: %s", util.FuncDescription(f.Handler))
		case InvalidTypeOfParameterError:
			return fmt.Errorf("invalid type of parameter, function usage: %s", util.FuncDescription(f.Handler))
		default:
			return err
		}
	}
	return nil
}

func (f *FunctionNode) WriteTo(out *binary.Encoder) {
	// TODO: move all WriteTo methods to a single method
}

func (f *FunctionNode) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("FunctionNode {")
	builder.IncIndent()

	builder.WriteLine("Name: ", f.Name)

	if f.Params.Size() > 0 {
		builder.WriteLine("Params: [")
		builder.IncIndent()
		for _, param := range f.Params {
			builder.WriteMultiLine(param.String(), true)
		}
		builder.DecIndent()
		builder.WriteLine("]")
	} else {
		builder.WriteLine("Params: []")
	}

	builder.WriteLine("Handler: ", fmt.Sprintf("%p", f.Handler))

	builder.DecIndent()
	builder.WriteWithIndent("}")
	return builder.String()
}

type Parameters []Parameter

func (p Parameters) Size() int {
	return len(p)
}

func (p Parameters) GetParameter(label string, idx int) (*Parameter, bool) {
	var param *Parameter

	for i := range p {
		if p[i].Name == label {
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
	Name  string
	Value Value
}

func (p *Parameter) Labeled() bool {
	return len(p.Name) > 0
}

func (p *Parameter) String() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("Parameter: {")
	builder.IncIndent()

	builder.WriteLine("Name: ", p.Name)
	builder.WriteLine("Value: ", p.Value.String())

	builder.DecIndent()
	builder.WriteWithIndent("}")

	return builder.String()
}

type Value interface {
	Type() ValueType
	// get converted value
	Get() interface{}
	String() string
}

// BaseParameterValue is the basic type set of parameters
type BaseParameterValue struct {
	ValueType ValueType
	Value     string
}

func NewBaseParameterValue(valueType ValueType, value string) Value {
	return &BaseParameterValue{ValueType: valueType, Value: value}
}

func (b *BaseParameterValue) Type() ValueType {
	return b.ValueType
}

func (b *BaseParameterValue) Get() interface{} {
	v, err := parseStringToValue(b.Value, b.ValueType)
	if err != nil {
		panic(err)
	}
	return v
}

func (b *BaseParameterValue) String() string {
	return "BaseParameterValue { Type: " + b.ValueType.String() + ", Value: " + b.Value + " }"
}

// DictParameterValue
type DictParameterValue struct {
	ValueOrderMap []string
	Value         map[string]Value
}

func NewDictParameterValue() Value {
	return &DictParameterValue{Value: make(map[string]Value)}
}

func (d *DictParameterValue) Type() ValueType {
	return DictValue
}

func (d *DictParameterValue) Size() int {
	return len(d.Value)
}

func (d *DictParameterValue) String() string {
	return "DictParameterValue { Type: dict, Value: " + fmt.Sprintf("%v", d.Value) + " }"
}

func (d *DictParameterValue) AddEntry(label string, value Value) {
	if len(label) == 0 {
		label = strconv.Itoa(len(d.Value))
	}
	d.Value[label] = value
	d.ValueOrderMap = append(d.ValueOrderMap, label)
}

func (d *DictParameterValue) GetValue(label string, idx int) (Value, bool) {
	v, ok := d.Value[label]
	if ok {
		return v, true
	}
	return d.GetValueByIndex(idx)
}

func (d *DictParameterValue) GetValueByLabel(label string) (Value, bool) {
	v, ok := d.Value[label]
	return v, ok
}

func (d *DictParameterValue) GetValueByIndex(idx int) (Value, bool) {
	if idx < 0 || idx >= d.Size() {
		return nil, false
	}
	label := d.ValueOrderMap[idx]
	return d.GetValueByLabel(label)
}

func (d *DictParameterValue) Get() interface{} {
	return d.Value
}

// ReferenceParameterValue
type ReferenceParameterValue struct {
	Name    string
	RefType ValueType
	Value   *ImmutableValue
}

func NewReferenceParameterValue(name string, value *ImmutableValue) Value {
	return &ReferenceParameterValue{Name: name, RefType: Unknown, Value: value}
}

func (r *ReferenceParameterValue) Type() ValueType {
	return ReferenceValue
}

func (r *ReferenceParameterValue) Get() interface{} {
	v, err := parseStringToValue(r.Value.SyncAndGet().(string), r.RefType)
	if err != nil {
		panic(err)
	}
	return v
}

func (r *ReferenceParameterValue) String() string {
	return "ReferenceParameterValue { Type: reference, Value: " + r.Value.String() + " }"
}

func (r *ReferenceParameterValue) GetAs(valueType ValueType) (interface{}, error) {
	return parseStringToValue(r.Value.SyncAndGet().(string), valueType)
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

func (v *ImmutableValue) String() string {
	return fmt.Sprintf("ImmutableValue { Value: %v, assigned: %v }",
		v.Value, v.assigned)
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
	Builtin         bool
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
		param, ok := params.GetParameter(paramDef.Name, idx)
		if !ok {
			return nil
		}
		// 3.validate parameter type
		if param.Value.Type() == ReferenceValue {
			// if parameter is a variable, check if function param type is dict
			if paramDef.Type == DictValue {
				return InvalidUsageOfVariableError
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

// Define function utils

func DefFuncs(funcDefList ...*FunctionDefinition) []*FunctionDefinition {
	return funcDefList
}

func DefLibFunc(name string, handler FunctionHandler, paramConstraint FuncParamConstraint) *FunctionDefinition {
	return &FunctionDefinition{Name: name, Builtin: false, Handler: handler, ParamConstraint: paramConstraint}
}

func DefBuiltinFunc(name string, handler FunctionHandler, paramConstraint FuncParamConstraint) *FunctionDefinition {
	return &FunctionDefinition{Name: name, Builtin: true, Handler: handler, ParamConstraint: paramConstraint}
}

func DefParams(paramDefList ...ParameterDefinition) FuncParamConstraint {
	return paramDefList
}

func DefParam(paramType ValueType, label string, optional bool, constValue ...interface{}) ParameterDefinition {
	label = strings.Trim(label, " \t\n\r")
	if len(label) == 0 {
		panic(InvalidFuncParamDefError)
	}
	return ParameterDefinition{Type: paramType, Name: label, Optional: optional, ConstValue: constValue}
}
