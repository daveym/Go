package main

import (
	"fmt"
)

func main() {

	var a = [5]int{1, 2, 3, 4, 5}
	var b = [5]rune{'a', 'b', 'c', 'd', 'e'}

	fmt.Println(len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		fmt.Println(b[i])
	}

}
