package output

import (
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
	"io/ioutil"
	"os"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefBuiltinFunc("out", outputToStdout, DefParams()),
		DefBuiltinFunc("out.file", outputToFile, DefParams(
			DefParam(StringValue, "name", false),
		)),
	)
}

// out(): output pipe data to stdout
func outputToStdout(_ Parameters, in io.Reader, _ io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(input)
	return err
}

// out.file(name: string): output pipe data to file
func outputToFile(params Parameters, in io.Reader, _ io.Writer) error {
	v, ok := params.GetParameter("type", 0)
	if !ok {
		return NotEnoughParameterError
	}
	filename := v.Value.Get().(string)
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, input, 0666)
}
