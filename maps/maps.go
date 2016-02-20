package main

import (
	"fmt"
)

// To create an empty map, use the builtin make: make(map[key-type]val-type).

func main() {

	bu := make(map[string]int)

	bu["EAYL"] = 7
	bu["Trainee"] = 6
	bu["Engineer"] = 5
	bu["Senior"] = 4
	bu["Tech Lead"] = 3
	bu["Solution Architect"] = 2
	bu["Principal"] = 1
	bu["Head"] = 0

	fmt.Println(len(bu))

	// Prints them unordered!
	fmt.Println(bu)
	delete(bu, "EAYL")

	fmt.Println(bu)

	// key, value - basically checks if it's present
	mapkey, mapvalue := bu["Principal"]
	fmt.Println(mapkey, mapvalue)

	// Init in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
