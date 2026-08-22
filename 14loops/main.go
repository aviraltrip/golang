package main

import "fmt"

func main() {
	days := []string{"Sun", "Tue", "Wed", "Fri", "Sat"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	for idx, day := range days {
		fmt.Printf("idx is %v and val is %v\n", idx, day)
	}
	lol := 1
	for lol < 10 {
		if lol == 5 {
			lol++
			continue
		}
		if lol == 2 {
			goto lmao
		}
		fmt.Println("val is: ", lol)
		lol++
	}
lmao:
	fmt.Println("Jumping at google.com")
}
