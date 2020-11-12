package output

import (
	"github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
	"os"
)

// out(type: [file|stdout] string, filename: string)
func Output(params functions.Parameters, in io.Reader, _ io.Writer) error {
	// validate parameters
	var opType = "stdout"
	var filename string
	v, ok := params.GetParameter("type", 0)
	if ok {
		opType = v.Value.Get().(string)
		if opType == "file" {
			v, ok = params.GetParameter("filename", 1)
			if !ok {
				return functions.NotEnoughParameterError
			}
			filename = v.Value.Get().(string)
		} else if opType != "stdout" {
			return functions.IllegalParameterError
		}
	}

	var input []byte
	var err error

	// read input
	input, err = ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	// output
	if opType == "file" {
		err = ioutil.WriteFile(filename, input, 0666)
	} else {
		_, err = os.Stdout.Write(input)
	}
	return err
}
