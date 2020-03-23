package pools

type ILocationPool interface {
}

type LocationPool struct {
	ID        string
	Broadcast chan LocationMessage
}

func NewLocationPool() ILocationPool {
	return &LocationPool{}
}
