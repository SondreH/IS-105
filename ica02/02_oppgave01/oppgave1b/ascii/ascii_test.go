package ascii

import "testing"

//Oppgave1b

func TestGreetingsASCII(t *testing.T) {
	ascii := GreetingsASCII()
	if !(isASCII(ascii)) {
		t.Fail()
	}
}

func isASCII(s string) bool {
	for _, C := range s {
		if C > 127 {
			return false
		}
	}
	return true
}
