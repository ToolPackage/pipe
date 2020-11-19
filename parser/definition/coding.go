package definition

import (
	"github.com/vipally/binary"
)

// pipe node type
const (
	variableNode = iota
	streamNode
	pipeNodeArray
	functionNode
)

const (
	FuncNameLenLimit            = 255
	FuncParamLenLimit           = 255
	FuncParamConstValueLenLimit = 255
	FuncVariableLenLimit        = 255
	VariableNameLenLimit        = 255
	FuncMultiPipeLenLimit       = 65535
	FuncMultiPipeNodeLenLimit   = 65535
)

func Serialize(funcDef *CompactFunction) []byte {
	encoder := binary.NewEncoder(funcDef.Size())
	writeCompactFunction(funcDef, encoder)
	return encoder.Buffer()
}

func writeCompactFunction(c *CompactFunction, out *binary.Encoder) {
	// write func name
	out.String(c.Name)
	// write func md5 (16)
	out.String(c.Md5)
	// write func param length
	out.Uint8(uint8(len(c.Name)))
	// write func params
	for idx := range c.Params {
		writeParameterDefinition(&c.Params[idx], out)
	}
	// write callable
	writeCompactFunctionCallable(c.Callable, out)
}

func writeParameterDefinition(p *ParameterDefinition, out *binary.Encoder) {
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
	constValueWriter := getBasicTypeValueWriter(p.Type)
	for _, v := range p.ConstValue {
		constValueWriter(out, v)
	}
}

func getBasicTypeValueWriter(valueType ValueType) func(out *binary.Encoder, i interface{}) {
	switch valueType {
	case IntegerValue:
		return func(out *binary.Encoder, i interface{}) {
			out.Int(i.(int))
		}
	case FloatValue:
		return func(out *binary.Encoder, i interface{}) {
			out.Float64(i.(float64))
		}
	case StringValue:
		return func(out *binary.Encoder, i interface{}) {
			out.String(i.(string))
		}
	case BoolValue:
		return func(out *binary.Encoder, i interface{}) {
			out.Bool(i.(bool))
		}
	default:
		panic("invalid const value type")
	}
}

func writeCompactFunctionCallable(c *CompactFunctionCallable, out *binary.Encoder) {
	writeMultiPipe(c.Pipes, out)
}

func writeMultiPipe(m *MultiPipe, out *binary.Encoder) {
	// write variables length
	out.Uint8(uint8(len(m.Variables)))
	// write variables
	for name := range m.Variables {
		// write variable name
		// no need to serialize ImmutableValue
		out.String(name)
	}
	// write pipe length
	out.Uint16(uint16(len(m.PipeList)), false)
	// write pipes
	for idx := range m.PipeList {
		writePipe(m.PipeList[idx], out)
	}
}

func writePipe(p *Pipe, out *binary.Encoder) {
	// write nodes len
	out.Uint16(uint16(len(p.Nodes)), false)
	// write nodes
	for idx := range p.Nodes {
		writePipeNode(p.Nodes[idx], out)
	}
}

func writePipeNode(p PipeNode, out *binary.Encoder) {
	if s, ok := p.(*StreamNode); ok {
		writeStreamNode(s, out)
		return
	}
	if arr, ok := p.(*PipeNodeArray); ok {
		writePipeNodeArray(arr, out)
		return
	}
}

func writeVariableNode(v *VariableNode, out *binary.Encoder) {
	// write node type
	out.Uint8(uint8(variableNode))
	// write variable name
	out.String(v.Name)
}

func writeStreamNode(s *StreamNode, out *binary.Encoder) {
	// write node type
	out.Uint8(uint8(streamNode))

	writeFunctionNode(s.Splitter, out)
	writePipe(s.Processor, out)
	writeFunctionNode(s.Collector, out)
}

func writePipeNodeArray(p *PipeNodeArray, out *binary.Encoder) {
	// write node type
	out.Uint8(uint8(pipeNodeArray))
	// write nodes len
	out.Uint8(uint8(len(p.Nodes)))
	// write nodes
	for idx := range p.Nodes {
		node := p.Nodes[idx]
		switch node.(type) {
		case *VariableNode:
			writeVariableNode(node.(*VariableNode), out)
		case *FunctionNode:
			writeFunctionNode(node.(*FunctionNode), out)
		default:
			panic("serialization error: invalid node type")
		}
	}
}

