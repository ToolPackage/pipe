package input

import (
	"io"
	"io/ioutil"
	"os"
)

func Input(_ []string, _ io.Reader, out io.Writer) {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	_, err = out.Write(input)
	if err != nil {
		panic(err)
	}
}
