package main

import "fmt"

const ASCII = "\x48\x65\x6c\x6c\x6f\x2c\x20"

//const kinesisk = "\x4e16x\754c"

func main() {
	fmt.Print(ASCII)
	//fmt.Print("kinesisk) //fungerer ikke, fordi
	//det ikke er en del av ASCII extended
	//(av de 255 tegnene, http://www.asciitable.com/).
	//Bruker derfor:
	fmt.Print("世界")
}
