// Package trie implements a rune Trie (https://en.wikipedia.org/wiki/Trie).
package trie

// Node is an element of the Trie.
type Node struct {
	Children map[rune]*Node
	// CompleteWord is empty string for most Node.
	// It only contains a value when the current Node contains the last character/rune of a complete entry in the Trie.
	CompleteWord string
}

// NewNode returns an initialized Node.
func NewNode() *Node {
	return &Node{
		Children:     map[rune]*Node{},
		CompleteWord: "",
	}
}

// FindMatchingNode starts from the Node as root,
// traverses it's children to find the input word, and finally
// returns the Node where the input word ends.
// It stops as soon as a it reaches the node that ends with input word,
// and does not care if that node contains a partial or complete entry in the Trie.
func (root *Node) FindMatchingNode(word string) *Node {
	if word == "" {
		return nil
	}

	result := root

	for _, c := range word {
		node := result.Children[c]

		if node == nil {
			return nil
		}

		result = node
	}

	return result
}

// IsCompleteWord returns true if the Node contains the last character of a complete entry in the Trie.
func (root *Node) IsCompleteWord() bool {
	return len(root.CompleteWord) != 0
}

// findLongestPrefixNode finds the node with similar longest prefix of the input.
func (root *Node) findLongestPrefixNode(prefix string) (matchedNode *Node, matchedPrefix string) {
	result := root
	matchedKeys := make([]rune, 0)

	for _, char := range prefix {
		node := result.Children[char]

		if node == nil && result == root {
			// No match at the beginning
			matchedNode = nil
			matchedPrefix = ""

			return
		}

		if node == nil {
			break
		} else {
			matchedKeys = append(matchedKeys, char)
			result = node
		}
	}

	matchedNode = result
	matchedPrefix = string(matchedKeys)

	return
}

// findCompleteWordsInChildren starts with the current Node and recursively traverse
// the Trie to find all the complete words from the closest children.
func (root *Node) findCompleteWordsInChildren() (result []string) {
	for key := range root.Children {
		node := root.Children[key]

		if node.IsCompleteWord() {
			result = append(result, node.CompleteWord)
		}

		completeWords := node.findCompleteWordsInChildren()
		result = append(result, completeWords...)
	}

	return result
}

// Trie is a prefix tree. There's one Node for each character/rune of a string, forming a tree like structure.
type Trie struct {
	Root *Node
}

// NewTrie returns an initialized Trie.
func NewTrie() *Trie {
	return &Trie{
		Root: NewNode(),
	}
}

// AddWord adds a word to the Trie.
func (trie *Trie) AddWord(word string) {
	tree := trie.Root

	for _, c := range word {
		if _, ok := tree.Children[c]; !ok {
			node := NewNode()
			tree.Children[c] = node
		}

		tree = tree.Children[c]
	}

	finalNode := tree
	finalNode.CompleteWord = word
}

// MatchPrefix finds all the words in a trie that starts with the prefix.
func (trie *Trie) MatchPrefix(prefix string) []string {
	result := make([]string, 0)

	if prefix == "" {
		return result
	}

	node, _ := trie.Root.findLongestPrefixNode(prefix)

	if node == nil {
		return result
	}

	result = node.findCompleteWordsInChildren()

	if node.IsCompleteWord() {
		result = append(result, node.CompleteWord)
	}

	return result
}

// MatchLongestCommonPrefix finds the node with similar longest prefix of the input,
// with additional information like remaining string it couldn't match.
// The returned isMatchCompleteWord is true when the match is also a complete entry in the Trie.
func (trie *Trie) MatchLongestCommonPrefix(prefix string) (matchedPrefix string, remaining string,
	isMatchCompleteWord bool, node *Node) {
	remaining = prefix

	if prefix == "" {
		return
	}

	node, matchedPrefix = trie.Root.findLongestPrefixNode(prefix)

	if node != nil {
		remaining = prefix[len(matchedPrefix):]
		isMatchCompleteWord = node.IsCompleteWord()
	}

	return
}


