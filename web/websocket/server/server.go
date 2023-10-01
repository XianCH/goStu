package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connecton upgradation:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err2 := conn.ReadMessage()
		if err2 != nil {
			log.Print("error during message reading", err2)
			break
		}
		log.Print("Received: %s", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index page")
}

func main() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
