package executor

import (
	"errors"
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
	"runtime"
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

func Execute(params []string, parallel bool) error {
	cmd := strings.Trim(strings.Join(params, ""), " ")
	if len(cmd) == 0 {
		registry.PrintFunctionUsages()
		return nil
	}

	multiPipe := parser.ParseMultiPipe(cmd)

	if parallel {
		runtime.GOMAXPROCS(runtime.NumCPU())

		syncList := make([]chan error, len(multiPipe.Pipes))
		for idx := range multiPipe.Pipes {
			syncList[idx] = make(chan error, 1)
			go runPipe(&multiPipe.Pipes[idx], syncList[idx])
		}

		errMsg := strings.Builder{}
		for idx, sync := range syncList {
			err := <-sync
			if err != nil {
				errMsg.WriteString(fmt.Sprintf("pipe [%d] error: %v\n", idx+1, err))
			}
		}
		if errMsg.Len() > 0 {
			return errors.New(errMsg.String())
		}
	} else {
		for idx := range multiPipe.Pipes {
			if err := runPipeSync(&multiPipe.Pipes[idx]); err != nil {
				return fmt.Errorf("pipe [%d] error: %v\n", idx+1, err)
			}
		}
	}

	return nil
}

func runPipe(pipe *Pipe, sync chan error) {
	var err error
	defer func() {
		if err != nil {
			sync <- err
			return
		}
		if msg := recover(); msg != nil {
			sync <- fmt.Errorf("unexpected error: %v", msg)
			return
		}
		sync <- nil
	}()
	err = runPipeSync(pipe)
}

func runPipeSync(pipe *Pipe) error {
	in := NewDualChannel()
	out := NewDualChannel()
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
	return nil
}
