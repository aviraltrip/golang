package main

import "fmt"

func main() {
	langs := make(map[string]string)

	langs["JS"] = "Javascript"
	langs["RB"] = "Ruby"
	langs["PY"] = "Python"

	fmt.Println("List of all lang: ", langs)
	fmt.Println("JS shorts for: ", langs["JS"])

	delete(langs, "RB")
	fmt.Println("List of all langs: ", langs)

	for key, value := range langs {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	// for _, value := range langs {
	// 	fmt.Printf("For key v, value is %v\n", value)
	// }
	// ignored using underscore

}
