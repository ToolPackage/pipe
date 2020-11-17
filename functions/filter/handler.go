package filter

import (
	"bufio"
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
	"regexp"
	"strings"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefFunc("filter.line.match", lineMatch, DefParams(
			DefParam(StringValue, "pattern", false),
		)),
		DefFunc("filter.line.contains", lineContains, DefParams(
			DefParam(StringValue, "substr", false),
		)),
	)
}

// filter.line.match(pattern: string): filter input lines those match provided regexp pattern
func lineMatch(params Parameters, in io.Reader, out io.Writer) error {
	v, _ := params.GetParameter("pattern", 0)
	pattern := v.Value.Get().(string)

	inBuf := bufio.NewReader(in)
	outBuf := bufio.NewWriter(out)
	for {
		line, err := inBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if ok, err := regexp.Match(pattern, []byte(line)); err != nil {
			return err
		} else if ok {
			_, err = outBuf.WriteString(line)
			err = outBuf.Flush()
			if err != nil {
				return err
			}
		}
	}
}

// filter.line.contains(substr: string): filter input lines those contain provided string pattern
func lineContains(params Parameters, in io.Reader, out io.Writer) error {
	v, _ := params.GetParameter("pattern", 0)
	substr := v.Value.Get().(string)

	inBuf := bufio.NewReader(in)
	outBuf := bufio.NewWriter(out)
	for {
		line, err := inBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if ok := strings.Contains(line, substr); ok {
			_, err = outBuf.WriteString(line)
			err = outBuf.Flush()
			if err != nil {
				return err
			}
		}
	}
}
