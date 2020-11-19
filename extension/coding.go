package extension

import (
	"fmt"
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/ToolPackage/pipe/registry"
	"github.com/vipally/binary"
)

// pipe node type
const (
	variableNode = iota
	streamNode
	pipeNodeArray
	functionNode
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
	out.Uint8(uint8(len(c.Params)))
	// write func params
	for idx := range c.Params {
		writeParameterDefinition(&c.Params[idx], out)
	}
	// write pipes
	writeMultiPipe(c.Pipes, out)
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
	if len(p.ConstValue) > 0 {
		// write param const value
		constValueWriter := getBasicTypeValueWriter(p.Type)
		for _, v := range p.ConstValue {
			constValueWriter(out, v)
		}
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
		if constValueSz > 0 {
			constValueReader := getBasicTypeValueReader(paramDef.Type)
			for j := uint8(0); j < constValueSz; j++ {
				paramDef.ConstValue[j] = constValueReader(in)
			}
		}
	}
	// read MultiPipe
	// read variables
	sz = in.Uint8()
	funcDef.Pipes = &MultiPipe{Variables: make(map[string]*ImmutableValue)}
	global.SetMultiPipe(funcDef.Pipes)
	for i := uint8(0); i < sz; i++ {
		name := in.String()
		funcDef.Pipes.Variables[name] = NewImmutableValue()
	}
	// read pipes
	pipeSz := in.Uint16(false)
	funcDef.Pipes.PipeList = make(Pipes, pipeSz)
	for i := uint16(0); i < pipeSz; i++ {
		funcDef.Pipes.PipeList[i] = readPipe(in)
	}

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
		panic(fmt.Sprintf("invalid const value type %d", valueType))
	}
}

var global CodingContext

type CodingContext struct {
	multiPipe *MultiPipe
}

func (c *CodingContext) SetMultiPipe(m *MultiPipe) {
	c.multiPipe = m
}

func (c *CodingContext) GetMultiPipe() *MultiPipe {
	return c.multiPipe
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
				node := &VariableNode{Name: in.String()}
				if v, ok := global.GetMultiPipe().Variables[node.Name]; ok {
					node.Value = v
				} else {
					node.Value = NewImmutableValue()
					global.GetMultiPipe().Variables[node.Name] = node.Value
				}
				n.Nodes[i] = node
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
	registry.LookupFunctionHandler(n)
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
		if v, ok := global.GetMultiPipe().Variables[r.Name]; ok {
			r.Value = v
		} else {
			panic(fmt.Sprintf("reference to undefined variable: %s", r.Name))
		}
		return r
	default:
		b := &BaseParameterValue{}
		b.ValueType = ValueType(valueType)
		b.Value = in.String()
		return b
	}
}
