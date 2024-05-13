package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type key int

const userKey key = 0

func main() {
	http.HandleFunc("/Hello", HelloHandler)
	log.Println("server start at 12345")
	http.ListenAndServe(":12345", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-user-id")
	ctx := context.WithValue(context.Background(), userKey, userID)
	GreetUser(w, ctx)

}

func GreetUser(w http.ResponseWriter, ctx context.Context) {
	userID := ctx.Value(userKey)

	if userID != nil {
		fmt.Fprintf(w, "Hello, user %s!", userID)
	} else {
		fmt.Fprint(w, "Hello, anonymous user!")
	}
}
