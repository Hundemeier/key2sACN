package main

import (
	"math/rand"
	"time"

	"github.com/Hundemeier/go-sacn/sacn"
)

var trans sacn.Transmitter

//stores the channels for the univeres
var universeChanMap = make(map[uint16]chan<- [512]byte)

//stores the data byte arrays for the universes
var dataMap = make(map[uint16][512]byte)

func initSacn(conf config) {
	cid := [16]byte{}
	rand.Seed(int64(time.Now().Nanosecond()))
	rand.Read(cid[:])
	transmitter, err := sacn.NewTransmitter("", cid, "key2sACN")
	logErr(err)
	trans = transmitter

	for _, out := range conf.Outputs {
		activateUniverse(out.Universe)
		trans.SetMulticast(out.Universe, out.Multicast)
		trans.SetDestinations(out.Universe, out.Destinations)
	}
}

func activateUniverse(universe uint16) error {
	ch, err := trans.Activate(universe)
	if err != nil {
		return err
	}
	universeChanMap[universe] = ch
	return nil
}

func deactivateUniverse(universe uint16) {
	close(universeChanMap[universe])
	universeChanMap[universe] = nil
}

func sendData(universe uint16, channel uint16, data byte) {
	dmx := dataMap[universe]
	dmx[channel] = data
	if universeChanMap[universe] != nil {
		universeChanMap[universe] <- dmx
	}
	dataMap[universe] = dmx
}

func getDestinations(universe uint16) []string {
	dests := trans.Destinations(universe)
	destString := make([]string, 0)
	for _, dest := range dests {
		destString = append(destString, dest.IP.String())
	}
	return destString
}
