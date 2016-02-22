package main

import "fmt"

func MyCounter() func() {
	thecount := 0

	increment := func() {
		thecount++
		fmt.Println("The count is ", thecount)
	}

	return increment

}

func main() {

	var increment func()
	increment = MyCounter()
	for i := 0; i < 5; i++ {
		increment()
	}

}
