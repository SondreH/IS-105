package lineshift

import (
	"fmt"
	"strings"
)

func Lineshift(etArray []byte) {
	i := len(etArray) - 2
	//Sjekke lengden på arrayet og flytte indeksen 2 plasser tilbake
	arrLinjeNestSist := etArray[i]
	siste := fmt.Sprintf("%X\n", arrLinjeNestSist) //For å sjekke om
	/*det nestsiste ASCII-symbolet har 0D eller 0A/annet.
	(https://en.wikipedia.org/wiki/Newline
	TrimSpace fjerner alt av whitespacing for så å sjekke om
	A tilsvarer å bevege seg en linje forover (LF, Line feed),
	hvorav D tilsvarer linjeskift (CR, Carriage Return).
	Hvis siste linje har linjeskift printes beskjeden)*/

	fmt.Printf("%X\n", etArray) //Enklere å dobbelsjekke & teste,
	//printer UTF-8 symbolene for hver av tegnene.

	if strings.TrimSpace(siste) == "D" {
		/*TrimSpace fjerner alt av whitespacing for så å sjekke om
		D tilsvarer linjeskift (CR, Carriage Return).
		Hvis siste linje har linjeskift printes beskjeden*/
		fmt.Println("Det er med linjeskift.")
	} else {
		fmt.Println("Det er ikke med linjeskift.")
	}

	/*etArrayString := string(etArray)
	var a = "%\n"
	containLineshift := fmt.Sprintf(a, etArray)
	fmt.Println("Inneholder denne filen linjeskift?", containLineshift)*/
}
