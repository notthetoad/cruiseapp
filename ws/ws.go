package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan string
}

func (c *Client) handleConn() {
	for {
		select {
		case message := <-c.send:
			nw, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.hub.unregister <- c
				return
			}
			_, err = nw.Write([]byte(message))
			if err != nil {
				c.hub.unregister <- c
				return
			}
			if err := nw.Close(); err != nil {
				c.hub.unregister <- c
				return
			}
		}
	}

}

type Hub struct {
	broadcast  chan string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func ServeWs(w http.ResponseWriter, r *http.Request, hub *Hub) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan string)}
	client.hub.register <- client

	go client.handleConn()
}
