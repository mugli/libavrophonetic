package main

import (
	"bufio"
	"github.com/mugli/libAvroPhonetic/pkg/phoneticconverter/dictionary/types"
	"log"
	"os"
	"path/filepath"
)

func saveBinWordData() {
	words := types.NewWords()

	wordDictionaryPath, err := filepath.Abs(filepath.Join(dataDirectory, wordsSourceFileName))
	if err != nil {
		log.Fatal(err)
	}

	sourceWords := loadSourceWordDictionary(wordDictionaryPath)
	for _, w := range sourceWords {
		words.Trie.AddWord(w)
	}

	generatedFileName, err := filepath.Abs(filepath.Join(dataDirectory, wordsBinFileName))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Generated word dictionary will be saved to: %s\n", generatedFileName)
	outFile, err := os.Create(generatedFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	log.Println("Saving word dictionary.")
	err = words.SaveToGob(outFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Saved generated word dictionary.")
}

func loadSourceWordDictionary(filepath string) (words []string) {
	log.Printf("Loading source word dictionary: %s\n", filepath)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded source word dictionary. Total: %d enties.\n", len(words))

	return
}