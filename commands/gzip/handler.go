package gzip

import (
	"bytes"
	"compress/gzip"
	"github.com/ToolPackage/pipe/commands"
	"io"
	"io/ioutil"
)

func Compress(_ commands.CommandParameters, in io.Reader, out io.Writer) error {
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

func Decompress(_ commands.CommandParameters, in io.Reader, out io.Writer) error {
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
