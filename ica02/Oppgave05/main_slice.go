package main

import (
	"fmt"

	"github.com/SondreH/IS-105/ica02/Oppgave05/slice"
)

func main() {

	//Deklarer og allokerer 2 slicer, med kapasitet 10,30
	var mySlice = slice.AllocateVar(10)
	var mySlice2 = slice.AllocateMake(30)

	mySlice[0] = 'a'
	//Setter plass 0 til 6 og
	mySlice2[10] = 'H'
	mySlice2[11] = 'E'
	mySlice2[12] = 'I'
	mySlice2[13] = 'S'
	mySlice2[14] = 'A'
	mySlice2[15] = 'N'
	mySlice2[16] = 'N'
	pointeren := &mySlice2 //Pointer = pointeren

	fmt.Println("Capacity mySlice2:")
	fmt.Println(cap(mySlice2))
	fmt.Println("Length mySlice2:")
	fmt.Println(len(mySlice2))
	fmt.Println("Length pointeren:")
	fmt.Println(len(*pointeren)) //Pointeren = mySlice2

	//Resliser fra kapasitet 30 til 190
	mySlice2 = slice.Reslice(mySlice2, 10, 200)

	fmt.Println("New length mySlice2:")
	fmt.Println(len(mySlice2))
	fmt.Println("Capacity new mySlice2:")
	fmt.Println(cap(mySlice2))
	fmt.Printf("%s\n", mySlice2)
	fmt.Println("Length pointeren:")
	fmt.Println(len(*pointeren)) //Pointeren er fortsatt lik mySlice2,
	//og selve pointeren har ikke blitt endret på; men har blitt større fordi mySlice2 har blitt større.
}
