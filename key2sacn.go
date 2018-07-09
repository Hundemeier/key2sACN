package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

var keyChan = make(chan KeyEvent)

func main() {
	port := flag.Uint("port", 8080, "the port on which the webinterface is listening. Only use port 80, when no other application is using this port!")

	flag.Parse()

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
		Addr: fmt.Sprintf(":%v", *port),
	}
	fmt.Println("Serving at:")
	addrs, _ := getMyInterfaceAddr()
	for _, addr := range addrs {
		fmt.Printf("\t%v%v\n", addr, server.Addr)
	}
	fmt.Println("Close with \033[47;30m Ctrl+C \033[0m")
	//http.Handle("/", http.FileServer(http.Dir("./webgui/build")))
	//http.Handle("/", http.FileServer(assetFS()))
	http.Handle("/", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "webgui/dist"}))
	http.HandleFunc("/websocket", handleWebsocket)
	log.Fatal(server.ListenAndServe())
}

func getMyInterfaceAddr() ([]net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	addresses := []net.IP{}
	for _, iface := range ifaces {

		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			addresses = append(addresses, ip)
		}
	}
	if len(addresses) == 0 {
		return nil, fmt.Errorf("no address Found, net.InterfaceAddrs: %v", addresses)
	}
	//only need first
	return addresses, nil
}
