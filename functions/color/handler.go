package color

import (
	"bufio"
	f "github.com/ToolPackage/pipe/functions"
	"io"
)

func Register() []*f.FunctionDefinition {
	return f.DefFuncs(
		f.DefFunc("color.json", json, f.DefParams()),
	)
}

// color.json(): colorize input in json syntax
func json(params f.Parameters, in io.Reader, out io.Writer) error {
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
