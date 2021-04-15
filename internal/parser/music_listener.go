// Code generated from Music.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Music

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MusicListener is a complete listener for a parse tree produced by MusicParser.
type MusicListener interface {
	antlr.ParseTreeListener

	// EnterMusic is called when entering the music production.
	EnterMusic(c *MusicContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterControlAssign is called when entering the controlAssign production.
	EnterControlAssign(c *ControlAssignContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterNoteList is called when entering the noteList production.
	EnterNoteList(c *NoteListContext)

	// EnterNoteLine is called when entering the noteLine production.
	EnterNoteLine(c *NoteLineContext)

	// EnterNotes is called when entering the notes production.
	EnterNotes(c *NotesContext)

	// EnterNoteWithSign is called when entering the noteWithSign production.
	EnterNoteWithSign(c *NoteWithSignContext)

	// ExitMusic is called when exiting the music production.
	ExitMusic(c *MusicContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitControlAssign is called when exiting the controlAssign production.
	ExitControlAssign(c *ControlAssignContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitNoteList is called when exiting the noteList production.
	ExitNoteList(c *NoteListContext)

	// ExitNoteLine is called when exiting the noteLine production.
	ExitNoteLine(c *NoteLineContext)

	// ExitNotes is called when exiting the notes production.
	ExitNotes(c *NotesContext)

	// ExitNoteWithSign is called when exiting the noteWithSign production.
	ExitNoteWithSign(c *NoteWithSignContext)
}
