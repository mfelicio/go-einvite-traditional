package ws

import "time"

type Message struct {
	IdFrom int64
	Data   []byte
	Time   time.Time
}
