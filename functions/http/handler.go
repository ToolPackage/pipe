package http

import (
	f "github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
	"net/http"
)

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("http.get", get, f.DefParams(
			f.DefParam(f.StringValue, "url", false),
			f.DefParam(f.DictValue, "headers", true),
			f.DefParam(f.StringValue, "outputMode", true, "body", "raw"),
		)),
	)
}

func get(params f.Parameters, in io.Reader, out io.Writer) error {
	url, _ := params.GetParameter("url", 0)
	req, err := http.NewRequest("GET", url.Value.Get().(string), nil)
	if err != nil {
		return err
	}

	// add headers
	v, ok := params.GetParameter("headers", 1)
	if ok {
		headers := v.Value.(*f.DictParameterValue)
		for key, value := range headers.Value {
			if value.Type() != f.StringValue {
				return f.InvalidTypeOfParameterError
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