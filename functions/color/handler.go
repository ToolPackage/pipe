package color

import (
	"bufio"
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefFunc("color.json", json, DefParams()),
	)
}

// color.json(): colorize input in json syntax
func json(_ Parameters, in io.Reader, out io.Writer) error {
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
