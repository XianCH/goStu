package main

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc 定义了处理 HTTP 请求的函数类型
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Middleware 定义了中间件函数类型
type Middleware func(HandlerFunc) HandlerFunc

// Chain 是中间件链
type Chain struct {
	middlewares []Middleware
}

// Use 添加中间件到链中
func (c *Chain) Use(middleware Middleware) {
	c.middlewares = append(c.middlewares, middleware)
}

// Finalize 将处理函数和中间件链组合
func (c *Chain) Finalize(finalHandler HandlerFunc) HandlerFunc {
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		finalHandler = c.middlewares[i](finalHandler)
	}
	return finalHandler
}

// LoggingMiddleware 日志中间件
func LoggingMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}

// AuthMiddleware 简单的认证中间件
func AuthMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "secret-token" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

// CompressMiddleware 压缩中间件 (示例，不实际压缩)
func CompressMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		next(w, r)
	}
}

// HelloHandler 最终处理器
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// 创建中间件链
	chain := &Chain{}

	// 添加中间件到链中
	chain.Use(LoggingMiddleware)
	chain.Use(AuthMiddleware)
	chain.Use(CompressMiddleware)

	// 最终处理器
	finalHandler := chain.Finalize(HelloHandler)

	// 启动 HTTP 服务器
	http.Handle("/", http.HandlerFunc(finalHandler))
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
