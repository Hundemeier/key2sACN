package main

import (
	"fmt"
	"log"
)

func checkUniverse(universe int) (uint16, error) {
	if universe < 0 || universe > 64000 {
		return 0, fmt.Errorf("the given universe %v was not in range [0-63999]", universe)
	}
	return uint16(universe), nil
}

func checkChannel(channel int) (uint16, error) {
	if channel < 0 || channel > 511 {
		return 0, fmt.Errorf("the given channel %v was not in range [0-511]", channel)
	}
	return uint16(channel), nil
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
