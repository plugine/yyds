// Code generated from Music.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Music

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 107,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 33, 10, 3, 12, 3, 14, 3,
	36, 11, 3, 3, 4, 3, 4, 5, 4, 40, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 5, 6, 60, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 68, 10, 7,
	12, 7, 14, 7, 71, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9, 78, 10, 9,
	13, 9, 14, 9, 79, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 6, 11, 87, 10, 11,
	13, 11, 14, 11, 88, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 105, 10, 12, 3, 12,
	2, 4, 4, 12, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 2, 2, 104,
	2, 24, 3, 2, 2, 2, 4, 27, 3, 2, 2, 2, 6, 39, 3, 2, 2, 2, 8, 41, 3, 2, 2,
	2, 10, 59, 3, 2, 2, 2, 12, 61, 3, 2, 2, 2, 14, 72, 3, 2, 2, 2, 16, 77,
	3, 2, 2, 2, 18, 81, 3, 2, 2, 2, 20, 86, 3, 2, 2, 2, 22, 104, 3, 2, 2, 2,
	24, 25, 5, 4, 3, 2, 25, 26, 7, 2, 2, 3, 26, 3, 3, 2, 2, 2, 27, 28, 8, 3,
	1, 2, 28, 29, 5, 6, 4, 2, 29, 34, 3, 2, 2, 2, 30, 31, 12, 3, 2, 2, 31,
	33, 5, 6, 4, 2, 32, 30, 3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2,
	2, 34, 35, 3, 2, 2, 2, 35, 5, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 40, 5,
	8, 5, 2, 38, 40, 5, 10, 6, 2, 39, 37, 3, 2, 2, 2, 39, 38, 3, 2, 2, 2, 40,
	7, 3, 2, 2, 2, 41, 42, 7, 8, 2, 2, 42, 43, 7, 9, 2, 2, 43, 44, 7, 3, 2,
	2, 44, 45, 7, 9, 2, 2, 45, 9, 3, 2, 2, 2, 46, 47, 7, 7, 2, 2, 47, 48, 7,
	10, 2, 2, 48, 49, 5, 16, 9, 2, 49, 50, 7, 11, 2, 2, 50, 60, 3, 2, 2, 2,
	51, 52, 7, 7, 2, 2, 52, 53, 7, 4, 2, 2, 53, 54, 5, 12, 7, 2, 54, 55, 7,
	5, 2, 2, 55, 56, 7, 10, 2, 2, 56, 57, 5, 16, 9, 2, 57, 58, 7, 11, 2, 2,
	58, 60, 3, 2, 2, 2, 59, 46, 3, 2, 2, 2, 59, 51, 3, 2, 2, 2, 60, 11, 3,
	2, 2, 2, 61, 62, 8, 7, 1, 2, 62, 63, 5, 14, 8, 2, 63, 69, 3, 2, 2, 2, 64,
	65, 12, 3, 2, 2, 65, 66, 7, 12, 2, 2, 66, 68, 5, 14, 8, 2, 67, 64, 3, 2,
	2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 13,
	3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 7, 9, 2, 2, 73, 74, 7, 3, 2, 2,
	74, 75, 7, 9, 2, 2, 75, 15, 3, 2, 2, 2, 76, 78, 5, 18, 10, 2, 77, 76, 3,
	2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80,
	17, 3, 2, 2, 2, 81, 82, 7, 9, 2, 2, 82, 83, 7, 13, 2, 2, 83, 84, 5, 20,
	11, 2, 84, 19, 3, 2, 2, 2, 85, 87, 5, 22, 12, 2, 86, 85, 3, 2, 2, 2, 87,
	88, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 21, 3, 2, 2,
	2, 90, 105, 7, 9, 2, 2, 91, 92, 7, 9, 2, 2, 92, 93, 7, 6, 2, 2, 93, 105,
	7, 9, 2, 2, 94, 95, 7, 9, 2, 2, 95, 96, 7, 10, 2, 2, 96, 97, 7, 9, 2, 2,
	97, 105, 7, 11, 2, 2, 98, 99, 7, 9, 2, 2, 99, 100, 7, 10, 2, 2, 100, 101,
	7, 9, 2, 2, 101, 102, 7, 11, 2, 2, 102, 103, 7, 6, 2, 2, 103, 105, 7, 9,
	2, 2, 104, 90, 3, 2, 2, 2, 104, 91, 3, 2, 2, 2, 104, 94, 3, 2, 2, 2, 104,
	98, 3, 2, 2, 2, 105, 23, 3, 2, 2, 2, 9, 34, 39, 59, 69, 79, 88, 104,
}
var literalNames = []string{
	"", "'='", "'('", "')'", "'*'", "'pattern'", "'set'", "", "'{'", "'}'",
	"','", "':'",
}
var symbolicNames = []string{
	"", "", "LEFT_RBRACKET", "RIGHT_RBRACKET", "MULTIPLE", "PATTERN", "SET",
	"TOKEN", "L_CURLY", "R_CURLY", "COMMA", "COLON", "WHITESPACE", "LINE_COMMENT",
}

