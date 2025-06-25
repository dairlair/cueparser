package cueparser

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func mkToken(typ TokenType, val string) Token {
	return Token{
		Typ: typ,
		Val: val,
	}
}

var (
	tEOF   = mkToken(tokenTypeEof, "EOF")
	tSpace = mkToken(tokenTypeSpace, " ")
)

type tokTest struct {
	name   string
	input  string
	tokens []Token
}

var tokTests = []tokTest{
	//{"empty", "", []Token{tEOF}},
	{"one space", " ", []Token{tSpace, tEOF}},
}

func TestTokenizer(t *testing.T) {
	for _, test := range tokTests {
		tokens := Tokenize(test.input)

		if diff := cmp.Diff(test.tokens, tokens); diff != "" {
			t.Errorf("Tokens mismatch (-want +got):\n%s", diff)
		}
	}
}
