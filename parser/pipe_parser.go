// Code generated from parser/Pipe.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 212,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 3, 2, 7, 2, 60, 10, 2, 12, 2, 14, 2, 63, 11, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 74, 10, 4, 12, 4, 14, 4,
	77, 11, 4, 5, 4, 79, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 85, 10, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 5, 9, 100, 10, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 7, 11, 111, 10, 11, 12, 11, 14, 11, 114, 11, 11, 3, 12, 3, 12, 3,
	12, 7, 12, 119, 10, 12, 12, 12, 14, 12, 122, 11, 12, 3, 13, 3, 13, 5, 13,
	126, 10, 13, 3, 14, 3, 14, 5, 14, 130, 10, 14, 3, 15, 3, 15, 3, 15, 7,
	15, 135, 10, 15, 12, 15, 14, 15, 138, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	7, 16, 144, 10, 16, 12, 16, 14, 16, 147, 11, 16, 3, 16, 5, 16, 150, 10,
	16, 5, 16, 152, 10, 16, 3, 16, 3, 16, 3, 17, 5, 17, 157, 10, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 164, 10, 17, 3, 18, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 174, 10, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 7, 21, 181, 10, 21, 12, 21, 14, 21, 184, 11, 21, 3, 22, 3,
	22, 3, 22, 5, 22, 189, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 5, 24, 198, 10, 24, 3, 25, 3, 25, 5, 25, 202, 10, 25, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 2, 2, 30, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 2, 2, 2, 207, 2, 61, 3, 2, 2, 2, 4, 64, 3,
	2, 2, 2, 6, 69, 3, 2, 2, 2, 8, 82, 3, 2, 2, 2, 10, 89, 3, 2, 2, 2, 12,
	91, 3, 2, 2, 2, 14, 93, 3, 2, 2, 2, 16, 95, 3, 2, 2, 2, 18, 104, 3, 2,
	2, 2, 20, 107, 3, 2, 2, 2, 22, 115, 3, 2, 2, 2, 24, 125, 3, 2, 2, 2, 26,
	127, 3, 2, 2, 2, 28, 131, 3, 2, 2, 2, 30, 139, 3, 2, 2, 2, 32, 156, 3,
	2, 2, 2, 34, 165, 3, 2, 2, 2, 36, 168, 3, 2, 2, 2, 38, 171, 3, 2, 2, 2,
	40, 177, 3, 2, 2, 2, 42, 188, 3, 2, 2, 2, 44, 192, 3, 2, 2, 2, 46, 197,
	3, 2, 2, 2, 48, 201, 3, 2, 2, 2, 50, 203, 3, 2, 2, 2, 52, 205, 3, 2, 2,
	2, 54, 207, 3, 2, 2, 2, 56, 209, 3, 2, 2, 2, 58, 60, 5, 4, 3, 2, 59, 58,
	3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2,
	62, 3, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 7, 3, 2, 2, 65, 66, 5, 28,
	15, 2, 66, 67, 5, 6, 4, 2, 67, 68, 5, 16, 9, 2, 68, 5, 3, 2, 2, 2, 69,
	78, 7, 4, 2, 2, 70, 75, 5, 8, 5, 2, 71, 72, 7, 5, 2, 2, 72, 74, 5, 8, 5,
	2, 73, 71, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76,
	3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 70, 3, 2, 2, 2,
	78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 81, 7, 6, 2, 2, 81, 7, 3, 2,
	2, 2, 82, 84, 5, 10, 6, 2, 83, 85, 5, 12, 7, 2, 84, 83, 3, 2, 2, 2, 84,
	85, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 7, 7, 2, 2, 87, 88, 5, 14,
	8, 2, 88, 9, 3, 2, 2, 2, 89, 90, 7, 25, 2, 2, 90, 11, 3, 2, 2, 2, 91, 92,
	7, 8, 2, 2, 92, 13, 3, 2, 2, 2, 93, 94, 7, 18, 2, 2, 94, 15, 3, 2, 2, 2,
	95, 99, 7, 9, 2, 2, 96, 97, 5, 20, 11, 2, 97, 98, 7, 10, 2, 2, 98, 100,
	3, 2, 2, 2, 99, 96, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 102, 5, 18, 10, 2, 102, 103, 7, 11, 2, 2, 103, 17, 3, 2, 2, 2,
	104, 105, 7, 12, 2, 2, 105, 106, 5, 22, 12, 2, 106, 19, 3, 2, 2, 2, 107,
	112, 5, 22, 12, 2, 108, 109, 7, 10, 2, 2, 109, 111, 5, 22, 12, 2, 110,
	108, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113,
	3, 2, 2, 2, 113, 21, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 120, 5, 24,
	13, 2, 116, 117, 7, 17, 2, 2, 117, 119, 5, 24, 13, 2, 118, 116, 3, 2, 2,
	2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121,
	23, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 126, 5, 26, 14, 2, 124, 126,
	5, 36, 19, 2, 125, 123, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 25, 3, 2,
	2, 2, 127, 129, 5, 28, 15, 2, 128, 130, 5, 30, 16, 2, 129, 128, 3, 2, 2,
	2, 129, 130, 3, 2, 2, 2, 130, 27, 3, 2, 2, 2, 131, 136, 7, 25, 2, 2, 132,
	133, 7, 13, 2, 2, 133, 135, 7, 25, 2, 2, 134, 132, 3, 2, 2, 2, 135, 138,
	3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 29, 3, 2,
	2, 2, 138, 136, 3, 2, 2, 2, 139, 151, 7, 4, 2, 2, 140, 145, 5, 32, 17,
	2, 141, 142, 7, 5, 2, 2, 142, 144, 5, 32, 17, 2, 143, 141, 3, 2, 2, 2,
	144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146,
	149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 150, 7, 5, 2, 2, 149, 148,
	3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151, 140, 3, 2,
	2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 7, 6, 2, 2,
	154, 31, 3, 2, 2, 2, 155, 157, 5, 34, 18, 2, 156, 155, 3, 2, 2, 2, 156,
	157, 3, 2, 2, 2, 157, 163, 3, 2, 2, 2, 158, 164, 5, 36, 19, 2, 159, 164,
	5, 38, 20, 2, 160, 164, 5, 48, 25, 2, 161, 164, 5, 54, 28, 2, 162, 164,
	5, 56, 29, 2, 163, 158, 3, 2, 2, 2, 163, 159, 3, 2, 2, 2, 163, 160, 3,
	2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 162, 3, 2, 2, 2, 164, 33, 3, 2, 2,
	2, 165, 166, 7, 25, 2, 2, 166, 167, 7, 7, 2, 2, 167, 35, 3, 2, 2, 2, 168,
	169, 7, 14, 2, 2, 169, 170, 7, 25, 2, 2, 170, 37, 3, 2, 2, 2, 171, 173,
	7, 15, 2, 2, 172, 174, 5, 40, 21, 2, 173, 172, 3, 2, 2, 2, 173, 174, 3,
	2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 7, 16, 2, 2, 176, 39, 3, 2, 2,
	2, 177, 182, 5, 42, 22, 2, 178, 179, 7, 5, 2, 2, 179, 181, 5, 42, 22, 2,
	180, 178, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182,
	183, 3, 2, 2, 2, 183, 41, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 186, 5,
	44, 23, 2, 186, 187, 7, 7, 2, 2, 187, 189, 3, 2, 2, 2, 188, 185, 3, 2,
	2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 5, 46, 24,
	2, 191, 43, 3, 2, 2, 2, 192, 193, 7, 25, 2, 2, 193, 45, 3, 2, 2, 2, 194,
	198, 5, 48, 25, 2, 195, 198, 5, 54, 28, 2, 196, 198, 5, 56, 29, 2, 197,
	194, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 196, 3, 2, 2, 2, 198, 47, 3,
	2, 2, 2, 199, 202, 5, 50, 26, 2, 200, 202, 5, 52, 27, 2, 201, 199, 3, 2,
	2, 2, 201, 200, 3, 2, 2, 2, 202, 49, 3, 2, 2, 2, 203, 204, 7, 19, 2, 2,
	204, 51, 3, 2, 2, 2, 205, 206, 7, 20, 2, 2, 206, 53, 3, 2, 2, 2, 207, 208,
	7, 21, 2, 2, 208, 55, 3, 2, 2, 2, 209, 210, 7, 22, 2, 2, 210, 57, 3, 2,
	2, 2, 22, 61, 75, 78, 84, 99, 112, 120, 125, 129, 136, 145, 149, 151, 156,
	163, 173, 182, 188, 197, 201,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'def'", "'('", "','", "')'", "':'", "'?'", "'{'", "';'", "'}'", "'return'",
	"'.'", "'$'", "'['", "']'", "", "", "", "", "", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Connect",
	"ValueType", "IntegerLiteral", "DecimalLiteral", "StringLiteral", "BooleanLiteral",
	"True", "False", "Identifier", "WS",
}

