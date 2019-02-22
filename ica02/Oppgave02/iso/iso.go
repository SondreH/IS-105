package iso

import "fmt"

// Oppgave 2a
// Lag selv en "string literal" med utvidet ASCII
// Blir det kun en slik "string literal" eller trenger man flere
// for å representere utvidet ASCII?

const ASCIIExt = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8A\x8B\x8C\x8D\x8E\x8F\x90\x91\x92" + //0x80 til 0x9F er kontroll-karakterer; ikke utskrivbare.
	"\x93\x94\x95\x96\x97\x98\x99\x9A\x9C\x9D\x9E\x9F\xA0\xA1\xA2\xA3\xA4\xA5\xA6" +
	"\xA7\xA8\xA9\xAA\xAB\xAC\xAD\xAE\xAF\xB0\xB1\xB2\xB3\xB4\xB5\xB6\xB7\xB8\xB9" +
	"\xBA\xBB\xBC\xBD\xBE\xBF\xC0\xC1\xC2\xC3\xC4\xC5\xC6\xC7\xC8\xC9\xCA\xCB\xCC" +
	"\xCD\xCE\xCF\xD0\xD1\xD2\xD3\xD4\xD5\xD6\xD7\xD8\xD8\xD9\xDA\xDB\xDC\xDD\xDE" +
	"\xDF\xE0\xE1\xE2\xE3\xE4\xE5\xE6\xE7\xE8\xE9\xEA\xEB\xEC\xED\xEE\xEF\xF0\xF1" +
	"\xF2\xF3\xF4\xF5\xF6\xF7\xF8\xF9\xFA\xFB\xFC\xFD\xFE\xFF" //ASCIIExt(ext = extended) tegnsettet inneholder 128 heksadesimale tegn

const GreetASCII = "æøå\x22\x53\x61\x6C\x75\x74\x2C\x20\xC7\x61\x20\x76\x61\x20\xB0\x2D\x29\x20\x80\xc7\x61\x20\x63\x6F\x75\x74\x65\x20\x35\x30\x22\n" //GreetASCII,
//hvorav euro-tegnet ikke er med i ASCII-Extended, men ble først introdusert i ISO/IEC 8859-15.
// IterateOverASCIIStringLiteral tar en "string literal" som INN-data

//sl = string literal. Iterer gjennom slice ASCIIExt (referert i iso_main.go) og printer %x i hexidesimal 16
//%c printer den korrisponderende Unicode kode-printen, hvorav %b printer i hexidesimal 2.
//
func GetExtendedASCIIStringLiteral() string {
	return ASCIIExt
}

func IterateOverExtendedASCIIStringLiteral(sl string) {
	// Kode for Oppgave 2a

	for i := 0; i < len(sl); i++ {
		fmt.Printf("%X %c %b \n", sl[i], sl[i], sl[i])
	}
}

// GreetingExtendedASCII returnerer en tekst-streng i
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtASCII(sl string) string {
	//Greetings

	for i := 0; i < len(sl); i++ {
		fmt.Printf("%c", sl[i])
	}
	return string(sl)
}
