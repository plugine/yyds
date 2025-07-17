// Code generated from Music.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Music

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMusicListener is a complete listener for a parse tree produced by MusicParser.
type BaseMusicListener struct{}

var _ MusicListener = &BaseMusicListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMusicListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMusicListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMusicListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMusicListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMusic is called when production music is entered.
func (s *BaseMusicListener) EnterMusic(ctx *MusicContext) {}

// ExitMusic is called when production music is exited.
func (s *BaseMusicListener) ExitMusic(ctx *MusicContext) {}

// EnterContent is called when production content is entered.
func (s *BaseMusicListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseMusicListener) ExitContent(ctx *ContentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMusicListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMusicListener) ExitExpression(ctx *ExpressionContext) {}

// EnterControlAssign is called when production controlAssign is entered.
func (s *BaseMusicListener) EnterControlAssign(ctx *ControlAssignContext) {}

// ExitControlAssign is called when production controlAssign is exited.
func (s *BaseMusicListener) ExitControlAssign(ctx *ControlAssignContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseMusicListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseMusicListener) ExitPattern(ctx *PatternContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseMusicListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseMusicListener) ExitArgs(ctx *ArgsContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseMusicListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseMusicListener) ExitAssign(ctx *AssignContext) {}

// EnterNoteList is called when production noteList is entered.
func (s *BaseMusicListener) EnterNoteList(ctx *NoteListContext) {}

// ExitNoteList is called when production noteList is exited.
func (s *BaseMusicListener) ExitNoteList(ctx *NoteListContext) {}

// EnterNoteLine is called when production noteLine is entered.
func (s *BaseMusicListener) EnterNoteLine(ctx *NoteLineContext) {}

// ExitNoteLine is called when production noteLine is exited.
func (s *BaseMusicListener) ExitNoteLine(ctx *NoteLineContext) {}

// EnterNotes is called when production notes is entered.
func (s *BaseMusicListener) EnterNotes(ctx *NotesContext) {}

// ExitNotes is called when production notes is exited.
func (s *BaseMusicListener) ExitNotes(ctx *NotesContext) {}

// EnterNoteWithSign is called when production noteWithSign is entered.
func (s *BaseMusicListener) EnterNoteWithSign(ctx *NoteWithSignContext) {}

// ExitNoteWithSign is called when production noteWithSign is exited.
func (s *BaseMusicListener) ExitNoteWithSign(ctx *NoteWithSignContext) {}
