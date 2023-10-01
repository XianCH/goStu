package main

import (
	"fmt"
	"sync"
)

func main() {
	var myMap sync.Map

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		myMap.Store(key, i)
	}

	fmt.Println("wating all gonrint down....")

	myMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}
