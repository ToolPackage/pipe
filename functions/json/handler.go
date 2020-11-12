package json

import (
	"bytes"
	"encoding/json"
	"github.com/ToolPackage/pipe/functions"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
)

// json.pretty()
func Pretty(_ functions.Parameters, in io.Reader, out io.Writer) error {
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

// json.get(path: string)
func Get(params functions.Parameters, in io.Reader, out io.Writer) error {
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
