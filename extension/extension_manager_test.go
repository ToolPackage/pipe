package extension

import (
	"github.com/ToolPackage/pipe/functions"
	"github.com/ToolPackage/pipe/parser"
	"github.com/ToolPackage/pipe/parser/definition"
	"testing"
)

func TestCoding(t *testing.T) {
	functions.Register()
	script := parser.ParseScript("./test.pipe")
	definition.Serialize(&script.Funcs[0])
}
