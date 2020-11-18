// Code generated from parser/Pipe.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Pipe

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PipeListener is a complete listener for a parse tree produced by PipeParser.
type PipeListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterFuncName is called when entering the funcName production.
	EnterFuncName(c *FuncNameContext)

	// EnterFuncParamsDef is called when entering the funcParamsDef production.
	EnterFuncParamsDef(c *FuncParamsDefContext)

	// EnterFuncParamDef is called when entering the funcParamDef production.
	EnterFuncParamDef(c *FuncParamDefContext)

	// EnterFuncParamName is called when entering the funcParamName production.
	EnterFuncParamName(c *FuncParamNameContext)

	// EnterOptionalParamFlag is called when entering the optionalParamFlag production.
	EnterOptionalParamFlag(c *OptionalParamFlagContext)

	// EnterFuncParamType is called when entering the funcParamType production.
	EnterFuncParamType(c *FuncParamTypeContext)

	// EnterFuncBody is called when entering the funcBody production.
	EnterFuncBody(c *FuncBodyContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterCmd is called when entering the cmd production.
	EnterCmd(c *CmdContext)

	// EnterMultiPipe is called when entering the multiPipe production.
	EnterMultiPipe(c *MultiPipeContext)

	// EnterPipe is called when entering the pipe production.
	EnterPipe(c *PipeContext)

	// EnterPipeNode is called when entering the pipeNode production.
	EnterPipeNode(c *PipeNodeContext)

	// EnterVariableNode is called when entering the variableNode production.
	EnterVariableNode(c *VariableNodeContext)

	// EnterStreamNode is called when entering the streamNode production.
	EnterStreamNode(c *StreamNodeContext)

	// EnterStreamSplitter is called when entering the streamSplitter production.
	EnterStreamSplitter(c *StreamSplitterContext)

	// EnterStreamCollector is called when entering the streamCollector production.
	EnterStreamCollector(c *StreamCollectorContext)

	// EnterPipeNodeArray is called when entering the pipeNodeArray production.
	EnterPipeNodeArray(c *PipeNodeArrayContext)

	// EnterPipeNodeElement is called when entering the pipeNodeElement production.
	EnterPipeNodeElement(c *PipeNodeElementContext)

	// EnterFunctionNode is called when entering the functionNode production.
	EnterFunctionNode(c *FunctionNodeContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterFunctionParameters is called when entering the functionParameters production.
	EnterFunctionParameters(c *FunctionParametersContext)

	// EnterFunctionParameter is called when entering the functionParameter production.
	EnterFunctionParameter(c *FunctionParameterContext)

	// EnterFunctionParameterLabel is called when entering the functionParameterLabel production.
	EnterFunctionParameterLabel(c *FunctionParameterLabelContext)

	// EnterFunctionParameterValue is called when entering the functionParameterValue production.
	EnterFunctionParameterValue(c *FunctionParameterValueContext)

	// EnterVariableValue is called when entering the variableValue production.
	EnterVariableValue(c *VariableValueContext)

	// EnterDictValue is called when entering the dictValue production.
	EnterDictValue(c *DictValueContext)

	// EnterDictEntries is called when entering the dictEntries production.
	EnterDictEntries(c *DictEntriesContext)

	// EnterDictEntry is called when entering the dictEntry production.
	EnterDictEntry(c *DictEntryContext)

	// EnterDictEntryLabel is called when entering the dictEntryLabel production.
	EnterDictEntryLabel(c *DictEntryLabelContext)

	// EnterDictEntryValue is called when entering the dictEntryValue production.
	EnterDictEntryValue(c *DictEntryValueContext)

	// EnterNumberValue is called when entering the numberValue production.
	EnterNumberValue(c *NumberValueContext)

	// EnterIntegerValue is called when entering the integerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterDecimalValue is called when entering the decimalValue production.
	EnterDecimalValue(c *DecimalValueContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitFuncName is called when exiting the funcName production.
	ExitFuncName(c *FuncNameContext)

	// ExitFuncParamsDef is called when exiting the funcParamsDef production.
	ExitFuncParamsDef(c *FuncParamsDefContext)

	// ExitFuncParamDef is called when exiting the funcParamDef production.
	ExitFuncParamDef(c *FuncParamDefContext)

	// ExitFuncParamName is called when exiting the funcParamName production.
	ExitFuncParamName(c *FuncParamNameContext)

	// ExitOptionalParamFlag is called when exiting the optionalParamFlag production.
	ExitOptionalParamFlag(c *OptionalParamFlagContext)

	// ExitFuncParamType is called when exiting the funcParamType production.
	ExitFuncParamType(c *FuncParamTypeContext)

	// ExitFuncBody is called when exiting the funcBody production.
	ExitFuncBody(c *FuncBodyContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitCmd is called when exiting the cmd production.
	ExitCmd(c *CmdContext)

	// ExitMultiPipe is called when exiting the multiPipe production.
	ExitMultiPipe(c *MultiPipeContext)

	// ExitPipe is called when exiting the pipe production.
	ExitPipe(c *PipeContext)

	// ExitPipeNode is called when exiting the pipeNode production.
	ExitPipeNode(c *PipeNodeContext)

	// ExitVariableNode is called when exiting the variableNode production.
	ExitVariableNode(c *VariableNodeContext)

	// ExitStreamNode is called when exiting the streamNode production.
	ExitStreamNode(c *StreamNodeContext)

	// ExitStreamSplitter is called when exiting the streamSplitter production.
	ExitStreamSplitter(c *StreamSplitterContext)

	// ExitStreamCollector is called when exiting the streamCollector production.
	ExitStreamCollector(c *StreamCollectorContext)

	// ExitPipeNodeArray is called when exiting the pipeNodeArray production.
	ExitPipeNodeArray(c *PipeNodeArrayContext)

	// ExitPipeNodeElement is called when exiting the pipeNodeElement production.
	ExitPipeNodeElement(c *PipeNodeElementContext)

	// ExitFunctionNode is called when exiting the functionNode production.
	ExitFunctionNode(c *FunctionNodeContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitFunctionParameters is called when exiting the functionParameters production.
	ExitFunctionParameters(c *FunctionParametersContext)

	// ExitFunctionParameter is called when exiting the functionParameter production.
	ExitFunctionParameter(c *FunctionParameterContext)

	// ExitFunctionParameterLabel is called when exiting the functionParameterLabel production.
	ExitFunctionParameterLabel(c *FunctionParameterLabelContext)

	// ExitFunctionParameterValue is called when exiting the functionParameterValue production.
	ExitFunctionParameterValue(c *FunctionParameterValueContext)

	// ExitVariableValue is called when exiting the variableValue production.
	ExitVariableValue(c *VariableValueContext)

	// ExitDictValue is called when exiting the dictValue production.
	ExitDictValue(c *DictValueContext)

	// ExitDictEntries is called when exiting the dictEntries production.
	ExitDictEntries(c *DictEntriesContext)

	// ExitDictEntry is called when exiting the dictEntry production.
	ExitDictEntry(c *DictEntryContext)

	// ExitDictEntryLabel is called when exiting the dictEntryLabel production.
	ExitDictEntryLabel(c *DictEntryLabelContext)

	// ExitDictEntryValue is called when exiting the dictEntryValue production.
	ExitDictEntryValue(c *DictEntryValueContext)

	// ExitNumberValue is called when exiting the numberValue production.
	ExitNumberValue(c *NumberValueContext)

	// ExitIntegerValue is called when exiting the integerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitDecimalValue is called when exiting the decimalValue production.
	ExitDecimalValue(c *DecimalValueContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)
}
