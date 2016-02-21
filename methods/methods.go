package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func (p *point) addtopoint(x, y int) {
	p.x += x
	p.y += y

}

func main() {

	p := point{1, 1}
	p.addtopoint(2, 2)

	fmt.Println(p)
}
