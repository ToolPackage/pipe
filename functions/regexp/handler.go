package regexp

import (
	"fmt"
	"github.com/ToolPackage/pipe/functions"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
)

// regexp.test(pattern: string)
func Test(params functions.FunctionParameters, in io.Reader, out io.Writer) error {
	pattern, ok := params.GetParameter("pattern", 0)
	if !ok {
		return functions.NotEnoughParameterError
	}
	if pattern.ValueType != functions.StringValue {
		return functions.InvalidTypeOfParameterError
	}

	// read input
	input, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	ok, err = regexp.Match(pattern.GetValue().(string), input)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(out, strconv.FormatBool(ok))
	return err
}
