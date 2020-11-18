package util

import "strings"

type StringBuilder struct {
	buf           strings.Builder
	indent        string
	indentSpacing string
}

func NewStringBuilder(indentSpacing string) *StringBuilder {
	return &StringBuilder{indentSpacing: indentSpacing}
}

func (b *StringBuilder) IncIndent() {
	b.indent = strings.Repeat(b.indentSpacing, len(b.indent)/len(b.indentSpacing)+1)
}

func (b *StringBuilder) DecIndent() {
	b.indent = strings.Repeat(b.indentSpacing, len(b.indent)/len(b.indentSpacing)-1)
}

func (b *StringBuilder) WriteIndent() {
	b.buf.WriteString(b.indent)
}

func (b *StringBuilder) NewLine() {
	b.buf.WriteRune('\n')
}

func (b *StringBuilder) WriteLine(s ...string) {
	b.WriteIndent()
	for _, str := range s {
		b.buf.WriteString(str)
	}
	b.NewLine()
}

func (b *StringBuilder) WriteMultiLine(s string, indentFirstLine bool) {
	arr := strings.Split(s, "\n")
	if indentFirstLine {
		b.WriteIndent()
	}
	b.WriteString(arr[0])
	b.NewLine()

	for i := 1; i < len(arr); i++ {
		b.WriteLine(arr[i])
	}
}

func (b *StringBuilder) WriteWithIndent(s ...string) {
	b.WriteIndent()
	for _, str := range s {
		b.buf.WriteString(str)
	}
}

func (b *StringBuilder) WriteString(s string) {
	b.buf.WriteString(s)
}

func (b *StringBuilder) WriteRune(r rune) {
	b.buf.WriteRune(r)
}

func (b *StringBuilder) String() string {
	return b.buf.String()
}
