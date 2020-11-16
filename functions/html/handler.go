package html

import (
	f "github.com/ToolPackage/pipe/functions"
	"github.com/yosssi/gohtml"
	"io"
	"io/ioutil"
)

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("html.pretty", pretty, f.DefParams()),
	)
}

// html.pretty()
func pretty(_ f.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	output := gohtml.Format(string(input))
	_, err = out.Write([]byte(output))
	return err
}
