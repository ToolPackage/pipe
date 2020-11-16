// Code generated from parser/Pipe.g4 by ANTLR 4.7.1. DO NOT EDIT.

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

// EnterFuncDef is called when production funcDef is entered.
func (s *BasePipeListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BasePipeListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterFuncName is called when production funcName is entered.
func (s *BasePipeListener) EnterFuncName(ctx *FuncNameContext) {}

// ExitFuncName is called when production funcName is exited.
func (s *BasePipeListener) ExitFuncName(ctx *FuncNameContext) {}

// EnterFuncParamsDef is called when production funcParamsDef is entered.
func (s *BasePipeListener) EnterFuncParamsDef(ctx *FuncParamsDefContext) {}

// ExitFuncParamsDef is called when production funcParamsDef is exited.
func (s *BasePipeListener) ExitFuncParamsDef(ctx *FuncParamsDefContext) {}

// EnterFuncParamDef is called when production funcParamDef is entered.
func (s *BasePipeListener) EnterFuncParamDef(ctx *FuncParamDefContext) {}

// ExitFuncParamDef is called when production funcParamDef is exited.
func (s *BasePipeListener) ExitFuncParamDef(ctx *FuncParamDefContext) {}

// EnterFuncParamName is called when production funcParamName is entered.
func (s *BasePipeListener) EnterFuncParamName(ctx *FuncParamNameContext) {}

// ExitFuncParamName is called when production funcParamName is exited.
func (s *BasePipeListener) ExitFuncParamName(ctx *FuncParamNameContext) {}

// EnterOptionalParamFlag is called when production optionalParamFlag is entered.
func (s *BasePipeListener) EnterOptionalParamFlag(ctx *OptionalParamFlagContext) {}

// ExitOptionalParamFlag is called when production optionalParamFlag is exited.
func (s *BasePipeListener) ExitOptionalParamFlag(ctx *OptionalParamFlagContext) {}

// EnterFuncParamType is called when production funcParamType is entered.
func (s *BasePipeListener) EnterFuncParamType(ctx *FuncParamTypeContext) {}

// ExitFuncParamType is called when production funcParamType is exited.
func (s *BasePipeListener) ExitFuncParamType(ctx *FuncParamTypeContext) {}

// EnterFuncBody is called when production funcBody is entered.
func (s *BasePipeListener) EnterFuncBody(ctx *FuncBodyContext) {}

// ExitFuncBody is called when production funcBody is exited.
func (s *BasePipeListener) ExitFuncBody(ctx *FuncBodyContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BasePipeListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BasePipeListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterMultiPipe is called when production multiPipe is entered.
func (s *BasePipeListener) EnterMultiPipe(ctx *MultiPipeContext) {}

// ExitMultiPipe is called when production multiPipe is exited.
func (s *BasePipeListener) ExitMultiPipe(ctx *MultiPipeContext) {}

// EnterPipe is called when production pipe is entered.
func (s *BasePipeListener) EnterPipe(ctx *PipeContext) {}

// ExitPipe is called when production pipe is exited.
func (s *BasePipeListener) ExitPipe(ctx *PipeContext) {}

// EnterPipeNode is called when production pipeNode is entered.
func (s *BasePipeListener) EnterPipeNode(ctx *PipeNodeContext) {}

// ExitPipeNode is called when production pipeNode is exited.
func (s *BasePipeListener) ExitPipeNode(ctx *PipeNodeContext) {}

// EnterVariableNode is called when production variableNode is entered.
func (s *BasePipeListener) EnterVariableNode(ctx *VariableNodeContext) {}

// ExitVariableNode is called when production variableNode is exited.
func (s *BasePipeListener) ExitVariableNode(ctx *VariableNodeContext) {}

// EnterFunctionNode is called when production functionNode is entered.
func (s *BasePipeListener) EnterFunctionNode(ctx *FunctionNodeContext) {}

// ExitFunctionNode is called when production functionNode is exited.
func (s *BasePipeListener) ExitFunctionNode(ctx *FunctionNodeContext) {}

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

// EnterFunctionParameterValue is called when production functionParameterValue is entered.
func (s *BasePipeListener) EnterFunctionParameterValue(ctx *FunctionParameterValueContext) {}

// ExitFunctionParameterValue is called when production functionParameterValue is exited.
func (s *BasePipeListener) ExitFunctionParameterValue(ctx *FunctionParameterValueContext) {}

// EnterVariableValue is called when production variableValue is entered.
func (s *BasePipeListener) EnterVariableValue(ctx *VariableValueContext) {}

// ExitVariableValue is called when production variableValue is exited.
func (s *BasePipeListener) ExitVariableValue(ctx *VariableValueContext) {}

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
