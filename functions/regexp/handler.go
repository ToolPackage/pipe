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
		DefBuiltinFunc("regexp.test", test, DefParams(
			DefParam(StringValue, "pattern", false),
		)),
		DefBuiltinFunc("regexp.replace", replace, DefParams(
			DefParam(StringValue, "pattern", false),
			DefParam(StringValue, "repl", false),
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

// regexp.replace(pattern: string, repl: string)
func replace(params Parameters, in io.Reader, out io.Writer) error {
	pattern, _ := params.GetParameter("pattern", 0)
	repl, _ := params.GetParameter("repl", 1)

	// read input
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	// TODO: regexp cannot match newline
	re := regexp.MustCompile(pattern.Value.Get().(string))
	s := re.ReplaceAllString(string(input), repl.Value.Get().(string))

	_, err = out.Write([]byte(s))
	return err
}
