package gzip

import (
	"bytes"
	"compress/gzip"
	"github.com/ToolPackage/pipe/command"
	"github.com/ToolPackage/pipe/core"
	"io"
	"io/ioutil"
)

func init() {
	_ = core.RegisterCommand("gzip.compress", Compress)
	_ = core.RegisterCommand("gzip.decompress", Decompress)
}

func Compress(_ command.CommandParameters, in io.Reader, out io.Writer) {
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

func Decompress(_ command.CommandParameters, in io.Reader, out io.Writer) {
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
