package rulebasedconv_test

import (
	"fmt"

	"github.com/mugli/libavrophonetic/pkg/rulebasedconv"
)

func ExampleConverter_ConvertWord() {
	converter := rulebasedconv.NewConverter()
	output := converter.ConvertWord("bangla")
	fmt.Println(output)

	// Output:
	// বাংলা
}
