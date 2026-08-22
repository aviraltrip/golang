package main

import "fmt"

func main() {
	loginC := 10
	var res string
	if loginC < 10 {
		res = "Regular user"
	} else if loginC > 10 {
		res = "Watch out"
	} else {
		res = "Exactly 10 login count"
	}

	fmt.Println(res)

	if num := 3; num < 10 {
		fmt.Println("Num less than 10")
	} else {
		fmt.Println("Num not less than 10")
	}

	// if err != nil {

	// }
}
