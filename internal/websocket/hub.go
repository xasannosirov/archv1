package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Connection struct {
	WS   *websocket.Conn
	Send chan []byte
}

type Hub struct {
	Connections map[*Connection]bool
	Broadcast   chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Connections: make(map[*Connection]bool),
		Broadcast:   make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		message := <-h.Broadcast
		for conn := range h.Connections {
			select {
			case conn.Send <- message:
			default:
				close(conn.Send)
				delete(h.Connections, conn)
			}
		}
	}
}

func HandleConnection(h *Hub, w http.ResponseWriter, r *http.Request) {
	socketUpgrade := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := socketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	connection := &Connection{
		WS:   conn,
		Send: make(chan []byte),
	}

	h.Connections[connection] = true

	go func() {
		defer func() {
			delete(h.Connections, connection)
			err := connection.WS.Close()
			if err != nil {
				return
			}
		}()

		for {
			_, message, err := connection.WS.ReadMessage()
			if err != nil {
				break
			}
			h.Broadcast <- message
		}
	}()

	go func() {
		for msg := range connection.Send {
			if err := connection.WS.WriteMessage(websocket.TextMessage, msg); err != nil {
				break
			}
		}
	}()
}
