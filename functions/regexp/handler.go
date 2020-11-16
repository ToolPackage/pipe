package regexp

import (
	"fmt"
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefFunc("regexp.test", test, DefParams(
			DefParam(StringValue, "pattern", false),
		)),
		DefFunc("regexp.replace", replace, DefParams(
			DefParam(StringValue, "pattern", false),
			DefParam(StringValue, "new", false),
		)),
	)
}

// regexp.test(pattern: string): test input with regexp pattern
func test(params Parameters, in io.Reader, out io.Writer) error {
	pattern, _ := params.GetParameter("pattern", 0)

	// read input
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	ok, err := regexp.Match(pattern.Value.Get().(string), input)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(out, strconv.FormatBool(ok))
	return err
}

// regexp.replace(pattern: string, new: string)
func replace(params Parameters, in io.Reader, out io.Writer) error {
	pattern, _ := params.GetParameter("pattern", 0)
	newS, _ := params.GetParameter("new", 1)

	// read input
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(pattern.Value.Get().(string))
	s := re.ReplaceAllString(string(input), newS.Value.Get().(string))

	_, err = out.Write([]byte(s))
	return err
}
