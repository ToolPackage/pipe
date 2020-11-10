package base64

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
)

func Base64Encode(args []string, in io.Reader, out io.Writer) {
	switch os.Args[1] {
	case "--encode":
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		encoded := base64.StdEncoding.EncodeToString(input)

		_, err = os.Stdout.Write([]byte(encoded))
		if err != nil {
			panic(err)
		}
	case "--decode":
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		decoded, err := base64.StdEncoding.DecodeString(string(input))
		if err != nil {
			panic(err)
		}

		_, err = os.Stdout.Write([]byte(decoded))
		if err != nil {
			panic(err)
		}
	default:
		panic("invalid argument")
	}
}
