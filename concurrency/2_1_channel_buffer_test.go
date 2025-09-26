package concurrency

import (
	"fmt"
	"testing"
)

/*
[Buffered] In a goroutine:
Can hold values WITHOUT waiting, up to a capacity
Allow when receive <= capacity of channel
Block when receive > capacity of channel
*/
func sendBufferedOneGoroutineSendNotReceive() {
	ch := make(chan string, 2)
	ch <- "hello"
}

func sendBufferedOneGoroutineReceiveNotSend() {
	ch := make(chan string, 2)
	fmt.Println(<-ch) //deadlock
}

func Test_channel_buffer(t *testing.T) {
	sendBufferedOneGoroutineSendNotReceive()
	//sendBufferedOneGoroutineReceiveNotSend()
}
