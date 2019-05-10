package slice

//import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere og lengden på slicen.
// Returnerer en slice av type []byte
//
func AllocateVar(b int) []byte {
	var slice []byte
	// make reserverer et minneområde av størrelse b.
	slice = make([]byte, b, b)
	// returnerer slicen vi har allokert plass til.
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

	// Hvis uidx er større enn kapasiteten på slicen, blir vi nødt til å allokere mer plass.
	if uidx > cap(slc) {
		// Vi lage en ny slice
		var newSlice = AllocateMake(uidx - lidx)
		// Og kopierer over innholdet fra den gamle slicen
		copy(newSlice, slc[lidx:])
		// Til slutt returneres den nye slicen.
		return newSlice
	} else {
		// Hvis vi har nok kapasitet i eksisterende slice, utvides/reduseres størrelsen innfor minneområdet som er allokert.
		return slc[lidx:uidx]
	}
}

// CopySlice ???
