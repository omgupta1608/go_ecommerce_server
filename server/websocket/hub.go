package ws

import (
	t "github.com/omgupta1608/aftershoot_task/types"
)

type hub t.Hub

var H = hub{
	Broadcast: make(chan t.Message),
	Register:  make(chan t.Subscription),
}

func (H *hub) Run() {
	for {
		select {
		case s := <-H.Register:
			if H.Connections == nil {
				H.Connections = make(map[*t.Connection]bool)
				// H.Rooms[s.Room] = connection
			}
			H.Connections[s.Conn] = true

		case m := <-H.Broadcast:
			connections := H.Connections

			for c := range connections {

				select {
				case c.Send <- m.Data:
				default:
					close(c.Send)
					delete(connections, c)

				}
			}
		}
	}
}
