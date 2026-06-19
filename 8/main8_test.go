package main

import (
	"sync/atomic"
	"testing"
)

func TestWaitGroup(t *testing.T) {

	wg := NewWaitGroup()
	wg.Add(3)

	var counter int32
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()

	if counter != 3 {
		t.Errorf("waitGroup(): expected 3 goroutines, have %d", counter)
	}

}
