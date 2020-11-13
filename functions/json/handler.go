package json

import (
	"bytes"
	"encoding/json"
	"github.com/ToolPackage/pipe/functions"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
)

import f "github.com/ToolPackage/pipe/functions"

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("json.pretty", pretty, f.DefParams()),
		f.DefFunc("json.get", get, f.DefParams(
			f.DefParam(f.StringValue, "path", false),
		)),
	)
}

// json.pretty(): pretty json input
func pretty(_ functions.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, input, "", "\t"); err != nil {
		return err
	}

	_, err = out.Write(prettyJSON.Bytes())
	return err
}

// json.get(path: string): get value from json input by path
func get(params functions.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	v, ok := params.GetParameter("path", 0)
	if !ok {
		return functions.NotEnoughParameterError
	}
	value := gjson.Get(string(input), v.Value.Get().(string))

	_, err = out.Write([]byte(value.String()))
	return err
}
