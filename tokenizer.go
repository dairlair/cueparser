package cueparser

import (
	"fmt"
	"unicode/utf8"
)

const eof = -1

type tokenizer struct {
	input  string
	pos    int
	line   int
	length int
	token  Token
}

func newTokenizer(input string) *tokenizer {
	return &tokenizer{
		input:  input,
		length: len([]rune(input)),
	}
}

type stateFn func(*tokenizer) stateFn

func tokenizerStart(t *tokenizer) stateFn {
	switch r := t.nextRune(); {
	case r == eof:
		return nil
	default:
		return t.errorf("unexpected rune %c", r)
	}
}

func (t *tokenizer) nextRune() rune {
	if t.pos >= len(t.input) {
		return eof
	}

	r, w := utf8.DecodeRuneInString(t.input[t.pos:])
	t.pos += w
	if r == '\n' {
		t.line++
	}
	return r
}

func (t *tokenizer) nextToken() Token {
	t.token = Token{tokenTypeEof, "EOF"}
	state := tokenizerStart
	for {
		state = state(t)
		if state == nil {
			return t.token
		}
	}
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating t.nextToken.
func (t *tokenizer) errorf(format string, args ...any) stateFn {
	t.token = Token{Typ: TokenTypeError, Val: fmt.Sprintf(format, args...)}
	return nil
}

func Tokenize(input string) (tokens []Token) {
	tr := newTokenizer(input)
	for {
		token := tr.nextToken()
		tokens = append(tokens, token)
		if token.Typ == tokenTypeEof || token.Typ == TokenTypeError {
			break
		}
	}
	return
}
