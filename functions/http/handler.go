package http

import (
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
	"io/ioutil"
	"net/http"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefFunc("http.get", get, DefParams(
			DefParam(StringValue, "url", false),
			DefParam(DictValue, "headers", true),
			DefParam(StringValue, "outputMode", true, "body", "raw"),
		)),
	)
}

// http.get(url: string, headers?: dict, outputMode?: 'body' | 'raw'): create http get request
func get(params Parameters, in io.Reader, out io.Writer) error {
	url, _ := params.GetParameter("url", 0)
	req, err := http.NewRequest("GET", url.Value.Get().(string), nil)
	if err != nil {
		return err
	}

	// add headers
	v, ok := params.GetParameter("headers", 1)
	if ok {
		headers := v.Value.(*DictParameterValue)
		for key, value := range headers.Value {
			if value.Type() != StringValue {
				return InvalidTypeOfParameterError
			}
			req.Header.Add(key, value.Get().(string))
		}
	}

	client := &http.Client{}
	rep, err := client.Do(req)
	if err != nil {
		return err
	}

	var output []byte
	outputRaw := false
	v, ok = params.GetParameter("outputMode", 2)
	if ok {
		outputRaw = v.Value.Get().(string) == "raw"
	}
	if outputRaw {
		// TODO:
	} else {
		if output, err = ioutil.ReadAll(rep.Body); err != nil {
			return err
		}
	}

	_, err = out.Write(output)
	return err
}
