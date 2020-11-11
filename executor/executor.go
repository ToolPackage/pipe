package executor

import (
	"fmt"
	"github.com/ToolPackage/pipe/commands"
	"github.com/ToolPackage/pipe/parser"
	"github.com/ToolPackage/pipe/util"
	"os"
	"strings"
)

func Execute() {
	cmds := parser.ParseCommands(strings.Join(os.Args[1:], ""))

	in := NewDualChannel()
	out := NewDualChannel()
	for _, cmd := range cmds {
		if err := cmd.Exec(in, out); err != nil {
			if err == commands.NotEnoughParameterError {
				fmt.Print("not enough parameters, usage: ", util.FuncDescription(cmd.Handler))
			} else {
				fmt.Println(err.Error())
			}
			os.Exit(1)
		}

		in, out = out, in
		out.Reset()
	}
}
