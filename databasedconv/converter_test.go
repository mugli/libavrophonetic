package databasedconv_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mugli/libavrophonetic/databasedconv"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input string
	want  []string
}

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

func buildTestCases() []testCase {
	return []testCase{
		{
			input: "sari",
			want:  []string{"শারি", "সারি", "শাড়ি", "শাড়ী"},
		},
		{
			input: "sar",
			want:  []string{"সার", "সাড়", "ষাঁড়"},
		},
		{
			input: "amra",
			want:  []string{"আমরা", "আমড়া"},
		},
		{
			input: "lalshak",
			want:  []string{"লালশাক"},
		},
		{
			input: "lalrong",
			want:  []string{"লালরং", "লালরঙ"},
		},
		{
			input: "ongshochched",
			want:  []string{"অংশচ্ছেদ"},
		},
		{
			input: "ongshocched",
			want:  []string{"অংশচ্ছেদ"},
		},
		{
			input: "shadhinota",
			want:  []string{"স্বাধীনতা"},
		},
		{
			input: "dukkho",
			want:  []string{"দুখ", "দুঃখ"},
		},
	}
}

func TestConverter_ConvertWord(t *testing.T) {
	converter, err := newConverterForTest()
	assert.NoError(t, err)

	testCases := buildTestCases()

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			got := converter.ConvertWord(testCase.input)
			assert.Subset(t, got, testCase.want)
		})
	}
}

func BenchmarkConverter_ConvertWord(b *testing.B) {
	converter, err := newConverterForTest()
	assert.NoError(b, err)

	testCases := buildTestCases()
	numTests := len(testCases) - 1

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		converter.ConvertWord(testCases[i%numTests].input)
	}
}
