package output

import (
	"github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
	"os"
)

func Register() []*functions.FunctionDefinition {
	return functions.DefFuncs(
		functions.DefFunc("out", outputToStdout, functions.DefParams()),
		functions.DefFunc("out.file", outputToFile, functions.DefParams(
			functions.DefParam(functions.StringValue, "name", false),
		)),
	)
}

// out(): output pipe data to stdout
func outputToStdout(_ functions.Parameters, in io.Reader, _ io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(input)
	return err
}

// out.file(name: string): output pipe data to file
func outputToFile(params functions.Parameters, in io.Reader, _ io.Writer) error {
	v, ok := params.GetParameter("type", 0)
	if !ok {
		return functions.NotEnoughParameterError
	}
	filename := v.Value.Get().(string)
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, input, 0666)
}
