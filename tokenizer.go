package cueparser

import (
	"bufio"
	"io"
	"strings"
)

type Tokenizer struct {
	r          bufio.Reader
	classifier runeClassifier
	tokens     []Token
}

func NewTokenizer(r io.Reader) *Tokenizer {
	input := bufio.NewReader(r)
	classifier := newRuneClassifier()
	return &Tokenizer{
		r:          *input,
		classifier: classifier,
		tokens:     make([]Token, 0),
	}
}

type tokenizerStateFn func(*Tokenizer) tokenizerStateFn

func tokenizerStateStart(t *Tokenizer) tokenizerStateFn {
	nextRune, _, err := t.r.ReadRune()
	if err != nil {
		if err == io.EOF {
			t.tokens = append(t.tokens, Token{Typ: tokenTypeEof, Val: ""})
		}

	}
}

func Tokenize(input string) []Token {
	t := NewTokenizer(strings.NewReader(input))
	for state := tokenizerStateStart; state != nil; {
		state = state(t)
	}
	return t.tokens
}
