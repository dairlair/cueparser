package cueparser

// runeClass is the type of UTF-8 character classification: A quote, space, escape, end of line, etc...
type runeClass int

// Classes of rune token
const (
	ruleClassUnknown runeClass = iota
	runeClassSpace
)

func (s runeClass) String() string {
	switch s {
	case ruleClassUnknown:
		return "ruleClassUnknown"
	case runeClassSpace:
		return "runeClassSpace"
	default:
		return "Unknown"
	}
}

// Named classes of UTF-8 runes
const (
	spaceRunes = " \t"
	//quoteRunes   = `"`
	//commentRunes = "#;"
	//eolRunes     = "\r\n"
)

type runeClassifier struct {
	mapping map[rune]runeClass
}

func (rc runeClassifier) addRuneClass(runes string, tokenType runeClass) {
	for _, runeChar := range runes {
		rc.mapping[runeChar] = tokenType
	}
}

// newRuneClassifier creates a new classifier for ASCII characters.
func newRuneClassifier() runeClassifier {
	rc := runeClassifier{
		mapping: map[rune]runeClass{},
	}
	rc.addRuneClass(spaceRunes, runeClassSpace)
	return rc
}

func (rc runeClassifier) ClassifyRune(runeVal rune) runeClass {
	return rc.mapping[runeVal]
}
