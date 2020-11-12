package input

import (
	"github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
	"os"
)

// in(type: [file|text|stdin] string, filename|value: string)
func Input(params functions.Parameters, _ io.Reader, out io.Writer) error {
	var input []byte
	var err error

	// validate parameters
	var opType = "stdin"
	var param2 string
	v, ok := params.GetParameter("type", 0)
	if ok {
		if v.Value.Type() != functions.StringValue {
			return functions.InvalidTypeOfParameterError
		}
		opType = v.Value.Get().(string)

		if opType == "file" {
			v, ok = params.GetParameter("filename", 1)
			if !ok {
				return functions.NotEnoughParameterError
			}
			if v.Value.Type() != functions.StringValue {
				return functions.InvalidTypeOfParameterError
			}
			param2 = v.Value.Get().(string)
		} else if opType == "text" {
			v, ok = params.GetParameter("value", 1)
			if !ok {
				return functions.NotEnoughParameterError
			}
			if v.Value.Type() != functions.StringValue {
				return functions.InvalidTypeOfParameterError
			}
			param2 = v.Value.Get().(string)
		} else if opType != "stdin" {
			return functions.IllegalParameterError
		}
	}

	// input
	if opType == "file" {
		input, err = ioutil.ReadFile(param2)
	} else if opType == "text" {
		input = []byte(param2)
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
	}

	if err != nil {
		return err
	}

	_, err = out.Write(input)
	return err
}
