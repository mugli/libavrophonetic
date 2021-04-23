// Package classic provides Avro Phonetic Classic converter to convert English string to similar sounding Bengali string
package classic

import (
	"unicode"
)

// Converter is a classic (fixed/static) Avro Phonetic converter
type Converter struct {
	patterns *patterns
}

// NewConverter returns an initialized classic (fixed/static) Avro Phonetic converter
func NewConverter() *Converter {
	return &Converter{
		patterns: newPatterns(),
	}
}

// ConvertWord transliterates from English string to similar sounding Bengali string
//nolint:funlen
//nolint:gocognit
func (phonetic *Converter) ConvertWord(word string) string {
	var output string

	input := fixCase(word)

	for i := 0; i < len(input); i++ {
		var rightPos, leftPos int
		startPos := i
		hasMatched := false

		for _, pattern := range *(phonetic.patterns) {
			rightPos = i + len(pattern.match)

			if (rightPos <= len(input)) && string(input[startPos:rightPos]) == pattern.match { //nolint:nestif
				leftPos = startPos - 1

				for _, exception := range pattern.exceptions {
					shouldReplace := true
					cursor := 0

					for _, matchCondition := range exception.ifAllMatch {
						if matchCondition.when == suffix {
							cursor = rightPos
						} else {
							cursor = leftPos
						}

						switch matchCondition.is {
						case punctuation:
							if ((cursor < 0 && matchCondition.when == prefix) ||
								(cursor >= len(input) && matchCondition.when == suffix) ||
								isPunctuation(input[cursor])) == matchCondition.isNot {
								shouldReplace = false

								break
							}

						case vowel:
							if (((cursor >= 0 && matchCondition.when == prefix) ||
								(cursor < len(input) && matchCondition.when == suffix)) &&
								isVowel(input[cursor])) == matchCondition.isNot {
								shouldReplace = false

								break
							}

						case consonant:
							if (((cursor >= 0 && matchCondition.when == prefix) ||
								(cursor < len(input) && matchCondition.when == suffix)) && isConsonant(input[cursor])) == matchCondition.isNot {
								shouldReplace = false

								break
							}

						case exactly:
							var s, e int
							if matchCondition.when == suffix {
								s = rightPos
								e = rightPos + len(matchCondition.value)
							} else {
								s = startPos - len(matchCondition.value)
								e = startPos
							}

							if !isExact(matchCondition.value, input, s, e, matchCondition.isNot) {
								shouldReplace = false

								break
							}
						}
					}

					if shouldReplace {
						output += exception.thenReplace
						i = rightPos - 1
						hasMatched = true

						break
					}
				}

				if hasMatched {
					break
				}

				// Default
				output += pattern.replace
				i = rightPos - 1
				hasMatched = true

				break
			}
		}

		if !hasMatched {
			output += string(input[i])
		}
	}

	return output
}

func isExact(needle string, heyStack []rune, startPos int, rightPos int, isNot bool) bool {
	return (startPos >= 0 && rightPos < len(heyStack) && (string(heyStack[startPos:rightPos]) == needle)) != isNot
}

func fixCase(input string) []rune {
	fixed := make([]rune, 0, len(input))

	for _, r := range input {
		if isCaseSensitive(r) {
			fixed = append(fixed, r)
		} else {
			fixed = append(fixed, unicode.ToLower(r))
		}
	}

	return fixed
}
