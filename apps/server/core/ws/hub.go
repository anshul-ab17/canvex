package ws

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.Register:
			roomID := client.RoomID

			room, ok := h.Rooms[roomID]
			if !ok {
				room = NewRoom(roomID)
				h.Rooms[roomID] = room
				go room.Run()
			}

			room.Clients[client] = true
			client.Room = room

		case client := <-h.Unregister:
			if client.Room != nil {
				delete(client.Room.Clients, client)
			}
		}
	}
}