var ruleNames = []string{
	"script", "funcDef", "funcParamsDef", "funcParamDef", "funcParamName",
	"optionalParamFlag", "funcParamType", "funcBody", "returnStatement", "multiPipe",
	"pipe", "pipeNode", "function", "functionName", "functionParameters", "functionParameter",
	"functionParameterLabel", "variable", "dictValue", "dictEntries", "dictEntry",
	"dictEntryLabel", "dictEntryValue", "numberValue", "integerValue", "decimalValue",
	"stringValue", "booleanValue",
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
	PipeParserEOF            = antlr.TokenEOF
	PipeParserT__0           = 1
	PipeParserT__1           = 2
	PipeParserT__2           = 3
	PipeParserT__3           = 4
	PipeParserT__4           = 5
	PipeParserT__5           = 6
	PipeParserT__6           = 7
	PipeParserT__7           = 8
	PipeParserT__8           = 9
	PipeParserT__9           = 10
	PipeParserT__10          = 11
	PipeParserT__11          = 12
	PipeParserT__12          = 13
	PipeParserT__13          = 14
	PipeParserConnect        = 15
	PipeParserValueType      = 16
	PipeParserIntegerLiteral = 17
	PipeParserDecimalLiteral = 18
	PipeParserStringLiteral  = 19
	PipeParserBooleanLiteral = 20
	PipeParserTrue           = 21
	PipeParserFalse          = 22
	PipeParserIdentifier     = 23
	PipeParserWS             = 24
)

