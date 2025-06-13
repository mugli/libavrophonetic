
# libavrophonetic



Go module providing classic (rule-based) and dictionary backed transliterators for [Avro Phonetic](https://en.wikipedia.org/wiki/Avro_Keyboard).

---

## Update (2025-06-13)

### Take a look at https://github.com/OpenBangla/upodesh for a better alternative.

Trie is fast, but can require a lot of memory. `OpenBangla/upodesh` uses Finite State Transducer (FST) instead to tackle the memory usage, and it is still fast (See this PR: https://github.com/OpenBangla/upodesh/pull/10)! 

---

## Optimizations

This implementation is also the fastest dictionary based suggestion generator as far as I know. Primarily because this does not scan through the dictionary looking for regular-expression match and use a **Trie** instead.

Comparing apples to oranges (because why not), this is **~100 times faster** than previous JavaScript and regular-expression based suggestion generator (tested in Node.js env).




## Demo

This module is intended to be used as a library.

However, for quickly checking the output there is a demo CLI. Run the following command:

```bash
go run ./cmd/avrophoneticdemo shadhinota
```

## Installation


```bash 
  go get -u github.com/mugli/libavrophonetic
```

## Usage/Examples

```go
package main

import (
	"fmt"
	"os"

	"github.com/mugli/libavrophonetic/databasedconv"
	"github.com/mugli/libavrophonetic/rulebasedconv"
)

func main() {
	input := "bangla"

	rulebasedConverter := rulebasedconv.NewConverter()
	databasedConverter, _ := databasedconv.NewConverter() // ignoring error for brevity

	rulebasedOutput := rulebasedConverter.ConvertWord(input)
	databasedOutput := databasedConverter.ConvertWord(input)

	fmt.Printf("(Rulebased conversion) %s = %s \n", input, rulebasedOutput) // বাংলা
	fmt.Printf("(Databased conversion) %s = %v \n", input, databasedOutput) // [বাংলা বাঙলা]
}
```

## API Documentation

https://pkg.go.dev/github.com/mugli/libavrophonetic

## Running Tests

To run tests/see coverage, run the following commands:

```bash
  make test
  make test-cover
```

## Data generation

Instead of using plain text data-files, this module uses a [gob encoded](https://blog.golang.org/gob) files for faster data loading (aka, Trie generation).

Also, the gob files gets embedded with the binary during compile time using [embed package](https://golang.org/pkg/embed/) introduced in Go 1.16.

If you change source files (having `soure-` prefix in the filenames) in `./data` directory, run the following command to re-generate the binary data files:

```bash
  make generate-data
```

## Authors and Acknowledgement

The git blame is not showing Tahmid's name in this repo because this was started from scratch, but both me and [@tahmidsadik](https://github.com/tahmidsadik/) had a lot of fun making the initial prototype for this.


## License

[MIT](https://choosealicense.com/licenses/mit/)

  
