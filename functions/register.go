package functions

import (
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
	"github.com/ToolPackage/pipe/functions/unix"
	"github.com/ToolPackage/pipe/functions/url"
	"github.com/ToolPackage/pipe/registry"
	"sync"
)

var o sync.Once

func LoadBuiltinFunctions() {
	o.Do(func() {
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
		registry.RegisterFunctions(unix.Register())
	})
}
