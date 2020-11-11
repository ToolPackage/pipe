package base64

import (
	"encoding/base64"
	"github.com/ToolPackage/pipe/commands"
	"io"
	"io/ioutil"
)

func Encode(_ commands.CommandParameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	encoded := base64.StdEncoding.EncodeToString(input)

	_, err = out.Write([]byte(encoded))
	return err
}

func Decode(_ commands.CommandParameters, in io.Reader, out io.Writer) error {
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
