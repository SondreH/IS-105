package main

import (
	"fmt"
	"github.com/karlosmartines/karls-windows/IS-105/ica02/oppgave4_unicode/unicode"
)
// Print result from unicode.Translate
func main() {
	fmt.Println(unicode.Translate("“nord og sør”", "is"))

	unicode.UnicodeCodePointDemo()
	}	