// PipeParser rules.
const (
	PipeParserRULE_script                 = 0
	PipeParserRULE_funcDef                = 1
	PipeParserRULE_funcParamsDef          = 2
	PipeParserRULE_funcParamDef           = 3
	PipeParserRULE_funcParamName          = 4
	PipeParserRULE_optionalParamFlag      = 5
	PipeParserRULE_funcParamType          = 6
	PipeParserRULE_funcBody               = 7
	PipeParserRULE_returnStatement        = 8
	PipeParserRULE_multiPipe              = 9
	PipeParserRULE_pipe                   = 10
	PipeParserRULE_pipeNode               = 11
	PipeParserRULE_function               = 12
	PipeParserRULE_functionName           = 13
	PipeParserRULE_functionParameters     = 14
	PipeParserRULE_functionParameter      = 15
	PipeParserRULE_functionParameterLabel = 16
	PipeParserRULE_variable               = 17
	PipeParserRULE_dictValue              = 18
	PipeParserRULE_dictEntries            = 19
	PipeParserRULE_dictEntry              = 20
	PipeParserRULE_dictEntryLabel         = 21
	PipeParserRULE_dictEntryValue         = 22
	PipeParserRULE_numberValue            = 23
	PipeParserRULE_integerValue           = 24
	PipeParserRULE_decimalValue           = 25
	PipeParserRULE_stringValue            = 26
	PipeParserRULE_booleanValue           = 27
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) AllFuncDef() []IFuncDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncDefContext)(nil)).Elem())
	var tst = make([]IFuncDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncDefContext)
		}
	}

	return tst
}

func (s *ScriptContext) FuncDef(i int) IFuncDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *PipeParser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PipeParserRULE_script)
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserT__0 {
		{
			p.SetState(56)
			p.FuncDef()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_funcDef
	return p
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) FunctionName() IFunctionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *FuncDefContext) FuncParamsDef() IFuncParamsDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamsDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamsDefContext)
}

func (s *FuncDefContext) FuncBody() IFuncBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncBodyContext)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFuncDef(s)
	}
}

func (p *PipeParser) FuncDef() (localctx IFuncDefContext) {
	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PipeParserRULE_funcDef)

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
		p.Match(PipeParserT__0)
	}
	{
		p.SetState(63)
		p.FunctionName()
	}
	{
		p.SetState(64)
		p.FuncParamsDef()
	}
	{
		p.SetState(65)
		p.FuncBody()
	}

	return localctx
}

// IFuncParamsDefContext is an interface to support dynamic dispatch.
type IFuncParamsDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamsDefContext differentiates from other interfaces.
	IsFuncParamsDefContext()
}

type FuncParamsDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamsDefContext() *FuncParamsDefContext {
	var p = new(FuncParamsDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_funcParamsDef
	return p
}

func (*FuncParamsDefContext) IsFuncParamsDefContext() {}

func NewFuncParamsDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamsDefContext {
	var p = new(FuncParamsDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_funcParamsDef

	return p
}

func (s *FuncParamsDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamsDefContext) AllFuncParamDef() []IFuncParamDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncParamDefContext)(nil)).Elem())
	var tst = make([]IFuncParamDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncParamDefContext)
		}
	}

	return tst
}

func (s *FuncParamsDefContext) FuncParamDef(i int) IFuncParamDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncParamDefContext)
}

func (s *FuncParamsDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamsDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamsDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFuncParamsDef(s)
	}
}

func (s *FuncParamsDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFuncParamsDef(s)
	}
}

func (p *PipeParser) FuncParamsDef() (localctx IFuncParamsDefContext) {
	localctx = NewFuncParamsDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PipeParserRULE_funcParamsDef)
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
		p.SetState(67)
		p.Match(PipeParserT__1)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserIdentifier {
		{
			p.SetState(68)
			p.FuncParamDef()
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PipeParserT__2 {
			{
				p.SetState(69)
				p.Match(PipeParserT__2)
			}
			{
				p.SetState(70)
				p.FuncParamDef()
			}

			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(78)
		p.Match(PipeParserT__3)
	}

	return localctx
}

// IFuncParamDefContext is an interface to support dynamic dispatch.
type IFuncParamDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamDefContext differentiates from other interfaces.
	IsFuncParamDefContext()
}

type FuncParamDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamDefContext() *FuncParamDefContext {
	var p = new(FuncParamDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_funcParamDef
	return p
}

func (*FuncParamDefContext) IsFuncParamDefContext() {}

func NewFuncParamDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamDefContext {
	var p = new(FuncParamDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_funcParamDef

	return p
}

func (s *FuncParamDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamDefContext) FuncParamName() IFuncParamNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamNameContext)
}

func (s *FuncParamDefContext) FuncParamType() IFuncParamTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncParamTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncParamTypeContext)
}

func (s *FuncParamDefContext) OptionalParamFlag() IOptionalParamFlagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalParamFlagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalParamFlagContext)
}

func (s *FuncParamDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFuncParamDef(s)
	}
}

func (s *FuncParamDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFuncParamDef(s)
	}
}

func (p *PipeParser) FuncParamDef() (localctx IFuncParamDefContext) {
	localctx = NewFuncParamDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PipeParserRULE_funcParamDef)
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
		p.SetState(80)
		p.FuncParamName()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserT__5 {
		{
			p.SetState(81)
			p.OptionalParamFlag()
		}

	}
	{
		p.SetState(84)
		p.Match(PipeParserT__4)
	}
	{
		p.SetState(85)
		p.FuncParamType()
	}

	return localctx
}

// IFuncParamNameContext is an interface to support dynamic dispatch.
type IFuncParamNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamNameContext differentiates from other interfaces.
	IsFuncParamNameContext()
}

type FuncParamNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamNameContext() *FuncParamNameContext {
	var p = new(FuncParamNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_funcParamName
	return p
}

func (*FuncParamNameContext) IsFuncParamNameContext() {}

func NewFuncParamNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamNameContext {
	var p = new(FuncParamNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_funcParamName

	return p
}

func (s *FuncParamNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, 0)
}

func (s *FuncParamNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFuncParamName(s)
	}
}

func (s *FuncParamNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFuncParamName(s)
	}
}

func (p *PipeParser) FuncParamName() (localctx IFuncParamNameContext) {
	localctx = NewFuncParamNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PipeParserRULE_funcParamName)

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
		p.SetState(87)
		p.Match(PipeParserIdentifier)
	}

	return localctx
}

// IOptionalParamFlagContext is an interface to support dynamic dispatch.
type IOptionalParamFlagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionalParamFlagContext differentiates from other interfaces.
	IsOptionalParamFlagContext()
}

type OptionalParamFlagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalParamFlagContext() *OptionalParamFlagContext {
	var p = new(OptionalParamFlagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_optionalParamFlag
	return p
}

func (*OptionalParamFlagContext) IsOptionalParamFlagContext() {}

func NewOptionalParamFlagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalParamFlagContext {
	var p = new(OptionalParamFlagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_optionalParamFlag

	return p
}

func (s *OptionalParamFlagContext) GetParser() antlr.Parser { return s.parser }
func (s *OptionalParamFlagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalParamFlagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalParamFlagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterOptionalParamFlag(s)
	}
}

func (s *OptionalParamFlagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitOptionalParamFlag(s)
	}
}

func (p *PipeParser) OptionalParamFlag() (localctx IOptionalParamFlagContext) {
	localctx = NewOptionalParamFlagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PipeParserRULE_optionalParamFlag)

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
		p.SetState(89)
		p.Match(PipeParserT__5)
	}

	return localctx
}

// IFuncParamTypeContext is an interface to support dynamic dispatch.
type IFuncParamTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncParamTypeContext differentiates from other interfaces.
	IsFuncParamTypeContext()
}

type FuncParamTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncParamTypeContext() *FuncParamTypeContext {
	var p = new(FuncParamTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_funcParamType
	return p
}

func (*FuncParamTypeContext) IsFuncParamTypeContext() {}

func NewFuncParamTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncParamTypeContext {
	var p = new(FuncParamTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_funcParamType

	return p
}

func (s *FuncParamTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncParamTypeContext) ValueType() antlr.TerminalNode {
	return s.GetToken(PipeParserValueType, 0)
}

func (s *FuncParamTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncParamTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFuncParamType(s)
	}
}

func (s *FuncParamTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFuncParamType(s)
	}
}

func (p *PipeParser) FuncParamType() (localctx IFuncParamTypeContext) {
	localctx = NewFuncParamTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PipeParserRULE_funcParamType)

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
		p.SetState(91)
		p.Match(PipeParserValueType)
	}

	return localctx
}

// IFuncBodyContext is an interface to support dynamic dispatch.
type IFuncBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncBodyContext differentiates from other interfaces.
	IsFuncBodyContext()
}

type FuncBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncBodyContext() *FuncBodyContext {
	var p = new(FuncBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_funcBody
	return p
}

func (*FuncBodyContext) IsFuncBodyContext() {}

func NewFuncBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncBodyContext {
	var p = new(FuncBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_funcBody

	return p
}

func (s *FuncBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncBodyContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *FuncBodyContext) MultiPipe() IMultiPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiPipeContext)
}

func (s *FuncBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFuncBody(s)
	}
}

func (s *FuncBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFuncBody(s)
	}
}

func (p *PipeParser) FuncBody() (localctx IFuncBodyContext) {
	localctx = NewFuncBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PipeParserRULE_funcBody)
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
		p.SetState(93)
		p.Match(PipeParserT__6)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserT__11 || _la == PipeParserIdentifier {
		{
			p.SetState(94)
			p.MultiPipe()
		}
		{
			p.SetState(95)
			p.Match(PipeParserT__7)
		}

	}
	{
		p.SetState(99)
		p.ReturnStatement()
	}
	{
		p.SetState(100)
		p.Match(PipeParserT__8)
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Pipe() IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *PipeParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PipeParserRULE_returnStatement)

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
		p.SetState(102)
		p.Match(PipeParserT__9)
	}
	{
		p.SetState(103)
		p.Pipe()
	}

	return localctx
}

// IMultiPipeContext is an interface to support dynamic dispatch.
type IMultiPipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPipeContext differentiates from other interfaces.
	IsMultiPipeContext()
}

type MultiPipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiPipeContext() *MultiPipeContext {
	var p = new(MultiPipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_multiPipe
	return p
}

func (*MultiPipeContext) IsMultiPipeContext() {}

func NewMultiPipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPipeContext {
	var p = new(MultiPipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_multiPipe

	return p
}

func (s *MultiPipeContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPipeContext) AllPipe() []IPipeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipeContext)(nil)).Elem())
	var tst = make([]IPipeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipeContext)
		}
	}

	return tst
}

func (s *MultiPipeContext) Pipe(i int) IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *MultiPipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterMultiPipe(s)
	}
}

func (s *MultiPipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitMultiPipe(s)
	}
}

func (p *PipeParser) MultiPipe() (localctx IMultiPipeContext) {
	localctx = NewMultiPipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PipeParserRULE_multiPipe)

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
		p.SetState(105)
		p.Pipe()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(106)
				p.Match(PipeParserT__7)
			}
			{
				p.SetState(107)
				p.Pipe()
			}

		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IPipeContext is an interface to support dynamic dispatch.
type IPipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeContext differentiates from other interfaces.
	IsPipeContext()
}

type PipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeContext() *PipeContext {
	var p = new(PipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_pipe
	return p
}

func (*PipeContext) IsPipeContext() {}

func NewPipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeContext {
	var p = new(PipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_pipe

	return p
}

func (s *PipeContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeContext) AllPipeNode() []IPipeNodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipeNodeContext)(nil)).Elem())
	var tst = make([]IPipeNodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipeNodeContext)
		}
	}

	return tst
}

func (s *PipeContext) PipeNode(i int) IPipeNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeNodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipeNodeContext)
}

func (s *PipeContext) AllConnect() []antlr.TerminalNode {
	return s.GetTokens(PipeParserConnect)
}

func (s *PipeContext) Connect(i int) antlr.TerminalNode {
	return s.GetToken(PipeParserConnect, i)
}

func (s *PipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterPipe(s)
	}
}

func (s *PipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitPipe(s)
	}
}

func (p *PipeParser) Pipe() (localctx IPipeContext) {
	localctx = NewPipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PipeParserRULE_pipe)
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
		p.SetState(113)
		p.PipeNode()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserConnect {
		{
			p.SetState(114)
			p.Match(PipeParserConnect)
		}
		{
			p.SetState(115)
			p.PipeNode()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPipeNodeContext is an interface to support dynamic dispatch.
type IPipeNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeNodeContext differentiates from other interfaces.
	IsPipeNodeContext()
}

type PipeNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeNodeContext() *PipeNodeContext {
	var p = new(PipeNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_pipeNode
	return p
}

func (*PipeNodeContext) IsPipeNodeContext() {}

func NewPipeNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeNodeContext {
	var p = new(PipeNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_pipeNode

	return p
}

func (s *PipeNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeNodeContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *PipeNodeContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *PipeNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterPipeNode(s)
	}
}

func (s *PipeNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitPipeNode(s)
	}
}

func (p *PipeParser) PipeNode() (localctx IPipeNodeContext) {
	localctx = NewPipeNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PipeParserRULE_pipeNode)

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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PipeParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Function()
		}

	case PipeParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Variable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FunctionName() IFunctionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *FunctionContext) FunctionParameters() IFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *PipeParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PipeParserRULE_function)
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
		p.SetState(125)
		p.FunctionName()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserT__1 {
		{
			p.SetState(126)
			p.FunctionParameters()
		}

	}

	return localctx
}

