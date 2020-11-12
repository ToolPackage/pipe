// Code generated from parser/Pipe.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Pipe

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PipeListener is a complete listener for a parse tree produced by PipeParser.
type PipeListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterFunctionParameters is called when entering the functionParameters production.
	EnterFunctionParameters(c *FunctionParametersContext)

	// EnterFunctionParameter is called when entering the functionParameter production.
	EnterFunctionParameter(c *FunctionParameterContext)

	// EnterFunctionParameterLabel is called when entering the functionParameterLabel production.
	EnterFunctionParameterLabel(c *FunctionParameterLabelContext)

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

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitFunctionParameters is called when exiting the functionParameters production.
	ExitFunctionParameters(c *FunctionParametersContext)

	// ExitFunctionParameter is called when exiting the functionParameter production.
	ExitFunctionParameter(c *FunctionParameterContext)

	// ExitFunctionParameterLabel is called when exiting the functionParameterLabel production.
	ExitFunctionParameterLabel(c *FunctionParameterLabelContext)

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
