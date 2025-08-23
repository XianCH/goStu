package threadmodel

import (
	"fmt"
	"time"
)

func worker(result chan<- int, jobs <-chan int, id int) {
	fmt.Printf("Worker %d started\n", id)
	for job := range jobs {
		time.Sleep(1 * time.Second)
		result <- job * 2
		fmt.Printf("Worker %d processed job %d\n", id, job)
	}
}

func CspTM() {
	jobs := make(chan int, 10)
	result := make(chan int, 10)

	for i := 0; i < 5; i++ {
		go worker(result, jobs, i)
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 5; i++ {
		res := <-result
		fmt.Printf("Result: %d\n", res)
	}

}
