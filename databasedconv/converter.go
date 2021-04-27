package databasedconv

import (
	"fmt"
	"io"

	"github.com/mugli/libavrophonetic/data"
	"github.com/mugli/libavrophonetic/internal/trie"
)

type Converter struct {
	patterns *Patterns
	words    *Words
}

func NewConverter() (*Converter, error) {
	wordsFile, err := data.DataFiles.Open("generated-words.gob")
	if err != nil {
		return nil, fmt.Errorf("failed to embedded datafile: %v", err)
	}

	patternsFile, err := data.DataFiles.Open("generated-patterns.gob")
	if err != nil {
		return nil, fmt.Errorf("failed to embedded datafile: %v", err)
	}

	return NewConverterFromGob(wordsFile, patternsFile)
}

func NewConverterFromGob(gobWords io.Reader, gobPatterns io.Reader) (*Converter, error) {
	patterns := NewPatterns()
	words := NewWords()

	err := words.LoadFromGob(gobWords)
	if err != nil {
		return nil, err
	}

	err = patterns.LoadFromGob(gobPatterns)
	if err != nil {
		return nil, err
	}

	return &Converter{
		patterns: patterns,
		words:    words,
	}, nil
}

func (converter *Converter) ConvertWord(input string) []string {
	result := make([]string, 0)

	word := fixString(input)
	matched, remaining, _, _ := converter.patterns.Trie.MatchLongestCommonPrefix(word)

	matchedPatterns := converter.patterns.Dict[matched].Transliterate
	sizeOfPatterns := len(matchedPatterns)
	matchedNodes := make([]*trie.Node, 0, sizeOfPatterns*len(converter.patterns.CommonOptionalPatterns))

	for _, p := range matchedPatterns {
		node := converter.words.Trie.Root.FindMatchingNode(p)
		if node != nil {
			matchedNodes = append(matchedNodes, node)
		}

		// Try matching optional patterns too
		additionalNodes := make([]*trie.Node, 0, len(matchedNodes)*len(converter.patterns.CommonOptionalPatterns))

		for _, matchedNode := range matchedNodes {
			for _, commonPattern := range converter.patterns.CommonOptionalPatterns {
				additionalNode := matchedNode.FindMatchingNode(commonPattern)
				if additionalNode != nil {
					additionalNodes = append(additionalNodes, additionalNode)
				}
			}
		}
		// Merge additional nodes with matchedNodes
		matchedNodes = append(matchedNodes, additionalNodes...)
	}

	for len(remaining) > 0 {
		newMatched, newRemaining, isCompleteWord, _ := converter.patterns.Trie.MatchLongestCommonPrefix(remaining)
		if !isCompleteWord {
			for i := len(remaining) - 1; i >= 0; i-- {
				newMatched, newRemaining, isCompleteWord, _ = converter.patterns.Trie.MatchLongestCommonPrefix(remaining[0:i])
				if isCompleteWord {
					remaining = remaining[i:]
					break
				}
			}
		} else {
			remaining = newRemaining
		}

		newMatchedPatterns := converter.patterns.Dict[newMatched].Transliterate
		sizeOfPatterns = len(newMatchedPatterns)
		newMatchedNodes := make([]*trie.Node, 0, sizeOfPatterns)

		for _, p := range newMatchedPatterns {
			for _, node := range matchedNodes {
				newNode := node.FindMatchingNode(p)
				if newNode != nil {
					newMatchedNodes = append(newMatchedNodes, newNode)
				}
			}
		}

		if converter.patterns.Dict[newMatched].EntireBlockOptional {
			// Entirely optional patterns like "([ওোঅ]|(অ্য)|(য়ো?))?" may not yield any result
			matchedNodes = append(matchedNodes, newMatchedNodes...)
		} else {
			matchedNodes = newMatchedNodes
		}

		// Try matching optional patterns too
		additionalNodes := make([]*trie.Node, 0, len(matchedNodes)*len(converter.patterns.CommonOptionalPatterns))

		for _, matchedNode := range matchedNodes {
			for _, commonPattern := range converter.patterns.CommonOptionalPatterns {
				additionalNode := matchedNode.FindMatchingNode(commonPattern)
				if additionalNode != nil {
					additionalNodes = append(additionalNodes, additionalNode)
				}
			}
		}

		// Merge additional nodes with matchedNodes
		matchedNodes = append(matchedNodes, additionalNodes...)
	}

	for _, node := range matchedNodes {
		if node.IsCompleteWord() {
			result = append(result, node.CompleteWord)
		}
	}

	result = unique(result)

	return result
}
