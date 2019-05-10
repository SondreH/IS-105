package main

import "fmt"

func main() {
	fmt.Print("Linje 1: ")
	fmt.Printf("%s \n", "\xC2\xBD\xb2\x3d\xb2\x20\xe2\x8c\x98")
	fmt.Print("Linje 2: ")
	fmt.Printf("%q \n", "\xC2\xBD\xb2\x3d\xb2\x20\xe2\x8c\x98")
	fmt.Print("Linje 3: ")
	fmt.Printf("%x \n", "\xC2\xBD\xb2\x3d\xb2\x20\xe2\x8c\x98")
	fmt.Print("Linje 4: ")
	fmt.Printf("%+q \n", "\xC2\xBD\xb2\x3d\xb2\x20\xe2\x8c\x98")
	fmt.Print("Linje 5: ")
	fmt.Printf("% x \n", "\xC2\xBD\xb2\x3d\xb2\x20\xe2\x8c\x98")
	fmt.Print("Linje 6: ")
	fmt.Printf("%c \n", "\xC2\xBD\xb2\x3d\xb2\x20\xe2\x8c\x98")
}
