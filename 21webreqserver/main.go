package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

}

func PerformGetRequest() {
	const url = "http://localhost:3000/get"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("Content Length is: ", res.ContentLength)

	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := resString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(resString.String())

	fmt.Println(string(content))
	fmt.Println(content)

}
