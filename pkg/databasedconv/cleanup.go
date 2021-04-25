package databasedconv

import "strings"

func isConvertibleChar(c uint8) bool {
	if ((c >= 'a') && (c <= 'z')) || ((c >= 'A') && (c <= 'Z')) {
		return true
	}

	return false
}

func fixString(input string) string {
	var result string

	input = strings.TrimSpace(input)

	for i := 0; i < len(input); i++ {
		makeLower := true
		char := input[i]

		if !isConvertibleChar(char) {
			continue
		}

		// Fix string for o. In the beginning, after punctuations etc it should be capital O
		if char == 'o' || char == 'O' {
			if i == 0 {
				makeLower = false
			} else if !isConvertibleChar(input[i-1]) {
				makeLower = false
			}
		}

		if makeLower {
			result += strings.ToLower(string(char))
		} else {
			result += strings.ToUpper(string(char))
		}
	}

	return result
}

func unique(input []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)

	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
