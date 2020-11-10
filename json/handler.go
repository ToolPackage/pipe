package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func JsonPretty(args []string, in io.Reader, out io.Writer) {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--json":
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, input, "", "\t"); err != nil {
				panic(err)
			}

			fmt.Println(string(prettyJSON.Bytes()))
		default:
			panic("invalid arguments")
		}
	} else {
		fmt.Println(string(input))
	}
}