var ruleNames = []string{
	"music", "content", "expression", "controlAssign", "pattern", "args", "assign",
	"noteList", "noteLine", "notes", "noteWithSign",
}

type MusicParser struct {
	*antlr.BaseParser
}

// NewMusicParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *MusicParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMusicParser(input antlr.TokenStream) *MusicParser {
	this := new(MusicParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Music.g4"

	return this
}

// MusicParser tokens.
const (
	MusicParserEOF            = antlr.TokenEOF
	MusicParserT__0           = 1
	MusicParserLEFT_RBRACKET  = 2
	MusicParserRIGHT_RBRACKET = 3
	MusicParserMULTIPLE       = 4
	MusicParserPATTERN        = 5
	MusicParserSET            = 6
	MusicParserTOKEN          = 7
	MusicParserL_CURLY        = 8
	MusicParserR_CURLY        = 9
	MusicParserCOMMA          = 10
	MusicParserCOLON          = 11
	MusicParserWHITESPACE     = 12
	MusicParserLINE_COMMENT   = 13
)

// MusicParser rules.
const (
	MusicParserRULE_music         = 0
	MusicParserRULE_content       = 1
	MusicParserRULE_expression    = 2
	MusicParserRULE_controlAssign = 3
	MusicParserRULE_pattern       = 4
	MusicParserRULE_args          = 5
	MusicParserRULE_assign        = 6
	MusicParserRULE_noteList      = 7
	MusicParserRULE_noteLine      = 8
	MusicParserRULE_notes         = 9
	MusicParserRULE_noteWithSign  = 10
)

// IMusicContext is an interface to support dynamic dispatch.
type IMusicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMusicContext differentiates from other interfaces.
	IsMusicContext()
}

type MusicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMusicContext() *MusicContext {
	var p = new(MusicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_music
	return p
}

func (*MusicContext) IsMusicContext() {}

func NewMusicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MusicContext {
	var p = new(MusicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_music

	return p
}

func (s *MusicContext) GetParser() antlr.Parser { return s.parser }

func (s *MusicContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *MusicContext) EOF() antlr.TerminalNode {
	return s.GetToken(MusicParserEOF, 0)
}

func (s *MusicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MusicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MusicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterMusic(s)
	}
}

func (s *MusicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitMusic(s)
	}
}

func (p *MusicParser) Music() (localctx IMusicContext) {
	localctx = NewMusicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MusicParserRULE_music)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.content(0)
	}
	{
		p.SetState(23)
		p.Match(MusicParserEOF)
	}

	return localctx
}

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_content
	return p
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ContentContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitContent(s)
	}
}

func (p *MusicParser) Content() (localctx IContentContext) {
	return p.content(0)
}

func (p *MusicParser) content(_p int) (localctx IContentContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewContentContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IContentContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, MusicParserRULE_content, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewContentContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, MusicParserRULE_content)
			p.SetState(28)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(29)
				p.Expression()
			}

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ControlAssign() IControlAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControlAssignContext)
}

