package databasedconv_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mugli/libavrophonetic/pkg/databasedconv"
	"github.com/stretchr/testify/assert"
)

func newConverterForTest() (*databasedconv.Converter, error) {
	wordsFile, err := os.Open("testdata/generated-words.gob")
	if err != nil {
		return nil, fmt.Errorf("failed to open test datafile: %v", err)
	}

	patternsFile, err := os.Open("testdata/generated-patterns.gob")
	if err != nil {
		return nil, fmt.Errorf("failed to open test datafile: %v", err)
	}

	return databasedconv.NewConverterFromGob(wordsFile, patternsFile)
}

func TestConvertWord(t *testing.T) {
	converter, err := newConverterForTest()
	assert.NoError(t, err)

	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "sari",
			expected: []string{"শারি", "সারি", "শাড়ি", "শাড়ী"},
		},
		{
			input:    "sar",
			expected: []string{"সার", "সাড়", "ষাঁড়"},
		},
		{
			input:    "amra",
			expected: []string{"আমরা", "আমড়া"},
		},
		{
			input:    "lalshak",
			expected: []string{"লালশাক"},
		},
		{
			input:    "lalrong",
			expected: []string{"লালরং", "লালরঙ"},
		},
		{
			input:    "ongshochched",
			expected: []string{"অংশচ্ছেদ"},
		},
		{
			input:    "ongshocched",
			expected: []string{"অংশচ্ছেদ"},
		},
		{
			input:    "shadhinota",
			expected: []string{"স্বাধীনতা"},
		},
		{
			input:    "dukkho",
			expected: []string{"দুখ", "দুঃখ"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			got := converter.ConvertWord(testCase.input)
			assert.Subset(t, got, testCase.expected)
		})
	}
}
