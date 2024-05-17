package main

import (
	"fmt"
	"time"
)

// Task 是一个表示任务的结构体
type Task struct {
	Name     string
	Priority int
}

// Scheduler 创建一个任务调度器，接受一个任务队列，返回一个函数来执行调度
func Scheduler(tasks []Task) func() {
	// 创建一个通道用于接收任务
	taskChannel := make(chan Task)

	// 将任务添加到任务通道
	go func() {
		for _, task := range tasks {
			taskChannel <- task
		}
		close(taskChannel)
	}()

	// 返回一个闭包函数来执行调度
	return func() {
		for task := range taskChannel {
			fmt.Printf("Executing task: %s (Priority: %d)\n", task.Name, task.Priority)
			time.Sleep(time.Duration(task.Priority) * time.Second)
			fmt.Printf("Task %s completed.\n", task.Name)
		}
	}
}

func taskTest() {
	// 定义一些任务
	tasks := []Task{
		{Name: "Task A", Priority: 3},
		{Name: "Task B", Priority: 2},
		{Name: "Task C", Priority: 1},
	}

	// 创建任务调度器
	schedule := Scheduler(tasks)

	// 执行任务调度器
	schedule()
}
