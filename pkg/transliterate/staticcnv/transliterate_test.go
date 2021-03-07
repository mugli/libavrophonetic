package staticcnv_test

import (
	"github.com/mugli/libAvroPhonetic/pkg/transliterate/staticcnv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertWord(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "amar",
			want:  "আমার",
		},
		{
			input: "a",
			want:  "আ",
		},
		{
			input: "ia",
			want:  "ইয়া",
		},
		{
			input: "a`",
			want:  "া",
		},
		{
			input: "R",
			want:  "ড়",
		},
		{
			input: "Rh",
			want:  "ঢ়",
		},
		{
			input: "bou",
			want:  "বউ",
		},
		{
			input: "bOU",
			want:  "বৌ",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.input, func(t *testing.T) {
			got := staticcnv.ConvertWord(testCase.input)
			assert.Equal(t, testCase.want, got)
		})
	}
}
