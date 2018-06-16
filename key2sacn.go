package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
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
	go server.ListenAndServe()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for _ = range c {
		err := writeConfig()
		logErr(err)
		os.Exit(0)
	}
}
