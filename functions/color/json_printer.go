package color

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
	"unicode"
)

// Ref: https://tools.ietf.org/html/rfc7159

type PrettifyPrinter interface {
	print(b int32)
}

// tokens
const (
	DOT      = '.'
	Add      = '+'
	Sub      = '-'
	Quote    = '"'
	Colon    = ':'
	Comma    = ','
	LBrace   = '{'
	RBrace   = '}'
	LBracket = '['
	RBracket = ']'

	// whitespace
	Space   = 0x20
	Tab     = 0x09
	Newline = 0x0A
	Return  = 0x0D
)

// scopes
const (
	ScopeNone = iota
	ScopeObject
	ScopeMemberName
	ScopeArray
	// value
	ScopeStringValue
	ScopeNumberValue
	ScopeLiteralValue
)

// highlight keys
const (
	KeyBrace   = "brace"
	KeyQuote   = "quote"
	KeyName    = "name"
	KeyColon   = "colon"
	KeyBracket = "bracket"
	KeyComma   = "comma"
	KeyNumber  = "number"
	KeyString  = "string"
	KeyLiteral = "literal"
)

type JsonPrinter struct {
	scopes       *Stack
	lastPopScope int
	lastToken    rune
	config       map[string]string
	out          io.Writer
}

func NewJsonPrinter(config map[string]string, out io.Writer) *JsonPrinter {
	j := &JsonPrinter{
		scopes:       NewStack(),
		lastPopScope: ScopeNone,
		lastToken:    -1,
		config:       config,
		out:          out,
	}
	return j
}

func (j *JsonPrinter) print(ch rune) {
	//fmt.Println(j.scopes.String())
	defer j.markToken(ch)

	if ch != Quote && !j.scopes.isEmpty() {
		scope := j.currentScope()
		if scope == ScopeMemberName {
			j.printWithColor(j.config[KeyName], ch)
			return
		}
		if scope == ScopeStringValue {
			j.printWithColor(j.config[KeyString], ch)
			return
		}
	}

	switch ch {
	case LBrace:
		j.enterScope(ScopeObject)
		j.printWithColor(j.config[KeyBrace], ch)
		return
	case RBrace:
		if j.currentScope() == ScopeObject {
			j.exitScope()
			j.printWithColor(j.config[KeyBrace], ch)
			return
		}
		if j.scopes.AtTop(1) == ScopeObject {
			j.scopes.Pop()
			j.exitScope()
			j.printWithColor(j.config[KeyBrace], ch)
			return
		}
	case LBracket:
		j.enterScope(ScopeArray)
		j.printWithColor(j.config[KeyBracket], ch)
		return
	case RBracket:
		scope := j.currentScope()
		if scope == ScopeArray {
			j.exitScope()
			j.printWithColor(j.config[KeyBracket], ch)
			return
		}
		if scope > ScopeStringValue {
			if j.scopes.AtTop(1) == ScopeArray {
				j.scopes.Pop()
				j.exitScope()
				j.printWithColor(j.config[KeyBracket], ch)
				return
			}
		}
	case Quote:
		j.printQuote()
		return
	case Colon:
		if j.lastPopScope == ScopeMemberName {
			j.printWithColor(j.config[KeyColon], ch)
			return
		}
	case Comma:
		if j.currentScope() > ScopeStringValue { // number & boolean
			j.exitScope()
			j.printWithColor(j.config[KeyComma], Comma)
			return
		}
		if j.lastToken == RBracket || j.lastToken == RBrace || j.lastPopScope == ScopeStringValue {
			j.printWithColor(j.config[KeyComma], Comma)
			return
		}
	case Add, Sub:
		if j.currentScope() == ScopeNumberValue {
			j.printWithColor(j.config[KeyNumber], ch)
			return
		}
	case Space, Tab, Newline, Return:
		break
	default:
		j.printOther(ch)
		return
	}
	j.skip(ch)
}

func (j *JsonPrinter) printQuote() {
	scope := j.currentScope()
	switch {
	case scope == ScopeArray || j.lastToken == Colon:
		j.enterScope(ScopeStringValue)
		j.printWithColor(j.config[KeyQuote], Quote)
	case scope == ScopeStringValue:
		j.exitScope()
		j.printWithColor(j.config[KeyQuote], Quote)
	case scope == ScopeObject:
		j.enterScope(ScopeMemberName)
		j.printWithColor(j.config[KeyQuote], Quote)
	case scope == ScopeMemberName:
		j.exitScope()
		j.printWithColor(j.config[KeyQuote], Quote)
	default:
		j.skip(Quote)
	}
}

func (j *JsonPrinter) printOther(ch rune) {
	scope := j.currentScope()

	switch {
	case ch == DOT:
		if scope == ScopeNumberValue {
			j.printWithColor(j.config[KeyNumber], ch)
			return
		}
	case unicode.IsLetter(ch):
		if ch == 'e' && scope == ScopeNumberValue {
			j.printWithColor(j.config[KeyNumber], ch)
			return
		}
		if j.lastToken == Colon || scope == ScopeArray {
			j.enterScope(ScopeLiteralValue)
			j.printWithColor(j.config[KeyLiteral], ch)
			return
		}
		if scope == ScopeLiteralValue {
			j.printWithColor(j.config[KeyLiteral], ch)
			return
		}
	case unicode.IsDigit(ch):
		if j.lastToken == Colon || scope == ScopeArray {
			j.enterScope(ScopeNumberValue)
			j.printWithColor(j.config[KeyNumber], ch)
			return
		}
		if scope == ScopeNumberValue {
			j.printWithColor(j.config[KeyNumber], ch)
			return
		}
	}

	j.skip(ch)
}

func (j *JsonPrinter) skip(ch rune) {
	//fmt.Print(j.scopes.Peek())
	fmt.Print(string(ch))
}

func (j *JsonPrinter) enterScope(scope int) {
	j.scopes.Push(scope)
}

func (j *JsonPrinter) currentScope() int {
	return j.scopes.Peek().(int)
}

func (j *JsonPrinter) exitScope() {
	j.lastPopScope = j.scopes.Pop().(int)
}

func (j *JsonPrinter) markToken(ch rune) {
	if ch == Space || ch == Tab || ch == Newline || ch == Return {
		return
	}
	j.lastToken = ch
}

func (j *JsonPrinter) printWithColor(color string, ch rune) {
	v := colorFuncMap[color](string(ch)).String()
	_, err := j.out.Write([]byte(v))
	if err != nil {
		panic(err)
	}
}

type PrintFunc func(arg interface{}) aurora.Value

var colorFuncMap = map[string]PrintFunc{
	"black":        aurora.Black,
	"red":          aurora.Red,
	"green":        aurora.Green,
	"yellow":       aurora.Yellow,
	"blue":         aurora.Blue,
	"magenta":      aurora.Magenta,
	"cyan":         aurora.Cyan,
	"white":        aurora.White,
	"brightBlack":  aurora.BrightBlack,
	"brightYellow": aurora.BrightYellow,
	"brightGreen":  aurora.BrightGreen,
	"brightBlue":   aurora.BrightBlue,
	"brightRed":    aurora.BrightRed,
	"brightCyan":   aurora.BrightCyan,
}
