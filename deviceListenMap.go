package main

import (
	"fmt"
	"sync"
	"syscall"

	evdev "github.com/gvalkov/golang-evdev"
)

//a set of all devices, that are listend on. Map for concurrent access
var listenedDevices = struct {
	sync.RWMutex
	m map[*evdev.InputDevice]struct{}
}{m: make(map[*evdev.InputDevice]struct{})}

func setListeningDevice(device *evdev.InputDevice, listen bool) {
	if device == nil {
		return
	}

	//we have to do this here or we have a deadlock because isListening also locks
	alreadyListening := isListening(device)

	listenedDevices.Lock()
	defer listenedDevices.Unlock()

	if listen && !alreadyListening {
		listenedDevices.m[device] = struct{}{}
		//grab a device, if it should be read from
		device.Grab()
	} else if !listen {
		//delete: search the *devices with the correct id and delete them
		for dev := range listenedDevices.m {
			if getID(dev) == getID(device) {
				dev.Release() //release device
				delete(listenedDevices.m, dev)
				dev.File.Close() //close the file, so every i/o is stopped
				syscall.Close(int(dev.File.Fd()))
				fmt.Println("setListen sets false")
			}
		}
	}
}

func setListeningID(id int, listen bool) error {
	//get device with the id
	devices, err := evdev.ListInputDevices()
	if err != nil {
		return fmt.Errorf("could not read the devices list")
	}
	var device *evdev.InputDevice
	for _, dev := range devices {
		if id == getID(dev) {
			device = dev
			break
		}
	}
	//if we found a device invoke the set function:
	if device != nil {
		setListeningDevice(device, listen)
		return nil
	}
	return fmt.Errorf("could not find device with id %v", id)
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

func getListeningID() []int {
	//convert slice of *evdev.InputDevice to slice with listening == true IDs
	listenDev := getListeningDevices()
	listenIDs := make([]int, 0, len(listenDev))
	for _, dev := range listenDev {
		listenIDs = append(listenIDs, getID(dev))
	}
	return listenIDs
}

func isListening(device *evdev.InputDevice) bool {
	if device == nil {
		return false
	}
	listenedDevices.RLock()
	defer listenedDevices.RUnlock()
	//search the set
	for listenDev := range listenedDevices.m {
		if getID(listenDev) == getID(device) {
			return true
		}
	}
	return false
}
