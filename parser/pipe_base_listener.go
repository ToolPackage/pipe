// Code generated from parser\Pipe.g4 by ANTLR 4.8. DO NOT EDIT.

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

// EnterCommands is called when production functions is entered.
func (s *BasePipeListener) EnterCommands(ctx *CommandsContext) {}

// ExitCommands is called when production functions is exited.
func (s *BasePipeListener) ExitCommands(ctx *CommandsContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BasePipeListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BasePipeListener) ExitPattern(ctx *PatternContext) {}

// EnterCommandName is called when production commandName is entered.
func (s *BasePipeListener) EnterCommandName(ctx *CommandNameContext) {}

// ExitCommandName is called when production commandName is exited.
func (s *BasePipeListener) ExitCommandName(ctx *CommandNameContext) {}

// EnterCommandArguments is called when production commandArguments is entered.
func (s *BasePipeListener) EnterCommandArguments(ctx *CommandArgumentsContext) {}

// ExitCommandArguments is called when production commandArguments is exited.
func (s *BasePipeListener) ExitCommandArguments(ctx *CommandArgumentsContext) {}

// EnterCommandArgument is called when production commandArgument is entered.
func (s *BasePipeListener) EnterCommandArgument(ctx *CommandArgumentContext) {}

// ExitCommandArgument is called when production commandArgument is exited.
func (s *BasePipeListener) ExitCommandArgument(ctx *CommandArgumentContext) {}

// EnterCommandArgumentLabel is called when production commandArgumentLabel is entered.
func (s *BasePipeListener) EnterCommandArgumentLabel(ctx *CommandArgumentLabelContext) {}

// ExitCommandArgumentLabel is called when production commandArgumentLabel is exited.
func (s *BasePipeListener) ExitCommandArgumentLabel(ctx *CommandArgumentLabelContext) {}

// EnterNumberValue is called when production numberValue is entered.
func (s *BasePipeListener) EnterNumberValue(ctx *NumberValueContext) {}

// ExitNumberValue is called when production numberValue is exited.
func (s *BasePipeListener) ExitNumberValue(ctx *NumberValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BasePipeListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BasePipeListener) ExitStringValue(ctx *StringValueContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BasePipeListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BasePipeListener) ExitBooleanValue(ctx *BooleanValueContext) {}
