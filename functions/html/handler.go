package html

import (
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/yosssi/gohtml"
	"io"
	"io/ioutil"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefFunc("html.pretty", pretty, DefParams()),
	)
}

// html.pretty()
func pretty(_ Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	output := gohtml.Format(string(input))
	_, err = out.Write([]byte(output))
	return err
}
