package cueparser

// runeTokenClass is the type of UTF-8 character classification: A quote, space, escape.
type runeTokenClass int

// Classes of rune token
const (
	unknownRuneClass runeTokenClass = iota
	spaceRuneClass
	quoteRuneClass
	commentRuneClass
	eolRuneClass
	eofRuneClass
)

// Named classes of UTF-8 runes
const (
	spaceRunes   = " \t"
	quoteRunes   = `"`
	commentRunes = "#;"
	eolRunes     = "\r\n"
)

type runeClassifier struct {
	mapping map[rune]runeTokenClass
}

func (rc runeClassifier) addRuneClass(runes string, tokenType runeTokenClass) {
	for _, runeChar := range runes {
		rc.mapping[runeChar] = tokenType
	}
}

// newRuneClassifier creates a new classifier for ASCII characters.
func newRuneClassifier() runeClassifier {
	rc := runeClassifier{
		mapping: map[rune]runeTokenClass{},
	}
	rc.addRuneClass(spaceRunes, spaceRuneClass)
	rc.addRuneClass(quoteRunes, quoteRuneClass)
	rc.addRuneClass(commentRunes, commentRuneClass)
	rc.addRuneClass(eolRunes, eolRuneClass)
	return rc
}

func (rc runeClassifier) ClassifyRune(runeVal rune) runeTokenClass {
	return rc.mapping[runeVal]
}
