package main

import (
	"sync"
	"testing"
)

func TestIncrement(t *testing.T) {
	counter := &Counter{}

	numWorking := 100
	var wg sync.WaitGroup
	wg.Add(numWorking)
	for i := 0; i < numWorking; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	expected := numWorking
	actual := counter.Value()

	if expected != actual {
		t.Errorf("Increment: Counter value is incorrect, got: %d, want: %d", actual, expected)
	}
}

func TestDecrement(t *testing.T) {
	counter := &Counter{}

	numWorking := 100
	var wg sync.WaitGroup
	wg.Add(numWorking)
	for i := 0; i < numWorking; i++ {
		go func() {
			defer wg.Done()
			counter.Decrement()
		}()
	}

	wg.Wait()

	expected := -numWorking
	actual := counter.Value()

	if expected != actual {
		t.Errorf("Decrement: Counter value is incorrect, got: %d, want: %d", actual, expected)
	}
}
