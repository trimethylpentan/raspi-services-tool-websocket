package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	upgrader.CheckOrigin = func(request *http.Request) bool {
		return true
	}

	http.HandleFunc("/ws/system-information", handleMessage)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleMessage(writer http.ResponseWriter, request *http.Request) {
	connection, error := upgrader.Upgrade(writer, request, nil)
	if error != nil {
		log.Println(error)
		return
	}

	go loop(connection)
}

func loop(connection *websocket.Conn) {
	for true {
		connection.WriteMessage(websocket.TextMessage, []byte("Test"))
		time.Sleep(5 * time.Second)
	}
}
