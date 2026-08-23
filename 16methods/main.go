package main

import "fmt"

func main() {
	//no inheritance in golang, no super or parent
	aviral := User{"Aviral", "aviral@gmail.com", true, 20}
	fmt.Println(aviral)
	fmt.Printf("details are: %+v\n", aviral) //+v sari cheeze  only v values sirf
	fmt.Printf("Name is %v & email is %v\n", aviral.Name, aviral.Email)
	aviral.getStatus()
	aviral.newMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) newMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email) // creates copy not actual me jake change
}
