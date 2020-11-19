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
	Parallel bool `long:"parallel" short:"p" description:"Run in parallel mode"`
}

func (r *RunSubCommands) Execute(args []string) error {
	return executor.Execute(args, r.Parallel)
}

type InstallSubCommands struct {
	Global bool   `long:"global" short:"g" description:"install script to central library"`
	File   string `long:"file" short:"f" description:"script file name" required:"false"`
}

func (i *InstallSubCommands) Execute(args []string) error {
	filename := strings.Trim(i.File, " ")
	if len(filename) == 0 {
		if len(args) == 0 {
			return fmt.Errorf("script filename was not specified")
		}
		filename = args[0]
	}
	return extension.Install(filename, i.Global)
}

type UsageSubCommands struct{}

func (u *UsageSubCommands) Execute(args []string) error {
	if len(args) > 0 {
		for _, funcName := range args {
			if err := registry.PrintFunctionUsage(funcName); err != nil {
				return err
			}
		}
	} else {
		registry.PrintFunctionUsages()
	}
	return nil
}

type VersionSubCommands struct{}

func (v *VersionSubCommands) Execute(args []string) error {
	fmt.Printf("pipe version %s-%s\n", version, buildId)
	return nil
}
