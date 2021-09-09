package pkg

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/snipextt/lets-talk-server/internal"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWS(rw http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println("Cannot switch protocals!")
	}
	defer ws.Close()
	clientInstance := &internal.Client{
		Conn: ws,
		// brodcastedMessages: make(chan Message, 100),
	}
	go clientInstance.WatchMessages()
}