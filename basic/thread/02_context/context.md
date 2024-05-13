## **create**

```go
ctx := context.Background() // 创建一个空的根Context
ctx, cancel := context.WithCancel(ctx) // 创建一个可取消的Context
ctx, cancel := context.WithTimeout(ctx, time.Second * 10) // 创建一个带有超时的Context
ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second * 30)) // 创建一个带有截止时间的Context
ctx := context.WithValue(context.Background(), myKey, "hello")
```

```go
func MyFunc(ctx context.Context) {
    // 使用select语句监听Context的Done通道
    select {
    case <-ctx.Done():
        // 如果Context被取消或者超时，则执行相应的操作
        fmt.Println("Context canceled or timed out")
    default:
        // 如果Context未被取消或者超时，则继续执行相应的操作
        fmt.Println("Context not canceled or timed out")
    }
}
```







## **HTTP请求中的Context**：

在Go语言的net/http包中，也经常使用Context来传递HTTP请求的相关信息，如请求的取消、超时等。

- **创建HTTP请求的Context**：通常在处理HTTP请求时，可以使用`http.Request`的`WithContext()`方法来创建一个带有Context的请求。

```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    // 创建一个带有请求的Context
    ctx := r.Context()
    // 调用其他函数时，传递Context
    SomeFunc(ctx)
}
```

- **取消HTTP请求的Context**：在处理HTTP请求时，如果需要取消请求，可以调用Context的`cancel()`方法。

```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    // 创建一个带有请求的Context
    ctx := r.Context()
    // 创建一个可取消的Context
    ctx, cancel := context.WithCancel(ctx)
    defer cancel() // 在函数退出时取消Context
    // 调用其他函数时，传递Context
    SomeFunc(ctx)
}
```

