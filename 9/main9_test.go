package main

import "testing"

func TestPipeLine(t *testing.T) {

	ch1 := make(chan uint8)
	var result []float64
	expected := []float64{1, 8, 27, 64, 125}

	go func() {
		defer close(ch1)
		for i := uint8(1); i <= 5; i++ {
			ch1 <- i
		}
	}()

	for v := range pipeLine(ch1) {
		result = append(result, v)
	}

	if len(result) != 5 {
		t.Errorf("pipeline() lenght = %d, want 5", len(result))
		return
	}

	for v := range result {
		if result[v] != expected[v] {
			t.Errorf("pipeline()[%d] = %f,want %f", v, result[v], expected[v])
		}
	}
}
