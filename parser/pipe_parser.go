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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 277,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3, 2, 7,
	2, 78, 10, 2, 12, 2, 14, 2, 81, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 93, 10, 4, 12, 4, 14, 4, 96, 11, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 7, 5, 102, 10, 5, 12, 5, 14, 5, 105, 11, 5, 5, 5,
	107, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 113, 10, 6, 3, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	128, 10, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 135, 10, 11, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 7, 13, 143, 10, 13, 12, 13, 14,
	13, 146, 11, 13, 3, 14, 3, 14, 3, 14, 7, 14, 151, 10, 14, 12, 14, 14, 14,
	154, 11, 14, 3, 15, 3, 15, 5, 15, 158, 10, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 5, 16, 166, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 7, 19, 177, 10, 19, 12, 19, 14, 19, 180, 11,
	19, 3, 20, 3, 20, 5, 20, 184, 10, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	5, 22, 191, 10, 22, 3, 23, 3, 23, 3, 23, 7, 23, 196, 10, 23, 12, 23, 14,
	23, 199, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 205, 10, 24, 12, 24,
	14, 24, 208, 11, 24, 3, 24, 5, 24, 211, 10, 24, 5, 24, 213, 10, 24, 3,
	24, 3, 24, 3, 25, 5, 25, 218, 10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 230, 10, 27, 3, 28, 3, 28, 3,
	28, 3, 29, 3, 29, 5, 29, 237, 10, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	7, 30, 244, 10, 30, 12, 30, 14, 30, 247, 11, 30, 3, 31, 3, 31, 3, 31, 5,
	31, 252, 10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 5, 33, 263, 10, 33, 3, 34, 3, 34, 5, 34, 267, 10, 34, 3, 35, 3,
	35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 2, 2, 39, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 2, 2, 2,
	270, 2, 79, 3, 2, 2, 2, 4, 84, 3, 2, 2, 2, 6, 89, 3, 2, 2, 2, 8, 97, 3,
	2, 2, 2, 10, 110, 3, 2, 2, 2, 12, 117, 3, 2, 2, 2, 14, 119, 3, 2, 2, 2,
	16, 121, 3, 2, 2, 2, 18, 123, 3, 2, 2, 2, 20, 132, 3, 2, 2, 2, 22, 136,
	3, 2, 2, 2, 24, 139, 3, 2, 2, 2, 26, 147, 3, 2, 2, 2, 28, 157, 3, 2, 2,
	2, 30, 159, 3, 2, 2, 2, 32, 169, 3, 2, 2, 2, 34, 171, 3, 2, 2, 2, 36, 173,
	3, 2, 2, 2, 38, 183, 3, 2, 2, 2, 40, 185, 3, 2, 2, 2, 42, 188, 3, 2, 2,
	2, 44, 192, 3, 2, 2, 2, 46, 200, 3, 2, 2, 2, 48, 217, 3, 2, 2, 2, 50, 221,
	3, 2, 2, 2, 52, 229, 3, 2, 2, 2, 54, 231, 3, 2, 2, 2, 56, 234, 3, 2, 2,
	2, 58, 240, 3, 2, 2, 2, 60, 251, 3, 2, 2, 2, 62, 255, 3, 2, 2, 2, 64, 262,
	3, 2, 2, 2, 66, 266, 3, 2, 2, 2, 68, 268, 3, 2, 2, 2, 70, 270, 3, 2, 2,
	2, 72, 272, 3, 2, 2, 2, 74, 274, 3, 2, 2, 2, 76, 78, 5, 4, 3, 2, 77, 76,
	3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2,
	80, 82, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 83, 7, 2, 2, 3, 83, 3, 3, 2,
	2, 2, 84, 85, 7, 3, 2, 2, 85, 86, 5, 6, 4, 2, 86, 87, 5, 8, 5, 2, 87, 88,
	5, 18, 10, 2, 88, 5, 3, 2, 2, 2, 89, 94, 7, 26, 2, 2, 90, 91, 7, 4, 2,
	2, 91, 93, 7, 26, 2, 2, 92, 90, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92,
	3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 7, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2,
	97, 106, 7, 5, 2, 2, 98, 103, 5, 10, 6, 2, 99, 100, 7, 6, 2, 2, 100, 102,
	5, 10, 6, 2, 101, 99, 3, 2, 2, 2, 102, 105, 3, 2, 2, 2, 103, 101, 3, 2,
	2, 2, 103, 104, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2,
	106, 98, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109,
	7, 7, 2, 2, 109, 9, 3, 2, 2, 2, 110, 112, 5, 12, 7, 2, 111, 113, 5, 14,
	8, 2, 112, 111, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2,
	114, 115, 7, 8, 2, 2, 115, 116, 5, 16, 9, 2, 116, 11, 3, 2, 2, 2, 117,
	118, 7, 26, 2, 2, 118, 13, 3, 2, 2, 2, 119, 120, 7, 9, 2, 2, 120, 15, 3,
	2, 2, 2, 121, 122, 7, 19, 2, 2, 122, 17, 3, 2, 2, 2, 123, 127, 7, 10, 2,
	2, 124, 125, 5, 24, 13, 2, 125, 126, 7, 18, 2, 2, 126, 128, 3, 2, 2, 2,
	127, 124, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129,
	130, 5, 20, 11, 2, 130, 131, 7, 11, 2, 2, 131, 19, 3, 2, 2, 2, 132, 134,
	7, 12, 2, 2, 133, 135, 5, 26, 14, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3,
	2, 2, 2, 135, 21, 3, 2, 2, 2, 136, 137, 5, 24, 13, 2, 137, 138, 7, 2, 2,
	3, 138, 23, 3, 2, 2, 2, 139, 144, 5, 26, 14, 2, 140, 141, 7, 18, 2, 2,
	141, 143, 5, 26, 14, 2, 142, 140, 3, 2, 2, 2, 143, 146, 3, 2, 2, 2, 144,
	142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 25, 3, 2, 2, 2, 146, 144, 3,
	2, 2, 2, 147, 152, 5, 28, 15, 2, 148, 149, 7, 17, 2, 2, 149, 151, 5, 28,
	15, 2, 150, 148, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2,
	152, 153, 3, 2, 2, 2, 153, 27, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 155, 158,
	5, 30, 16, 2, 156, 158, 5, 36, 19, 2, 157, 155, 3, 2, 2, 2, 157, 156, 3,
	2, 2, 2, 158, 29, 3, 2, 2, 2, 159, 160, 7, 13, 2, 2, 160, 161, 5, 32, 17,
	2, 161, 162, 7, 6, 2, 2, 162, 165, 5, 26, 14, 2, 163, 164, 7, 6, 2, 2,
	164, 166, 5, 34, 18, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	167, 3, 2, 2, 2, 167, 168, 7, 13, 2, 2, 168, 31, 3, 2, 2, 2, 169, 170,
	5, 42, 22, 2, 170, 33, 3, 2, 2, 2, 171, 172, 5, 42, 22, 2, 172, 35, 3,
	2, 2, 2, 173, 178, 5, 38, 20, 2, 174, 175, 7, 6, 2, 2, 175, 177, 5, 38,
	20, 2, 176, 174, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2,
	178, 179, 3, 2, 2, 2, 179, 37, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 181, 184,
	5, 40, 21, 2, 182, 184, 5, 42, 22, 2, 183, 181, 3, 2, 2, 2, 183, 182, 3,
	2, 2, 2, 184, 39, 3, 2, 2, 2, 185, 186, 7, 14, 2, 2, 186, 187, 7, 26, 2,
	2, 187, 41, 3, 2, 2, 2, 188, 190, 5, 44, 23, 2, 189, 191, 5, 46, 24, 2,
	190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 43, 3, 2, 2, 2, 192, 197,
	7, 26, 2, 2, 193, 194, 7, 4, 2, 2, 194, 196, 7, 26, 2, 2, 195, 193, 3,
	2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2,
	2, 198, 45, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 212, 7, 5, 2, 2, 201,
	206, 5, 48, 25, 2, 202, 203, 7, 6, 2, 2, 203, 205, 5, 48, 25, 2, 204, 202,
	3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2,
	2, 2, 207, 210, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 211, 7, 6, 2, 2,
	210, 209, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 213, 3, 2, 2, 2, 212,
	201, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 215,
	7, 7, 2, 2, 215, 47, 3, 2, 2, 2, 216, 218, 5, 50, 26, 2, 217, 216, 3, 2,
	2, 2, 217, 218, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 220, 5, 52, 27,
	2, 220, 49, 3, 2, 2, 2, 221, 222, 7, 26, 2, 2, 222, 223, 7, 8, 2, 2, 223,
	51, 3, 2, 2, 2, 224, 230, 5, 54, 28, 2, 225, 230, 5, 56, 29, 2, 226, 230,
	5, 66, 34, 2, 227, 230, 5, 72, 37, 2, 228, 230, 5, 74, 38, 2, 229, 224,
	3, 2, 2, 2, 229, 225, 3, 2, 2, 2, 229, 226, 3, 2, 2, 2, 229, 227, 3, 2,
	2, 2, 229, 228, 3, 2, 2, 2, 230, 53, 3, 2, 2, 2, 231, 232, 7, 14, 2, 2,
	232, 233, 7, 26, 2, 2, 233, 55, 3, 2, 2, 2, 234, 236, 7, 15, 2, 2, 235,
	237, 5, 58, 30, 2, 236, 235, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 238,
	3, 2, 2, 2, 238, 239, 7, 16, 2, 2, 239, 57, 3, 2, 2, 2, 240, 245, 5, 60,
	31, 2, 241, 242, 7, 6, 2, 2, 242, 244, 5, 60, 31, 2, 243, 241, 3, 2, 2,
	2, 244, 247, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246,
	59, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248, 249, 5, 62, 32, 2, 249, 250,
	7, 8, 2, 2, 250, 252, 3, 2, 2, 2, 251, 248, 3, 2, 2, 2, 251, 252, 3, 2,
	2, 2, 252, 253, 3, 2, 2, 2, 253, 254, 5, 64, 33, 2, 254, 61, 3, 2, 2, 2,
	255, 256, 7, 26, 2, 2, 256, 63, 3, 2, 2, 2, 257, 263, 5, 66, 34, 2, 258,
	263, 5, 70, 36, 2, 259, 263, 5, 72, 37, 2, 260, 263, 5, 74, 38, 2, 261,
	263, 5, 54, 28, 2, 262, 257, 3, 2, 2, 2, 262, 258, 3, 2, 2, 2, 262, 259,
	3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 262, 261, 3, 2, 2, 2, 263, 65, 3, 2,
	2, 2, 264, 267, 5, 68, 35, 2, 265, 267, 5, 70, 36, 2, 266, 264, 3, 2, 2,
	2, 266, 265, 3, 2, 2, 2, 267, 67, 3, 2, 2, 2, 268, 269, 7, 20, 2, 2, 269,
	69, 3, 2, 2, 2, 270, 271, 7, 21, 2, 2, 271, 71, 3, 2, 2, 2, 272, 273, 7,
	22, 2, 2, 273, 73, 3, 2, 2, 2, 274, 275, 7, 23, 2, 2, 275, 75, 3, 2, 2,
	2, 27, 79, 94, 103, 106, 112, 127, 134, 144, 152, 157, 165, 178, 183, 190,
	197, 206, 210, 212, 217, 229, 236, 245, 251, 262, 266,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'def'", "'.'", "'('", "','", "')'", "':'", "'?'", "'{'", "'}'", "'return'",
	"'/'", "'$'", "'['", "']'", "", "", "", "", "", "", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Connect",
	"PipeSeparator", "ValueType", "IntegerLiteral", "DecimalLiteral", "StringLiteral",
	"BooleanLiteral", "True", "False", "Identifier", "WS",
}

