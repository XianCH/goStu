package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/var", fooHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)

	s := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: fooHandler(),
	}
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "成功连接")
	fmt.Println(r.Method)
	w.Write([]byte("hello"))
}
