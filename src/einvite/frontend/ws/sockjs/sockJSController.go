package wssockjs

import (
	"einvite/frontend/ws"
	sockjs "github.com/fzzy/sockjs-go/sockjs"
	"log"
)

var manager = ws.NewConnectionManager()

func HandleSockjsSession(session sockjs.Session) {

	connection := &connection{
		session:     session,
		inputStream: make(chan (*ws.Message)),
	}

	connection.OnClosed(func(reasonCode int) {

		var reason string
		switch reasonCode {
		case 1:
			reason = "closed by client"
		case 2:
			reason = "closed by server"
		case 3:
			reason = "closed by timeout"
		}

		log.Println("Connection", connection.Id(), "closed. Reason:", reason)
	})

	//blocking
	manager.Handle(connection)
}
