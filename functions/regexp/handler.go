package regexp

import (
	"fmt"
	f "github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
)

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("regexp.test", test, f.DefParams(
			f.DefParam(f.StringValue, "pattern", false),
		)),
	)
}

// regexp.test(pattern: string): test input with regexp pattern
func test(params f.Parameters, in io.Reader, out io.Writer) error {
	pattern, ok := params.GetParameter("pattern", 0)
	if !ok {
		return f.NotEnoughParameterError
	}
	if pattern.Value.Type() != f.StringValue {
		return f.InvalidTypeOfParameterError
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
