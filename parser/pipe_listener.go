// Code generated from parser/Pipe.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Pipe

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PipeListener is a complete listener for a parse tree produced by PipeParser.
type PipeListener interface {
	antlr.ParseTreeListener

	// EnterCommands is called when entering the commands production.
	EnterCommands(c *CommandsContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterCommandName is called when entering the commandName production.
	EnterCommandName(c *CommandNameContext)

	// EnterCommandArguments is called when entering the commandArguments production.
	EnterCommandArguments(c *CommandArgumentsContext)

	// EnterCommandArgument is called when entering the commandArgument production.
	EnterCommandArgument(c *CommandArgumentContext)

	// EnterCommandArgumentLabel is called when entering the commandArgumentLabel production.
	EnterCommandArgumentLabel(c *CommandArgumentLabelContext)

	// EnterNumberValue is called when entering the numberValue production.
	EnterNumberValue(c *NumberValueContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// ExitCommands is called when exiting the commands production.
	ExitCommands(c *CommandsContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitCommandName is called when exiting the commandName production.
	ExitCommandName(c *CommandNameContext)

	// ExitCommandArguments is called when exiting the commandArguments production.
	ExitCommandArguments(c *CommandArgumentsContext)

	// ExitCommandArgument is called when exiting the commandArgument production.
	ExitCommandArgument(c *CommandArgumentContext)

	// ExitCommandArgumentLabel is called when exiting the commandArgumentLabel production.
	ExitCommandArgumentLabel(c *CommandArgumentLabelContext)

	// ExitNumberValue is called when exiting the numberValue production.
	ExitNumberValue(c *NumberValueContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)
}
