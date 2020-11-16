package gzip

import (
	"bytes"
	"compress/gzip"
	"github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
)

func Register() []*functions.FunctionDefinition {
	return functions.DefFuncs(
		functions.DefFunc("gzip.compress", compress, functions.DefParams()),
		functions.DefFunc("gzip.decompress", decompress, functions.DefParams()),
	)
}

// gzip.compress()
func compress(_ functions.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	w := gzip.NewWriter(out)
	if _, err = w.Write(input); err != nil {
		return err
	}
	return w.Close()
}

// gzip.decompress()
func decompress(_ functions.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	r, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return err
	}
	uncompressed, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	_, err = out.Write(uncompressed)
	return err
}
