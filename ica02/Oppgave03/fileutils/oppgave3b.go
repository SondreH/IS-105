package main

import (
	"fmt"
	"./fileutils"
)

func main() {
	
	fmt.Println("lang01")
	fileutils.FileToByteslice("../files/lang01.wl")

	fmt.Println("lang02")
	fileutils.FileToByteslice("../files/lang02.wl")

	fmt.Println("lang03")
	fileutils.FileToByteslice("../files/lang03.wl")
}
