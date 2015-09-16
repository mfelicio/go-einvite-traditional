package ws

import ()

type Connection interface {
	Id() int64
	SetId(id int64)
	Close(reason string)

	Send(data string)

	GoListen() chan (*Message)

	OnClosed(handler ClosedHandler)
}

type ClosedHandler func(int)

const (
	CONNECTION_CLOSEDBYCLIENT = 1
	CONNECTION_CLOSEDBYSERVER = 2
	CONNECTION_TIMEOUT        = 3
)
