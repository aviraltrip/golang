package main

import (
	"fmt"
	"net/url"
	neturl "net/url"
)

const urlString = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	// parsing
	res, err := neturl.Parse(urlString)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Printf("type of query params: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
