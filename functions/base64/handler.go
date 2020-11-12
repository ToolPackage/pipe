package base64

import (
	"encoding/base64"
	"github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
)

// base64.encode()
func Encode(_ functions.FunctionParameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	encoded := base64.StdEncoding.EncodeToString(input)

	_, err = out.Write([]byte(encoded))
	return err
}

// base64.decode()
func Decode(_ functions.FunctionParameters, in io.Reader, out io.Writer) error {
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
