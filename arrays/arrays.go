package main

import (
	"fmt"
)

func main() {

	var a = [5]int{1, 2, 3, 4, 5}
	var b = [5]byte{'a', 'b', 'c', 'd', 'e'}

	fmt.Println(len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		fmt.Println(string(b[i]))
	}

	// multidimensional, woooooo
	c := [3][3]int{{0, 0, 0}, {0, 0, 0}}
	fmt.Println(c)

}
