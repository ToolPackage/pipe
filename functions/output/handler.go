package output

import (
	"github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
	"os"
)
import f "github.com/ToolPackage/pipe/functions"

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("out", outputToStdout, f.DefParams()),
		f.DefFunc("out.file", outputToFile, f.DefParams(
			f.DefParam(f.StringValue, "name", false),
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
