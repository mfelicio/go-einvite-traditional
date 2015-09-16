package ws

import ()

type connectionManager struct {
	nextConnectionId int64
	connections      map[int64]Connection
}

func NewConnectionManager() ConnectionManager {

	return &connectionManager{
		0,
		make(map[int64]Connection),
	}
}

type ConnectionManager interface {
	Handle(conn Connection)

	GetConnection(id int64) (Connection, bool)

	GetNewId() int64
	TotalConnections() int
}

func (this *connectionManager) GetConnection(id int64) (c Connection, ok bool) {

	c, ok = this.connections[id]

	return
}

func (this *connectionManager) TotalConnections() int {

	return len(this.connections)
}

func (this *connectionManager) GetNewId() int64 {
	//TODO synchronize
	this.nextConnectionId++
	return this.nextConnectionId
}

func (this *connectionManager) Handle(conn Connection) {

	conn.SetId(this.GetNewId())

	this.AddConnection(conn)
	defer this.RemoveConnection(conn)

	this.Listen(conn)
}

func (this *connectionManager) AddConnection(conn Connection) {

	this.connections[conn.Id()] = conn
}

func (this *connectionManager) RemoveConnection(conn Connection) {

	delete(this.connections, conn.Id())
}

func (this *connectionManager) Listen(conn Connection) {

	inputStream := conn.GoListen()

	var msg *Message
	var ok bool

	for {

		msg, ok = <-inputStream

		if ok {

			go _processMessage(conn, msg)

		} else {

			break
		}
	}
}

func _processMessage(conn Connection, msg *Message) {
	//conn.onMessage(msg)
}
