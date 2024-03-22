package contextBase

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(ctx context.Context, ch chan int, wg *sync.WaitGroup, logchan chan string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			logchan <- "Producer: worker is down...\n"
			return
		default:
			num := rand.Intn(100)
			ch <- num
			logchan <- fmt.Sprintf("Producer: has sent num: %v\n", num)
			time.Sleep(time.Second * 3)
		}
	}
}

func consumer(ctx context.Context, ch chan int, wg *sync.WaitGroup, logchan chan string, mu *sync.Mutex) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			logchan <- "Consumer: worker is down...\n"
			return
		case num := <-ch:
			mu.Lock()
			fmt.Printf("Consumer: received num: %v\n", num)
			mu.Unlock()
			logchan <- fmt.Sprintf("Consumer: has received num: %v\n", num)
		}
	}
}

func logger(ctx context.Context, logchan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Logger: done...\n")
			return
		case logstring, ok := <-logchan:
			if !ok {
				fmt.Printf("Logger: logchan closed\n")
				return
			}
			fmt.Println(logstring)
		}
	}
}

func Demo_01() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	ch := make(chan int)
	logchan := make(chan string)
	defer close(logchan)
	defer close(ch)
	deadline := time.Now().Add(time.Second * 10)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	wg.Add(1)
	go producer(ctx, ch, &wg, logchan)
	wg.Add(1)
	go consumer(ctx, ch, &wg, logchan, &mu)
	wg.Add(1)
	go logger(ctx, logchan, &wg)

	wg.Wait()
	fmt.Println("All jobs done")
	time.Sleep(time.Second * 5)
}

