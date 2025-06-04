package cueparser

import (
	"bufio"
	"fmt"
	"io"
)

// The internal state used by the tokenizer state machine.
type tokenizerState int

// Tokenizer state machine states.
const (
	startState   tokenizerState = iota // No runes have been seen
	inWordState                        // We are processing regular runes in a word
	quotingState                       // we are within a quoted string
	commentState                       // we are within a quoted string
)

// Tokenizer turns an input stream into a sequence of typed tokens
type Tokenizer struct {
	input      bufio.Reader
	classifier runeClassifier
}

// NewTokenizer creates a new tokenizer from an input stream.
func NewTokenizer(r io.Reader) *Tokenizer {
	input := bufio.NewReader(r)
	classifier := newRuneClassifier()
	return &Tokenizer{
		input:      *input,
		classifier: classifier,
	}
}

// scanStream scans the stream for the next token using the internal state machine.
// It will panic if it encounters a rune which it does not know how to handle.
func (t *Tokenizer) scanStream() (*Token, error) {
	// The initial state of FSM
	state := startState

	var tokenType TokenType
	// It is our buffer where we are storing consumed runes and flush it when token finished.
	var value []rune
	var nextRune rune
	var nextRuneClassifiedClass runeTokenClass
	var err error

	for {
		nextRune, _, err = t.input.ReadRune()
		nextRuneClassifiedClass = t.classifier.ClassifyRune(nextRune)

		if err == io.EOF {
			nextRuneClassifiedClass = eofRuneClass
		} else if err != nil {
			return nil, err
		}

		switch state {
		case startState: // no runes read yet
			{
				switch nextRuneClassifiedClass {
				case eofRuneClass:
					{
						// We read nothing and the first rune we read - is End OF life :( We can return only nothing.
						return nil, io.EOF
					}
				case spaceRuneClass:
					{
						// It is ok, space runes are fine, just skip them and wait for REAL words begin!
					}

				case commentRuneClass:
					{
						// Good practice to begin any file with comments, let's switch to the `commentState` and
						// set the Type to CommentToken
						tokenType = CommentToken
						state = commentState
					}
				default:
					{
						tokenType = WordToken
						value = append(value, nextRune)
						state = inWordState
					}
				}
			}
		case inWordState:
			{
				switch nextRuneClassifiedClass {
				case eofRuneClass:
					// During the word runes consuming we met the EOL, ok, lets flush read Value with not error.
					{
						token := &Token{
							Type:  tokenType,
							Value: string(value),
						}
						return token, nil
					}
				case spaceRuneClass:
					{
						token := &Token{
							Type:  tokenType,
							Value: string(value),
						}
						return token, err
					}
				default:
					{
						// Just add read Value to buffer and read further.
						value = append(value, nextRune)
					}
				}
			}
		//case quotingState: // in non-escaping single quotes
		//	{
		//		switch nextRuneClassifiedClass {
		//		case eofRuneClass:
		//			{
		//				err = fmt.Errorf("EOF found when expecting closing quote")
		//				token := &Token{
		//					Type: Type,
		//					Value:     string(Value)}
		//				return token, err
		//			}
		//		default:
		//			{
		//				Value = append(Value, nextRune)
		//			}
		//		}
		//	}
		//case commentState: // in a comment
		//	{
		//		switch nextRuneClassifiedClass {
		//		case eofRuneClass:
		//			{
		//				token := &Token{
		//					Type: Type,
		//					Value:     string(Value)}
		//				return token, err
		//			}
		//		case spaceRuneClass:
		//			{
		//				if nextRune == '\n' {
		//					state = startState
		//					token := &Token{
		//						Type: Type,
		//						Value:     string(Value)}
		//					return token, err
		//				} else {
		//					Value = append(Value, nextRune)
		//				}
		//			}
		//		default:
		//			{
		//				Value = append(Value, nextRune)
		//			}
		//		}
		//	}
		default:
			{
				return nil, fmt.Errorf("unexpected state: %v", state)
			}
		}
	}
}

// Next returns the next token in the stream.
func (t *Tokenizer) Next() (*Token, error) {
	return t.scanStream()
}
