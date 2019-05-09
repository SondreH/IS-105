package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//Path på linje 67 må endres til din lokale path.

//Basert på https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
func skrivFil(lines [5]int, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close() //ikke nødvendig, men grei kontroll.

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	//Array med 256 plasser for hvert ASCII-symbol
	var a256 [256]int
	var a5 [5]int
	//Leser fila, setter teksten inn i en array
	filText, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Kan ikke lese filen", os.Args[1])
		log.Fatal(err)
	}
	//str := string(filText) //Setter teksten til string istedenfor []byte

	//Legger hvert symbol til sin respektive plass mellom 0-256
	for i := 0; i <= len(filText)-1; i++ {
		a256[filText[i]]++
	}

	//Printer antall antall newlines fra symbol-arrayet
	fmt.Println("Ant newline/linjeskift: ", a256[10])
	//Går gjennom a256. Finner høyeste verdi i lista
	//Printer plassen, symbolet og antall ganger den blir skrevet.
	//Setter verdien i den nåværende plassen til 0
	//(for å finne verdien lavere enn høyeste). Gjentas fem ganger.
	for i := 0; i < 5; i++ {
		antll := 0
		tegn := 0
		for i := 0; i < len(a256); i++ {
			if antll < a256[i] {
				antll = a256[i]
				tegn = i
			}
		}
		fmt.Printf("%d", i+1) //Nummer x
		fmt.Print(". Rune: ")
		fmt.Printf("%q", tegn) //Hvilken rune
		fmt.Println(", Counts: ", antll)
		a5[i] = antll //For å skrive top 5 til ny fil. Plass 0 står ubrukt, kan gjøres til a10, men enklere å lese.
		a256[tegn] = 0
	}
	antFiler, _ := ioutil.ReadDir("C:/Users/Sondre/go/src/github.com/SondreH/IS-105/ica03/Oppgave3/frequence") //MÅ BYTTE TIL DIN PATH.
	//fmt.Println(len(antFiler)) //hvis du ønsker å se ant filer i folderen.
	var pathNummer = strconv.Itoa(len(antFiler)) //Nummeret vil variere ut ifra antall filer totalt i mappen.
	//Bruker strconv.Itoa for å konvertere fra int til string
	//(blir et '♦'-tegn ved vanlig parsing, altså hvis: pathNummer := string(len(antFiler)))
	fmt.Println("Antall filer i mappen:", pathNummer)
	//En ulempe er et kan skape skape problemer etter sletting av filer ved mange opprettinger.
	skrivFil(a5, "noe"+pathNummer+".txt")
	fmt.Print("Navnet til filen er: 'noe", pathNummer, ".txt'\n")
}

//Tidligere forsøk:

/*func tellAntRunes(s string, r rune) int { //for å telle antall linjeskift (newline)
//og de fem mest brukte tegnene
antNL := 0
for _, x := range s {
	if x == r {
		antNL++
	}
}
return antNL
}

/*fil, err := os.Open("pg100.txt") //Åpner en fil, her pg100.txt
if err != nil {
	log.Fatal(err)
}
fildata, err  := ioutil.ReadAll(fil)
if err != nil{
	log.Fatal(err)
}*/

//fmt.Printf("I Hexadecimal: %x\n", fildata)

/*
	fmt.Println("Antall tegn:")
	for i := 0; i <= len(filText); i++ {
		var ant int = tellAntRunes(str, rune(i))
		if ant > 0 {
			fmt.Print("Verdi: ")
			fmt.Printf("%c ", i)
			fmt.Print("forekommer: ", ant, " antall ganger \n")
			//Lister antall ganger hver ASCII-decimal blir brukt (sjekker ASCII-extended)
		}
	}

	/*fmt.Println("Totalt antall runes:", utf8.RuneCountInString(str))
	fmt.Println("Antall linjer:", tellAntRunes(str, 10)+1) //+1 fordi
	//antlinjeskift+1 fordi det ikke er linjeskift før den første linja.

	/*
		var a = "a"
		var b = "a"
		var c = "a"
		var d = "a"
		var e = "a"

			func storstTop5 int {
				var ant int = 0
				var storst int = 0
				var storst2 int = 0
				var storst3 int = 0
				var storst4 int = 0
				var storst5 int = 0
				for _, ant := range s {
				if storst == storst2 {
					ant++
					storst := ant
				}
				if storst2 > storst3 && b != c {

				}
				return storst, storst2, storst3, storst4, storst5
				}
			}
			fmt.Println("De fem hyppigst brukte karakterene er:", a, b, c, d, e, "og de blir gjentatt respektivt", storstTop5, "ganger hver")
*/
