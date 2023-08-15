package types

import "github.com/gorilla/websocket"

type Message struct {
	Data []byte
	// Room string
}

type Connection struct {
	Ws *websocket.Conn
	Send chan []byte
}

type Subscription struct {
	Conn *Connection
}

type Hub struct {
	Connections map[*Connection]bool
	Register chan Subscription
	Broadcast chan Message
}

type SenderBody struct {
	Type string
}