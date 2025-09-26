package concurrency

import (
	"fmt"
	"testing"
	"time"
)

/*
Channel is a pipe that connects goroutines, so they can send and receive values safely without extra locks
ch := make(chan int)
+ send: ch <- value
+ receive: v := <- ch
+ close: close(ch) sender side usually closes
*/

/*
[Unbuffered]
*/

/*
In a goroutine: block main goroutine if any case
*/
func sendUnbufferedOneGoroutineSendNotReceive() {
	ch := make(chan string)
	ch <- "hello"     // block
	fmt.Println("Hi") // no run
}

func sendUnbufferedOneGoroutineReceiveNotSend() {
	ch := make(chan string)
	fmt.Println(<-ch) // block
	fmt.Println("Hi") // no run
}

/*
In many goroutine:
+ if A send B not receive -> A block but B continues and exit
+ if A not send B receive -> block in B, A done -> fatal error: all goroutines are asleep - deadlock!
+ Receive blocks until another goroutine sends
*/
func sendUnbufferedManyGoroutine() {
	ch := make(chan string)
	go func() {
		fmt.Println("[A]: Sending:", 88)
		ch <- "hello"
		fmt.Println("[A]: Sent!")
	}()
	time.Sleep(time.Second)
	fmt.Println("[B]: Receiving")
	msg := <-ch
	fmt.Println("[B]: Got", msg)
}

func Test_Channel(t *testing.T) {
	//sendUnbufferedOneGoroutineSendNotReceive()
	//sendUnbufferedOneGoroutineReceiveNotSend()
	sendUnbufferedManyGoroutine()
}
