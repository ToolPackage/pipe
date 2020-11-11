package base64

import (
	"encoding/base64"
	"github.com/ToolPackage/pipe/command"
	"github.com/ToolPackage/pipe/core"
	"io"
	"io/ioutil"
)

func init() {
	_ = core.RegisterCommand("base64.encode", Encode)
	_ = core.RegisterCommand("base64.decode", Decode)
}

func Encode(_ command.CommandParameters, in io.Reader, out io.Writer) {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	encoded := base64.StdEncoding.EncodeToString(input)

	_, err = out.Write([]byte(encoded))
	if err != nil {
		panic(err)
	}
}

func Decode(_ command.CommandParameters, in io.Reader, out io.Writer) {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	decoded, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		panic(err)
	}

	_, err = out.Write([]byte(decoded))
	if err != nil {
		panic(err)
	}
}
