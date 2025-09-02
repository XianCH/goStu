package contextBase

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Worker Cancel1")
		return
	default:
		fmt.Println("Worker 12js3143working7....")
		time.Sleep(time.Second)
	}

}

func Base01() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 3)
}

func fetchDate(ctx context.Context, ch chan int) {
	sum := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("fetchDate Cancel")
			return
		default:
			sum += 1
			select {
			case ch <- sum:
			case <-ctx.Done():
				return
			}
		}
	}
}

func fetchDataMain() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	arr := make(chan int, 10)

	go fetchDate(ctx, arr)

	for {
		select {
		case val, ok := <-arr:
			if !ok {
				fmt.Println("arr closed")
				return
			}
			fmt.Println("Received:", val)
		case <-ctx.Done():
			fmt.Println("Main Cancel")
			close(arr)
			return
		}
	}

}
