package ws

import "github.com/gorilla/websocket"

type Client struct {
	Conn   *websocket.Conn
	Send   chan []byte
	Hub    *Hub
	Room   *Room
	RoomID string
}

func NewClient(conn *websocket.Conn, hub *Hub, roomID string) *Client {
	return &Client{
		Conn:   conn,
		Send:   make(chan []byte),
		Hub:    hub,
		RoomID: roomID,
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		if c.Room != nil {
			c.Room.Broadcast <- msg
		}
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()

	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
