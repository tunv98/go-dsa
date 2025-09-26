package concurrency

import (
	"fmt"
	"testing"
)

func miniTask() {
	ch := make(chan string)
	go func() {
		ch <- "[A] Hello"
	}()
	go func() {
		ch <- "[B] Hello"
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("[C] Received")
}

func Test_miniTask(t *testing.T) {
	miniTask()
}
