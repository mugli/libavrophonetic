package classic_test

import (
	"fmt"

	"github.com/mugli/libAvroPhonetic/pkg/phoneticconverter/classic"
)

func ExampleConvertWord() {
	converter := classic.NewConverter()
	output := converter.ConvertWord("bangla")
	fmt.Println(output)

	// Output:
	// বাংলা
}
