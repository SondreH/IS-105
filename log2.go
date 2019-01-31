package main

import (
"fmt"
"math"
"os"
)

func log2cli (){
	ettArgument := os.Args
	annetArgument :=os.Args[1:]

	arg := os.Args[3]

	fmt.Println(ettArgument)
	fmt.Println(annetArgument)
	fmt.Println(arg)
}

func main (){
fmt.Println("Resultatet er: ", math.Log2(2))
}
