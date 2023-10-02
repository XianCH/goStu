package main

import (
	"fmt"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r http.Request) {
	fmt.Println(r.Host)
	fmt.Println(r.Header)
	fmt.Println(r.Method)

}

func main() {

}
