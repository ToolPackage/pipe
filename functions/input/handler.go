package input

import (
	f "github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
	"os"
)

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("in", readFromStdin, f.DefParams(
			f.DefParam(f.StringValue, "type", true, "stdin"),
		)),
		f.DefFunc("in.file", readFromFile, f.DefParams(
			f.DefParam(f.StringValue, "name", false),
		)),
		f.DefFunc("in.text", readFromText, f.DefParams(
			f.DefParam(f.StringValue, "value", false),
		)),
	)
}

// in(): read input from stdin
func readFromStdin(_ f.Parameters, _ io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	_, err = out.Write(input)
	return err
}

// in.file(name: string): read input from file
func readFromFile(params f.Parameters, _ io.Reader, out io.Writer) error {
	v, ok := params.GetParameter("name", 0)
	if !ok {
		return f.NotEnoughParameterError
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
func readFromText(params f.Parameters, _ io.Reader, out io.Writer) error {
	v, ok := params.GetParameter("value", 0)
	if !ok {
		return f.NotEnoughParameterError
	}
	value := v.Value.Get().(string)
	_, err := out.Write([]byte(value))
	return err
}
