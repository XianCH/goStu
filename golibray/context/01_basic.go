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
