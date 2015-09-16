package wssockjs

import (
	"einvite/frontend/ws"
	"github.com/fzzy/sockjs-go/sockjs"
	"sync"
	"time"
)

type connection struct {
	id          int64
	session     sockjs.Session
	inputStream chan (*ws.Message)
	closeOnce   sync.Once

	onClosed ws.ClosedHandler
}

func (this *connection) Close(reason string) {

	this.TryClose(ws.CONNECTION_CLOSEDBYSERVER, reason)
}

func (this *connection) TryClose(code int, reason string) {
	this.closeOnce.Do(func() {

		close(this.inputStream)

		if code == ws.CONNECTION_CLOSEDBYSERVER {

			this.session.Close(3000, reason)
		}

		this.onClosed(code)
	})
}

func (this *connection) GoListen() chan (*ws.Message) {

	go func() {
		var data []byte

		id := this.Id()

		for {

			data = this.session.Receive()

			if data != nil {

				this.inputStream <- &ws.Message{id, data, time.Now()}

			} else {

				go this.TryClose(ws.CONNECTION_CLOSEDBYCLIENT, "")
				break
			}
		}
	}()

	return this.inputStream
}

func (this *connection) Send(data string) {

	this.session.Send([]byte(data))
}

func (this *connection) OnClosed(handler ws.ClosedHandler) {

	this.onClosed = handler
}

func (this *connection) Id() int64 {
	return this.id
}

func (this *connection) SetId(id int64) {
	this.id = id
}
