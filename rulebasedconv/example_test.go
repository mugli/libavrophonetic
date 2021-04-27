package rulebasedconv_test

import (
	"fmt"

	"github.com/mugli/libavrophonetic/rulebasedconv"
)

func ExampleConverter_ConvertWord() {
	converter := rulebasedconv.NewConverter()
	output := converter.ConvertWord("bangla")
	fmt.Println(output)

	// Output:
	// বাংলা
}
