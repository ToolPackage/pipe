package input

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Input(args []string, _ io.Reader, out io.Writer) {
	var input []byte
	if len(args) > 0 {
		if args[0] == "file" {
			if len(args) < 2 {
				panic(fmt.Errorf("not enough arguments, expect: input(file, filename)"))
			}
			input = inputFile(args[1])
		} else if args[0] == "stdin" {
			input = inputStdin()
		} else {
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
