package cueparser

// TokenType is a top-level token classification: A word, space, comment, unknown.

type TokenType int

// Classes of lexic tokens
const (
	UnknownToken TokenType = iota
	WordToken
	CommentToken
)

// Token is a (type, Value) pair representing a lexic token.
type Token struct {
	Type  TokenType
	Value string
}
