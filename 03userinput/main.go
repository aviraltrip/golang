package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome bro"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for sabji:")

	// comma ok || comma err

	input, _ := reader.ReadString('\n') // similar to try catch
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}
