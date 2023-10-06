package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	id      int
	message string
}

func main() {

	dataChannel := make(chan Data, 10)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			data := Data{
				id:      id,
				message: fmt.Sprintf("message from %d", i),
			}
			dataChannel <- data
		}(i)
	}

	wg.Wait()
	// 启动一个协程来处理从通道接收到的数据
	go func() {
		for data := range dataChannel {
			processData(data)
		}
	}()
	close(dataChannel)

	time.Sleep(1 * time.Second)
}

// 处理数据的函数
func processData(data Data) {
	fmt.Printf("Processing data with ID %d: %s\n", data.id, data.message)
	time.Sleep(500 * time.Millisecond) // 模拟处理时间
}