// IFunctionNameContext is an interface to support dynamic dispatch.
type IFunctionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionNameContext differentiates from other interfaces.
	IsFunctionNameContext()
}

type FunctionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNameContext() *FunctionNameContext {
	var p = new(FunctionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_functionName
	return p
}

func (*FunctionNameContext) IsFunctionNameContext() {}

func NewFunctionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNameContext {
	var p = new(FunctionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_functionName

	return p
}

func (s *FunctionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNameContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(PipeParserIdentifier)
}

func (s *FunctionNameContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, i)
}

func (s *FunctionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFunctionName(s)
	}
}

func (s *FunctionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFunctionName(s)
	}
}

func (p *PipeParser) FunctionName() (localctx IFunctionNameContext) {
	localctx = NewFunctionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PipeParserRULE_functionName)
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
		p.SetState(129)
		p.Match(PipeParserIdentifier)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserT__10 {
		{
			p.SetState(130)
			p.Match(PipeParserT__10)
		}
		{
			p.SetState(131)
			p.Match(PipeParserIdentifier)
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionParametersContext is an interface to support dynamic dispatch.
type IFunctionParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParametersContext differentiates from other interfaces.
	IsFunctionParametersContext()
}

type FunctionParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParametersContext() *FunctionParametersContext {
	var p = new(FunctionParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_functionParameters
	return p
}

func (*FunctionParametersContext) IsFunctionParametersContext() {}

func NewFunctionParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParametersContext {
	var p = new(FunctionParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_functionParameters

	return p
}

func (s *FunctionParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParametersContext) AllFunctionParameter() []IFunctionParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem())
	var tst = make([]IFunctionParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionParameterContext)
		}
	}

	return tst
}

func (s *FunctionParametersContext) FunctionParameter(i int) IFunctionParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterContext)
}

func (s *FunctionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFunctionParameters(s)
	}
}

func (s *FunctionParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFunctionParameters(s)
	}
}

func (p *PipeParser) FunctionParameters() (localctx IFunctionParametersContext) {
	localctx = NewFunctionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PipeParserRULE_functionParameters)
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
		p.SetState(137)
		p.Match(PipeParserT__1)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PipeParserT__11)|(1<<PipeParserT__12)|(1<<PipeParserIntegerLiteral)|(1<<PipeParserDecimalLiteral)|(1<<PipeParserStringLiteral)|(1<<PipeParserBooleanLiteral)|(1<<PipeParserIdentifier))) != 0 {
		{
			p.SetState(138)
			p.FunctionParameter()
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(139)
					p.Match(PipeParserT__2)
				}
				{
					p.SetState(140)
					p.FunctionParameter()
				}

			}
			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PipeParserT__2 {
			{
				p.SetState(146)
				p.Match(PipeParserT__2)
			}

		}

	}
	{
		p.SetState(151)
		p.Match(PipeParserT__3)
	}

	return localctx
}

// IFunctionParameterContext is an interface to support dynamic dispatch.
type IFunctionParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParameterContext differentiates from other interfaces.
	IsFunctionParameterContext()
}

type FunctionParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterContext() *FunctionParameterContext {
	var p = new(FunctionParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_functionParameter
	return p
}

func (*FunctionParameterContext) IsFunctionParameterContext() {}

func NewFunctionParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterContext {
	var p = new(FunctionParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_functionParameter

	return p
}

func (s *FunctionParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunctionParameterContext) DictValue() IDictValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictValueContext)
}

func (s *FunctionParameterContext) NumberValue() INumberValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberValueContext)
}

func (s *FunctionParameterContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *FunctionParameterContext) BooleanValue() IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *FunctionParameterContext) FunctionParameterLabel() IFunctionParameterLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterLabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterLabelContext)
}

func (s *FunctionParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFunctionParameter(s)
	}
}

func (s *FunctionParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFunctionParameter(s)
	}
}

func (p *PipeParser) FunctionParameter() (localctx IFunctionParameterContext) {
	localctx = NewFunctionParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PipeParserRULE_functionParameter)
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserIdentifier {
		{
			p.SetState(153)
			p.FunctionParameterLabel()
		}

	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PipeParserT__11:
		{
			p.SetState(156)
			p.Variable()
		}

	case PipeParserT__12:
		{
			p.SetState(157)
			p.DictValue()
		}

	case PipeParserIntegerLiteral, PipeParserDecimalLiteral:
		{
			p.SetState(158)
			p.NumberValue()
		}

	case PipeParserStringLiteral:
		{
			p.SetState(159)
			p.StringValue()
		}

	case PipeParserBooleanLiteral:
		{
			p.SetState(160)
			p.BooleanValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionParameterLabelContext is an interface to support dynamic dispatch.
type IFunctionParameterLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParameterLabelContext differentiates from other interfaces.
	IsFunctionParameterLabelContext()
}

type FunctionParameterLabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterLabelContext() *FunctionParameterLabelContext {
	var p = new(FunctionParameterLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_functionParameterLabel
	return p
}

func (*FunctionParameterLabelContext) IsFunctionParameterLabelContext() {}

func NewFunctionParameterLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterLabelContext {
	var p = new(FunctionParameterLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_functionParameterLabel

	return p
}

func (s *FunctionParameterLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterLabelContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, 0)
}

func (s *FunctionParameterLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFunctionParameterLabel(s)
	}
}

func (s *FunctionParameterLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFunctionParameterLabel(s)
	}
}

