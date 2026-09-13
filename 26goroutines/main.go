package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointer

func main() {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		wg.Add(1)
		go getStatusCode(web)
	}
	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Oops endpoint:", endpoint, err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
