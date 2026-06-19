package main

import "fmt"

type WaitGroup struct {
	ch    chan struct{}
	count int
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		ch: make(chan struct{}),
	}
}

func (w *WaitGroup) Add(n int) {
	w.count += n
}

func (w *WaitGroup) Done() {
	w.ch <- struct{}{}
}

func (w *WaitGroup) Wait() {
	for i := 0; i < w.count; i++ {
		<-w.ch
	}
}

func main() {

	wg := NewWaitGroup()
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("горутина %d работает\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("все горутины завершились")
}
