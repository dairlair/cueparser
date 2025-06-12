package cueparser

import (
	"fmt"
)

// TokenType is a top-level token classification: A word, space, comment, unknown.

type TokenType int

// Classes of lexic tokens
const (
	TokenTypeUnknown TokenType = iota
	TokenTypeWord
	TokenTypeComment
	// TokenTypeEndOfLine is being used as dedicated type because:
	// > Each line in the cue sheet file defines a command, such as FILE, TRACK, INDEX, TITLE, PERFORMER, etc.
	TokenTypeEndOfLine // The main
)

// Token is a (type, Value) pair representing a lexic token.
type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	switch {
	case t.Type == TokenTypeWord:
		return fmt.Sprintf("%.10q...", t.Value)
	}
	return fmt.Sprintf("%q", t.Value)
}
