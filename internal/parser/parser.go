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
	if args == nil {
		return
	}
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

		lineTempo16Count := 0
		lastNote := &music.Note{Symbol: "", Tempo16Count: 0}
		for _, n := range noteLineContext.notesContext.GetChildren() {
			noteToken := n.(*NoteWithSignContext)
			tempo16Count := 16 // 默认为全音符
			multiple := 1      // 默认不重复
			if noteToken.noteSignContext != nil {
				signNum, _ := strconv.Atoi(noteToken.noteSignContext.GetText())
				if signNum > 0 {
					tempo16Count = 16 / signNum
				}
			}
			if noteToken.multipleContext != nil {
				multipleNum, _ := strconv.Atoi(noteToken.multipleContext.GetText())
				if multipleNum > 0 {
					multiple = multipleNum
				}
			}
			tokenSymbol := music.Symbol(noteToken.GetTokenContext().GetText())

			if tokenSymbol == music.ContinueSymbol {
				// 保持延音音符的引用
				lastNote.Tempo16Count += tempo16Count * multiple
				for t := 0; t < tempo16Count*multiple; t++ {
					pattern.TempoMap[t+lineTempo16Count] = append(pattern.TempoMap[t+lineTempo16Count], lastNote)
				}
				lineTempo16Count += tempo16Count * multiple
				continue
			} else if tokenSymbol == music.SilentSymbol {
				lineTempo16Count += tempo16Count * multiple
				continue
			}
			for m := 0; m < multiple; m++ {
				note := &music.Note{Symbol: tokenSymbol, Tempo16Count: tempo16Count}
				lastNote = note
				for t := 0; t < tempo16Count; t++ {
					pattern.TempoMap[t+lineTempo16Count] = append(pattern.TempoMap[t+lineTempo16Count], note)
				}
				lineTempo16Count += tempo16Count
			}
		}
		if pattern.MaxTempo16Count < lineTempo16Count {
			pattern.MaxTempo16Count = lineTempo16Count
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
