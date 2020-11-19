package definition

import "github.com/vipally/binary"

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
