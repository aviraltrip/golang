package main

import "fmt"

func main() {
	//no inheritance in golang, no super or parent
	aviral := User{"Aviral", "aviral@gmail.com", true, 20}
	fmt.Println(aviral)
	fmt.Printf("details are: %+v\n", aviral) //+v sari cheeze  only v values sirf
	fmt.Printf("Name is %v & email is %v", aviral.Name, aviral.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
