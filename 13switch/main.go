package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNo := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNo)

	switch diceNo {
	case 1:
		fmt.Println("Dice value 1 & u can open")

	case 2:
		fmt.Println("u can more 2")
	case 3:
		fmt.Println("u can move 3")
		fallthrough
	case 4:
		fmt.Println("u can move 4")
		fallthrough //force 2 cont to very next case ignoring whether case is t or f
	case 5:
		fmt.Println("u can move 5")
	case 6:
		fmt.Println("u can move 6 & roll again")
	default:
		fmt.Println("lmao")
	}
}
