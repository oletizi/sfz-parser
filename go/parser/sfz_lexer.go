// Code generated from Sfz.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 73, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 6, 5, 29, 10, 5, 13, 5, 14, 5, 30, 3, 6, 3, 6, 3, 6, 5, 6,
	36, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 44, 10, 8, 12, 8,
	14, 8, 47, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	7, 9, 58, 10, 9, 12, 9, 14, 9, 61, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 7,
	10, 67, 10, 10, 12, 10, 14, 10, 70, 11, 10, 3, 10, 3, 10, 3, 45, 2, 11,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 5, 6,
	2, 11, 12, 15, 15, 34, 34, 62, 64, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11,
	34, 34, 2, 77, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2,
	2, 7, 25, 3, 2, 2, 2, 9, 28, 3, 2, 2, 2, 11, 35, 3, 2, 2, 2, 13, 37, 3,
	2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 53, 3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21,
	22, 7, 62, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7, 64, 2, 2, 24, 6, 3, 2, 2,
	2, 25, 26, 7, 63, 2, 2, 26, 8, 3, 2, 2, 2, 27, 29, 10, 2, 2, 2, 28, 27,
	3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2,
	31, 10, 3, 2, 2, 2, 32, 33, 7, 15, 2, 2, 33, 36, 7, 12, 2, 2, 34, 36, 9,
	3, 2, 2, 35, 32, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 12, 3, 2, 2, 2, 37,
	38, 9, 4, 2, 2, 38, 14, 3, 2, 2, 2, 39, 40, 7, 49, 2, 2, 40, 41, 7, 44,
	2, 2, 41, 45, 3, 2, 2, 2, 42, 44, 11, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44,
	47, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 48, 3, 2, 2,
	2, 47, 45, 3, 2, 2, 2, 48, 49, 7, 44, 2, 2, 49, 50, 7, 49, 2, 2, 50, 51,
	3, 2, 2, 2, 51, 52, 8, 8, 2, 2, 52, 16, 3, 2, 2, 2, 53, 54, 7, 49, 2, 2,
	54, 55, 7, 49, 2, 2, 55, 59, 3, 2, 2, 2, 56, 58, 10, 3, 2, 2, 57, 56, 3,
	2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60,
	62, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 8, 9, 2, 2, 63, 18, 3, 2, 2,
	2, 64, 68, 7, 37, 2, 2, 65, 67, 10, 3, 2, 2, 66, 65, 3, 2, 2, 2, 67, 70,
	3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 71, 3, 2, 2, 2,
	70, 68, 3, 2, 2, 2, 71, 72, 8, 10, 2, 2, 72, 20, 3, 2, 2, 2, 8, 2, 30,
	35, 45, 59, 68, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'<'", "'>'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "STRING", "NEWLINE", "WHITESPACE", "BLOCK_COMMENT", "LINE_COMMENT",
	"HASH_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "STRING", "NEWLINE", "WHITESPACE", "BLOCK_COMMENT",
	"LINE_COMMENT", "HASH_COMMENT",
}

type SfzLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSfzLexer(input antlr.CharStream) *SfzLexer {

	l := new(SfzLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Sfz.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SfzLexer tokens.
const (
	SfzLexerT__0          = 1
	SfzLexerT__1          = 2
	SfzLexerT__2          = 3
	SfzLexerSTRING        = 4
	SfzLexerNEWLINE       = 5
	SfzLexerWHITESPACE    = 6
	SfzLexerBLOCK_COMMENT = 7
	SfzLexerLINE_COMMENT  = 8
	SfzLexerHASH_COMMENT  = 9
)
