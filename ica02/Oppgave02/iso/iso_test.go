package iso

import (
	"fmt"
	"testing"
)

func Test_greetingExtendedASCII(t *testing.T) {
	//Greetings
	sl := GetExtendedASCIIStringLiteral()
	var erDetASCII = true
	for _, i := range sl {
		if i < 127 {
			erDetASCII = false
		}
	}
	if erDetASCII == false {
		fmt.Println("Dette er ikke ASCII.")
		t.Fail()
	}
}
