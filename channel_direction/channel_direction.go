package main

import "fmt"

func sendping(pingchannel chan<- string, msg string) {
	pingchannel <- msg // Add the message to the ping channel
}

func recievepong(pingchannel <-chan string, pongchannel chan<- string) {
	msg := <-pingchannel // take the message from ping channel and store it in msg
	pongchannel <- msg   // move the message from one channel to the other.
}

func main() {
	pingchannel := make(chan string, 1) // For sending
	pongchannel := make(chan string, 1) // For recieving

	sendping(pingchannel, "passed message")
	recievepong(pingchannel, pongchannel)

	fmt.Println(<-pongchannel)

}
