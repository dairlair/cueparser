package cueparser

import (
	"github.com/google/go-cmp/cmp"
	"io"
	"testing"
)

var tokenizerTests = map[string]struct {
	input          string
	expectedTokens []Token
}{
	//"oneWord": {
	//	input: "hello",
	//	expectedTokens: []Token{
	//		{
	//			Type:  TokenTypeWord,
	//			Value: "hello",
	//		},
	//	},
	//},
	//"oneWordWithTrailingWhitespacesAndTabs": {
	//	input: "\t   \t hello",
	//	expectedTokens: []Token{
	//		{
	//			Type:  TokenTypeWord,
	//			Value: "hello",
	//		},
	//	},
	//},
	//"twoWordsWithTrailingWhitespacesAndTabs": {
	//	input: "\t   \t hello world! \t ",
	//	expectedTokens: []Token{
	//		{
	//			Type:  TokenTypeWord,
	//			Value: "hello",
	//		},
	//		{
	//			Type:  TokenTypeWord,
	//			Value: "world!",
	//		},
	//	},
	//},
	"one letter comment starting with semicolon": {
		input: ";x",
		expectedTokens: []Token{
			{
				Type:  TokenTypeComment,
				Value: "x",
			},
		},
	},
}

func TestTokenizer(t *testing.T) {
	for name, test := range tokenizerTests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			tokens, err := SplitToTokens(test.input)
			if err != nil && err != io.EOF {
				t.Error("unexpected error:", err)
			}

			if diff := cmp.Diff(test.expectedTokens, tokens); diff != "" {
				t.Errorf("Tokens mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
