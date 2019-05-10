package main

import (
	"fmt"

	"github.com/SondreH/IS-105/ica02/Oppgave04/unicode"
)

// Printer her resultat fra unicode.Translate og
//bruker funksjonen UnicodeCodePointDemo for å printe unicode symbolene
func main() {
	fmt.Println(unicode.Translate("“nord og sør”", "is"))

	unicode.UnicodeCodePointDemo()
}
