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

func SplitToTokens(s string) ([]Token, error) {
	return []Token{}, nil
}

func tokenizerStateStart(t *Tokenizer) tokenizerStateFn {
	return nil
}

func Tokenize(input string) []Token {
	t := NewTokenizer(strings.NewReader(input))
	for state := tokenizerStateStart; state != nil; {
		state = state(t)
	}
	return t.tokens
}
