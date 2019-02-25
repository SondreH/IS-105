package main

import (
	"fmt"
	"./unicode"
)
// Print result from unicode.Translate
func main() {
	fmt.Println(unicode.Translate("“nord og sør”", "is"))

	unicode.UnicodeCodePointDemo()
	}	