var ruleNames = []string{
	"script", "funcDef", "funcName", "funcParamsDef", "funcParamDef", "funcParamName",
	"optionalParamFlag", "funcParamType", "funcBody", "returnStatement", "cmd",
	"multiPipe", "pipe", "pipeNode", "streamNode", "streamSplitter", "streamCollector",
	"pipeNodeArray", "pipeNodeElement", "variableNode", "functionNode", "functionName",
	"functionParameters", "functionParameter", "functionParameterLabel", "functionParameterValue",
	"variableValue", "dictValue", "dictEntries", "dictEntry", "dictEntryLabel",
	"dictEntryValue", "numberValue", "integerValue", "decimalValue", "stringValue",
	"booleanValue",
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
	PipeParserPipeSeparator  = 16
	PipeParserValueType      = 17
	PipeParserIntegerLiteral = 18
	PipeParserDecimalLiteral = 19
	PipeParserStringLiteral  = 20
	PipeParserBooleanLiteral = 21
	PipeParserTrue           = 22
	PipeParserFalse          = 23
	PipeParserIdentifier     = 24
	PipeParserWS             = 25
)

// PipeParser rules.
const (
	PipeParserRULE_script                 = 0
	PipeParserRULE_funcDef                = 1
	PipeParserRULE_funcName               = 2
	PipeParserRULE_funcParamsDef          = 3
	PipeParserRULE_funcParamDef           = 4
	PipeParserRULE_funcParamName          = 5
	PipeParserRULE_optionalParamFlag      = 6
	PipeParserRULE_funcParamType          = 7
	PipeParserRULE_funcBody               = 8
	PipeParserRULE_returnStatement        = 9
	PipeParserRULE_cmd                    = 10
	PipeParserRULE_multiPipe              = 11
	PipeParserRULE_pipe                   = 12
	PipeParserRULE_pipeNode               = 13
	PipeParserRULE_streamNode             = 14
	PipeParserRULE_streamSplitter         = 15
	PipeParserRULE_streamCollector        = 16
	PipeParserRULE_pipeNodeArray          = 17
	PipeParserRULE_pipeNodeElement        = 18
	PipeParserRULE_variableNode           = 19
	PipeParserRULE_functionNode           = 20
	PipeParserRULE_functionName           = 21
	PipeParserRULE_functionParameters     = 22
	PipeParserRULE_functionParameter      = 23
	PipeParserRULE_functionParameterLabel = 24
	PipeParserRULE_functionParameterValue = 25
	PipeParserRULE_variableValue          = 26
	PipeParserRULE_dictValue              = 27
	PipeParserRULE_dictEntries            = 28
	PipeParserRULE_dictEntry              = 29
	PipeParserRULE_dictEntryLabel         = 30
	PipeParserRULE_dictEntryValue         = 31
	PipeParserRULE_numberValue            = 32
	PipeParserRULE_integerValue           = 33
	PipeParserRULE_decimalValue           = 34
	PipeParserRULE_stringValue            = 35
	PipeParserRULE_booleanValue           = 36
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

func (s *ScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(PipeParserEOF, 0)
}

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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserT__0 {
		{
			p.SetState(74)
			p.FuncDef()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(PipeParserEOF)
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

func (s *FuncDefContext) FuncName() IFuncNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncNameContext)
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
		p.SetState(82)
		p.Match(PipeParserT__0)
	}
	{
		p.SetState(83)
		p.FuncName()
	}
	{
		p.SetState(84)
		p.FuncParamsDef()
	}
	{
		p.SetState(85)
		p.FuncBody()
	}

	return localctx
}

