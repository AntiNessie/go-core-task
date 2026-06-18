package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
	}()

	merged := mergeChannels(ch1, ch2)

	var result []int
	for v := range merged {
		result = append(result, v)
	}

	if len(result) != 10 {
		t.Errorf("mergeChannels() length = %d, want 10", len(result))
		return
	}

	for i := 1; i <= 10; i++ {
		found := false
		for _, v := range result {
			if v == i {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("mergeChannels() missing value %d", i)
		}
	}

}
