package main

import (
	"context"
	"fmt"
	"math/rand"
)

func randomGenerator(ctx context.Context) <-chan int {

	ch := make(chan int)

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- rand.Intn(100):
			}
		}
	}()

	return ch

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := randomGenerator(ctx)

	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println(num)
	}
}
