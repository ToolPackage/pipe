package output

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Output(args []string, in io.Reader, _ io.Writer) {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	if len(args) > 0 {
		if args[0] == "file" {
			if len(args) < 2 {
				panic(fmt.Errorf("not enough arguments, expect: output(file, filePath)"))
			}
			filename := args[1]
			outputFile(filename, input)
			return
		} else if args[0] == "stdout" {
			outputStdout(input)
		} else {
			panic("invalid argument, expect: output([file|stdout],...)")
		}
	}
	outputStdout(input)
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
