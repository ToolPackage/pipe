package url

import (
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
	"io/ioutil"
	urlib "net/url"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefBuiltinFunc("url.encode", encode, DefParams()),
		DefBuiltinFunc("url.decode", decode, DefParams()),
	)
}

// url.encode()
func encode(_ Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	decodedValue := urlib.QueryEscape(string(input))
	_, err = out.Write([]byte(decodedValue))
	return err
}

// url.decode()
func decode(_ Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	decodedValue, err := urlib.QueryUnescape(string(input))
	if err != nil {
		return err
	}
	_, err = out.Write([]byte(decodedValue))
	return err
}
