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

	syncList := make([]chan error, len(pipe.Pipes))
	for idx, pipe := range pipe.Pipes {
		syncList[idx] = make(chan error, 0)
		go runPipe(&pipe, syncList[idx])
	}

	for idx, sync := range syncList {

	}

	return nil
}

func runPipe(pipe *Pipe, sync chan error) {
	in := NewDualChannel()
	out := NewDualChannel()
	for _, node := range pipe.Nodes {
		if err := node.Exec(in, out); err != nil {
			f, ok := node.(*FunctionNode)
			if ok {
				switch err {
				case NotEnoughParameterError:
					err = fmt.Errorf("not enough parameters, function usage: %s", util.FuncDescription(f.Handler))
				case InvalidTypeOfParameterError:
					err = fmt.Errorf("invalid type of parameter, function usage: %s", util.FuncDescription(f.Handler))
				}
			}
			sync <- err
			return
		}
		in, out = out, in
		out.Reset()
	}
}

func executeStreamPipe(pipe *MultiPipe) error {
	return nil
}
