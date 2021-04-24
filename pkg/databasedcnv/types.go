package databasedcnv

import (
	"encoding/gob"
	"fmt"
	"io"

	"github.com/mugli/libavrophonetic/internal/pkg/trie"
)

type PatternBlock struct {
	Transliterate       []string
	EntireBlockOptional bool
}

type Patterns struct {
	Trie                   *trie.Trie
	Dict                   map[string]PatternBlock
	CommonOptionalPatterns []string
}

func NewPatterns() *Patterns {
	return &Patterns{
		Trie:                   trie.NewTrie(),
		Dict:                   make(map[string]PatternBlock),
		CommonOptionalPatterns: make([]string, 0),
	}
}

// LoadFromGob deserializes data encoded as gob ("encoding/gob").
func (patterns *Patterns) LoadFromGob(r io.Reader) error {
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(patterns)

	if err != nil {
		return fmt.Errorf("failed to deserialize: %w", err)
	}

	return nil
}

// SaveToGob serializes data as gob ("encoding/gob").
func (patterns *Patterns) SaveToGob(w io.Writer) error {
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(patterns)

	if err != nil {
		return fmt.Errorf("failed to serialize: %w", err)
	}

	return nil
}

type Words struct {
	Trie *trie.Trie
}

func NewWords() *Words {
	return &Words{
		Trie: trie.NewTrie(),
	}
}

// LoadFromGob deserializes data encoded as gob ("encoding/gob").
func (words *Words) LoadFromGob(r io.Reader) error {
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(words)

	if err != nil {
		return fmt.Errorf("failed to deserialize: %w", err)
	}

	return nil
}

// SaveToGob serializes data as gob ("encoding/gob").
func (words *Words) SaveToGob(w io.Writer) error {
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(words)

	if err != nil {
		return fmt.Errorf("failed to serialize: %w", err)
	}

	return nil
}
