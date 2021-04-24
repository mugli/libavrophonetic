package main

import (
	"flag"
	"log"
)

var dataDirectory string

const (
	wordsSourceFileName       = "source-words.txt"
	wordsBinFileName          = "generated-words.gob"
	patternsSourceFileName    = "source-regex-patterns.json"
	patternsPreprocessedFileName = "preprocessed-patterns.json"
	patternsBinFileName       = "generated-patterns.gob"
)

func init() {
	const (
		defaultDataDirectory   = "."
		dataDirectoryPathUsage = "the path to data directory"
	)
	flag.StringVar(&dataDirectory, "data-directory", defaultDataDirectory, dataDirectoryPathUsage)
}

func main() {
	flag.Parse()
	saveBinWordData()
	log.Println("----------------------------------------------------")
	saveRegexPatternsFromSource()
	log.Println("----------------------------------------------------")
	saveBinPatternsData()
}
