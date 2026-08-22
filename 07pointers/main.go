package main

import "fmt"

func main() {
	fmt.Println("Pointers in go")

	// var ptr *int
	// fmt.Println("val of ptr is ", ptr)

	myNo := 23
	var ptr = &myNo

	fmt.Println("Val of actual ptr is ", ptr)
	fmt.Println("val of actual pte is ", *ptr)
	*ptr = *ptr + 2
	fmt.Println("New val: ", myNo)
}
