package extension

import (
	"fmt"
	"github.com/ToolPackage/pipe/parser"
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/ToolPackage/pipe/registry"
	"github.com/vipally/binary"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const storageDirName = ".pipe"
const funcBinarySuffix = ".pipe"

var scriptStoragePath = getStoragePath()
var funcDefMapping = make(map[string]*CompactFunction)

// pipe node type
const (
	variableNode = iota
	streamNode
	pipeNodeArray
	functionNode
)

func init() {
	if err := loadLibraries(); err != nil {
		panic(err)
	}
}

func loadLibraries() error {
	files, err := ioutil.ReadDir(scriptStoragePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, fileInfo := range files {
		if strings.HasSuffix(fileInfo.Name(), funcBinarySuffix) {
			bytes, err := ioutil.ReadFile(filepath.Join(scriptStoragePath, fileInfo.Name()))
			if err != nil {
				return err
			}
			funcDef := deserializeFunction(bytes)
			// register function to registry
			registry.RegisterFunction(DefLibFunc(funcDef.Name, funcDef.Callable.Exec, funcDef.Params))
			// add function to mapping for convenience to query function md5
			funcDefMapping[funcDef.Name] = funcDef
		}
	}

	return nil
}

func Install(filename string) error {
	pipeScript := parser.ParseScript(filename)
	// TODO: loading animation
	for idx := range pipeScript.Funcs {
		funcDef := pipeScript.Funcs[idx]
		old, ok := funcDefMapping[funcDef.Name]
		// save function when:
		// - if old version definition is found and their md5 are not equal
		// - no old version definition is found
		if ok && funcDef.Md5 == old.Md5 {
			fmt.Println("no update on function", funcDef.Name)
			continue
		}
		// convert to byte code and save to file
		bytes := serializeFunction(&funcDef)
		if err := ioutil.WriteFile(filepath.Join(storageDirName, funcDef.Name+funcBinarySuffix),
			bytes, 0666); err != nil {
			return err
		}
	}
	return nil
}

func serializeFunction(funcDef *CompactFunction) []byte {
	encoder := binary.NewEncoder(0)
	writeCompactFunction(funcDef, encoder)
	return encoder.Buffer()
}

func writeCompactFunction(c *CompactFunction, out *binary.Encoder) {
	// write func name length
	out.Uint8(uint8(len(c.Name)))
	// write func name
	out.String(c.Name)
	// write func param length
	out.Uint8(uint8(len(c.Name)))
	// write func params
	for idx := range c.Params {
		writeParameterDefinition(&c.Params[idx], out)
	}
	// write func md5 (16)
	out.String(c.Md5)
	// write callable
	writeCompactFunctionCallable(c.Callable, out)
}

func writeParameterDefinition(p *ParameterDefinition, out *binary.Encoder) {
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

func writeCompactFunctionCallable(c *CompactFunctionCallable, out *binary.Encoder) {
	writeMultiPipe(c.Pipes, out)
}

func writeMultiPipe(m *MultiPipe, out *binary.Encoder) {
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
		writePipe(&m.Pipes[idx], out)
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
	// write variable name len
	out.Uint8(uint8(len(v.Name)))
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
	// write func name len
	out.Uint8(uint8(len(f.Name)))
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
	// write param name len
	out.Uint8(uint8(len(p.Name)))
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
		for name, value := range d.Value {
			// write name len
			out.Uint8(uint8(len(name)))
			// write name
			out.String(name)
			// write value
			writeValue(value, out)
		}
	case ReferenceValue:
		r := v.(*ReferenceParameterValue)
		// write reference type
		out.Uint8(uint8(r.RefType))
		// write reference name len
		out.Uint8(uint8(len(r.Name)))
		// write reference name
		out.String(r.Name)
	default: // base type
		b := v.(*BaseParameterValue)
		out.Uint8(uint8(len(b.Value)))
		out.String(b.Value)
	}
}

func deserializeFunction(bytes []byte) *CompactFunction {
	decoder := binary.NewDecoder(bytes)
	funcDef := &CompactFunction{}
	funcDef.ReadFrom(decoder)
	return funcDef
}

func getStoragePath() string {
	return filepath.Join(getUserHomeDir(), storageDirName)
}

func getUserHomeDir() string {
	home := os.Getenv("HOME")
	if home != "" {
		return home
	}

	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
		if home != "" {
			return home
		}

		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home != "" {
			return home
		}
	}

	panic("could not detect home directory")
}