func (p *PipeParser) FunctionParameterLabel() (localctx IFunctionParameterLabelContext) {
	localctx = NewFunctionParameterLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PipeParserRULE_functionParameterLabel)

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
		p.SetState(163)
		p.Match(PipeParserIdentifier)
	}
	{
		p.SetState(164)
		p.Match(PipeParserT__4)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *PipeParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PipeParserRULE_variable)

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
		p.SetState(166)
		p.Match(PipeParserT__11)
	}
	{
		p.SetState(167)
		p.Match(PipeParserIdentifier)
	}

	return localctx
}

// IDictValueContext is an interface to support dynamic dispatch.
type IDictValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictValueContext differentiates from other interfaces.
	IsDictValueContext()
}

type DictValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictValueContext() *DictValueContext {
	var p = new(DictValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_dictValue
	return p
}

func (*DictValueContext) IsDictValueContext() {}

func NewDictValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictValueContext {
	var p = new(DictValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_dictValue

	return p
}

func (s *DictValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DictValueContext) DictEntries() IDictEntriesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictEntriesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictEntriesContext)
}

func (s *DictValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterDictValue(s)
	}
}

func (s *DictValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitDictValue(s)
	}
}

func (p *PipeParser) DictValue() (localctx IDictValueContext) {
	localctx = NewDictValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PipeParserRULE_dictValue)
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
		p.SetState(169)
		p.Match(PipeParserT__12)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PipeParserIntegerLiteral)|(1<<PipeParserDecimalLiteral)|(1<<PipeParserStringLiteral)|(1<<PipeParserBooleanLiteral)|(1<<PipeParserIdentifier))) != 0 {
		{
			p.SetState(170)
			p.DictEntries()
		}

	}
	{
		p.SetState(173)
		p.Match(PipeParserT__13)
	}

	return localctx
}

// IDictEntriesContext is an interface to support dynamic dispatch.
type IDictEntriesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictEntriesContext differentiates from other interfaces.
	IsDictEntriesContext()
}

type DictEntriesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictEntriesContext() *DictEntriesContext {
	var p = new(DictEntriesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_dictEntries
	return p
}

func (*DictEntriesContext) IsDictEntriesContext() {}

func NewDictEntriesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictEntriesContext {
	var p = new(DictEntriesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_dictEntries

	return p
}

func (s *DictEntriesContext) GetParser() antlr.Parser { return s.parser }

func (s *DictEntriesContext) AllDictEntry() []IDictEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDictEntryContext)(nil)).Elem())
	var tst = make([]IDictEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDictEntryContext)
		}
	}

	return tst
}

func (s *DictEntriesContext) DictEntry(i int) IDictEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDictEntryContext)
}

func (s *DictEntriesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictEntriesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictEntriesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterDictEntries(s)
	}
}

func (s *DictEntriesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitDictEntries(s)
	}
}

func (p *PipeParser) DictEntries() (localctx IDictEntriesContext) {
	localctx = NewDictEntriesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PipeParserRULE_dictEntries)
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
		p.SetState(175)
		p.DictEntry()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserT__2 {
		{
			p.SetState(176)
			p.Match(PipeParserT__2)
		}
		{
			p.SetState(177)
			p.DictEntry()
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDictEntryContext is an interface to support dynamic dispatch.
type IDictEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictEntryContext differentiates from other interfaces.
	IsDictEntryContext()
}

type DictEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictEntryContext() *DictEntryContext {
	var p = new(DictEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_dictEntry
	return p
}

func (*DictEntryContext) IsDictEntryContext() {}

func NewDictEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictEntryContext {
	var p = new(DictEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_dictEntry

	return p
}

func (s *DictEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *DictEntryContext) DictEntryValue() IDictEntryValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictEntryValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictEntryValueContext)
}

func (s *DictEntryContext) DictEntryLabel() IDictEntryLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictEntryLabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictEntryLabelContext)
}

func (s *DictEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterDictEntry(s)
	}
}

func (s *DictEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitDictEntry(s)
	}
}

func (p *PipeParser) DictEntry() (localctx IDictEntryContext) {
	localctx = NewDictEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PipeParserRULE_dictEntry)
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
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserIdentifier {
		{
			p.SetState(183)
			p.DictEntryLabel()
		}
		{
			p.SetState(184)
			p.Match(PipeParserT__4)
		}

	}
	{
		p.SetState(188)
		p.DictEntryValue()
	}

	return localctx
}

