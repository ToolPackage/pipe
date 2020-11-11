package gzip

import (
	"bytes"
	"compress/gzip"
	"github.com/ToolPackage/pipe/commands"
	"io"
	"io/ioutil"
)

func Compress(_ commands.CommandParameters, in io.Reader, out io.Writer) {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	w := gzip.NewWriter(out)
	if _, err = w.Write(input); err != nil {
		panic(err)
	}
	if err = w.Close(); err != nil {
		panic(err)
	}
}

func Decompress(_ commands.CommandParameters, in io.Reader, out io.Writer) {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	r, err := gzip.NewReader(bytes.NewReader(input))
	uncompressed, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	_, err = out.Write(uncompressed)
	if err != nil {
		panic(err)
	}
}
