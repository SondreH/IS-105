package main

import (
	"github.com/shammers95/IS-105/ica03/Oppgave1/fileutils"
)

func main() {
	fileutils.FileToByteslice("./fileutils/text1.txt")
	fileutils.FileToByteslice("./fileutils/text2.txt")
	//Kaller metoden FileToByteslice med
	//parameter text1.txt og text2.txt fra mappen fileutils.
}
