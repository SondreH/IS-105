package lineshift

import (
	"fmt"
	"strings"
)

func Lineshift(etArray []byte) {
	i := len(etArray) - 2
	i2 := len(etArray) - 1
	//Sjekke i sjekker lengden på arrayet og flytte indeksen 2 plasser tilbake.
	arrLinjeNestSist := etArray[i]
	arrLinjeSist := etArray[i2]
	nestSiste := fmt.Sprintf("%X\n", arrLinjeNestSist)
	siste := fmt.Sprintf("%X\n", arrLinjeSist) //For å sjekke om
	/*det nestsiste og siste ASCII-symbolet har 0D eller 0A, eller annet.
	(https://en.wikipedia.org/wiki/Newline */
	//var minTekst = fmt.Printf("%X\n", etArray)
	fmt.Printf("%X\n", etArray) //Enklere å dobbelsjekke & teste,
	//printer ASCII symbolene for hver av tegnene.
	fmt.Println("Trimspace returnerer for nestsiste linje:", strings.TrimSpace(nestSiste))
	fmt.Println("Trimspace returnerer for siste linje:", strings.TrimSpace(siste))

	/*if strings.Contains(minTekst, "A") {
		fmt.Println("Dette er A")
	} else if strings.Contains(minTekst, "D") {
		fmt.Println("Dette er D")
	} else {
		fmt.Println("Dette er noe annet")
	}*/

	/*TrimSpace fjerner alt av whitespacing for så å sjekke om
	A tilsvarer å bevege seg en linje forover (LF, Line feed),
	hvorav D tilsvarer linjeskift (CR, Carriage Return).
	Hvis siste linje har linjeskift printes beskjeden)*/
	if strings.TrimSpace(nestSiste) == "D" || strings.TrimSpace(siste) == "D" {
		/*TrimSpace fjerner alt av whitespacing for så å sjekke om
		D tilsvarer linjeskift (CR, Carriage Return). Dersom den tilsvarer A inneholder den en ny linje (LF, Linefeed)*/
		fmt.Println("Det er med linjeskift (CR, Carriage Return).")
	} else if (strings.TrimSpace(siste) == "A" || strings.TrimSpace(nestSiste) == "A") && (strings.TrimSpace(siste) != "D" || strings.TrimSpace(nestSiste) != "D") {
		fmt.Println("Dette er med linefeed (LF) uten CR")
	} else {
		fmt.Println("Dette er i et annet format enn CR / LF.")
	}
}
