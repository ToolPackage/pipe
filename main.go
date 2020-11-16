package main

import (
	"github.com/ToolPackage/pipe/cmd"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	_, err := flags.Parse(&cmd.Opts)
	if err != nil {
		os.Exit(1)
	}
}
