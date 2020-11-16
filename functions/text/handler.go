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
		DefFunc("text.join", join, DefParams(
			DefParam(DictValue, "s", false),
			DefParam(StringValue, "sep", true),
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

// text.join(elems: []string, sep: string)
func join(params Parameters, _ io.Reader, out io.Writer) error {
	elems, _ := params.GetParameter("elems", 0)
	dict := elems.Value.(*DictParameterValue)
	items := make([]string, 0)
	for i := 0; i < dict.Size(); i++ {
		v, _ := dict.GetValueByIndex(i)
		if v.Type() == ReferenceValue {
			tmp, err := v.(*ReferenceParameterValue).GetAs(StringValue)
			if err != nil {
				return InvalidTypeOfParameterError
			}
			items = append(items, tmp.(string))
		} else if v.Type() != StringValue {
			return InvalidTypeOfParameterError
		} else {
			items = append(items, v.Get().(string))
		}
	}

	separator := ""
	sep, ok := params.GetParameter("sep", 1)
	if ok {
		separator = sep.Value.Get().(string)
	}
	newS := strings.Join(items, separator)

	_, err := out.Write([]byte(newS))
	return err
}
