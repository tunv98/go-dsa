package concurrency

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
)

func Test_errorGroup(t *testing.T) {
	g, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 3; i++ {
		g.Go(func() error {
			// Do some work
			if ctx.Err() != nil {
				return ctx.Err()
			}
			if i == 1 || i == 2 {
				return fmt.Errorf("error at %d", i)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}
