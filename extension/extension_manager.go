package extension

import (
	"github.com/ToolPackage/pipe/parser"
)

func Install(filename string) error {
	err := parser.ParseScript(filename)
	if err != nil {
		return err
	}

	return nil
}
