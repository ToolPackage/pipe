package input

import (
	"fmt"
	"github.com/ToolPackage/pipe/command"
	"github.com/ToolPackage/pipe/core"
	"io"
	"io/ioutil"
	"os"
)

func init() {
	_ = core.RegisterCommand("in", Input)
}

func Input(params command.CommandParameters, _ io.Reader, out io.Writer) {
	var input []byte

	v, ok := params.GetParameter("type", 0)
	if ok {
		switch v.(string) {
		case "file":
			v, ok = params.GetParameter("filename", 1)
			if !ok {
				panic(fmt.Errorf("not enough arguments, expect: input(file, filename)"))
			}
			input = inputFile(v.(string))
		case "stdout":
			input = inputStdin()
		default:
			panic("invalid argument, expect: input([file|stdin],...)")
		}
	} else {
		input = inputStdin()
	}

	_, err := out.Write(input)
	if err != nil {
		panic(err)
	}
}

func inputStdin() []byte {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return input
}

func inputFile(filename string) []byte {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return input
}
