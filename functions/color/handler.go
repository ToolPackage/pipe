package color

import (
	"bufio"
	"github.com/ToolPackage/pipe/functions"
	"io"
)

func Register() []*functions.FunctionDefinition {
	return functions.DefFuncs(
		functions.DefFunc("color.json", json, functions.DefParams()),
	)
}

// color.json(): colorize input in json syntax
func json(params functions.Parameters, in io.Reader, out io.Writer) error {
	config := map[string]string{
		KeyBrace:   "brightBlack",
		KeyQuote:   "green",
		KeyName:    "brightCyan",
		KeyColon:   "brightYellow",
		KeyBracket: "brightBlack",
		KeyNumber:  "brightBlue",
		KeyComma:   "brightBlack",
		KeyString:  "brightGreen",
		KeyLiteral: "brightRed",
	}

	printer := NewJsonPrinter(config, out)
	input := bufio.NewReader(in)
	for {
		c, _, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		printer.print(c)
	}

	return nil
}
