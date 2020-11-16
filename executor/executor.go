package executor

import (
	"fmt"
	f "github.com/ToolPackage/pipe/functions"
	"github.com/ToolPackage/pipe/functions/base64"
	"github.com/ToolPackage/pipe/functions/color"
	"github.com/ToolPackage/pipe/functions/filter"
	"github.com/ToolPackage/pipe/functions/gzip"
	"github.com/ToolPackage/pipe/functions/html"
	"github.com/ToolPackage/pipe/functions/http"
	"github.com/ToolPackage/pipe/functions/input"
	"github.com/ToolPackage/pipe/functions/json"
	"github.com/ToolPackage/pipe/functions/output"
	"github.com/ToolPackage/pipe/functions/regexp"
	"github.com/ToolPackage/pipe/functions/text"
	"github.com/ToolPackage/pipe/functions/url"
	"github.com/ToolPackage/pipe/parser"
	"github.com/ToolPackage/pipe/registry"
	"github.com/ToolPackage/pipe/util"
	"os"
	"strings"
)

func init() {
	registry.RegisterFunctions(input.Register())
	registry.RegisterFunctions(output.Register())
	registry.RegisterFunctions(base64.Register())
	registry.RegisterFunctions(filter.Register())
	registry.RegisterFunctions(gzip.Register())
	registry.RegisterFunctions(color.Register())
	registry.RegisterFunctions(json.Register())
	registry.RegisterFunctions(regexp.Register())
	registry.RegisterFunctions(http.Register())
	registry.RegisterFunctions(url.Register())
	registry.RegisterFunctions(text.Register())
	registry.RegisterFunctions(html.Register())
}

func Execute(params []string, streamMode bool) {
	cmd := strings.Trim(strings.Join(params, ""), " ")
	if len(cmd) == 0 {
		registry.PrintFunctionUsages()
		return
	}

	pipe := parser.ParseMultiPipe(cmd)

	if streamMode {
		executeStreamPipe(pipe)
		return
	}

	in := NewDualChannel()
	out := NewDualChannel()
	for _, function := range pipe {
		if err := function.Exec(in, out); err != nil {
			switch err {
			case f.NotEnoughParameterError:
				fmt.Print("not enough parameters, function usage: ", util.FuncDescription(function.Handler))
			case f.InvalidTypeOfParameterError:
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

func executeStreamPipe(pipe []f.FunctionNode) {

}
