package extension

import (
	"github.com/ToolPackage/pipe/functions"
	"github.com/ToolPackage/pipe/parser"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCoding(t *testing.T) {
	functions.LoadBuiltinFunctions()
	script := parser.ParseScript("./test.pipe")
	bytes := make([][]byte, 2)
	expectation := strings.Builder{}
	for idx, funcDef := range script.Funcs {
		bytes[idx] = Serialize(&funcDef)
		expectation.WriteString(funcDef.String())
	}
	result := strings.Builder{}
	for idx := range bytes {
		funcDef := Deserialize(bytes[idx])
		result.WriteString(funcDef.String())
	}
	assert.Equal(t, expectation.String(), result.String())
}