func (s *ExpressionContext) Pattern() IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MusicParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MusicParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MusicParserSET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.ControlAssign()
		}

	case MusicParserPATTERN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.Pattern()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IControlAssignContext is an interface to support dynamic dispatch.
type IControlAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsControlAssignContext differentiates from other interfaces.
	IsControlAssignContext()
}

type ControlAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  antlr.Token
}

func NewEmptyControlAssignContext() *ControlAssignContext {
	var p = new(ControlAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_controlAssign
	return p
}

func (*ControlAssignContext) IsControlAssignContext() {}

func NewControlAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlAssignContext {
	var p = new(ControlAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_controlAssign

	return p
}

func (s *ControlAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlAssignContext) GetKey() antlr.Token { return s.key }

func (s *ControlAssignContext) GetValue() antlr.Token { return s.value }

func (s *ControlAssignContext) SetKey(v antlr.Token) { s.key = v }

func (s *ControlAssignContext) SetValue(v antlr.Token) { s.value = v }

func (s *ControlAssignContext) SET() antlr.TerminalNode {
	return s.GetToken(MusicParserSET, 0)
}

func (s *ControlAssignContext) AllTOKEN() []antlr.TerminalNode {
	return s.GetTokens(MusicParserTOKEN)
}

func (s *ControlAssignContext) TOKEN(i int) antlr.TerminalNode {
	return s.GetToken(MusicParserTOKEN, i)
}

func (s *ControlAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterControlAssign(s)
	}
}

func (s *ControlAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitControlAssign(s)
	}
}

func (p *MusicParser) ControlAssign() (localctx IControlAssignContext) {
	localctx = NewControlAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MusicParserRULE_controlAssign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(MusicParserSET)
	}
	{
		p.SetState(40)

		var _m = p.Match(MusicParserTOKEN)

		localctx.(*ControlAssignContext).key = _m
	}
	{
		p.SetState(41)
		p.Match(MusicParserT__0)
	}
	{
		p.SetState(42)

		var _m = p.Match(MusicParserTOKEN)

		localctx.(*ControlAssignContext).value = _m
	}

	return localctx
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() INoteListContext

	// GetArgsContext returns the argsContext rule contexts.
	GetArgsContext() IArgsContext

	// SetList sets the list rule contexts.
	SetList(INoteListContext)

	// SetArgsContext sets the argsContext rule contexts.
	SetArgsContext(IArgsContext)

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	list        INoteListContext
	argsContext IArgsContext
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_pattern
	return p
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) GetList() INoteListContext { return s.list }

func (s *PatternContext) GetArgsContext() IArgsContext { return s.argsContext }

func (s *PatternContext) SetList(v INoteListContext) { s.list = v }

func (s *PatternContext) SetArgsContext(v IArgsContext) { s.argsContext = v }

func (s *PatternContext) PATTERN() antlr.TerminalNode {
	return s.GetToken(MusicParserPATTERN, 0)
}

func (s *PatternContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(MusicParserL_CURLY, 0)
}

func (s *PatternContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(MusicParserR_CURLY, 0)
}

func (s *PatternContext) NoteList() INoteListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoteListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoteListContext)
}

func (s *PatternContext) LEFT_RBRACKET() antlr.TerminalNode {
	return s.GetToken(MusicParserLEFT_RBRACKET, 0)
}

func (s *PatternContext) RIGHT_RBRACKET() antlr.TerminalNode {
	return s.GetToken(MusicParserRIGHT_RBRACKET, 0)
}

