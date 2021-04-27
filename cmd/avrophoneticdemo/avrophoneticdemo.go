package main

import (
	"fmt"
	"os"

	"github.com/mugli/libavrophonetic/databasedconv"
	"github.com/mugli/libavrophonetic/rulebasedconv"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	arg := os.Args[1]

	rulebasedConverter := rulebasedconv.NewConverter()

	databasedConverter, err := databasedconv.NewConverter()
	if err != nil {
		fmt.Printf("Error reading datafiles %v\n", err)
		os.Exit(1)
	}

	rulebasedConversion := rulebasedConverter.ConvertWord(arg)
	databasedConversion := databasedConverter.ConvertWord(arg)

	fmt.Printf("(Rulebased conversion) %s = %s \n", arg, rulebasedConversion)
	fmt.Printf("(Databased conversion) %s = %v \n", arg, databasedConversion)
}

func printUsage() {
	help := `
Pass a word as argument:

avrophoneticdemo shadhinota`

	fmt.Println(help)
}
