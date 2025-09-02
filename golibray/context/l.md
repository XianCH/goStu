```go
//with time out
ctx := context.Background()
ctx , cancel := context.WithCancel(context.Background())
ctx , cancel := context.WithTimeout(context.Background(),time.Second * 5)

//create context with deadline
deadline := time.Now.Add(time.Second * 10)
ctx , cancel := context.WithDeadline(context.Background(),deadline)
defer cancel()

// pass value by context
ctxWithValue := contextWithValue(ctx,key,value)
```

## 01_demo 
我们将创建一个简单的生产者-消费者模式，其中生产者将随机生成数字，并将其发送到通道中，而消费者则从通道中接收数字，
并进行相应的处理。在这个过程中，我们将使用 context 进行任务的取消，并使用互斥锁保护共享资源（日志记录）。



