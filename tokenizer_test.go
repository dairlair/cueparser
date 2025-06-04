package cueparser

import (
	"github.com/google/go-cmp/cmp"
	"io"
	"strings"
	"testing"
)

var tokenizerTests = map[string]struct {
	input          string
	expectedTokens []Token
}{
	"oneWord": {
		input: "hello",
		expectedTokens: []Token{
			{
				Type:  WordToken,
				Value: "hello",
			},
		},
	},
	"oneWordWithTrailingWhitespacesAndTabs": {
		input: "\t   \t hello",
		expectedTokens: []Token{
			{
				Type:  WordToken,
				Value: "hello",
			},
		},
	},
	"twoWordsWithTrailingWhitespacesAndTabs": {
		input: "\t   \t hello world! \t ",
		expectedTokens: []Token{
			{
				Type:  WordToken,
				Value: "hello",
			},
			{
				Type:  WordToken,
				Value: "world!",
			},
		},
	},
	"twoWordsWithCommentBySemicolon": {
		input: "\t   \t alice bob \t ;is is a comment",
		expectedTokens: []Token{
			{
				Type:  WordToken,
				Value: "alice",
			},
			{
				Type:  WordToken,
				Value: "bob",
			},
			{
				Type:  CommentToken,
				Value: "is is a comment",
			},
		},
	},
}

func TestTokenizer(t *testing.T) {
	for name, test := range tokenizerTests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			tokenizer := NewTokenizer(strings.NewReader(test.input))

			tokens := make([]Token, 0)
			for {
				token, err := tokenizer.Next()
				if err != nil {
					if err == io.EOF {
						break
					}
					t.Error(err)
				}

				if token == nil {
					t.Fatalf("The nil token returned when not nil was expected")
				}

				tokens = append(tokens, *token)
			}

			if diff := cmp.Diff(test.expectedTokens, tokens); diff != "" {
				t.Errorf("Tokens mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