func writeFunctionNode(f *FunctionNode, out *binary.Encoder) {
	// write node type
	out.Uint8(uint8(functionNode))
	// write func name
	out.String(f.Name)
	// write parameter len
	out.Uint8(uint8(len(f.Params)))
	// write parameters
	for idx := range f.Params {
		writeParameter(f.Params[idx], out)
	}
}

func writeParameter(p *Parameter, out *binary.Encoder) {
	// write param name
	out.String(p.Name)
	// write value
	writeValue(p.Value, out)
}

func writeValue(v Value, out *binary.Encoder) {
	// write value type
	out.Uint8(uint8(v.Type()))

	switch v.Type() {
	case DictValue:
		d := v.(*DictParameterValue)
		// write entry len
		out.Uint8(uint8(len(d.Value)))
		for name, value := range d.Value {
			// write name
			out.String(name)
			// write value
			writeValue(value, out)
		}
	case ReferenceValue:
		r := v.(*ReferenceParameterValue)
		// write reference type
		out.Uint8(uint8(r.RefType))
		// write reference name
		out.String(r.Name)
	default: // base type
		b := v.(*BaseParameterValue)
		out.String(b.Value)
	}
}

func Deserialize(bytes []byte) *CompactFunction {
	decoder := binary.NewDecoder(bytes)
	return readCompactFunction(decoder)
}

func readCompactFunction(in *binary.Decoder) *CompactFunction {
	funcDef := &CompactFunction{}
	funcDef.Name = in.String()
	funcDef.Md5 = in.String()
	// read params
	sz := in.Uint8()
	funcDef.Params = make(FunctionParameters, sz)
	for i := uint8(0); i < sz; i++ {
		paramDef := &funcDef.Params[i]
		paramDef.Name = in.String()
		paramDef.Type = ValueType(in.Uint8())
		paramDef.Optional = in.Uint8() == 1
		// read const value
		constValueSz := in.Uint8()
		paramDef.ConstValue = make([]interface{}, constValueSz)
		constValueReader := getBasicTypeValueReader(paramDef.Type)
		for j := uint8(0); j < constValueSz; j++ {
			paramDef.ConstValue[j] = constValueReader(in)
		}
	}
	// read callable
	funcDef.Callable = readCompactFunctionCallable(in)

	return funcDef
}

func getBasicTypeValueReader(valueType ValueType) func(in *binary.Decoder) interface{} {
	switch valueType {
	case IntegerValue:
		return func(in *binary.Decoder) interface{} {
			return in.Int()
		}
	case FloatValue:
		return func(in *binary.Decoder) interface{} {
			return in.Float64()
		}
	case StringValue:
		return func(in *binary.Decoder) interface{} {
			return in.String()
		}
	case BoolValue:
		return func(in *binary.Decoder) interface{} {
			return in.Bool()
		}
	default:
		panic("invalid const value type")
	}
}

func readCompactFunctionCallable(in *binary.Decoder) *CompactFunctionCallable {
	c := &CompactFunctionCallable{}
	// read MultiPipe
	// read variables
	sz := in.Uint8()
	c.Pipes = &MultiPipe{Variables: make(map[string]*ImmutableValue)}
	for i := uint8(0); i < sz; i++ {
		name := in.String()
		c.Pipes.Variables[name] = NewImmutableValue()
	}
	// read pipes
	pipeSz := in.Uint16(false)
	c.Pipes.PipeList = make(Pipes, pipeSz)
	for i := uint16(0); i < pipeSz; i++ {
		c.Pipes.PipeList[i] = readPipe(in)
	}
	return c
}

func readPipe(in *binary.Decoder) *Pipe {
	pipe := &Pipe{}
	// read nodes
	nodeSz := in.Uint16(false)
	pipe.Nodes = make(PipeNodes, nodeSz)
	for j := uint16(0); j < nodeSz; j++ {
		pipe.Nodes[j] = readPipeNode(in)
	}
	return pipe
}

func readPipeNode(in *binary.Decoder) PipeNode {
	switch in.Uint8() {
	case streamNode:
		n := &StreamNode{}
		n.Splitter = readFunctionNode(in)
		n.Processor = readPipe(in)
		n.Collector = readFunctionNode(in)
		return n
	case pipeNodeArray:
		n := &PipeNodeArray{}
		nodeSz := in.Uint8()
		n.Nodes = make(PipeNodes, nodeSz)
		for i := uint8(0); i < nodeSz; i++ {
			switch in.Uint8() {
			case variableNode:
				n.Nodes[i] = &VariableNode{Name: in.String()}
			case functionNode:
				n.Nodes[i] = readFunctionNode(in)
			default:
				panic("deserialization error: invalid node type")
			}
		}
		return n
	default:
		panic("invalid node type")
	}
}

