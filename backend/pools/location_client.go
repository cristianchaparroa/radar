package pools

import (
	"fmt"
	"github.com/google/uuid"
	"radar/entities"
	"radar/providers/websocket"
)

type ILocationClient interface {
	Read()
	GetLocation() *entities.Location
	PublishMessage(m interface{}) error
}

type LocationClient struct {
	ID         string
	connection websocket.IGConnection
	pool       ILocationPool
	location   *entities.Location
}

func NewLocationClient(pool ILocationPool, gconnection websocket.IGConnection, location *entities.Location) ILocationClient {

	return &LocationClient{
		ID:         uuid.New().String(),
		pool:       pool,
		connection: gconnection,
		location:   location,
	}
}

func (c *LocationClient) PublishMessage(m interface{}) error {
	return nil
}

func (c *LocationClient) GetLocation() *entities.Location {
	return c.location
}

func (c *LocationClient) Read() {

	defer func() {
		unregisterChannel := c.pool.GetUnregisterChannel()
		unregisterChannel <- c
		c.connection.Close()
	}()

	for {
		var message LocationMessage
		err := c.connection.ReadJSON(&message)

		if err != nil {
			fmt.Println(err)
		}
		broadcastChannel := c.pool.GetBroadcastChannel()
		broadcastChannel <- message
	}
}