// IFuncNameContext is an interface to support dynamic dispatch.
type IFuncNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncNameContext differentiates from other interfaces.
	IsFuncNameContext()
}

type FuncNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameContext() *FuncNameContext {
	var p = new(FuncNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_funcName
	return p
}

func (*FuncNameContext) IsFuncNameContext() {}

func NewFuncNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameContext {
	var p = new(FuncNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_funcName

	return p
}

func (s *FuncNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncNameContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(PipeParserIdentifier)
}

func (s *FuncNameContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, i)
}

func (s *FuncNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFuncName(s)
	}
}

func (s *FuncNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFuncName(s)
	}
}

func (p *PipeParser) FuncName() (localctx IFuncNameContext) {
	localctx = NewFuncNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PipeParserRULE_funcName)
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
		p.SetState(87)
		p.Match(PipeParserIdentifier)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserT__1 {
		{
			p.SetState(88)
			p.Match(PipeParserT__1)
		}
		{
			p.SetState(89)
			p.Match(PipeParserIdentifier)
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 6, PipeParserRULE_funcParamsDef)
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
		p.SetState(95)
		p.Match(PipeParserT__2)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserIdentifier {
		{
			p.SetState(96)
			p.FuncParamDef()
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PipeParserT__3 {
			{
				p.SetState(97)
				p.Match(PipeParserT__3)
			}
			{
				p.SetState(98)
				p.FuncParamDef()
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(106)
		p.Match(PipeParserT__4)
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
	p.EnterRule(localctx, 8, PipeParserRULE_funcParamDef)
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
		p.SetState(108)
		p.FuncParamName()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserT__6 {
		{
			p.SetState(109)
			p.OptionalParamFlag()
		}

	}
	{
		p.SetState(112)
		p.Match(PipeParserT__5)
	}
	{
		p.SetState(113)
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
	p.EnterRule(localctx, 10, PipeParserRULE_funcParamName)

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
		p.SetState(115)
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
	p.EnterRule(localctx, 12, PipeParserRULE_optionalParamFlag)

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
		p.SetState(117)
		p.Match(PipeParserT__6)
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
	p.EnterRule(localctx, 14, PipeParserRULE_funcParamType)

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
		p.SetState(119)
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

func (s *FuncBodyContext) PipeSeparator() antlr.TerminalNode {
	return s.GetToken(PipeParserPipeSeparator, 0)
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
	p.EnterRule(localctx, 16, PipeParserRULE_funcBody)
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
		p.SetState(121)
		p.Match(PipeParserT__7)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PipeParserT__10)|(1<<PipeParserT__11)|(1<<PipeParserIdentifier))) != 0 {
		{
			p.SetState(122)
			p.MultiPipe()
		}
		{
			p.SetState(123)
			p.Match(PipeParserPipeSeparator)
		}

	}
	{
		p.SetState(127)
		p.ReturnStatement()
	}
	{
		p.SetState(128)
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
	p.EnterRule(localctx, 18, PipeParserRULE_returnStatement)
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
		p.SetState(130)
		p.Match(PipeParserT__9)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PipeParserT__10)|(1<<PipeParserT__11)|(1<<PipeParserIdentifier))) != 0 {
		{
			p.SetState(131)
			p.Pipe()
		}

	}

	return localctx
}

// ICmdContext is an interface to support dynamic dispatch.
type ICmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCmdContext differentiates from other interfaces.
	IsCmdContext()
}

type CmdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdContext() *CmdContext {
	var p = new(CmdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_cmd
	return p
}

func (*CmdContext) IsCmdContext() {}

func NewCmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdContext {
	var p = new(CmdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_cmd

	return p
}

func (s *CmdContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdContext) MultiPipe() IMultiPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiPipeContext)
}

func (s *CmdContext) EOF() antlr.TerminalNode {
	return s.GetToken(PipeParserEOF, 0)
}

func (s *CmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterCmd(s)
	}
}

func (s *CmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitCmd(s)
	}
}

func (p *PipeParser) Cmd() (localctx ICmdContext) {
	localctx = NewCmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PipeParserRULE_cmd)

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
		p.SetState(134)
		p.MultiPipe()
	}
	{
		p.SetState(135)
		p.Match(PipeParserEOF)
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

func (s *MultiPipeContext) AllPipeSeparator() []antlr.TerminalNode {
	return s.GetTokens(PipeParserPipeSeparator)
}

func (s *MultiPipeContext) PipeSeparator(i int) antlr.TerminalNode {
	return s.GetToken(PipeParserPipeSeparator, i)
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
	p.EnterRule(localctx, 22, PipeParserRULE_multiPipe)

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
		p.Pipe()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(138)
				p.Match(PipeParserPipeSeparator)
			}
			{
				p.SetState(139)
				p.Pipe()
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 24, PipeParserRULE_pipe)
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
		p.SetState(145)
		p.PipeNode()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserConnect {
		{
			p.SetState(146)
			p.Match(PipeParserConnect)
		}
		{
			p.SetState(147)
			p.PipeNode()
		}

		p.SetState(152)
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

func (s *PipeNodeContext) StreamNode() IStreamNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStreamNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStreamNodeContext)
}

