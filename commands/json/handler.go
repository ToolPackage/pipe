package json

import (
	"bytes"
	"encoding/json"
	"github.com/ToolPackage/pipe/commands"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
)

func Pretty(_ commands.CommandParameters, in io.Reader, out io.Writer) error {
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

// json.get(path)
func Get(params commands.CommandParameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	v, ok := params.GetParameter("path", 0)
	if !ok {
		return commands.NotEnoughParameterError
	}
	value := gjson.Get(string(input), v.(string))

	_, err = out.Write([]byte(value.String()))
	return err
}
