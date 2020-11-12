// Code generated from parser/Pipe.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Pipe

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePipeListener is a complete listener for a parse tree produced by PipeParser.
type BasePipeListener struct{}

var _ PipeListener = &BasePipeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePipeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePipeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePipeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePipeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BasePipeListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BasePipeListener) ExitScript(ctx *ScriptContext) {}

// EnterFunction is called when production function is entered.
func (s *BasePipeListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BasePipeListener) ExitFunction(ctx *FunctionContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BasePipeListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BasePipeListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterFunctionParameters is called when production functionParameters is entered.
func (s *BasePipeListener) EnterFunctionParameters(ctx *FunctionParametersContext) {}

// ExitFunctionParameters is called when production functionParameters is exited.
func (s *BasePipeListener) ExitFunctionParameters(ctx *FunctionParametersContext) {}

// EnterFunctionParameter is called when production functionParameter is entered.
func (s *BasePipeListener) EnterFunctionParameter(ctx *FunctionParameterContext) {}

// ExitFunctionParameter is called when production functionParameter is exited.
func (s *BasePipeListener) ExitFunctionParameter(ctx *FunctionParameterContext) {}

// EnterFunctionParameterLabel is called when production functionParameterLabel is entered.
func (s *BasePipeListener) EnterFunctionParameterLabel(ctx *FunctionParameterLabelContext) {}

// ExitFunctionParameterLabel is called when production functionParameterLabel is exited.
func (s *BasePipeListener) ExitFunctionParameterLabel(ctx *FunctionParameterLabelContext) {}

// EnterDictValue is called when production dictValue is entered.
func (s *BasePipeListener) EnterDictValue(ctx *DictValueContext) {}

// ExitDictValue is called when production dictValue is exited.
func (s *BasePipeListener) ExitDictValue(ctx *DictValueContext) {}

// EnterDictEntries is called when production dictEntries is entered.
func (s *BasePipeListener) EnterDictEntries(ctx *DictEntriesContext) {}

// ExitDictEntries is called when production dictEntries is exited.
func (s *BasePipeListener) ExitDictEntries(ctx *DictEntriesContext) {}

// EnterDictEntry is called when production dictEntry is entered.
func (s *BasePipeListener) EnterDictEntry(ctx *DictEntryContext) {}

// ExitDictEntry is called when production dictEntry is exited.
func (s *BasePipeListener) ExitDictEntry(ctx *DictEntryContext) {}

// EnterDictEntryLabel is called when production dictEntryLabel is entered.
func (s *BasePipeListener) EnterDictEntryLabel(ctx *DictEntryLabelContext) {}

// ExitDictEntryLabel is called when production dictEntryLabel is exited.
func (s *BasePipeListener) ExitDictEntryLabel(ctx *DictEntryLabelContext) {}

// EnterDictEntryValue is called when production dictEntryValue is entered.
func (s *BasePipeListener) EnterDictEntryValue(ctx *DictEntryValueContext) {}

// ExitDictEntryValue is called when production dictEntryValue is exited.
func (s *BasePipeListener) ExitDictEntryValue(ctx *DictEntryValueContext) {}

// EnterNumberValue is called when production numberValue is entered.
func (s *BasePipeListener) EnterNumberValue(ctx *NumberValueContext) {}

// ExitNumberValue is called when production numberValue is exited.
func (s *BasePipeListener) ExitNumberValue(ctx *NumberValueContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BasePipeListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BasePipeListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterDecimalValue is called when production decimalValue is entered.
func (s *BasePipeListener) EnterDecimalValue(ctx *DecimalValueContext) {}

// ExitDecimalValue is called when production decimalValue is exited.
func (s *BasePipeListener) ExitDecimalValue(ctx *DecimalValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BasePipeListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BasePipeListener) ExitStringValue(ctx *StringValueContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BasePipeListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BasePipeListener) ExitBooleanValue(ctx *BooleanValueContext) {}
