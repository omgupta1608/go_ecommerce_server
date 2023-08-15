package ws

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/omgupta1608/aftershoot_task/types"
)

var addr = flag.String("addr", "localhost:5000", "http service address")

func SendWsMessage(msg types.SenderBody) {
	msgStr, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	fmt.Println(string(msgStr))
	err = c.WriteMessage(websocket.TextMessage, []byte(string(msgStr)))

	if err != nil {
		log.Println("error:", err)
		return
	}

	c.Close()
}
