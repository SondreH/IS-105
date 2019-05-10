package ascii

import "fmt"

//Oppgave1b
func GreetingsASCII() string {
	t1 := []byte("\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29")
	for i := 0; i < len(t1); i++ {
		fmt.Printf("%c", t1[i])
	}
	retur := string(t1)
	return retur
}
