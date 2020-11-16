package yaml

import (
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefFunc("yaml.get", get, DefParams(
			DefParam(StringValue, "path", false),
		)),
	)
}

// yaml.get()
func get(params Parameters, in io.Reader, out io.Writer) error {
	// TODO
	//path, _ := params.GetParameter("path", 0)
	//input, err := ioutil.ReadAll(in)
	//if err != nil {
	//	return err
	//}
	return nil
}
