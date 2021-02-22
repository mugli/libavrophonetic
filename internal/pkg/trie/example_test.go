package trie_test

import (
	"fmt"
	"github.com/mugli/libAvroPhonetic/internal/pkg/trie"
)

func Example() {
	// Create a new list and put some numbers in it.
	t := trie.NewTrie()
	t.AddWord("বাংলা")
	t.AddWord("বাংলাদেশ")
	t.AddWord("বাংলাদেশের")
	t.AddWord("বাঙালি")

	// To find all the words beginning with a prefix:
	result := t.MatchPrefix("বা")
	for _, word := range result {
		fmt.Println(word)
	}

	// Output:
	// বাংলা
	// বাংলাদেশ
	// বাংলাদেশের
	// বাঙালি
}

func Example_matchLongestCommonPrefix() {
	// Create a new list and put some numbers in it.
	t := trie.NewTrie()
	t.AddWord("বাংলা")
	t.AddWord("বাংলাদেশ")
	t.AddWord("বাংলাদেশের")
	t.AddWord("বাঙালি")

	// To find the node with similar longest prefix of the input:
	matchedPrefix, remaining, isMatchCompleteWord, _ := t.MatchLongestCommonPrefix("বাংলার")
	fmt.Println("matchedPrefix ==", matchedPrefix)
	fmt.Println("remaining ==", remaining)
	fmt.Println("isMatchCompleteWord ==", isMatchCompleteWord) // isMatchCompleteWord is true when the match is also a complete entry in the Trie

	// Output:
	// matchedPrefix == বাংলা
	// remaining == র
	// isMatchCompleteWord == true
}
