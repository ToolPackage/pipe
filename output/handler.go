package output

import (
	"io"
	"io/ioutil"
	"os"
)

func Output(_ []string, in io.Reader, _ io.Writer) {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	_, err = os.Stdout.Write(input)
	if err != nil {
		panic(err)
	}
}
