package output

import (
	"fmt"
	"github.com/ToolPackage/pipe/commands"
	"io"
	"io/ioutil"
	"os"
)

func Output(params commands.CommandParameters, in io.Reader, _ io.Writer) {
	v, ok := params.GetParameter("type", 0)
	if ok {
		switch v.(string) {
		case "file":
			v, ok = params.GetParameter("filename", 1)
			if !ok {
				panic(fmt.Errorf("not enough arguments, expect: output(file, filename)"))
			}
			outputFile(v.(string), read(in))
		case "stdout":
			outputStdout(read(in))
		default:
			panic("invalid argument, expect: output([file|stdout],...)")
		}
	} else {
		outputStdout(read(in))
	}
}

func read(in io.Reader) []byte {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}
	return data
}

func outputStdout(data []byte) {
	_, err := os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}

func outputFile(filename string, data []byte) {
	if err := ioutil.WriteFile(filename, data, 0666); err != nil {
		panic(err)
	}
}
