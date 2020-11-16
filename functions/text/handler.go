package text

import (
	"errors"
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
	"io/ioutil"
	"strings"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefFunc("text.cut", cut, DefParams(
			DefParam(IntegerValue, "start", false),
			DefParam(IntegerValue, "end", true),
		)),
		DefFunc("text.replace", replace, DefParams(
			DefParam(StringValue, "old", false),
			DefParam(StringValue, "new", false),
		)),
		DefFunc("text.repeat", repeat, DefParams(
			DefParam(IntegerValue, "n", false),
		)),
	)
}

// text.cut(start: int, end?: int): extract substring
func cut(params Parameters, in io.Reader, out io.Writer) error {
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

// text.replace(old: string, new: string): replace substring
func replace(params Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	oldS, _ := params.GetParameter("old", 0)
	newS, _ := params.GetParameter("new", 1)
	newString := strings.Replace(string(input),
		oldS.Value.Get().(string), newS.Value.Get().(string), -1)

	_, err = out.Write([]byte(newString))
	return err
}

// text.repeat(n: int): repeat input n times
func repeat(params Parameters, in io.Reader, out io.Writer) error {
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	n, _ := params.GetParameter("n", 0)

	newString := strings.Repeat(string(input), n.Value.Get().(int))

	_, err = out.Write([]byte(newString))
	return err
}
