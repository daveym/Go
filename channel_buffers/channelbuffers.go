// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.

package main

import "fmt"

func main() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 10)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "1"
	messages <- "2"
	messages <- "3"
	messages <- "4"
	messages <- "5"
	messages <- "6"
	messages <- "7"
	messages <- "8"
	messages <- "9"
	messages <- "10"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
