package fileutils

import (
	"IS-105/ica03/Oppgave1/lineshift"
	"fmt"
	"io"
	"log"
	"os"
)

/*FileToByteSlice : The input is the name of the file, output is the slice itself.*/
func FileToByteslice(filnavn string) []byte {
	// Open file for reading
	file, err := os.Open(filnavn)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sizeOfSlice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, sizeOfSlice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%c", byteSlice)
	fmt.Println("")

	lineshift.Lineshift(byteSlice)

	return byteSlice

}
