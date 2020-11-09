package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
)

func main() {
	switch os.Args[1] {
	case "--compress":
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		w := gzip.NewWriter(os.Stdout)
		if _, err = w.Write(input); err != nil {
			panic(err)
		}
		if err = w.Close(); err != nil {
			panic(err)
		}
	case "--uncompress":
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		r, err := gzip.NewReader(bytes.NewReader(input))
		uncompressed, err := ioutil.ReadAll(r)
		if err != nil {
			panic(err)
		}

		_, err = os.Stdout.Write(uncompressed)
		if err != nil {
			panic(err)
		}
	default:
		panic("invalid argument")
	}
}
