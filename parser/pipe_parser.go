// Code generated from parser/Pipe.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Pipe

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 74, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 7, 2, 24, 10, 2, 12,
	2, 14, 2, 27, 11, 2, 3, 3, 3, 3, 5, 3, 31, 10, 3, 3, 4, 3, 4, 3, 4, 7,
	4, 36, 10, 4, 12, 4, 14, 4, 39, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 45,
	10, 5, 12, 5, 14, 5, 48, 11, 5, 3, 5, 5, 5, 51, 10, 5, 5, 5, 53, 10, 5,
	3, 5, 3, 5, 3, 6, 5, 6, 58, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 63, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 2, 2, 2, 73, 2, 20, 3, 2, 2, 2, 4, 28, 3,
	2, 2, 2, 6, 32, 3, 2, 2, 2, 8, 40, 3, 2, 2, 2, 10, 57, 3, 2, 2, 2, 12,
	64, 3, 2, 2, 2, 14, 67, 3, 2, 2, 2, 16, 69, 3, 2, 2, 2, 18, 71, 3, 2, 2,
	2, 20, 25, 5, 4, 3, 2, 21, 22, 7, 8, 2, 2, 22, 24, 5, 4, 3, 2, 23, 21,
	3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2,
	26, 3, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 30, 5, 6, 4, 2, 29, 31, 5, 8,
	5, 2, 30, 29, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 37,
	7, 14, 2, 2, 33, 34, 7, 3, 2, 2, 34, 36, 7, 14, 2, 2, 35, 33, 3, 2, 2,
	2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 7, 3,
	2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 52, 7, 4, 2, 2, 41, 46, 5, 10, 6, 2, 42,
	43, 7, 5, 2, 2, 43, 45, 5, 10, 6, 2, 44, 42, 3, 2, 2, 2, 45, 48, 3, 2,
	2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46,
	3, 2, 2, 2, 49, 51, 7, 5, 2, 2, 50, 49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2,
	51, 53, 3, 2, 2, 2, 52, 41, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 3,
	2, 2, 2, 54, 55, 7, 6, 2, 2, 55, 9, 3, 2, 2, 2, 56, 58, 5, 12, 7, 2, 57,
	56, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 62, 3, 2, 2, 2, 59, 63, 5, 14,
	8, 2, 60, 63, 5, 16, 9, 2, 61, 63, 5, 18, 10, 2, 62, 59, 3, 2, 2, 2, 62,
	60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 11, 3, 2, 2, 2, 64, 65, 7, 14,
	2, 2, 65, 66, 7, 7, 2, 2, 66, 13, 3, 2, 2, 2, 67, 68, 7, 9, 2, 2, 68, 15,
	3, 2, 2, 2, 69, 70, 7, 10, 2, 2, 70, 17, 3, 2, 2, 2, 71, 72, 7, 11, 2,
	2, 72, 19, 3, 2, 2, 2, 10, 25, 30, 37, 46, 50, 52, 57, 62,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "'('", "','", "')'", "':'", "", "", "", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "CommandSeparator", "NumberLiteral", "StringLiteral",
	"BooleanLiteral", "True", "False", "Identifier", "WS",
}

var ruleNames = []string{
	"commands", "pattern", "commandName", "commandArguments", "commandArgument",
	"commandArgumentLabel", "numberValue", "stringValue", "booleanValue",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PipeParser struct {
	*antlr.BaseParser
}

func NewPipeParser(input antlr.TokenStream) *PipeParser {
	this := new(PipeParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Pipe.g4"

	return this
}

// PipeParser tokens.
const (
	PipeParserEOF              = antlr.TokenEOF
	PipeParserT__0             = 1
	PipeParserT__1             = 2
	PipeParserT__2             = 3
	PipeParserT__3             = 4
	PipeParserT__4             = 5
	PipeParserCommandSeparator = 6
	PipeParserNumberLiteral    = 7
	PipeParserStringLiteral    = 8
	PipeParserBooleanLiteral   = 9
	PipeParserTrue             = 10
	PipeParserFalse            = 11
	PipeParserIdentifier       = 12
	PipeParserWS               = 13
)

// PipeParser rules.
const (
	PipeParserRULE_commands             = 0
	PipeParserRULE_pattern              = 1
	PipeParserRULE_commandName          = 2
	PipeParserRULE_commandArguments     = 3
	PipeParserRULE_commandArgument      = 4
	PipeParserRULE_commandArgumentLabel = 5
	PipeParserRULE_numberValue          = 6
	PipeParserRULE_stringValue          = 7
	PipeParserRULE_booleanValue         = 8
)

// ICommandsContext is an interface to support dynamic dispatch.
type ICommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandsContext differentiates from other interfaces.
	IsCommandsContext()
}

type CommandsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandsContext() *CommandsContext {
	var p = new(CommandsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_commands
	return p
}

func (*CommandsContext) IsCommandsContext() {}

func NewCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandsContext {
	var p = new(CommandsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_commands

	return p
}

func (s *CommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandsContext) AllPattern() []IPatternContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatternContext)(nil)).Elem())
	var tst = make([]IPatternContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatternContext)
		}
	}

	return tst
}

