```go
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
