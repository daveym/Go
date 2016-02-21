package main

import "fmt"

// Interfaces can handles a collection of methods, not fields
// in this case, the animal interface will expose canMoo
type animal interface {
	canEat(food string) bool
}

type cow struct {
	Breed string
}

// canEat method is applied to a cow
func (c cow) canEat(food string) bool {

	if food == "grass" {
		return true
	}

	return false
}

func main() {

	var daisy animal
	daisy = cow{"Aberdeen Angus"}
	fmt.Println("Can eat grass? - ", daisy.canEat("grass"))
	fmt.Println("Can eat mars bars? - ", daisy.canEat("mars bars"))
}
