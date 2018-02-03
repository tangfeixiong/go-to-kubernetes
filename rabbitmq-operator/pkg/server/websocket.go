package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/tangfeixiong/go-to-kubernetes/rabbitmq-operator/pb"
)

//var upgrader = websocket.Upgrader{} // use default options

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ctl *controller) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ws handler: %v", r)

	header := http.Header{
		"Access-Control-Allow-Origin":  []string{"*"},
		"Access-Control-Allow-Methods": []string{"GET, POST, PUT, PATCH, DELETE"},
	}

	c, err := upgrader.Upgrade(w, r, header)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %+v", string(message))

		content := new(pb.CrdRecipient)
		err = json.Unmarshal(message, content)
		if err != nil {
			log.Println("unmarshal:", err)
			break
		}
		log.Printf("type: %v, content: %+v", mt, content)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
