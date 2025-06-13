package cueparser

// TokenType is a top-level token classification: A word, space, comment, unknown.

type TokenType int

// Classes of lexic tokens
const (
	TokenTypeUnknown TokenType = iota
	TokenTypeWord
	TokenTypeComment
	// tokenTypeEof is being used as dedicated type because:
	// > Each line in the cue sheet file defines a command, such as FILE, TRACK, INDEX, TITLE, PERFORMER, etc.
	tokenTypeEof // The main
)

// Token is a (type, Value) pair representing a lexic token.
type Token struct {
	Typ TokenType
	Val string
}

// @TODO Rewrite it to return `Token{{Typ: 3, Val: "x"}`.
//func (t Token) String() string {
//	switch {
//	case t.Typ == TokenTypeWord:
//		return fmt.Sprintf("word %.10q...", t.Val)
//	case t.Typ == tokenTypeEof:
//		return "EOF"
//	}
//
//	return fmt.Sprintf("unknown token type %q", t.Val)
//}
