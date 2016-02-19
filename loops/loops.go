package main

import "fmt"

func main() {

	count := 10

	// standard loop
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	// 'while' loop variant
	x := 20
	for x > count {
		fmt.Println(x)
		x--
	}

	// range, with index
	for j, utfchar := range "hello world" {
		fmt.Println("Index", j, " utf character", utfchar)

	}

	// range, no index
	for _, utfchar := range "hello world" {
		fmt.Println("utf character", utfchar)

	}

	// string iteration
	sentence := "howdy dowdy"

	for i, r := range sentence {
		fmt.Println("index ", i, "character ", string(r))

	}

}
