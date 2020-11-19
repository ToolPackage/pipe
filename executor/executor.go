package executor

import (
	"errors"
	"fmt"
	"github.com/ToolPackage/pipe/extension"
	"github.com/ToolPackage/pipe/functions"
	"github.com/ToolPackage/pipe/parser"
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/ToolPackage/pipe/registry"
	"github.com/ToolPackage/pipe/util"
	"runtime"
	"strings"
)

func init() {
	functions.LoadBuiltinFunctions()
	extension.LoadLibraries()
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

		syncList := make([]chan error, len(multiPipe.PipeList))
		for idx := range multiPipe.PipeList {
			syncList[idx] = make(chan error, 1)
			go runPipe(multiPipe.PipeList[idx], syncList[idx])
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
		for idx := range multiPipe.PipeList {
			if err := runPipeSync(multiPipe.PipeList[idx]); err != nil {
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
	in := util.NewDualChannel()
	out := util.NewDualChannel()
	return pipe.Exec(in, out)
}
