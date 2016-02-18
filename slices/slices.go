package main

import (
	"bytes"
	"fmt"
)

type path []byte

// AddOneToEachElement increments the slice
func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

// SubtractOneFromLength reduces the size of the slice
func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

// PtrSubtractOneFromLength modify the length of a slice
func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

// TruncateAtFinalSlash removes a final slash
func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func main() {

	var buffer [10]byte
	fmt.Println("buffer", buffer)

	slice := buffer[0:10]
	emptyslice := buffer[0:0]
	fmt.Println("initial slice", slice)

	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	fmt.Println("before", slice)
	AddOneToEachElement(slice)

	fmt.Println("after", slice)
	fmt.Println("Before length : len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After length: len(original slice) =", len(slice))
	fmt.Println("After length new slice: len(newSlice) =", len(newSlice))

	// Slice a slice
	slice2 := slice[0:4]
	fmt.Println("sliced slice", slice2)
	fmt.Println("Len slice2", len(slice2))
	PtrSubtractOneFromLength(&slice2)
	fmt.Println("Trimmed slice", len(slice2))

	fmt.Println("Cap of slice 2", cap(slice2))
	fmt.Println("Cap of empty slice", cap(emptyslice))

	pathName := path("/usr/bin/tso") // Conversion from string to path.
	fmt.Printf("%s\n", pathName)

	pathName.TruncateAtFinalSlash()

	fmt.Printf("%s\n", pathName)

	newslice := make([]int, 10, 15)
	fmt.Printf("len: %v, cap: %v\n", len(slice), cap(newslice))

}
