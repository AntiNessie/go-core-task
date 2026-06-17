package main

import (
	"context"
	"testing"
	"time"
)

func TestRandomGenerator(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ch := randomGenerator(ctx)

	for i := 0; i < 10; i++ {
		num := <-ch
		if num >= 100 || num < 0 {
			t.Errorf("randomGenerator() = %d, want 0-99", num)
		}
	}

}

func TestRandomGenerator_ClosedCh(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	ch := randomGenerator(ctx)
	<-ch
	cancel()
	time.Sleep(1 * time.Second)

	_, ok := <-ch
	if ok {
		t.Errorf("randomGenerator() - канал должен быть закрыт после контекста")
	}

}
