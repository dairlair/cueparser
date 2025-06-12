package cueparser

type tokenizer struct {
}

type tokenizerStateFn func(*tokenizer) tokenizerStateFn

func SplitToTokens(s string) ([]Token, error) {
	return []Token{}, nil
}
