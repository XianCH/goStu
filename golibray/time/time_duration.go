package main

import (
	"fmt"
	"time"
)

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

func add() {
	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)
}

func sub() {
	now := time.Now()
	later := now.Add(time.Hour)
	d := now.Sub(later)
	fmt.Println(d)
}

func Equal() {
	now := time.Now()
	later := now.Add(time.Minute)
	b := now.Equal(later)
	fmt.Println(b)
}

func tickDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}
