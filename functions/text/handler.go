package text

import (
	"errors"
	"github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
)

func Register() []*functions.FunctionDefinition {
	return functions.DefFuncs(
		functions.DefFunc("text.cut", cut, functions.DefParams(
			functions.DefParam(functions.IntegerValue, "start", false),
			functions.DefParam(functions.IntegerValue, "end", true),
		)),
	)
}

// text.cut(start: int, end?: int): extract substring
func cut(params functions.Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	var start int
	var end int
	v, _ := params.GetParameter("start", 0)
	start = v.Value.Get().(int)
	v, ok := params.GetParameter("end", 1)
	if ok {
		end = v.Value.Get().(int)
	} else {
		end = len(input)
	}

	if start < 0 {
		start += len(input)
	}
	if end < 0 {
		end += len(input)
	}

	if start >= end || start < 0 ||
		end <= start || end > len(input) {
		return errors.New("index out of bound")
	}

	data := string(input)[start:end]

	_, err = out.Write([]byte(data))
	return err
}
