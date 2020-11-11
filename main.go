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
	core.RegisterCommand("base64.encode", base64.Encode)
	core.RegisterCommand("base64.decode", base64.Decode)

	core.RegisterCommand("gzip.compress", gzip.Compress)
	core.RegisterCommand("gzip.decompress", gzip.Decompress)

	core.RegisterCommand("json.pretty", json.Pretty)
	core.RegisterCommand("json.get", json.Get)

	core.RegisterCommand("in", input.Input)
	core.RegisterCommand("out", output.Output)
}
