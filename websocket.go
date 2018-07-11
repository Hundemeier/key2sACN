package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var event = struct {
	Type  Type
	Error string
	Data  interface{}
	Time  string
}{ //Default values:
	Type: UNKNOWN,
}

type Type int

const (
	UNKNOWN Type = iota
	KEY_EVENT
	CONFIG_WROTE
	CONFIG_DELETE
)

func setWebsocketEvent(Type Type, err string, data interface{}) {
	new := event
	new.Type = Type
	new.Error = err
	new.Data = data
	new.Time = time.Now().Format(time.RFC3339)
	event = new
}

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { //always anser request no matter what the origin is
		return true //always return true, so that every request is allowed
	},
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logErr(err)
		return
	}
	defer c.Close()
	//store the last event that happended
	last := event
	for {
		time.Sleep(1 * time.Millisecond)
		if last != event {
			err = c.WriteJSON(event)
			if err != nil {
				if websocket.IsCloseError(err) {
					break
				}
				logErr(err)
				break
			}
		}
		last = event
	}
}
