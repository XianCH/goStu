package main

import (
	"fmt"
	"time"
)

func main() {
	bufferChan := make(chan int)
	unBufferChan := make(chan int, 3)

	go func() {
		unBufferChan <- 1
		fmt.Println("unBuffer is block")
	}()

	go func() {
		bufferChan <- 2
		fmt.Println("bufferChanIsNotBlock")
	}()

	time.Sleep(1 * time.Second)

	go func() {
		data := <-unBufferChan
		fmt.Println("Received data from unbuffered channel:", data)
	}()

	go func() {
		data := <-bufferChan
		fmt.Println("Received data from buffer channel:", data)
	}()

	time.Sleep(1 * time.Second)
}
