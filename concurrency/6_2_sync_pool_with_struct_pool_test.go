package concurrency

import (
	"sync"
	"time"
)

type Request struct {
	ID        string
	userID    int
	Data      map[string]any
	Headers   map[string]string
	Body      []byte
	Timestamp time.Time
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
	Error      error
}

var (
	requestPool = sync.Pool{
		New: func() any {
			return &Request{
				Data:    make(map[string]any),
				Headers: make(map[string]string),
				Body:    make([]byte, 0, 1024), // preallocate 1KB
			}
		},
	}
	responsePool = sync.Pool{
		New: func() any {
			return &Response{
				Headers: make(map[string]string),
				Body:    make([]byte, 0, 1024), // preallocate 1KB
			}
		},
	}
)
