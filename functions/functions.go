package main

import "fmt"

func reverse(sentence string) (reversedSentence string, isError bool) {
	isError = false

	n := len(sentence)
	runes := make([]rune, n)
	for _, rune := range sentence {
		n--
		runes[n] = rune
	}
	return string(runes[n:]), isError

}

func main() {

	sentence := "abcdefghijklmnopqrstuvwxyz"
	reversed, error := reverse(sentence)

	fmt.Println(reversed, error)
}
