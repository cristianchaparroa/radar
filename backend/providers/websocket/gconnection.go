package websocket

import "github.com/gorilla/websocket"

// IGConnection defines the methods to use throught
// a websocket connection
type IGConnection interface {
	WriteJSON(o interface{}) error
	ReadJSON(o interface{}) error
	Close() error
}

// GConnection is a wrapper  gorilla/websocket.Conn to be
// sure that implements IGConnection
type GConnection struct {
	// Conn is  pointer to websocket.Conn
	Conn *websocket.Conn
}

// NewConnection generates a pointer to GConnection
func NewConnection(c *websocket.Conn) *GConnection {
	return &GConnection{Conn: c}
}

// WriteJSON writes a json in the websocket connection
func (c *GConnection) WriteJSON(o interface{}) error {
	return c.Conn.WriteJSON(o)
}

// ReadJSON reads the data in the websocket connection
func (c *GConnection) ReadJSON(o interface{}) error {
	return c.Conn.ReadJSON(o)
}

// Close ends the websocket connection
func (c *GConnection) Close() error {
	return c.Conn.Close()
}
