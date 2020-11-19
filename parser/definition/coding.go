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
	for _, v := range p.ConstValue {
		// write value
		switch v.(type) {
		case int:
			out.Int(v.(int))
		case float64:
			out.Float64(v.(float64))
		case string:
			out.String(v.(string))
		case bool:
			out.Bool(v.(bool))
			// TODO:
		}
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
		writePipe(&m.PipeList[idx], out)
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
	if v, ok := p.(*VariableNode); ok {
		writeVariableNode(v, out)
		return
	}
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
	writePipe(&s.Processor, out)
	writeFunctionNode(s.Collector, out)
}

func writePipeNodeArray(p *PipeNodeArray, out *binary.Encoder) {
	// write node type
	out.Uint8(uint8(pipeNodeArray))
	// write nodes len
	out.Uint8(uint8(len(p.Nodes)))
	// write nodes
	for idx := range p.Nodes {
		writePipeNode(p.Nodes[idx], out)
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
		writeParameter(&f.Params[idx], out)
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
	// TODO:
	funcDef := &CompactFunction{}
	//len := in.Uint8()
	return funcDef
}

// implement BinarySizer

func (c *CompactFunction) Size() int {
	return binary.Sizeof(c.Name) + binary.Sizeof(c.Md5) + c.Params.Size() + c.Callable.Size()
}

func (fp FunctionParameters) Size() int {
	var sz = 1
	for idx := range fp {
		sz += fp[idx].Size()
	}
	return sz
}

func (p *ParameterDefinition) Size() int {
	var sz int
	// name sz + type byte + optional type + const value len byte
	sz += binary.Sizeof(p.Name) + 1 + 1 + 1
	for idx := range p.ConstValue {
		sz += binary.Sizeof(p.ConstValue[idx])
	}
	return sz
}

func (c *CompactFunctionCallable) Size() int {
	return c.Pipes.Size()
}

func (m *MultiPipe) Size() int {
	sz := 1 // variables len byte
	for name := range m.Variables {
		sz += binary.Sizeof(name)
	}
	sz += m.PipeList.Size()
	return sz
}

func (p Pipes) Size() int {
	sz := 2 // pipe len byte
	for idx := range p {
		sz += p[idx].Size()
	}
	return sz
}

func (p *Pipe) Size() int {
	return p.Nodes.Size()
}

func (p PipeNodes) Size() int {
	sz := 2 // len byte
	for idx := range p {
		sz += p[idx].Size()
	}
	return sz
}

func (v *VariableNode) Size() int {
	return 1 + binary.Sizeof(v.Name)
}

func (s *StreamNode) Size() int {
	return 1 + s.Splitter.Size() + s.Processor.Size() + s.Collector.Size()
}

func (m *PipeNodeArray) Size() int {
	return 1 + m.Nodes.Size()
}

func (f *FunctionNode) Size() int {
	return binary.Sizeof(f.Name) + f.Params.Size()
}

func (p Parameters) Size() int {
	sz := 1
	for idx := range p {
		sz += p[idx].Size()
	}
	return sz
}

func (p *Parameter) Size() int {
	// name + type byte + value
	return binary.Sizeof(p.Name) + 1 + p.Value.Size()
}

func (b *BaseParameterValue) Size() int {
	return binary.Sizeof(b.Value)
}

func (d *DictParameterValue) Size() int {
	sz := 1
	for name, value := range d.Value {
		sz += binary.Sizeof(name) + value.Size()
	}
	return sz
}

func (r *ReferenceParameterValue) Size() int {
	// refType byte + name
	return 1 + binary.Sizeof(r.Name)
}