func (s *PatternContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (p *MusicParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MusicParserRULE_pattern)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Match(MusicParserPATTERN)
		}
		{
			p.SetState(45)
			p.Match(MusicParserL_CURLY)
		}
		{
			p.SetState(46)

			var _x = p.NoteList()

			localctx.(*PatternContext).list = _x
		}
		{
			p.SetState(47)
			p.Match(MusicParserR_CURLY)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(MusicParserPATTERN)
		}
		{
			p.SetState(50)
			p.Match(MusicParserLEFT_RBRACKET)
		}
		{
			p.SetState(51)

			var _x = p.args(0)

			localctx.(*PatternContext).argsContext = _x
		}
		{
			p.SetState(52)
			p.Match(MusicParserRIGHT_RBRACKET)
		}
		{
			p.SetState(53)
			p.Match(MusicParserL_CURLY)
		}
		{
			p.SetState(54)

			var _x = p.NoteList()

			localctx.(*PatternContext).list = _x
		}
		{
			p.SetState(55)
			p.Match(MusicParserR_CURLY)
		}

	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArgsContext returns the argsContext rule contexts.
	GetArgsContext() IArgsContext

	// GetAssignContext returns the assignContext rule contexts.
	GetAssignContext() IAssignContext

	// SetArgsContext sets the argsContext rule contexts.
	SetArgsContext(IArgsContext)

	// SetAssignContext sets the assignContext rule contexts.
	SetAssignContext(IAssignContext)

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	argsContext   IArgsContext
	assignContext IAssignContext
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) GetArgsContext() IArgsContext { return s.argsContext }

func (s *ArgsContext) GetAssignContext() IAssignContext { return s.assignContext }

func (s *ArgsContext) SetArgsContext(v IArgsContext) { s.argsContext = v }

func (s *ArgsContext) SetAssignContext(v IAssignContext) { s.assignContext = v }

func (s *ArgsContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ArgsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MusicParserCOMMA, 0)
}

func (s *ArgsContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *MusicParser) Args() (localctx IArgsContext) {
	return p.args(0)
}

func (p *MusicParser) args(_p int) (localctx IArgsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArgsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArgsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, MusicParserRULE_args, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)

		var _x = p.Assign()

		localctx.(*ArgsContext).assignContext = _x
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArgsContext(p, _parentctx, _parentState)
			localctx.(*ArgsContext).argsContext = _prevctx
			p.PushNewRecursionContext(localctx, _startState, MusicParserRULE_args)
			p.SetState(62)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(63)
				p.Match(MusicParserCOMMA)
			}
			{
				p.SetState(64)

				var _x = p.Assign()

				localctx.(*ArgsContext).assignContext = _x
			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  antlr.Token
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) GetKey() antlr.Token { return s.key }

func (s *AssignContext) GetValue() antlr.Token { return s.value }

func (s *AssignContext) SetKey(v antlr.Token) { s.key = v }

func (s *AssignContext) SetValue(v antlr.Token) { s.value = v }

func (s *AssignContext) AllTOKEN() []antlr.TerminalNode {
	return s.GetTokens(MusicParserTOKEN)
}

func (s *AssignContext) TOKEN(i int) antlr.TerminalNode {
	return s.GetToken(MusicParserTOKEN, i)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *MusicParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MusicParserRULE_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)

		var _m = p.Match(MusicParserTOKEN)

		localctx.(*AssignContext).key = _m
	}
	{
		p.SetState(71)
		p.Match(MusicParserT__0)
	}
	{
		p.SetState(72)

		var _m = p.Match(MusicParserTOKEN)

		localctx.(*AssignContext).value = _m
	}

	return localctx
}

// INoteListContext is an interface to support dynamic dispatch.
type INoteListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoteListContext differentiates from other interfaces.
	IsNoteListContext()
}

type NoteListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoteListContext() *NoteListContext {
	var p = new(NoteListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_noteList
	return p
}

func (*NoteListContext) IsNoteListContext() {}

func NewNoteListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteListContext {
	var p = new(NoteListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_noteList

	return p
}

func (s *NoteListContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteListContext) AllNoteLine() []INoteLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INoteLineContext)(nil)).Elem())
	var tst = make([]INoteLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INoteLineContext)
		}
	}

	return tst
}

func (s *NoteListContext) NoteLine(i int) INoteLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoteLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INoteLineContext)
}

func (s *NoteListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterNoteList(s)
	}
}

func (s *NoteListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitNoteList(s)
	}
}

