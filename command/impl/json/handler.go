package json

import (
	"bytes"
	"encoding/json"
	"github.com/ToolPackage/pipe/command"
	"github.com/ToolPackage/pipe/core"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
)

func init() {
	_ = core.RegisterCommand("json.pretty", Pretty)
	_ = core.RegisterCommand("json.get", Get)
}

func Pretty(_ command.CommandParameters, in io.Reader, out io.Writer) {
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

func Get(params command.CommandParameters, in io.Reader, out io.Writer) {
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
