package color

import (
	"os"
	"testing"
)

func TestNewJsonPrinter(t *testing.T) {
	config := make(map[string]string)
	config[KeyBrace] = "brightBlack"
	config[KeyQuote] = "green"
	config[KeyName] = "brightCyan"
	config[KeyColon] = "brightYellow"
	config[KeyBracket] = "brightBlack"
	config[KeyNumber] = "brightBlue"
	config[KeyComma] = "brightBlack"
	config[KeyString] = "brightGreen"
	config[KeyLiteral] = "brightRed"
	printer := NewJsonPrinter(config, os.Stdout)

	data := `[{
    "a": {
        "name": "Windows",
        "b": [1, "asd", true, 1.0, 1.1e-10, false, null],
        "c": true,
        "d": null,
        "x": 123
    }
}, {}]
`
	for _, char := range data {
		printer.print(char)
	}
}
