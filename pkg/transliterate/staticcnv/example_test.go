package staticcnv_test

import (
	"fmt"
	"github.com/mugli/libAvroPhonetic/pkg/transliterate/staticcnv"
)

func ExampleConvertWord() {
	bengali := staticcnv.ConvertWord("amar")
	fmt.Println(bengali)

	// Output:
	// আমার
}