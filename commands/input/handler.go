package input

import (
	"github.com/ToolPackage/pipe/commands"
	"io"
	"io/ioutil"
	"os"
)

// in(file|text|stdin, filename|value)
func Input(params commands.CommandParameters, _ io.Reader, out io.Writer) error {
	var input []byte
	var err error

	// validate parameters
	var opType = "stdin"
	var param2 string
	v, ok := params.GetParameter("type", 0)
	if ok {
		opType = v.(string)
		if opType == "file" {
			v, ok = params.GetParameter("filename", 1)
			if !ok {
				return commands.NotEnoughParameterError
			}
			param2 = v.(string)
		} else if opType == "text" {
			v, ok = params.GetParameter("value", 1)
			if !ok {
				return commands.NotEnoughParameterError
			}
			param2 = v.(string)
		} else if opType != "stdin" {
			return commands.IllegalParameterError
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
