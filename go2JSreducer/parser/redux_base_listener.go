// Generated from Redux.g4 by ANTLR 4.7.

package parser // Redux

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseReduxListener is a complete listener for a parse tree produced by ReduxParser.
type BaseReduxListener struct{}

var _ ReduxListener = &BaseReduxListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseReduxListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseReduxListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseReduxListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseReduxListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseReduxListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseReduxListener) ExitProg(ctx *ProgContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseReduxListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseReduxListener) ExitStat(ctx *StatContext) {}

// EnterReducer is called when production reducer is entered.
func (s *BaseReduxListener) EnterReducer(ctx *ReducerContext) {}

// ExitReducer is called when production reducer is exited.
func (s *BaseReduxListener) ExitReducer(ctx *ReducerContext) {}