package main

import (
	"fmt"
	"net/http"
)

var keyChan = make(chan KeyEvent)

func main() {
	fmt.Println("Reading config file")
	conf := readConfig()
	fmt.Println("Starting...")

	initSacn(conf)
	initKeylogger(conf)
	initMapping(conf)

	//Start the listener for all keys that sends the events via sACN:
	go func() {
		for event := range keyChan {
			//fmt.Println(event)
			sendViaMap(event)
		}
	}()

	fmt.Println("Init grapqhl...")
	initGraphql()
	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
