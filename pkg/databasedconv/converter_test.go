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
		return nil, fmt.Errorf("failed to test datafile: %v", err)
	}

	patternsFile, err := os.Open("testdata/generated-patterns.gob")
	if err != nil {
		return nil, fmt.Errorf("failed to test datafile: %v", err)
	}

	return databasedconv.NewConverterFromGob(wordsFile, patternsFile)
}

func TestConvertWord(t *testing.T) {
	converter, err := newConverterForTest()
	assert.NoError(t, err)

	expected := []string{"শারি", "সারি", "শাড়ি", "শাড়ী"}
	got := converter.ConvertWord("sari")
	assert.Subset(t, got, expected)
}
