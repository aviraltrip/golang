package main

import "fmt"

func main() {
	greeter()
	res := adder(3, 5)
	fmt.Println("Res: ", res)
	proRes, msg := proAdder(2, 3, 4, 5)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro msg is: ", msg)
}

func greeter() {
	fmt.Println("Good morning pineapple looking very good very nice")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	tot := 0
	for _, val := range values {
		tot += val
	}
	return tot, "Hi Aviral"
}
