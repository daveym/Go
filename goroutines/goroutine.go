package main

import (
	"fmt"
	"sync"
)

// using waitgrouops
var wg sync.WaitGroup

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}

	wg.Done()
}

func main() {

	fmt.Println("Start")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	wg.Add(1)
	go f("goroutine")

	fmt.Println("Done")
	wg.Wait()
}
