package main

import (
	"bytes"
	"strings"
)

func execute(commands []Command) {
	buf := make([]byte, 0)
	in := bytes.NewReader(buf)
	out := new(strings.Builder)

	for _, cmd := range commands {
		cmd.handler(cmd.args, in, out)
		buf = []byte(out.String())
		in = bytes.NewReader(buf)
		out = new(strings.Builder)
	}
}
