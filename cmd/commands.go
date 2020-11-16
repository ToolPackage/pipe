package cmd

import (
	"fmt"
	"github.com/ToolPackage/pipe/executor"
	"github.com/ToolPackage/pipe/extension"
	"github.com/ToolPackage/pipe/registry"
	"strings"
)

var Opts PipeCommands

type PipeCommands struct {
	Run     RunSubCommands     `command:"run" description:"run pipe script in command line"`
	Install InstallSubCommands `command:"install" description:"install pipe script as extension"`
	Usage   UsageSubCommands   `command:"usage" description:"Display function usages"`
	Version VersionSubCommands `command:"version" description:"Display the version"`
}

type RunSubCommands struct {
	Stream bool `long:"stream" short:"s" description:"Run in stream mode"`
}

func (r *RunSubCommands) Execute(args []string) error {
	return executor.Execute(args, r.Stream)
}

type InstallSubCommands struct {
	File string `long:"file" short:"f" description:"script file name" required:"false"`
}

func (i *InstallSubCommands) Execute(args []string) error {
	filename := strings.Trim(i.File, " ")
	if len(filename) == 0 {
		if len(args) == 0 {
			return fmt.Errorf("extension filename was not specified")
		}
		filename = args[0]
	}
	return extension.Install(filename)
}

type UsageSubCommands struct{}

func (u *UsageSubCommands) Execute(args []string) error {
	registry.PrintFunctionUsages()
	return nil
}

type VersionSubCommands struct{}

func (v *VersionSubCommands) Execute(args []string) error {
	// TODO:
	return nil
}
