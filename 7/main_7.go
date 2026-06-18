package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	result := make(chan int)

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				result <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 4; i <= 6; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 7; i <= 10; i++ {
			ch3 <- i
		}
	}()

	merge := mergeChannels(ch1, ch2, ch3)

	for v := range merge {
		fmt.Println(v)
	}
}
