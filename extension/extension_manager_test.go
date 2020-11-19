package extension

import (
	"fmt"
	"github.com/ToolPackage/pipe/functions"
	"github.com/ToolPackage/pipe/parser"
	"github.com/ToolPackage/pipe/parser/definition"
	"testing"
)

func TestCoding(t *testing.T) {
	functions.Register()
	script := parser.ParseScript("./test.pipe")
	bytes := make([][]byte, 0)
	for idx, funcDef := range script.Funcs {
		bytes[idx] = definition.Serialize(&funcDef)
		fmt.Println(funcDef.String())
	}
	result := make([]*definition.CompactFunction, 0)
	for idx := range bytes {
		funcDef := definition.Deserialize(bytes[idx])
		result = append(result, funcDef)
		fmt.Println(funcDef.String())
	}
	// TODO: compare result with input
}
