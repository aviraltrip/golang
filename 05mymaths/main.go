package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Maths in go")
	// var myNoOne int = 2
	// var myNoTwo float64 = 4.5

	// fmt.Println("Sum is: ", myNoOne+int(myNoTwo))

	//random no
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random from crypto
	randomNo, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNo)
}
