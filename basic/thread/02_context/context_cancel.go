package main

import (
	"context"
	"fmt"
	"log"
)

var loggerChan chan string

func dowork()

func worker(ctx context.Context) {
}

func startWork(ctx context.Context) {
	for i := 0; i < 10; i++ {
		loggerChan <- fmt.Sprintf("woker:%s start work", i)
		ctx = context.WithValue(ctx, i)
		go worker(ctx)
	}
}

func startCancelTest() {
	loggerChan := make(chan string)
	ctx := context.Background()

	go func() {
		select {
		case logPrint := <-loggerChan:
			log.Println(logPrint)
		case <-ctx.Done():
			log.Println("logger end")
		}
	}()
}
