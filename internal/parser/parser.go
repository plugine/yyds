package parser

import (
	"strconv"
	"yyds/internal/music"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type listener struct {
	*BaseMusicListener
	m *music.Music
}

func (l *listener) EnterControlAssign(ctx *ControlAssignContext) {
	switch ctx.key.GetText() {
	case "tempo":
		l.m.Tempo, _ = strconv.Atoi(ctx.value.GetText())
	}
}

func (l *listener) parsePatternArgs(pattern *music.Pattern, args IArgsContext) {
	for _, child := range args.GetChildren() {
		switch ctx := child.(type) {
		case IArgsContext:
			l.parsePatternArgs(pattern, ctx)
		case IAssignContext:
			assign := ctx
			switch assign.GetKey().GetText() {
			case "repeat":
				pattern.Repeat, _ = strconv.Atoi(assign.GetValue().GetText())
			case "after_delay":
				pattern.AfterDelay, _ = strconv.Atoi(assign.GetValue().GetText())
			}
		}
	}
}

func (l *listener) EnterPattern(ctx *PatternContext) {
	pattern := &music.Pattern{TempoMap: make(map[int][]*music.Note)}

	l.parsePatternArgs(pattern, ctx.GetArgsContext())

	for _, c := range ctx.list.GetChildren() {
		noteLineContext := c.(*NoteLineContext)

		lineTempoCount := len(noteLineContext.notesContext.GetChildren())
		if pattern.MaxTempoCount < lineTempoCount {
			pattern.MaxTempoCount = lineTempoCount
		}
		lastNote := &music.Note{Symbol: "", TempoCount: 0}
		for j, n := range noteLineContext.notesContext.GetChildren() {
			noteToken := n.(*NoteTokenContext)
			tokenSymbol := music.Symbol(noteToken.GetText())
			if tokenSymbol == music.ContinueSymbol {
				// 保持延音音符的引用
				lastNote.TempoCount += 1
				pattern.TempoMap[j] = append(pattern.TempoMap[j], lastNote)
				continue
			}
			if tokenSymbol == music.SilentSymbol {
				continue
			}
			note := &music.Note{Symbol: tokenSymbol, TempoCount: 1}
			lastNote = note
			pattern.TempoMap[j] = append(pattern.TempoMap[j], note)
		}
	}
	l.m.Patterns = append(l.m.Patterns, *pattern)
}

func Parse(content string) *music.Music {
	is := antlr.NewInputStream(content)

	lexer := NewMusicLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listener := &listener{m: &music.Music{}}
	p := NewMusicParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Music())
	return listener.m
}