func (s *PipeNodeContext) PipeNodeArray() IPipeNodeArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeNodeArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeNodeArrayContext)
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
	p.EnterRule(localctx, 26, PipeParserRULE_pipeNode)

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

	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PipeParserT__10:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.StreamNode()
		}

	case PipeParserT__11, PipeParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.PipeNodeArray()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStreamNodeContext is an interface to support dynamic dispatch.
type IStreamNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreamNodeContext differentiates from other interfaces.
	IsStreamNodeContext()
}

type StreamNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreamNodeContext() *StreamNodeContext {
	var p = new(StreamNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_streamNode
	return p
}

func (*StreamNodeContext) IsStreamNodeContext() {}

func NewStreamNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreamNodeContext {
	var p = new(StreamNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_streamNode

	return p
}

func (s *StreamNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *StreamNodeContext) StreamSplitter() IStreamSplitterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStreamSplitterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStreamSplitterContext)
}

func (s *StreamNodeContext) Pipe() IPipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPipeContext)
}

func (s *StreamNodeContext) StreamCollector() IStreamCollectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStreamCollectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStreamCollectorContext)
}

func (s *StreamNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreamNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StreamNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterStreamNode(s)
	}
}

func (s *StreamNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitStreamNode(s)
	}
}

func (p *PipeParser) StreamNode() (localctx IStreamNodeContext) {
	localctx = NewStreamNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PipeParserRULE_streamNode)
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
		p.SetState(157)
		p.Match(PipeParserT__10)
	}
	{
		p.SetState(158)
		p.StreamSplitter()
	}
	{
		p.SetState(159)
		p.Match(PipeParserT__3)
	}
	{
		p.SetState(160)
		p.Pipe()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserT__3 {
		{
			p.SetState(161)
			p.Match(PipeParserT__3)
		}
		{
			p.SetState(162)
			p.StreamCollector()
		}

	}
	{
		p.SetState(165)
		p.Match(PipeParserT__10)
	}

	return localctx
}

// IStreamSplitterContext is an interface to support dynamic dispatch.
type IStreamSplitterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreamSplitterContext differentiates from other interfaces.
	IsStreamSplitterContext()
}

type StreamSplitterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreamSplitterContext() *StreamSplitterContext {
	var p = new(StreamSplitterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_streamSplitter
	return p
}

func (*StreamSplitterContext) IsStreamSplitterContext() {}

func NewStreamSplitterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreamSplitterContext {
	var p = new(StreamSplitterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_streamSplitter

	return p
}

func (s *StreamSplitterContext) GetParser() antlr.Parser { return s.parser }

func (s *StreamSplitterContext) FunctionNode() IFunctionNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNodeContext)
}

func (s *StreamSplitterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreamSplitterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StreamSplitterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterStreamSplitter(s)
	}
}

func (s *StreamSplitterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitStreamSplitter(s)
	}
}

func (p *PipeParser) StreamSplitter() (localctx IStreamSplitterContext) {
	localctx = NewStreamSplitterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PipeParserRULE_streamSplitter)

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
		p.SetState(167)
		p.FunctionNode()
	}

	return localctx
}

// IStreamCollectorContext is an interface to support dynamic dispatch.
type IStreamCollectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreamCollectorContext differentiates from other interfaces.
	IsStreamCollectorContext()
}

type StreamCollectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreamCollectorContext() *StreamCollectorContext {
	var p = new(StreamCollectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_streamCollector
	return p
}

func (*StreamCollectorContext) IsStreamCollectorContext() {}

func NewStreamCollectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreamCollectorContext {
	var p = new(StreamCollectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_streamCollector

	return p
}

func (s *StreamCollectorContext) GetParser() antlr.Parser { return s.parser }

func (s *StreamCollectorContext) FunctionNode() IFunctionNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNodeContext)
}

func (s *StreamCollectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreamCollectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StreamCollectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterStreamCollector(s)
	}
}

func (s *StreamCollectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitStreamCollector(s)
	}
}

func (p *PipeParser) StreamCollector() (localctx IStreamCollectorContext) {
	localctx = NewStreamCollectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PipeParserRULE_streamCollector)

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
		p.FunctionNode()
	}

	return localctx
}

// IPipeNodeArrayContext is an interface to support dynamic dispatch.
type IPipeNodeArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeNodeArrayContext differentiates from other interfaces.
	IsPipeNodeArrayContext()
}

type PipeNodeArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeNodeArrayContext() *PipeNodeArrayContext {
	var p = new(PipeNodeArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_pipeNodeArray
	return p
}

func (*PipeNodeArrayContext) IsPipeNodeArrayContext() {}

func NewPipeNodeArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeNodeArrayContext {
	var p = new(PipeNodeArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_pipeNodeArray

	return p
}

func (s *PipeNodeArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeNodeArrayContext) AllPipeNodeElement() []IPipeNodeElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPipeNodeElementContext)(nil)).Elem())
	var tst = make([]IPipeNodeElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPipeNodeElementContext)
		}
	}

	return tst
}

func (s *PipeNodeArrayContext) PipeNodeElement(i int) IPipeNodeElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPipeNodeElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPipeNodeElementContext)
}

func (s *PipeNodeArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeNodeArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeNodeArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterPipeNodeArray(s)
	}
}

func (s *PipeNodeArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitPipeNodeArray(s)
	}
}

func (p *PipeParser) PipeNodeArray() (localctx IPipeNodeArrayContext) {
	localctx = NewPipeNodeArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PipeParserRULE_pipeNodeArray)

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
		p.SetState(171)
		p.PipeNodeElement()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(172)
				p.Match(PipeParserT__3)
			}
			{
				p.SetState(173)
				p.PipeNodeElement()
			}

		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IPipeNodeElementContext is an interface to support dynamic dispatch.
type IPipeNodeElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPipeNodeElementContext differentiates from other interfaces.
	IsPipeNodeElementContext()
}

type PipeNodeElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeNodeElementContext() *PipeNodeElementContext {
	var p = new(PipeNodeElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_pipeNodeElement
	return p
}

func (*PipeNodeElementContext) IsPipeNodeElementContext() {}

func NewPipeNodeElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeNodeElementContext {
	var p = new(PipeNodeElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_pipeNodeElement

	return p
}

func (s *PipeNodeElementContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeNodeElementContext) VariableNode() IVariableNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNodeContext)
}

func (s *PipeNodeElementContext) FunctionNode() IFunctionNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNodeContext)
}

func (s *PipeNodeElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeNodeElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeNodeElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterPipeNodeElement(s)
	}
}

func (s *PipeNodeElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitPipeNodeElement(s)
	}
}

