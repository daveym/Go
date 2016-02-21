package main

import (
	"fmt"
)

type point struct {
	x, y int
}

type multipoint struct {
	name string
	point
}

func main() {

	var p point
	p.x = 3
	p.y = 5

	fmt.Println(p.x, p.y)

	p2 := point{5, 6}
	fmt.Println(p2)

	// named initialisation
	p3 := point{x: 1}
	fmt.Println(p3)

	// nested struct
	first := multipoint{
		"first",
		p2,
	}

	fmt.Println(first)
}
