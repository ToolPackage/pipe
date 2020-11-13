package base64

import (
	"encoding/base64"
	f "github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
)

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("base64.encode", encode, f.DefParams()),
		f.DefFunc("base64.decode", decode, f.DefParams()),
	)
}

// base64.encode()
func encode(_ f.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	encoded := base64.StdEncoding.EncodeToString(input)

	_, err = out.Write([]byte(encoded))
	return err
}

// base64.decode()
func decode(_ f.Parameters, in io.Reader, out io.Writer) error {
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
