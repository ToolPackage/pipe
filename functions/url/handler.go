package url

import (
	f "github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
	urlib "net/url"
)

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("url.encode", encode, f.DefParams()),
		f.DefFunc("url.decode", decode, f.DefParams()),
	)
}

// url.encode()
func encode(_ f.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	decodedValue := urlib.QueryEscape(string(input))
	_, err = out.Write([]byte(decodedValue))
	return err
}

// url.decode()
func decode(_ f.Parameters, in io.Reader, out io.Writer) error {
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
