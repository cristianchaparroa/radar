package pools

import "github.com/google/uuid"

type ILocationPool interface {
	Start()
	GetRegisterChannel() chan ILocationClient
	GetUnregisterChannel() chan ILocationClient
	GetBroadcastChannel() chan LocationMessage
}

type LocationPool struct {
	ID string
	clients          map[ILocationClient]bool
	broadcast        chan LocationMessage
	registerClient   chan ILocationClient
	unRegisterClient chan ILocationClient
}

func NewLocationPool() ILocationPool {
	return &LocationPool{
		ID:               uuid.New().String(),
		registerClient:   make(chan ILocationClient),
		unRegisterClient: make(chan ILocationClient),
		broadcast:        make(chan LocationMessage),
		clients:          make(map[ILocationClient]bool),
	}
}

func (p *LocationPool) Start() {

	for {

		select {
		case message := <-p.broadcast:
			p.broadcastMessage(message)
			break

		case client := <-p.registerClient:
			p.register(client)
			break
		case client := <-p.unRegisterClient:
			p.unregister(client)
			break

		}
	}
}

func (p *LocationPool) broadcastMessage(m LocationMessage) []error {
	es := make([]error, 0)

	for c := range p.clients {
		err := c.PublishMessage(m)
		if err != nil {
			es = append(es, err)
		}
	}

	return es
}

func (p *LocationPool) register(c ILocationClient) [] error {
	p.clients[c] = true

	es := make([]error, 0)

	for c:= range p.clients{

		l := c.GetLocation()
		m := NewLocationMessage(l.UserID, l.Latitude, l.Longitude)
		err := c.PublishMessage(m)

		if err != nil {
			es = append(es, err)
		}
	}

	return es
}

func (p *LocationPool) unregister(c ILocationClient) []error{
	es := make([]error, 0)

	delete(p.clients, c)

	return es
}


func (p *LocationPool) GetRegisterChannel() chan ILocationClient {
	return p.registerClient
}

func (p *LocationPool) GetUnregisterChannel() chan ILocationClient {
	return p.unRegisterClient
}

func (p *LocationPool) GetBroadcastChannel() chan LocationMessage {
	return p.broadcast
}