package cueparser

import "testing"

func TestClassifier(t *testing.T) {
	classifier := newRuneClassifier()
	tests := map[rune]runeClass{
		' ':  runeClassSpace,
		'\t': runeClassSpace,
		'a':  ruleClassUnknown,
		//'"':  quoteRuneClass,
		//'#':  commentRuneClass,
		//';':  commentRuneClass,
		//'\n': runeClassEol,
		//'\r': runeClassEol,
	}
	for runeChar, want := range tests {
		got := classifier.ClassifyRune(runeChar)
		if got != want {
			t.Errorf("ClassifyRune(%v) -> %v. Want: %v", runeChar, got, want)
		}
	}
}
