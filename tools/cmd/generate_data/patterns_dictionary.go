package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mugli/libAvroPhonetic/pkg/phoneticconverter/dictionary/types"
)

func saveBinPatternsData() {
	patterns := generatePatternTrie()

	outFilePath, err := filepath.Abs(filepath.Join(dataDirectory, patternsBinFileName))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Generated pattern dictionary will be saved to: %s\n", outFilePath)
	outFile, err := os.Create(outFilePath)

	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	log.Println("Saving pattern dictionary.")
	err = patterns.SaveToGob(outFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Saved generated pattern dictionary.")
}

func loadGeneratedJSONPatterns() (dataPatterns map[string]patternBlockPreprocessed) {
	inputFilePath, err := filepath.Abs(filepath.Join(dataDirectory, patternsGeneratedFileName))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loading pre-processed JSON: %s\n", inputFilePath)
	jsonFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteVal, &dataPatterns)

	log.Printf("Loaded patterns from pre-processed JSON. Total: %d \n", len(dataPatterns))

	return
}

func generatePatternTrie() *types.Patterns {
	retval := types.NewPatterns()

	dataPatterns := loadGeneratedJSONPatterns()

	for key, p := range dataPatterns {
		retval.Trie.AddWord(key)
		retval.Dict[key] = types.PatternBlock{
			Transliterate:       p.Transliterate,
			EntireBlockOptional: p.EntireBlockOptional,
		}
	}

	return retval
}
