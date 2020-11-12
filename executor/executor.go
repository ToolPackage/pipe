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

func defineFunc(name string, handler functions.FunctionHandler, paramConstraint functions.FuncParamConstraint) *functions.FunctionDefinition {
	return &functions.FunctionDefinition{Name: name, Handler: handler, ParamConstraint: paramConstraint}
}

func defineParams(paramDefSeqList ...functions.ParamDefSequence) functions.FuncParamConstraint {
	return paramDefSeqList
}

func paramSequence(paramDefList ...functions.ParameterDefinition) functions.ParamDefSequence {
	return paramDefList
}

func param(paramType functions.ParameterType, optional bool, constValue ...string) functions.ParameterDefinition {
	return functions.ParameterDefinition{Type: paramType, Optional: optional, ConstValue: constValue}
}

func init() {
	registry.RegisterFunction(defineFunc("in", input.Input, defineParams(
		paramSequence(),
		paramSequence(
			param(functions.StringValue, false, "stdin"),
		),
		paramSequence(
			param(functions.StringValue, false, "file", "text"),
			param(functions.StringValue, false),
		),
	)))
	registry.RegisterFunction(defineFunc("out", output.Output, defineParams(
		paramSequence(),
		paramSequence(
			param(functions.StringValue, false, "stdin"),
		),
		paramSequence(
			param(functions.StringValue, false, "file"),
			param(functions.StringValue, false),
		),
	)))

	registry.RegisterFunction(defineFunc("base64.encode", base64.Encode, defineParams()))
	registry.RegisterFunction(defineFunc("base64.decode", base64.Decode, defineParams()))

	registry.RegisterFunction(defineFunc("gzip.compress", gzip.Compress, defineParams()))
	registry.RegisterFunction(defineFunc("gzip.decompress", gzip.Decompress, defineParams()))

	registry.RegisterFunction(defineFunc("json.pretty", json.Pretty, defineParams()))
	registry.RegisterFunction(defineFunc("json.get", json.Get, defineParams(
		paramSequence(
			param(functions.StringValue, false, "stdin"),
		),
	)))

	registry.RegisterFunction(defineFunc("regexp.test", regexp.Test, defineParams(
		paramSequence(
			param(functions.StringValue, false, "stdin"),
		),
	)))
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
