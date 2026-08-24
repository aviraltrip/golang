package main

import "fmt"

func main() {

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// Hello 43210Two One World
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// main ke } se just pahle imagine ki defer lagake aata hai
// LIFO
