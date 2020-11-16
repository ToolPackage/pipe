package base64

import (
	"encoding/base64"
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
	"io/ioutil"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefFunc("base64.encode", encode, DefParams()),
		DefFunc("base64.decode", decode, DefParams()),
	)
}

// base64.encode()
func encode(_ Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	encoded := base64.StdEncoding.EncodeToString(input)

	_, err = out.Write([]byte(encoded))
	return err
}

// base64.decode()
func decode(_ Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	decoded, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		panic(err)
	}

	_, err = out.Write([]byte(decoded))
	return err
}
