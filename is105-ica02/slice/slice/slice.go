package slice

//import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere og lengden på slicen.
// Returnerer en slice av type []byte
//
func AllocateVar(b int) []byte {
	// make reserverer et minneområde av størrelse b.
	var slice = make([]byte, b, b)
	return slice
}

//
// AllocateMake tar lengde og kapasitet som b og returnerer en ny slice direkte
//
func AllocateMake(b int) []byte {
	// Kode for Oppgave 5a
	return make([]byte, b)
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	// kode her for 5b
	// vi antar at lidx er innenfor slc sin range
	// sjekker om det er nødvendig med ny minneallokering

	// Hvis
	if uidx > cap(slc) {
		var newSlice = AllocateMake(uidx - lidx)
		copy(newSlice, slc[lidx:])
		return newSlice
	} else {
		return slc[lidx:uidx]
	}
}

// CopySlice ???
