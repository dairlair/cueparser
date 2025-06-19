package cueparser

// TokenType is a top-level token classification: A KeyWord, space, comment, unknown.

type TokenType int

// Classes of lexic tokens
const (
	TokenTypeError TokenType = iota
	//// tokenTypeEol is being used as dedicated type because:
	//// > Each line in the cue sheet file defines a command, such as FILE, TRACK, INDEX, TITLE, PERFORMER, etc.
	//tokenTypeEol // The main
	tokenTypeEof // The main
)

// Token is a (type, Value) pair representing a lexic token.
type Token struct {
	Typ TokenType
	Val string
}