func (p *PipeParser) PipeNodeElement() (localctx IPipeNodeElementContext) {
	localctx = NewPipeNodeElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PipeParserRULE_pipeNodeElement)

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

	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PipeParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.VariableNode()
		}

	case PipeParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.FunctionNode()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableNodeContext is an interface to support dynamic dispatch.
type IVariableNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableNodeContext differentiates from other interfaces.
	IsVariableNodeContext()
}

type VariableNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableNodeContext() *VariableNodeContext {
	var p = new(VariableNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_variableNode
	return p
}

func (*VariableNodeContext) IsVariableNodeContext() {}

func NewVariableNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableNodeContext {
	var p = new(VariableNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_variableNode

	return p
}

func (s *VariableNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableNodeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, 0)
}

func (s *VariableNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterVariableNode(s)
	}
}

func (s *VariableNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitVariableNode(s)
	}
}

func (p *PipeParser) VariableNode() (localctx IVariableNodeContext) {
	localctx = NewVariableNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PipeParserRULE_variableNode)

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
		p.SetState(183)
		p.Match(PipeParserT__11)
	}
	{
		p.SetState(184)
		p.Match(PipeParserIdentifier)
	}

	return localctx
}

// IFunctionNodeContext is an interface to support dynamic dispatch.
type IFunctionNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionNodeContext differentiates from other interfaces.
	IsFunctionNodeContext()
}

type FunctionNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionNodeContext() *FunctionNodeContext {
	var p = new(FunctionNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_functionNode
	return p
}

func (*FunctionNodeContext) IsFunctionNodeContext() {}

func NewFunctionNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionNodeContext {
	var p = new(FunctionNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_functionNode

	return p
}

func (s *FunctionNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionNodeContext) FunctionName() IFunctionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionNameContext)
}

func (s *FunctionNodeContext) FunctionParameters() IFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *FunctionNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFunctionNode(s)
	}
}

func (s *FunctionNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFunctionNode(s)
	}
}

func (p *PipeParser) FunctionNode() (localctx IFunctionNodeContext) {
	localctx = NewFunctionNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PipeParserRULE_functionNode)
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
		p.SetState(186)
		p.FunctionName()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserT__2 {
		{
			p.SetState(187)
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
	p.EnterRule(localctx, 42, PipeParserRULE_functionName)
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
		p.SetState(190)
		p.Match(PipeParserIdentifier)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserT__1 {
		{
			p.SetState(191)
			p.Match(PipeParserT__1)
		}
		{
			p.SetState(192)
			p.Match(PipeParserIdentifier)
		}

		p.SetState(197)
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
	p.EnterRule(localctx, 44, PipeParserRULE_functionParameters)
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
		p.SetState(198)
		p.Match(PipeParserT__2)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PipeParserT__11)|(1<<PipeParserT__12)|(1<<PipeParserIntegerLiteral)|(1<<PipeParserDecimalLiteral)|(1<<PipeParserStringLiteral)|(1<<PipeParserBooleanLiteral)|(1<<PipeParserIdentifier))) != 0 {
		{
			p.SetState(199)
			p.FunctionParameter()
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(200)
					p.Match(PipeParserT__3)
				}
				{
					p.SetState(201)
					p.FunctionParameter()
				}

			}
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PipeParserT__3 {
			{
				p.SetState(207)
				p.Match(PipeParserT__3)
			}

		}

	}
	{
		p.SetState(212)
		p.Match(PipeParserT__4)
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

func (s *FunctionParameterContext) FunctionParameterValue() IFunctionParameterValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterValueContext)
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
	p.EnterRule(localctx, 46, PipeParserRULE_functionParameter)
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
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserIdentifier {
		{
			p.SetState(214)
			p.FunctionParameterLabel()
		}

	}
	{
		p.SetState(217)
		p.FunctionParameterValue()
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
	p.EnterRule(localctx, 48, PipeParserRULE_functionParameterLabel)

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
		p.SetState(219)
		p.Match(PipeParserIdentifier)
	}
	{
		p.SetState(220)
		p.Match(PipeParserT__5)
	}

	return localctx
}

// IFunctionParameterValueContext is an interface to support dynamic dispatch.
type IFunctionParameterValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionParameterValueContext differentiates from other interfaces.
	IsFunctionParameterValueContext()
}

type FunctionParameterValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionParameterValueContext() *FunctionParameterValueContext {
	var p = new(FunctionParameterValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_functionParameterValue
	return p
}

func (*FunctionParameterValueContext) IsFunctionParameterValueContext() {}

func NewFunctionParameterValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterValueContext {
	var p = new(FunctionParameterValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_functionParameterValue

	return p
}

func (s *FunctionParameterValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterValueContext) VariableValue() IVariableValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableValueContext)
}

