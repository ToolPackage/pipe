package main

import (
	"github.com/ToolPackage/pipe/executor"
	"github.com/ToolPackage/pipe/registry"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var opts struct {
		Stream  bool `long:"stream" short:"s" description:"Run in stream mode"`
		Usage   bool `long:"usage" short:"u" description:"Display function usages"`
		Version bool `long:"version" short:"v" description:"Display the version"`
	}

	params, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	if opts.Usage {
		registry.PrintFunctionUsages()
		return
	}

	if opts.Version {

	}

	executor.Execute(params, opts.Stream)
}
