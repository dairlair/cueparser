package cueparser

import (
	"github.com/google/go-cmp/cmp"
	"io"
	"testing"
)

func mkToken(typ TokenType, val string) Token {
	return Token{
		Typ: typ,
		Val: val,
	}
}

var (
	tEOF = mkToken(tokenTypeEof, "")
)

type tokTest struct {
	name   string
	input  string
	tokens []Token
}

var tokTests = []tokTest{
	{"empty", "", []Token{tEOF}},
}

func TestTokenizer(t *testing.T) {
	for _, test := range tokTests {
		tokens, err := SplitToTokens(test.input)
		if err != nil && err != io.EOF {
			t.Error("unexpected error:", err)
		}

		if diff := cmp.Diff(test.tokens, tokens); diff != "" {
			t.Errorf("Tokens mismatch (-want +got):\n%s", diff)
		}
	}
}
