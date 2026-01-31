package ws

type Room struct {
	ID        string
	Clients   map[*Client]bool
	Broadcast chan []byte
}

func NewRoom(id string) *Room {
	return &Room{
		ID:        id,
		Clients:   make(map[*Client]bool),
		Broadcast: make(chan []byte),
	}
}

func (r *Room) Run() {
	for msg := range r.Broadcast {
		for c := range r.Clients {
			c.Send <- msg
		}
	}
}
