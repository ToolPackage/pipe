package json

import (
	"bytes"
	"encoding/json"
	"github.com/ToolPackage/pipe/commands"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
)

func Pretty(_ commands.CommandParameters, in io.Reader, out io.Writer) {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, input, "", "\t"); err != nil {
		panic(err)
	}

	_, err = out.Write(prettyJSON.Bytes())
	if err != nil {
		panic(err)
	}
}

func Get(params commands.CommandParameters, in io.Reader, out io.Writer) {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	v, ok := params.GetParameter("path", 0)
	if !ok {
		panic("not enough arguments, expect: json.get(path)")
	}
	value := gjson.Get(string(input), v.(string))

	_, err = out.Write([]byte(value.String()))
	if err != nil {
		panic(err)
	}
}
