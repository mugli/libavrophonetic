package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wzshiming/crun"
)

type patternBlockSource struct {
	Transliterate       string `json:"transliterate"`
	EntireBlockOptional bool   `json:"entireBlockOptional"`
}

type patternBlockPreprocessed struct {
	Transliterate       []string `json:"transliterate"`
	Count               int      `json:"count"`
	EntireBlockOptional bool     `json:"entireBlockOptional,omitempty"`
}

type patternSource map[string]patternBlockSource
type patternPreprocessed map[string]patternBlockPreprocessed

func uniqueSlice(input []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func saveRegexPatternsFromSource() {
	sourceFileName, err := filepath.Abs(filepath.Join(dataDirectory, patternsSourceFileName))
	if err != nil {
		log.Fatal(err)
	}

	destFileName, err := filepath.Abs(filepath.Join(dataDirectory, patternsGeneratedFileName))
	if err != nil {
		log.Fatal(err)
	}

	var pSource patternSource
	var pDest patternPreprocessed

	log.Printf("Loading source regex patterns: %s\n", sourceFileName)
	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	byteVal, err := ioutil.ReadAll(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteVal, &pSource)

	log.Println("Pre-processing possible patterns from regex.")
	pDest = make(map[string]patternBlockPreprocessed, len(pSource))

	for key, pattern := range pSource {
		cs, err := crun.Compile(pattern.Transliterate)
		if err != nil {
			log.Fatal(err)
		}

		transliterate := []string{}
		cs.Range(func(generated string) bool {
			if generated != "" {
				transliterate = append(transliterate, generated)
			}
			return true
		})

		transliterate = uniqueSlice(transliterate)
		count := len(transliterate)

		pDest[key] = patternBlockPreprocessed{
			Transliterate:       transliterate,
			EntireBlockOptional: pattern.EntireBlockOptional,
			Count:               count,
		}
	}

	log.Println("Generating JSON output for pre-processed patterns.")
	jsonString, err := json.MarshalIndent(&pDest, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Saving pre-processed patterns to file: %s\n", destFileName)
	err = ioutil.WriteFile(destFileName, jsonString, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Saved pre-processed patterns.")
}
