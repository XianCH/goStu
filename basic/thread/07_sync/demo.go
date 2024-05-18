package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go hello()
	wg.Wait()
	fmt.Println("main goroutine done!")
}

func hello() {
	defer wg.Done()
	fmt.Println("Hello Goroutine")
}