func (s *FunctionParameterValueContext) DictValue() IDictValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDictValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDictValueContext)
}

func (s *FunctionParameterValueContext) NumberValue() INumberValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberValueContext)
}

func (s *FunctionParameterValueContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *FunctionParameterValueContext) BooleanValue() IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *FunctionParameterValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterFunctionParameterValue(s)
	}
}

func (s *FunctionParameterValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitFunctionParameterValue(s)
	}
}

func (p *PipeParser) FunctionParameterValue() (localctx IFunctionParameterValueContext) {
	localctx = NewFunctionParameterValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PipeParserRULE_functionParameterValue)

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

	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PipeParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.VariableValue()
		}

	case PipeParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.DictValue()
		}

	case PipeParserIntegerLiteral, PipeParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(224)
			p.NumberValue()
		}

	case PipeParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(225)
			p.StringValue()
		}

	case PipeParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(226)
			p.BooleanValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableValueContext is an interface to support dynamic dispatch.
type IVariableValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableValueContext differentiates from other interfaces.
	IsVariableValueContext()
}

type VariableValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableValueContext() *VariableValueContext {
	var p = new(VariableValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PipeParserRULE_variableValue
	return p
}

func (*VariableValueContext) IsVariableValueContext() {}

func NewVariableValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableValueContext {
	var p = new(VariableValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeParserRULE_variableValue

	return p
}

func (s *VariableValueContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableValueContext) Identifier() antlr.TerminalNode {
	return s.GetToken(PipeParserIdentifier, 0)
}

func (s *VariableValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.EnterVariableValue(s)
	}
}

func (s *VariableValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeListener); ok {
		listenerT.ExitVariableValue(s)
	}
}

func (p *PipeParser) VariableValue() (localctx IVariableValueContext) {
	localctx = NewVariableValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PipeParserRULE_variableValue)

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
		p.SetState(229)
		p.Match(PipeParserT__11)
	}
	{
		p.SetState(230)
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
	p.EnterRule(localctx, 54, PipeParserRULE_dictValue)
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
		p.SetState(232)
		p.Match(PipeParserT__12)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PipeParserT__11)|(1<<PipeParserIntegerLiteral)|(1<<PipeParserDecimalLiteral)|(1<<PipeParserStringLiteral)|(1<<PipeParserBooleanLiteral)|(1<<PipeParserIdentifier))) != 0 {
		{
			p.SetState(233)
			p.DictEntries()
		}

	}
	{
		p.SetState(236)
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
	p.EnterRule(localctx, 56, PipeParserRULE_dictEntries)
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
		p.SetState(238)
		p.DictEntry()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PipeParserT__3 {
		{
			p.SetState(239)
			p.Match(PipeParserT__3)
		}
		{
			p.SetState(240)
			p.DictEntry()
		}

		p.SetState(245)
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
	p.EnterRule(localctx, 58, PipeParserRULE_dictEntry)
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
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PipeParserIdentifier {
		{
			p.SetState(246)
			p.DictEntryLabel()
		}
		{
			p.SetState(247)
			p.Match(PipeParserT__5)
		}

	}
	{
		p.SetState(251)
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
	p.EnterRule(localctx, 60, PipeParserRULE_dictEntryLabel)

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
		p.SetState(253)
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

func (s *DictEntryValueContext) DecimalValue() IDecimalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalValueContext)
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

func (s *DictEntryValueContext) VariableValue() IVariableValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableValueContext)
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
	p.EnterRule(localctx, 62, PipeParserRULE_dictEntryValue)

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

	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(255)
			p.NumberValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.DecimalValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(257)
			p.StringValue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(258)
			p.BooleanValue()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(259)
			p.VariableValue()
		}

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
	p.EnterRule(localctx, 64, PipeParserRULE_numberValue)

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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PipeParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.IntegerValue()
		}

	case PipeParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
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
	p.EnterRule(localctx, 66, PipeParserRULE_integerValue)

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
		p.SetState(266)
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
	p.EnterRule(localctx, 68, PipeParserRULE_decimalValue)

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
		p.SetState(268)
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
	p.EnterRule(localctx, 70, PipeParserRULE_stringValue)

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
		p.SetState(270)
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
	p.EnterRule(localctx, 72, PipeParserRULE_booleanValue)

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
		p.SetState(272)
		p.Match(PipeParserBooleanLiteral)
	}

	return localctx
}