// IDictEntryLabelContext is an interface to support dynamic dispatch.
type IDictEntryLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictEntryLabelContext differentiates from other interfaces.
	IsDictEntryLabelContext()
}

type DictEntryLabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictEntryLabelContext() *DictEntryLabelContext {
	var p = new(DictEntryLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_dictEntryLabel
	return p
}

func (*DictEntryLabelContext) IsDictEntryLabelContext() {}

func NewDictEntryLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictEntryLabelContext {
	var p = new(DictEntryLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_dictEntryLabel

	return p
}

func (s *DictEntryLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *DictEntryLabelContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, 0)
}

func (s *DictEntryLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictEntryLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictEntryLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterDictEntryLabel(s)
	}
}

func (s *DictEntryLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitDictEntryLabel(s)
	}
}

func (p *PipeParser) DictEntryLabel() (localctx IDictEntryLabelContext) {
	localctx = NewDictEntryLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PipeParserRULE_dictEntryLabel)

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
		p.SetState(190)
		p.Match(PipeParserIdentifier)
	}

	return localctx
}

// IDictEntryValueContext is an interface to support dynamic dispatch.
type IDictEntryValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDictEntryValueContext differentiates from other interfaces.
	IsDictEntryValueContext()
}

type DictEntryValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictEntryValueContext() *DictEntryValueContext {
	var p = new(DictEntryValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_dictEntryValue
	return p
}

func (*DictEntryValueContext) IsDictEntryValueContext() {}

func NewDictEntryValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictEntryValueContext {
	var p = new(DictEntryValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_dictEntryValue

	return p
}

func (s *DictEntryValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DictEntryValueContext) NumberValue() INumberValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberValueContext)
}

func (s *DictEntryValueContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *DictEntryValueContext) BooleanValue() IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *DictEntryValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictEntryValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictEntryValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterDictEntryValue(s)
	}
}

func (s *DictEntryValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitDictEntryValue(s)
	}
}

func (p *PipeParser) DictEntryValue() (localctx IDictEntryValueContext) {
	localctx = NewDictEntryValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PipeParserRULE_dictEntryValue)

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

	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PipeParserIntegerLiteral, PipeParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.NumberValue()
		}

	case PipeParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.StringValue()
		}

	case PipeParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.BooleanValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *NumberValueContext) IntegerValue() IIntegerValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
}

func (s *NumberValueContext) DecimalValue() IDecimalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalValueContext)
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
	p.EnterRule(localctx, 46, PipeParserRULE_numberValue)

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

	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PipeParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.IntegerValue()
		}

	case PipeParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.DecimalValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerValueContext is an interface to support dynamic dispatch.
type IIntegerValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerValueContext differentiates from other interfaces.
	IsIntegerValueContext()
}

type IntegerValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerValueContext() *IntegerValueContext {
	var p = new(IntegerValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_integerValue
	return p
}

func (*IntegerValueContext) IsIntegerValueContext() {}

func NewIntegerValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerValueContext {
	var p = new(IntegerValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_integerValue

	return p
}

func (s *IntegerValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerValueContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(PipeParserIntegerLiteral, 0)
}

func (s *IntegerValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterIntegerValue(s)
	}
}

func (s *IntegerValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitIntegerValue(s)
	}
}

func (p *PipeParser) IntegerValue() (localctx IIntegerValueContext) {
	localctx = NewIntegerValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PipeParserRULE_integerValue)

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
		p.SetState(201)
		p.Match(PipeParserIntegerLiteral)
	}

	return localctx
}

// IDecimalValueContext is an interface to support dynamic dispatch.
type IDecimalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalValueContext differentiates from other interfaces.
	IsDecimalValueContext()
}

type DecimalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalValueContext() *DecimalValueContext {
	var p = new(DecimalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_decimalValue
	return p
}

func (*DecimalValueContext) IsDecimalValueContext() {}

func NewDecimalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalValueContext {
	var p = new(DecimalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_decimalValue

	return p
}

func (s *DecimalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalValueContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(PipeParserDecimalLiteral, 0)
}

func (s *DecimalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterDecimalValue(s)
	}
}

func (s *DecimalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitDecimalValue(s)
	}
}

func (p *PipeParser) DecimalValue() (localctx IDecimalValueContext) {
	localctx = NewDecimalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PipeParserRULE_decimalValue)

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
		p.SetState(203)
		p.Match(PipeParserDecimalLiteral)
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
	p.EnterRule(localctx, 52, PipeParserRULE_stringValue)

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
		p.SetState(205)
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
	p.EnterRule(localctx, 54, PipeParserRULE_booleanValue)

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
		p.SetState(207)
		p.Match(PipeParserBooleanLiteral)
	}

	return localctx
}
