package main

/* In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts
with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C */
import (
	"errors"
	"fmt"
)

func checkspelling(word string) (bool, error) {

	if word == "d4vey" {
		return true, nil
	}

	return false, errors.New("Something didnt match up")

}

func main() {

	if m, e := checkspelling("davey"); e != nil {
		fmt.Println("Failed:", e)
	} else {
		fmt.Println("Worked:", m)
	}

}
