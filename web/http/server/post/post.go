package main

import (
	"fmt"
	"net/http"

	"github.com/x14n/goStu/web/http/server/handler"
)

func main() {
	JsonTest()
}

func PostTest() {
	http.HandleFunc("/api/postdata", handler.PostDataHandler)
	http.ListenAndServe(":8080", nil)
}

func JsonTest() {
	http.HandleFunc("/api/json", handler.JsonHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动失败", err)
	}
}
