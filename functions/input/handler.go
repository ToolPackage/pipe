package input

import (
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
	"io/ioutil"
	"os"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefBuiltinFunc("in", readFromStdin, DefParams(
			DefParam(StringValue, "type", true, "stdin"),
		)),
		DefBuiltinFunc("in.file", readFromFile, DefParams(
			DefParam(StringValue, "name", false),
		)),
		DefBuiltinFunc("in.text", readFromText, DefParams(
			DefParam(StringValue, "value", false),
		)),
	)
}

// in(): read input from stdin
func readFromStdin(_ Parameters, _ io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	_, err = out.Write(input)
	return err
}

// in.file(name: string): read input from file
func readFromFile(params Parameters, _ io.Reader, out io.Writer) error {
	v, ok := params.GetParameter("name", 0)
	if !ok {
		return NotEnoughParameterError
	}
	filename := v.Value.Get().(string)
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	_, err = out.Write(input)
	return err
}

// in.text(value: string): use text value as input
func readFromText(params Parameters, _ io.Reader, out io.Writer) error {
	v, ok := params.GetParameter("value", 0)
	if !ok {
		return NotEnoughParameterError
	}
	value := v.Value.Get().(string)
	_, err := out.Write([]byte(value))
	return err
}
