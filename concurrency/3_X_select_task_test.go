package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
Goal: Simulate querying 3 “servers” in parallel and print whichever one responds first.
Requirements:
Create three goroutines.
Each should wait for a random time between 100ms and 1s.
Then send a message to a shared channel like "server 1 done".

In main, use select to:
Receive the first message to arrive.
If no message arrives within 600ms, print "timeout".
End the program after printing either the first response or the timeout message.
*/

func fastestResponse() {
	ch := make(chan string)
	for i := 1; i <= 3; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(88)))
			ch <- fmt.Sprintf("server %d done", i)
		}(i)
	}
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(600 * time.Millisecond):
		fmt.Println("timeout!")
	}
}

func Test_selectTask(t *testing.T) {
	fastestResponse()
}
