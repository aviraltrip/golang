package main

import (
	"fmt"
	"sort"
)

func main() {
	var sabji = []string{"lauki", "bhindi", "aalu"}
	fmt.Printf("type is %T\n", sabji)

	sabji = append(sabji, "tamatar", "torai")
	fmt.Println(sabji)

	sabji = append(sabji[1:3], sabji[3:]...)
	fmt.Println(sabji)

	scores := make([]int, 4)

	scores[0] = 234
	scores[1] = 100
	scores[2] = 454
	scores[3] = 104
	scores = append(scores, 900, 899, 500)
	fmt.Println(scores)
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	// removing val from slice based on idx
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var idx int = 2
	courses = append(courses[:idx], courses[idx+1:]...)
	fmt.Println(courses)
}
