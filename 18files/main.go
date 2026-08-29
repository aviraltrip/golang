package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Isko file me daalna hai"
	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	defer file.Close()
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is:", length)
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Data inside file\n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
