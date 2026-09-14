package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in go")
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	wg.Add(2)

	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
