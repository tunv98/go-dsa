package concurrency

import (
	"fmt"
	"testing"
	"time"
)

/*
select lets a goroutine wait on multiple channel operations at once
like switch but each case is a channel send or receive
*/

/*
This prints got from ch2 because that goroutine sends first (time sleep)
*/
func learnSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Sending: I'm A"
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "Sending: I'm B"
	}()
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received: %s\n", msg2)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout!")
	}
}

func Test_select(t *testing.T) {
	learnSelect()
}
