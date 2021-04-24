package rulebased_test

import (
	"fmt"

	"github.com/mugli/libAvroPhonetic/pkg/phoneticconverter/rulebased"
)

func ExampleConvertWord() {
	converter := rulebased.NewConverter()
	output := converter.ConvertWord("bangla")
	fmt.Println(output)

	// Output:
	// বাংলা
}
