package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	// limit random to only within the scope if the if statement
	if random := rand.Intn(10); random < 5 {
		fmt.Println("your number is less than 5, awesome")
	} else if random > 5 {
		fmt.Println("your number is greater than 5.")
	}
}