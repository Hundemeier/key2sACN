package main

import (
	"sync"

	evdev "github.com/gvalkov/golang-evdev"
)

//a set of all devices, that are listend on. Map for concurrent access
var listenedDevices = struct {
	sync.RWMutex
	m map[*evdev.InputDevice]struct{}
}{m: make(map[*evdev.InputDevice]struct{})}

func setListeningDevice(device *evdev.InputDevice, listen bool) {
	listenedDevices.Lock()
	defer listenedDevices.Unlock()

	if listen {
		listenedDevices.m[device] = struct{}{}
		//grab a device, if it should be read from
		device.Grab()
	} else {
		delete(listenedDevices.m, device)
		//release device
		device.Release()
	}
}

func getListeningDevices() []*evdev.InputDevice {
	listenedDevices.RLock()
	defer listenedDevices.RUnlock()
	//create slice with the correct capacity
	list := make([]*evdev.InputDevice, 0, len(listenedDevices.m))
	for device := range listenedDevices.m {
		list = append(list, device)
	}
	return list
}