func (s *CommandsContext) Pattern(i int) IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *CommandsContext) AllCommandSeparator() []antlr.TerminalNode {
	return s.GetTokens(PipeParserCommandSeparator)
}

func (s *CommandsContext) CommandSeparator(i int) antlr.TerminalNode {
	return s.GetToken(PipeParserCommandSeparator, i)
}

func (s *CommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterCommands(s)
	}
}

func (s *CommandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitCommands(s)
	}
}

func (p *PipeParser) Commands() (localctx ICommandsContext) {
	localctx = NewCommandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PipeParserRULE_commands)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Pattern()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserCommandSeparator {
		{
			p.SetState(19)
			p.Match(PipeParserCommandSeparator)
		}
		{
			p.SetState(20)
			p.Pattern()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_pattern
	return p
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) CommandName() ICommandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandNameContext)
}

func (s *PatternContext) CommandArguments() ICommandArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandArgumentsContext)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (p *PipeParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PipeParserRULE_pattern)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.CommandName()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserT__1 {
		{
			p.SetState(27)
			p.CommandArguments()
		}

	}

	return localctx
}

// ICommandNameContext is an interface to support dynamic dispatch.
type ICommandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandNameContext differentiates from other interfaces.
	IsCommandNameContext()
}

type CommandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandNameContext() *CommandNameContext {
	var p = new(CommandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_commandName
	return p
}

func (*CommandNameContext) IsCommandNameContext() {}

func NewCommandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandNameContext {
	var p = new(CommandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_commandName

	return p
}

func (s *CommandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandNameContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(PipeParserIdentifier)
}

func (s *CommandNameContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, i)
}

func (s *CommandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterCommandName(s)
	}
}

func (s *CommandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitCommandName(s)
	}
}

func (p *PipeParser) CommandName() (localctx ICommandNameContext) {
	localctx = NewCommandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PipeParserRULE_commandName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(PipeParserIdentifier)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserT__0 {
		{
			p.SetState(31)
			p.Match(PipeParserT__0)
		}
		{
			p.SetState(32)
			p.Match(PipeParserIdentifier)
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICommandArgumentsContext is an interface to support dynamic dispatch.
type ICommandArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandArgumentsContext differentiates from other interfaces.
	IsCommandArgumentsContext()
}

type CommandArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandArgumentsContext() *CommandArgumentsContext {
	var p = new(CommandArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_commandArguments
	return p
}

func (*CommandArgumentsContext) IsCommandArgumentsContext() {}

func NewCommandArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandArgumentsContext {
	var p = new(CommandArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_commandArguments

	return p
}

func (s *CommandArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandArgumentsContext) AllCommandArgument() []ICommandArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandArgumentContext)(nil)).Elem())
	var tst = make([]ICommandArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandArgumentContext)
		}
	}

	return tst
}

func (s *CommandArgumentsContext) CommandArgument(i int) ICommandArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandArgumentContext)
}

func (s *CommandArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterCommandArguments(s)
	}
}

func (s *CommandArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitCommandArguments(s)
	}
}

func (p *PipeParser) CommandArguments() (localctx ICommandArgumentsContext) {
	localctx = NewCommandArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PipeParserRULE_commandArguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(PipeParserT__1)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PipeParserNumberLiteral)|(1<<PipeParserStringLiteral)|(1<<PipeParserBooleanLiteral)|(1<<PipeParserIdentifier))) != 0 {
		{
			p.SetState(39)
			p.CommandArgument()
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(40)
					p.Match(PipeParserT__2)
				}
				{
					p.SetState(41)
					p.CommandArgument()
				}

			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PipeParserT__2 {
			{
				p.SetState(47)
				p.Match(PipeParserT__2)
			}

		}

	}
	{
		p.SetState(52)
		p.Match(PipeParserT__3)
	}

	return localctx
}

// ICommandArgumentContext is an interface to support dynamic dispatch.
type ICommandArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandArgumentContext differentiates from other interfaces.
	IsCommandArgumentContext()
}

type CommandArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandArgumentContext() *CommandArgumentContext {
	var p = new(CommandArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_commandArgument
	return p
}

func (*CommandArgumentContext) IsCommandArgumentContext() {}

func NewCommandArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandArgumentContext {
	var p = new(CommandArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_commandArgument

	return p
}

func (s *CommandArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandArgumentContext) NumberValue() INumberValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberValueContext)
}

func (s *CommandArgumentContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *CommandArgumentContext) BooleanValue() IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *CommandArgumentContext) CommandArgumentLabel() ICommandArgumentLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandArgumentLabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandArgumentLabelContext)
}

func (s *CommandArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterCommandArgument(s)
	}
}

func (s *CommandArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitCommandArgument(s)
	}
}

func (p *PipeParser) CommandArgument() (localctx ICommandArgumentContext) {
	localctx = NewCommandArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PipeParserRULE_commandArgument)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserIdentifier {
		{
			p.SetState(54)
			p.CommandArgumentLabel()
		}

	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PipeParserNumberLiteral:
		{
			p.SetState(57)
			p.NumberValue()
		}

	case PipeParserStringLiteral:
		{
			p.SetState(58)
			p.StringValue()
		}

	case PipeParserBooleanLiteral:
		{
			p.SetState(59)
			p.BooleanValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommandArgumentLabelContext is an interface to support dynamic dispatch.
type ICommandArgumentLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandArgumentLabelContext differentiates from other interfaces.
	IsCommandArgumentLabelContext()
}

type CommandArgumentLabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandArgumentLabelContext() *CommandArgumentLabelContext {
	var p = new(CommandArgumentLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_commandArgumentLabel
	return p
}

func (*CommandArgumentLabelContext) IsCommandArgumentLabelContext() {}

func NewCommandArgumentLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandArgumentLabelContext {
	var p = new(CommandArgumentLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_commandArgumentLabel

	return p
}

func (s *CommandArgumentLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandArgumentLabelContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, 0)
}

func (s *CommandArgumentLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandArgumentLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandArgumentLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterCommandArgumentLabel(s)
	}
}

func (s *CommandArgumentLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitCommandArgumentLabel(s)
	}
}

func (p *PipeParser) CommandArgumentLabel() (localctx ICommandArgumentLabelContext) {
	localctx = NewCommandArgumentLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PipeParserRULE_commandArgumentLabel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(PipeParserIdentifier)
	}
	{
		p.SetState(63)
		p.Match(PipeParserT__4)
	}

	return localctx
}

// INumberValueContext is an interface to support dynamic dispatch.
type INumberValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberValueContext differentiates from other interfaces.
	IsNumberValueContext()
}

type NumberValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberValueContext() *NumberValueContext {
	var p = new(NumberValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_numberValue
	return p
}

func (*NumberValueContext) IsNumberValueContext() {}

func NewNumberValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberValueContext {
	var p = new(NumberValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_numberValue

	return p
}

func (s *NumberValueContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberValueContext) NumberLiteral() antlr.TerminalNode {
	return s.GetToken(PipeParserNumberLiteral, 0)
}

func (s *NumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterNumberValue(s)
	}
}

func (s *NumberValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitNumberValue(s)
	}
}

func (p *PipeParser) NumberValue() (localctx INumberValueContext) {
	localctx = NewNumberValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PipeParserRULE_numberValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(PipeParserNumberLiteral)
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(PipeParserStringLiteral, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *PipeParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PipeParserRULE_stringValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(PipeParserStringLiteral)
	}

	return localctx
}

// IBooleanValueContext is an interface to support dynamic dispatch.
type IBooleanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanValueContext differentiates from other interfaces.
	IsBooleanValueContext()
}

type BooleanValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanValueContext() *BooleanValueContext {
	var p = new(BooleanValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_booleanValue
	return p
}

func (*BooleanValueContext) IsBooleanValueContext() {}

func NewBooleanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueContext {
	var p = new(BooleanValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_booleanValue

	return p
}

func (s *BooleanValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanValueContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(PipeParserBooleanLiteral, 0)
}

func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterBooleanValue(s)
	}
}

func (s *BooleanValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitBooleanValue(s)
	}
}

func (p *PipeParser) BooleanValue() (localctx IBooleanValueContext) {
	localctx = NewBooleanValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PipeParserRULE_booleanValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(PipeParserBooleanLiteral)
	}

	return localctx
}
