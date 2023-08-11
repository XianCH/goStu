package main

import (
	"fmt"
	"time"
)

func Hello(){
	fmt.Println("hello")
}

func main() {
	go Hello()
	fmt.Println("goroutine has been done!")
	time.Sleep(time.Second)
}