func (p *MusicParser) NoteList() (localctx INoteListContext) {
	localctx = NewNoteListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MusicParserRULE_noteList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == MusicParserTOKEN {
		{
			p.SetState(74)
			p.NoteLine()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INoteLineContext is an interface to support dynamic dispatch.
type INoteLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSignContext returns the signContext token.
	GetSignContext() antlr.Token

	// SetSignContext sets the signContext token.
	SetSignContext(antlr.Token)

	// GetNotesContext returns the notesContext rule contexts.
	GetNotesContext() INotesContext

	// SetNotesContext sets the notesContext rule contexts.
	SetNotesContext(INotesContext)

	// IsNoteLineContext differentiates from other interfaces.
	IsNoteLineContext()
}

type NoteLineContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	signContext  antlr.Token
	notesContext INotesContext
}

func NewEmptyNoteLineContext() *NoteLineContext {
	var p = new(NoteLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_noteLine
	return p
}

func (*NoteLineContext) IsNoteLineContext() {}

func NewNoteLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteLineContext {
	var p = new(NoteLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_noteLine

	return p
}

func (s *NoteLineContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteLineContext) GetSignContext() antlr.Token { return s.signContext }

func (s *NoteLineContext) SetSignContext(v antlr.Token) { s.signContext = v }

func (s *NoteLineContext) GetNotesContext() INotesContext { return s.notesContext }

func (s *NoteLineContext) SetNotesContext(v INotesContext) { s.notesContext = v }

func (s *NoteLineContext) COLON() antlr.TerminalNode {
	return s.GetToken(MusicParserCOLON, 0)
}

func (s *NoteLineContext) TOKEN() antlr.TerminalNode {
	return s.GetToken(MusicParserTOKEN, 0)
}

func (s *NoteLineContext) Notes() INotesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotesContext)
}

func (s *NoteLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterNoteLine(s)
	}
}

func (s *NoteLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitNoteLine(s)
	}
}

func (p *MusicParser) NoteLine() (localctx INoteLineContext) {
	localctx = NewNoteLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MusicParserRULE_noteLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)

		var _m = p.Match(MusicParserTOKEN)

		localctx.(*NoteLineContext).signContext = _m
	}
	{
		p.SetState(80)
		p.Match(MusicParserCOLON)
	}
	{
		p.SetState(81)

		var _x = p.Notes()

		localctx.(*NoteLineContext).notesContext = _x
	}

	return localctx
}

// INotesContext is an interface to support dynamic dispatch.
type INotesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotesContext differentiates from other interfaces.
	IsNotesContext()
}

type NotesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotesContext() *NotesContext {
	var p = new(NotesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_notes
	return p
}

func (*NotesContext) IsNotesContext() {}

func NewNotesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotesContext {
	var p = new(NotesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_notes

	return p
}

func (s *NotesContext) GetParser() antlr.Parser { return s.parser }

func (s *NotesContext) AllNoteWithSign() []INoteWithSignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INoteWithSignContext)(nil)).Elem())
	var tst = make([]INoteWithSignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INoteWithSignContext)
		}
	}

	return tst
}

func (s *NotesContext) NoteWithSign(i int) INoteWithSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoteWithSignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INoteWithSignContext)
}

func (s *NotesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterNotes(s)
	}
}

func (s *NotesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitNotes(s)
	}
}

func (p *MusicParser) Notes() (localctx INotesContext) {
	localctx = NewNotesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MusicParserRULE_notes)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(83)
				p.NoteWithSign()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// INoteWithSignContext is an interface to support dynamic dispatch.
type INoteWithSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTokenContext returns the tokenContext token.
	GetTokenContext() antlr.Token

	// GetMultipleContext returns the multipleContext token.
	GetMultipleContext() antlr.Token

	// GetNoteSignContext returns the noteSignContext token.
	GetNoteSignContext() antlr.Token

	// SetTokenContext sets the tokenContext token.
	SetTokenContext(antlr.Token)

	// SetMultipleContext sets the multipleContext token.
	SetMultipleContext(antlr.Token)

	// SetNoteSignContext sets the noteSignContext token.
	SetNoteSignContext(antlr.Token)

	// IsNoteWithSignContext differentiates from other interfaces.
	IsNoteWithSignContext()
}

type NoteWithSignContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	tokenContext    antlr.Token
	multipleContext antlr.Token
	noteSignContext antlr.Token
}

