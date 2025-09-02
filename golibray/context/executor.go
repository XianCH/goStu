package contextBase

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Task interface {
	Run(ctx context.Context) error
}

type DownloadTask struct {
	URL string
}

func (dt *DownloadTask) Run(ctx context.Context) error {
	log.Printf("Starting download from %s\n", dt.URL)
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Download completed from", dt.URL)
		return nil
	case <-ctx.Done():
		return fmt.Errorf("Download canceled from %s: %v", dt.URL, ctx.Err())
	}
}

type ComputeTask struct {
	JobID int
}

func (ct *ComputeTask) Run(ctx context.Context) error {
	log.Printf("Starting computation for job %d\n", ct.JobID)
	select {
	case <-time.After(3 * time.Second):
		result := 100
		fmt.Printf("Computation completed for job %d, result: %d\n", ct.JobID, result)
		return nil
	case <-ctx.Done():
		return fmt.Errorf("Computation canceled for job %d: %v", ct.JobID, ctx.Err())
	}
}

type Executor struct {
	tasks []Task
}

func NewExecutor(tasks ...Task) *Executor {
	return &Executor{tasks: tasks}
}

func (e *Executor) RunAll(ctx context.Context) {
	var wg sync.WaitGroup
	for _, task := range e.tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			if err := t.Run(ctx); err != nil {
				fmt.Println("Task error:", err)
			}
		}(task)
	}
	wg.Wait()
}

func ExecutorMain() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)

	executor := NewExecutor(
		&DownloadTask{URL: "http://example.com/file1"},
		&ComputeTask{JobID: 42},
	)

	fmt.Println("Starting all tasks...")
	go executor.RunAll(ctx)
	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("All tasks completed or timed out.")
}
