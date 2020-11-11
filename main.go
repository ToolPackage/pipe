package main

import (
	"github.com/ToolPackage/pipe/commands/base64"
	"github.com/ToolPackage/pipe/commands/gzip"
	"github.com/ToolPackage/pipe/commands/input"
	"github.com/ToolPackage/pipe/commands/json"
	"github.com/ToolPackage/pipe/commands/output"
	"github.com/ToolPackage/pipe/executor"
	"github.com/ToolPackage/pipe/registry"
)

func main() {
	executor.Execute()
}

func init() {
	registry.RegisterCommand("base64.encode", base64.Encode)
	registry.RegisterCommand("base64.decode", base64.Decode)

	registry.RegisterCommand("gzip.compress", gzip.Compress)
	registry.RegisterCommand("gzip.decompress", gzip.Decompress)

	registry.RegisterCommand("json.pretty", json.Pretty)
	registry.RegisterCommand("json.get", json.Get)

	registry.RegisterCommand("in", input.Input)
	registry.RegisterCommand("out", output.Output)
}
