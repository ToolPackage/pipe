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
	)
}

// regexp.test(pattern: string): test input with regexp pattern
func test(params Parameters, in io.Reader, out io.Writer) error {
	pattern, ok := params.GetParameter("pattern", 0)
	if !ok {
		return NotEnoughParameterError
	}
	if pattern.Value.Type() != StringValue {
		return InvalidTypeOfParameterError
	}

	// read input
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	ok, err = regexp.Match(pattern.Value.Get().(string), input)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(out, strconv.FormatBool(ok))
	return err
}
