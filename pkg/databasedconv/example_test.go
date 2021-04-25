package databasedconv_test

import (
	"fmt"

	"github.com/mugli/libavrophonetic/pkg/databasedconv"
)

func ExampleConverter_ConvertWord() {
	converter, err := databasedconv.NewConverter()
	if err != nil {
		fmt.Printf("Error reading datafiles %v\n", err)
		return
	}

	output := converter.ConvertWord("bangla")
	for _, word := range output {
		fmt.Println(word)
	}

	// Output:
	// বাংলা
	// বাঙলা
}
