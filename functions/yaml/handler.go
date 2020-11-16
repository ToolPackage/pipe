package yaml

import (
	f "github.com/ToolPackage/pipe/functions"
	"io"
)

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("yaml.get", get, f.DefParams(
			f.DefParam(f.StringValue, "path", false),
		)),
	)
}

// yaml.get()
func get(params f.Parameters, in io.Reader, out io.Writer) error {
	// TODO
	//path, _ := params.GetParameter("path", 0)
	//input, err := ioutil.ReadAll(in)
	//if err != nil {
	//	return err
	//}
	return nil
}
