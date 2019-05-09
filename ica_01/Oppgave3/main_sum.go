package main

import "fmt"
import "os"
import "strconv"
import "bufio"
import sum "github.com/SondreH/IS-105/ica_01/Oppgave3/sum_pass"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	var text2 string
	fmt.Println("Skriv inn tall for å summere, skriv x for å avslutte")

	for text != "x" {
		fmt.Print("Tall nummer 1: ")
		scanner.Scan()
		text = scanner.Text()
		fmt.Print("Tall nummer 2: ")
		scanner.Scan()
		text2 = scanner.Text()
		if text != "x" {

			tall1, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				fmt.Println("Tallet kan ikke være: " + text + ", bruk int64 verdi.")
			} else {
				tall2, err := strconv.ParseInt(text2, 10, 64)
				if err != nil {
					fmt.Println("Tallet kan ikke være: " + text2 + ", bruk int64 verdi.")
				} else {
					fmt.Println(text+" + "+text2+" =", sum.SumInt64(tall1, tall2))
				}
			}
		}
	}
}
