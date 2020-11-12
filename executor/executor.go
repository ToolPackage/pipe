package executor

import (
	"fmt"
	"github.com/ToolPackage/pipe/functions"
	"github.com/ToolPackage/pipe/functions/base64"
	"github.com/ToolPackage/pipe/functions/gzip"
	"github.com/ToolPackage/pipe/functions/input"
	"github.com/ToolPackage/pipe/functions/json"
	"github.com/ToolPackage/pipe/functions/output"
	"github.com/ToolPackage/pipe/functions/regexp"
	"github.com/ToolPackage/pipe/parser"
	"github.com/ToolPackage/pipe/registry"
	"github.com/ToolPackage/pipe/util"
	"os"
	"strings"
)

func init() {
	registry.RegisterFunction("in", input.Input)
	registry.RegisterFunction("out", output.Output)

	registry.RegisterFunction("base64.encode", base64.Encode)
	registry.RegisterFunction("base64.decode", base64.Decode)

	registry.RegisterFunction("gzip.compress", gzip.Compress)
	registry.RegisterFunction("gzip.decompress", gzip.Decompress)

	registry.RegisterFunction("json.pretty", json.Pretty)
	registry.RegisterFunction("json.get", json.Get)

	registry.RegisterFunction("regexp.test", regexp.Test)
}

func Execute(params []string, streamMode bool) {
	pipe := parser.ParseScript(strings.Join(params, ""))

	in := NewDualChannel()
	out := NewDualChannel()
	for _, function := range pipe {
		if err := function.Exec(in, out); err != nil {
			switch err {
			case functions.NotEnoughParameterError:
				fmt.Print("not enough parameters, function usage: ", util.FuncDescription(function.Handler))
			case functions.InvalidTypeOfParameterError:
				fmt.Print("invalid type of parameter, function usage: ", util.FuncDescription(function.Handler))
			default:
				fmt.Println(err.Error())
			}
			os.Exit(1)
		}

		in, out = out, in
		out.Reset()
	}
}
