package trie_test

import (
	"testing"

	"github.com/mugli/libAvroPhonetic/internal/pkg/trie"
	"github.com/stretchr/testify/assert"
)

func buildTrie(entries []string) *trie.Trie {
	trie := trie.NewTrie()

	for _, entry := range entries {
		trie.AddWord(entry)
	}

	return trie
}

func TestNode_FindMatchingNode(t *testing.T) {
	tree := buildTrie([]string{
		"ক", "কখ", "কখগঘঙচছ",
	})

	n1 := tree.Root.FindMatchingNode("ক")
	assert.NotNil(t, n1)

	n2 := n1.FindMatchingNode("খ")
	assert.NotNil(t, n2)

	n3 := n2.FindMatchingNode("গঘ")
	assert.NotNil(t, n3)

	n4 := tree.Root.FindMatchingNode("কখগঘ")
	assert.NotNil(t, n4)
}

func TestNode_IsCompleteWord(t *testing.T) {
	tree := buildTrie([]string{
		"ক", "কখ", "কখগঘঙচছ",
	})

	n1 := tree.Root.FindMatchingNode("ক")
	assert.True(t, n1.IsCompleteWord())

	n2 := n1.FindMatchingNode("খ")
	assert.True(t, n2.IsCompleteWord())

	n3 := n2.FindMatchingNode("গঘ")
	assert.False(t, n3.IsCompleteWord())

	n4 := tree.Root.FindMatchingNode("কখগঘ")
	assert.False(t, n4.IsCompleteWord())
}

func TestTrie_MatchPrefix(t *testing.T) {
	tree := buildTrie([]string{
		"ক", "কখগ", "কখগঘঙ",
		"চ", "চছজ", "চছজঝঞ",
		"১",
	})

	tests := []struct {
		prefix string
		want   []string
	}{
		{
			prefix: "ক",
			want:   []string{"ক", "কখগ", "কখগঘঙ"},
		},
		{
			prefix: "কখ",
			want:   []string{"কখগ", "কখগঘঙ"},
		},
		{
			prefix: "চছজঝঞ",
			want:   []string{"চছজঝঞ"},
		},
		{
			prefix: "২",
			want:   []string{},
		},
		{
			prefix: "",
			want:   []string{},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.prefix, func(t *testing.T) {
			got := tree.MatchPrefix(testCase.prefix)
			assert.ElementsMatch(t, testCase.want, got)
		})
	}
}

func TestTrie_MatchLongestCommonPrefix(t *testing.T) {
	tree := buildTrie([]string{
		"ক", "কখগ", "কখগঘঙ",
		"চ", "চছজ", "চছজঝঞ",
		"১",
	})

	tests := []struct {
		prefix               string
		matchedPrefixWant    string
		remainingWant        string
		shouldBeCompleteWord bool
		nodeIsNil            bool
	}{
		{
			prefix:               "ক",
			matchedPrefixWant:    "ক",
			remainingWant:        "",
			shouldBeCompleteWord: true,
			nodeIsNil:            false,
		},
		{
			prefix:               "ক1234",
			matchedPrefixWant:    "ক",
			remainingWant:        "1234",
			shouldBeCompleteWord: true,
			nodeIsNil:            false,
		},
		{
			prefix:               "1234",
			matchedPrefixWant:    "",
			remainingWant:        "1234",
			shouldBeCompleteWord: false,
			nodeIsNil:            true,
		},
		{
			prefix:               "কখগঘঙচছজঝঞ",
			matchedPrefixWant:    "কখগঘঙ",
			remainingWant:        "চছজঝঞ",
			shouldBeCompleteWord: true,
			nodeIsNil:            false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.prefix, func(t *testing.T) {
			matchedPrefixGot, remainingGot, gotCompleteWord, node := tree.MatchLongestCommonPrefix(testCase.prefix)
			assert.Equal(t, testCase.matchedPrefixWant, matchedPrefixGot)
			assert.Equal(t, testCase.remainingWant, remainingGot)
			assert.Equal(t, testCase.shouldBeCompleteWord, gotCompleteWord)
			assert.Equal(t, testCase.nodeIsNil, node == nil)
		})
	}
}

func BenchmarkTrie_MatchLongestCommonPrefix(b *testing.B) {
	tree := buildTrie([]string{
		"ক", "কখগ", "কখগঘঙ",
		"চ", "চছজ", "চছজঝঞ",
		"১",
	})

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		for i := 1; i < 2000; i++ {
			tree.MatchLongestCommonPrefix("কখগঘঙচছজঝঞ")
		}
	}
}

func BenchmarkNode_FindMatchingNode(b *testing.B) {
	tree := buildTrie([]string{
		"ক", "কখগ", "কখগঘঙ",
		"চ", "চছজ", "চছজঝঞ",
		"১",
	})

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		for i := 1; i < 2000; i++ {
			tree.Root.FindMatchingNode("কখগঘঙ")
		}
	}
}
