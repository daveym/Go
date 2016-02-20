package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {

	var p point
	p.x = 3
	p.y = 5

	fmt.Println(p.x, p.y)

	p2 := point{5, 6}
	fmt.Println(p2)

}
