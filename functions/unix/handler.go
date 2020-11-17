package unix

import (
	. "github.com/ToolPackage/pipe/parser/definition"
	"io"
)

func Register() []*FunctionDefinition {
	return DefFuncs(
		DefFunc("unix.cat", cat, DefParams()),
	)
}

// unix.cat(args: []string): reads files sequentially, writing them to standard output
//   available arguments:
//     -b or --number-nonblank: number non-blank output lines
//     -e: implies -v but also display end-of-line characters as $ (GNU only: -E the same, but without implying -v)
//     -n or --number: number all output lines
//     -s or --squeeze-blank: squeeze multiple adjacent blank lines
//     -t: implies -v, but also display tabs as ^I (GNU: -T the same, but without implying -v)
//     -u: use unbuffered I/O for stdout. POSIX does not specify the behavior without this option.
//     -v or --show-nonprinting:
func cat(_ Parameters, in io.Reader, out io.Writer) error {
	// TODO:
	return nil
}
