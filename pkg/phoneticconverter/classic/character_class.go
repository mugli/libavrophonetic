package classic

import (
	"unicode"
)

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

var consonants = map[rune]bool{
	'b': true,
	'c': true,
	'd': true,
	'f': true,
	'g': true,
	'h': true,
	'j': true,
	'k': true,
	'l': true,
	'm': true,
	'n': true,
	'p': true,
	'q': true,
	'r': true,
	's': true,
	't': true,
	'v': true,
	'w': true,
	'x': true,
	'y': true,
	'z': true,
}

var caseSensitive = map[rune]bool{
	'o': true,
	'i': true,
	'u': true,
	'd': true,
	'g': true,
	'j': true,
	'n': true,
	'r': true,
	's': true,
	't': true,
	'y': true,
	'z': true,
}

func isVowel(r rune) bool {
	return vowels[unicode.ToLower(r)]
}

func isConsonant(r rune) bool {
	return consonants[unicode.ToLower(r)]
}

func isPunctuation(r rune) bool {
	return !isVowel(r) && !isConsonant(r)
}

func isCaseSensitive(r rune) bool {
	return caseSensitive[unicode.ToLower(r)]
}
