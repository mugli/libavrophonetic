package rulebasedcnv_test

import (
	"fmt"

	"github.com/mugli/libAvroPhonetic/pkg/rulebasedcnv"
)

func ExampleConvertWord() {
	converter := rulebasedcnv.NewConverter()
	output := converter.ConvertWord("bangla")
	fmt.Println(output)

	// Output:
	// বাংলা
}