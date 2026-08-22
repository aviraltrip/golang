package main

import "fmt"

func main() {
	var sabji [4]string
	sabji[0] = "aalu"
	sabji[1] = "baingan"
	sabji[3] = "lauki"
	fmt.Println("Sabji list is ", sabji)
	fmt.Println("len is ", len(sabji)) // len 4 aaegi coz declare 4 ki, even though andar 3 items hai

	var fruits = [5]string{"banana", "apple", "mango"}
	fmt.Println("Fruits are: ", fruits)
	fmt.Println("len", len(fruits))
}
