package json

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

func Pretty(_ []string, in io.Reader, out io.Writer) {
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
