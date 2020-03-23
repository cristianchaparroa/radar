package pools

import (
	"radar/dataprovider/websocket"
	"radar/entities"
)

type ILocationClient interface {
	Read()
	PublishMessage(m interface{}) error
}

type LocationClient struct {
	ID         string
	connection websocket.IGConnection
	pool       ILocationPool
	profile    *entities.Profile
	location   *entities.Location
}

func NewLocationClient() ILocationClient {
	return &LocationClient{}
}
