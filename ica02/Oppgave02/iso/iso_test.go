package iso

import (
	"fmt"
	"testing"
	"unicode"
)

func Test_greetingExtendedASCII(t *testing.T) {
	//Greetings
	//sl := GetExtendedASCIIStringLiteral()
	sl := "\x84\x86\x41"
	for _, i := range sl {
		if i > unicode.MaxASCII {
			fmt.Println("Dette er ASCII.")
		} else {
			fmt.Println("Dette er ikke ASCII.")
			t.Fail()
		}
	}
}
