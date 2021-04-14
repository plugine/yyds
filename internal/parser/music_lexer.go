// Code generated from Music.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 15, 80, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 6, 8, 51, 10, 8, 13, 8, 14, 8, 52, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 13, 6, 13, 64, 10, 13, 13, 13, 14, 13, 65, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 74, 10, 14, 12, 14, 14, 14, 77,
	11, 14, 3, 14, 3, 14, 2, 2, 15, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 3, 2, 5, 7, 2, 45, 45,
	50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12,
	12, 15, 15, 2, 82, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2,
	2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 5, 31, 3,
	2, 2, 2, 7, 33, 3, 2, 2, 2, 9, 35, 3, 2, 2, 2, 11, 37, 3, 2, 2, 2, 13,
	45, 3, 2, 2, 2, 15, 50, 3, 2, 2, 2, 17, 54, 3, 2, 2, 2, 19, 56, 3, 2, 2,
	2, 21, 58, 3, 2, 2, 2, 23, 60, 3, 2, 2, 2, 25, 63, 3, 2, 2, 2, 27, 69,
	3, 2, 2, 2, 29, 30, 7, 63, 2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7, 94, 2, 2,
	32, 6, 3, 2, 2, 2, 33, 34, 7, 42, 2, 2, 34, 8, 3, 2, 2, 2, 35, 36, 7, 43,
	2, 2, 36, 10, 3, 2, 2, 2, 37, 38, 7, 114, 2, 2, 38, 39, 7, 99, 2, 2, 39,
	40, 7, 118, 2, 2, 40, 41, 7, 118, 2, 2, 41, 42, 7, 103, 2, 2, 42, 43, 7,
	116, 2, 2, 43, 44, 7, 112, 2, 2, 44, 12, 3, 2, 2, 2, 45, 46, 7, 117, 2,
	2, 46, 47, 7, 103, 2, 2, 47, 48, 7, 118, 2, 2, 48, 14, 3, 2, 2, 2, 49,
	51, 9, 2, 2, 2, 50, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 50, 3, 2, 2,
	2, 52, 53, 3, 2, 2, 2, 53, 16, 3, 2, 2, 2, 54, 55, 7, 125, 2, 2, 55, 18,
	3, 2, 2, 2, 56, 57, 7, 127, 2, 2, 57, 20, 3, 2, 2, 2, 58, 59, 7, 46, 2,
	2, 59, 22, 3, 2, 2, 2, 60, 61, 7, 60, 2, 2, 61, 24, 3, 2, 2, 2, 62, 64,
	9, 3, 2, 2, 63, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2,
	65, 66, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 8, 13, 2, 2, 68, 26, 3,
	2, 2, 2, 69, 70, 7, 49, 2, 2, 70, 71, 7, 49, 2, 2, 71, 75, 3, 2, 2, 2,
	72, 74, 10, 4, 2, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3,
	2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78,
	79, 8, 14, 3, 2, 79, 28, 3, 2, 2, 2, 6, 2, 52, 65, 75, 4, 8, 2, 2, 2, 3,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'\\'", "'('", "')'", "'pattern'", "'set'", "", "'{'", "'}'",
	"','", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "LEFT_RBRACKET", "RIGHT_RBRACKET", "PATTERN", "SET", "TOKEN",
	"L_CURLY", "R_CURLY", "COMMA", "COLON", "WHITESPACE", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "LEFT_RBRACKET", "RIGHT_RBRACKET", "PATTERN", "SET", "TOKEN",
	"L_CURLY", "R_CURLY", "COMMA", "COLON", "WHITESPACE", "LINE_COMMENT",
}

type MusicLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewMusicLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *MusicLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMusicLexer(input antlr.CharStream) *MusicLexer {
	l := new(MusicLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Music.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MusicLexer tokens.
const (
	MusicLexerT__0           = 1
	MusicLexerT__1           = 2
	MusicLexerLEFT_RBRACKET  = 3
	MusicLexerRIGHT_RBRACKET = 4
	MusicLexerPATTERN        = 5
	MusicLexerSET            = 6
	MusicLexerTOKEN          = 7
	MusicLexerL_CURLY        = 8
	MusicLexerR_CURLY        = 9
	MusicLexerCOMMA          = 10
	MusicLexerCOLON          = 11
	MusicLexerWHITESPACE     = 12
	MusicLexerLINE_COMMENT   = 13
)
