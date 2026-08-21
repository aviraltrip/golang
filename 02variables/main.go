// go mod init variables
// go run main.go
// fp - shortcut for fmp Println
package main

import "fmt"

const LoginToken string = "lmaololmeow" //capital L = public

func main() {
	var username string = "Aviral"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T\n", isLoggedIn)

	var smallVal uint8 = 255 // 256 mein error
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T\n", smallVal)

	var smallFloat float64 = 255.5467675756756
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T\n", smallFloat)

	// default values & some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T\n", anotherVariable)

	//implicit type
	var website = "https://www.wikipedia.org/"
	fmt.Println(website)

	// no var style - walrus operator
	noOfUser := 30000
	fmt.Println(noOfUser)
	fmt.Println(LoginToken)
	fmt.Printf("Variable type is %T\n", LoginToken)
}
