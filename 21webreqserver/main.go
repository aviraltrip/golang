package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerformPostJsonRequest()
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

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"
	// nakli json payload
	requestBody := strings.NewReader(`
	{
		"coursename": "learning go",
		"price": 0,
		"platform": "https://aviraltrip.vercel.app/"
	}`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const hiurl = "http://localhost:3000/platform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "aviral")
	data.Add("lastname", "tripathi")
	data.Add("email", "aviral@go.dev")

	res, err := http.PostForm(hiurl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
