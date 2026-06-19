package main

import "fmt"

func pipeLine(ch1 <-chan uint8) <-chan float64 {

	ch2 := make(chan float64)

	go func() {
		defer close(ch2)
		for v := range ch1 {
			ch2 <- float64(v * v * v)
		}
	}()

	return ch2
}

func main() {

	ch1 := make(chan uint8)

	go func() {
		defer close(ch1)
		for i := 1; i <= 10; i++ {
			ch1 <- uint8(i)
		}
	}()

	for v := range pipeLine(ch1) {
		fmt.Println(v)
	}

}
