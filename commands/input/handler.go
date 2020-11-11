package input

import (
	"github.com/ToolPackage/pipe/commands"
	"io"
	"io/ioutil"
	"os"
)

// in([file|stdin], filename)
func Input(params commands.CommandParameters, _ io.Reader, out io.Writer) error {
	var input []byte
	var err error

	// validate parameters
	var opType = "stdin"
	var filename string
	v, ok := params.GetParameter("type", 0)
	if ok {
		opType = v.(string)
		if opType == "file" {
			v, ok = params.GetParameter("filename", 1)
			if !ok {
				return commands.NotEnoughParameterError
			}
			filename = v.(string)
		} else if opType != "stdin" {
			return commands.IllegalParameterError
		}
	}

	// input
	if opType == "file" {
		input, err = ioutil.ReadFile(filename)
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
	}

	if err != nil {
		return err
	}

	_, err = out.Write(input)
	return err
}
