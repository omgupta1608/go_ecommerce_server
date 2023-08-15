package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	t "github.com/omgupta1608/aftershoot_task/types"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// utils.LogError("Error occurred while connecting to room: " + roomId)
		return
	}
	// userCount := len(H.Rooms[roomId])

	c := &t.Connection{
		Ws:   socket,
		Send: make(chan []byte, 256),
	}
	s := t.Subscription{
		Conn: c,
	}

	H.Register <- s

	// if userCount == 1 {
	// 	message := t.JoinMessage{
	// 		UserName: userName,
	// 	}

	// 	m := t.Message{
	// 		Data: message.ConvertToBytes(),
	// 		Room: s.Room,
	// 	}
	// 	H.Broadcast <- m
	// }

	_s := subscription(s)
	go _s.readPump()
	go _s.writePump()
}
