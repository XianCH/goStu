package main

import (
	"fmt"
	"time"
)

// 创建一个计数器函数，返回一个函数，每次调用计数器加1
func Counter() func() int {
	count := 0
	// 返回一个闭包函数
	return func() int {
		count++
		return count
	}
}

// 创建一个简单的延迟执行模式，将任务加入到队列，并在指定时间后执行
func DelayExecution(delay time.Duration, task func()) {
	go func() {
		time.Sleep(delay)
		task()
	}()
}

func counterTest() {
	// 创建计数器
	counter := Counter()

	// 使用计数器进行计数
	fmt.Println("Counter 1:", counter()) // 输出：Counter 1: 1
	fmt.Println("Counter 2:", counter()) // 输出：Counter 2: 2
	fmt.Println("Counter 3:", counter()) // 输出：Counter 3: 3

	// 创建一个延迟执行的任务
	DelayExecution(2*time.Second, func() {
		fmt.Println("Delayed task executed")
	})

	fmt.Println("Waiting for delayed task...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
