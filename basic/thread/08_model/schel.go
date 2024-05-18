package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task 接口定义了任务的行为
type Task interface {
	Execute() error // 执行任务的方法
}

// Scheduler 接口定义了调度器的行为
type Scheduler interface {
	AddTask(task Task) error // 添加任务到调度器
	Run() error              // 运行调度器
}

// Monitor 接口定义了监控器的行为
type Monitor interface {
	Start() error           // 启动监控器
	Stop() error            // 停止监控器
	GetMonitorData() string // 获取监控数据
}

// Worker 接口定义了工作节点的行为
type Worker interface {
	Run(wg *sync.WaitGroup) // 运行工作节点
}

type TaskImpl struct {
	ID   int
	Data string
}

func (t *TaskImpl) Execute() error {
	// 模拟任务执行
	fmt.Printf("Executing task %d: %s\n", t.ID, t.Data)
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Printf("Task %d completed\n", t.ID)
	return nil
}

type WorkerImpl struct {
	ID       int
	TaskChan chan Task
	Busy     bool
}

func (w *WorkerImpl) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range w.TaskChan {
		task.Execute()
		w.Busy = false
	}
}

type SchedulerImpl struct {
	WorkerPool []*WorkerImpl
	TaskChan   chan Task
	done       chan bool
}

func NewScheduler() *SchedulerImpl {
	return &SchedulerImpl{
		WorkerPool: make([]*WorkerImpl, numWorkers),
		TaskChan:   make(chan Task),
		done:       make(chan bool),
	}
}

func (s *SchedulerImpl) Run() {
	var wg sync.WaitGroup
	wg.Add(len(s.WorkerPool))

	// 启动工作节点
	for i := 0; i < numWorkers; i++ {
		worker := &WorkerImpl{
			ID:       i,
			TaskChan: make(chan Task),
			Busy:     false,
		}
		s.WorkerPool[i] = worker
		go worker.Run(&wg)
	}

	// 监听任务通道和工作节点
	go func() {
		defer wg.Done()
		for {
			select {
			case task := <-s.TaskChan:
				// 分配任务给空闲的工作节点
				for _, worker := range s.WorkerPool {
					if !worker.Busy {
						worker.TaskChan <- task
						worker.Busy = true
						break
					}
				}
			case <-s.done:
				// 关闭所有工作节点
				for _, worker := range s.WorkerPool {
					close(worker.TaskChan)
				}
				return
			}
		}
	}()
}

func main() {
	scheduler := NewScheduler()
	go scheduler.Run()

	// 生成任务
	go func() {
		taskId := 0
		for {
			select {
			case <-time.Tick(taskInterval):
				task := &TaskImpl{
					ID:   taskId,
					Data: fmt.Sprintf("Task %d", taskId),
				}
				scheduler.TaskChan <- task
				taskId++
				if taskId >= maxTasks {
					scheduler.done <- true
					return
				}
			}
		}
	}()

	// 等待调度器完成
	<-scheduler.done
}