func NewEmptyNoteWithSignContext() *NoteWithSignContext {
	var p = new(NoteWithSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MusicParserRULE_noteWithSign
	return p
}

func (*NoteWithSignContext) IsNoteWithSignContext() {}

func NewNoteWithSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteWithSignContext {
	var p = new(NoteWithSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MusicParserRULE_noteWithSign

	return p
}

func (s *NoteWithSignContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteWithSignContext) GetTokenContext() antlr.Token { return s.tokenContext }

func (s *NoteWithSignContext) GetMultipleContext() antlr.Token { return s.multipleContext }

func (s *NoteWithSignContext) GetNoteSignContext() antlr.Token { return s.noteSignContext }

func (s *NoteWithSignContext) SetTokenContext(v antlr.Token) { s.tokenContext = v }

func (s *NoteWithSignContext) SetMultipleContext(v antlr.Token) { s.multipleContext = v }

func (s *NoteWithSignContext) SetNoteSignContext(v antlr.Token) { s.noteSignContext = v }

func (s *NoteWithSignContext) AllTOKEN() []antlr.TerminalNode {
	return s.GetTokens(MusicParserTOKEN)
}

func (s *NoteWithSignContext) TOKEN(i int) antlr.TerminalNode {
	return s.GetToken(MusicParserTOKEN, i)
}

func (s *NoteWithSignContext) MULTIPLE() antlr.TerminalNode {
	return s.GetToken(MusicParserMULTIPLE, 0)
}

func (s *NoteWithSignContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(MusicParserL_CURLY, 0)
}

func (s *NoteWithSignContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(MusicParserR_CURLY, 0)
}

func (s *NoteWithSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteWithSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteWithSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.EnterNoteWithSign(s)
	}
}

func (s *NoteWithSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MusicListener); ok {
		listenerT.ExitNoteWithSign(s)
	}
}

func (p *MusicParser) NoteWithSign() (localctx INoteWithSignContext) {
	localctx = NewNoteWithSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MusicParserRULE_noteWithSign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)

			var _m = p.Match(MusicParserTOKEN)

			localctx.(*NoteWithSignContext).tokenContext = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)

			var _m = p.Match(MusicParserTOKEN)

			localctx.(*NoteWithSignContext).tokenContext = _m
		}
		{
			p.SetState(90)
			p.Match(MusicParserMULTIPLE)
		}
		{
			p.SetState(91)

			var _m = p.Match(MusicParserTOKEN)

			localctx.(*NoteWithSignContext).multipleContext = _m
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)

			var _m = p.Match(MusicParserTOKEN)

			localctx.(*NoteWithSignContext).tokenContext = _m
		}
		{
			p.SetState(93)
			p.Match(MusicParserL_CURLY)
		}
		{
			p.SetState(94)

			var _m = p.Match(MusicParserTOKEN)

			localctx.(*NoteWithSignContext).noteSignContext = _m
		}
		{
			p.SetState(95)
			p.Match(MusicParserR_CURLY)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)

			var _m = p.Match(MusicParserTOKEN)

			localctx.(*NoteWithSignContext).tokenContext = _m
		}
		{
			p.SetState(97)
			p.Match(MusicParserL_CURLY)
		}
		{
			p.SetState(98)

			var _m = p.Match(MusicParserTOKEN)

			localctx.(*NoteWithSignContext).noteSignContext = _m
		}
		{
			p.SetState(99)
			p.Match(MusicParserR_CURLY)
		}
		{
			p.SetState(100)
			p.Match(MusicParserMULTIPLE)
		}
		{
			p.SetState(101)

			var _m = p.Match(MusicParserTOKEN)

			localctx.(*NoteWithSignContext).multipleContext = _m
		}

	}

	return localctx
}

func (p *MusicParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ContentContext = nil
		if localctx != nil {
			t = localctx.(*ContentContext)
		}
		return p.Content_Sempred(t, predIndex)

	case 5:
		var t *ArgsContext = nil
		if localctx != nil {
			t = localctx.(*ArgsContext)
		}
		return p.Args_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MusicParser) Content_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MusicParser) Args_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
