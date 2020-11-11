package main

import (
	"github.com/ToolPackage/pipe/commands/base64"
	"github.com/ToolPackage/pipe/commands/gzip"
	"github.com/ToolPackage/pipe/commands/input"
	"github.com/ToolPackage/pipe/commands/json"
	"github.com/ToolPackage/pipe/commands/output"
	"github.com/ToolPackage/pipe/core"
)

func main() {
	core.Execute()
}

func init() {
	_ = core.RegisterCommand("base64.encode", base64.Encode)
	_ = core.RegisterCommand("base64.decode", base64.Decode)

	_ = core.RegisterCommand("gzip.compress", gzip.Compress)
	_ = core.RegisterCommand("gzip.decompress", gzip.Decompress)

	_ = core.RegisterCommand("json.pretty", json.Pretty)
	_ = core.RegisterCommand("json.get", json.Get)

	_ = core.RegisterCommand("in", input.Input)
	_ = core.RegisterCommand("out", output.Output)
}
