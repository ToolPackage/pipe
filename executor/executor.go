package executor

import (
	"fmt"
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
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/ToolPackage/pipe/registry"
	"github.com/ToolPackage/pipe/util"
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

func Execute(params []string, streamMode bool) error {
	cmd := strings.Trim(strings.Join(params, ""), " ")
	if len(cmd) == 0 {
		registry.PrintFunctionUsages()
		return nil
	}

	pipe := parser.ParseMultiPipe(cmd)

	if streamMode {
		return executeStreamPipe(pipe)
	}

	in := NewDualChannel()
	out := NewDualChannel()
	for _, pipe := range pipe.Pipes {
		for _, node := range pipe.Nodes {
			if err := node.Exec(in, out); err != nil {
				f, ok := node.(*FunctionNode)
				if ok {
					switch err {
					case NotEnoughParameterError:
						return fmt.Errorf("not enough parameters, function usage: %s", util.FuncDescription(f.Handler))
					case InvalidTypeOfParameterError:
						return fmt.Errorf("invalid type of parameter, function usage: %s", util.FuncDescription(f.Handler))
					}
				}
				return err
			}
			in, out = out, in
			out.Reset()
		}
	}

	return nil
}

func executeStreamPipe(pipe *MultiPipe) error {
	return nil
}