func readFunctionNode(in *binary.Decoder) *FunctionNode {
	n := &FunctionNode{}
	n.Name = in.String()
	// read parameters
	sz := in.Uint8()
	n.Params = make(Parameters, sz)
	for i := uint8(0); i < sz; i++ {
		n.Params[i] = readParameter(in)
	}
	return n
}

func readParameter(in *binary.Decoder) *Parameter {
	p := &Parameter{}
	p.Name = in.String()
	p.Value = readValue(in)
	return p
}

func readValue(in *binary.Decoder) Value {
	valueType := in.Uint8()
	switch valueType {
	case DictValue:
		d := &DictParameterValue{Value: make(map[string]Value)}
		sz := in.Uint8()
		for i := uint8(0); i < sz; i++ {
			name := in.String()
			d.Value[name] = readValue(in)
		}
		return d
	case ReferenceValue:
		r := &ReferenceParameterValue{}
		// read reference type
		r.RefType = ValueType(in.Uint8())
		// read reference name
		r.Name = in.String()
		// TODO: refer to variables
		return r
	default:
		b := &BaseParameterValue{}
		b.ValueType = ValueType(valueType)
		b.Value = in.String()
		return b
	}
}

// implement BinarySizer

func (c *CompactFunction) Size() int {
	// func name + md5 + params + callable
	return binary.Sizeof(c.Name) + binary.Sizeof(c.Md5) + c.Params.Size() + c.Callable.Size()
}

func (fp FunctionParameters) Size() int {
	// param len byte + params
	sz := 1
	for idx := range fp {
		sz += fp[idx].Size()
	}
	return sz
}

func (p *ParameterDefinition) Size() int {
	// name sz + type byte + optional type + const value len byte
	sz := binary.Sizeof(p.Name) + 1 + 1 + 1
	for idx := range p.ConstValue {
		sz += binary.Sizeof(p.ConstValue[idx])
	}
	return sz
}

func (c *CompactFunctionCallable) Size() int {
	return c.Pipes.Size()
}

func (m *MultiPipe) Size() int {
	// variable len byte + variables + pipes
	sz := 1
	for name := range m.Variables {
		sz += binary.Sizeof(name)
	}
	sz += m.PipeList.Size()
	return sz
}

func (p Pipes) Size() int {
	// len byte 2 + pipes
	sz := 2
	for idx := range p {
		sz += p[idx].Size()
	}
	return sz
}

func (p *Pipe) Size() int {
	return p.Nodes.Size()
}

func (p PipeNodes) Size() int {
	// len byte 2 + nodes
	sz := 2
	for idx := range p {
		sz += p[idx].Size()
	}
	return sz
}

func (v *VariableNode) Size() int {
	// node type + variable name
	return 1 + binary.Sizeof(v.Name)
}

func (s *StreamNode) Size() int {
	// node type + slitter + processor + collector
	return 1 + s.Splitter.Size() + s.Processor.Size() + s.Collector.Size()
}

func (m *PipeNodeArray) Size() int {
	// node type + nodes
	return 1 + m.Nodes.Size()
}

func (f *FunctionNode) Size() int {
	// node type + func name + params
	return 1 + binary.Sizeof(f.Name) + f.Params.Size()
}

func (p Parameters) Size() int {
	// param len byte + params
	sz := 1
	for idx := range p {
		sz += p[idx].Size()
	}
	return sz
}

func (p *Parameter) Size() int {
	// name + type byte + value
	return binary.Sizeof(p.Name) + p.Value.Size()
}

func (b *BaseParameterValue) Size() int {
	// value type + value
	return 1 + binary.Sizeof(b.Value)
}

func (d *DictParameterValue) Size() int {
	// value type + entry len byte + entries
	sz := 2
	for name, value := range d.Value {
		sz += binary.Sizeof(name) + value.Size()
	}
	return sz
}

func (r *ReferenceParameterValue) Size() int {
	// value type + refType byte + name
	return 2 + binary.Sizeof(r.Name)
}
