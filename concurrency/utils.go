package concurrency

import "fmt"

func hi(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hi %s! (%d)\n", name, i+1)
	}
}
