package cueparser

import (
	"fmt"
	"unicode/utf8"
)

const eof = -1

type tokenizer struct {
	input  string
	pos    int
	start  int
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
	case isSpace(r):
		return tokenizeSpace
	case r == eof:
		return nil
		// This default logic is not correct, fix it as well
	default:
		return t.errorf("unexpected rune %c", r)
	}
}

func tokenizeSpace(t *tokenizer) stateFn {
	switch r := t.nextRune(); {
	case isSpace(r):
		return tokenizeSpace
	case r == eof:
		return nil
	default:
		// With any another rune (non-space and non-eof) we emit the `tokenTypeSpace`
		return t.emit(tokenTypeSpace)
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

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating t.nextToken.
func (t *tokenizer) errorf(format string, args ...any) stateFn {
	t.token = Token{Typ: tokenTypeError, Val: fmt.Sprintf(format, args...)}
	return nil
}

// thisItem returns the item at the current input point with the specified type
// and advances the input.
func (t *tokenizer) thisItem(typ TokenType) Token {
	i := Token{typ, t.input[t.start:t.pos]}
	t.start = t.pos
	return i
}

// emit passes the trailing text as an item back to the parser.
func (t *tokenizer) emit(typ TokenType) stateFn {
	return t.emitItem(t.thisItem(typ))
}

// emitItem passes the specified item to the parser.
func (t *tokenizer) emitItem(tk Token) stateFn {
	t.token = tk
	return nil
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

func Tokenize(input string) (tokens []Token) {
	tr := newTokenizer(input)
	for {
		token := tr.nextToken()
		tokens = append(tokens, token)
		if token.Typ == tokenTypeEof || token.Typ == tokenTypeError {
			break
		}
	}
	return
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}
