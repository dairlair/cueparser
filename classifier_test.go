package cueparser

import "testing"

func TestClassifier(t *testing.T) {
	classifier := newRuneClassifier()
	tests := map[rune]runeTokenClass{
		' ':  spaceRuneClass,
		'\t': spaceRuneClass,
		'"':  quoteRuneClass,
		'#':  commentRuneClass,
		';':  commentRuneClass,
		'\n': eolRuneClass,
		'\r': eolRuneClass,
	}
	for runeChar, want := range tests {
		got := classifier.ClassifyRune(runeChar)
		if got != want {
			t.Errorf("ClassifyRune(%v) -> %v. Want: %v", runeChar, got, want)
		}
	}
}
