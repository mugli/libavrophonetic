package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mugli/libavrophonetic/pkg/databasedcnv"
)

func saveBinPatternsData() {
	patterns := generatePatternTrie()

	outFilePath, err := filepath.Abs(filepath.Join(dataDirectory, patternsBinFileName))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Generated pattern db will be saved to: %s\n", outFilePath)
	outFile, err := os.Create(outFilePath)

	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	log.Println("Saving pattern db.")
	err = patterns.SaveToGob(outFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Saved generated pattern db.")
}

func loadCommonPatterns() (commonPatterns []string) {
	inputFilePath, err := filepath.Abs(filepath.Join(dataDirectory, commonPatternsSourceFileName))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loading common patterns: %s\n", inputFilePath)
	jsonFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteVal, &commonPatterns)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func loadGeneratedJSONPatterns() (dataPatterns map[string]patternBlockPreprocessed) {
	inputFilePath, err := filepath.Abs(filepath.Join(dataDirectory, patternsPreprocessedFileName))
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

	err = json.Unmarshal(byteVal, &dataPatterns)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded patterns from pre-processed JSON. Total: %d \n", len(dataPatterns))

	return
}

func generatePatternTrie() *databasedcnv.Patterns {
	retval := databasedcnv.NewPatterns()

	dataPatterns := loadGeneratedJSONPatterns()

	for key, p := range dataPatterns {
		retval.Trie.AddWord(key)
		retval.Dict[key] = databasedcnv.PatternBlock{
			Transliterate:       p.Transliterate,
			EntireBlockOptional: p.EntireBlockOptional,
		}
	}

	retval.CommonOptionalPatterns = loadCommonPatterns()

	return retval
}
