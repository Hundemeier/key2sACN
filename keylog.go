package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	evdev "github.com/gvalkov/golang-evdev"
)

//KeyEvent is an abstract event for storing the events information
//If Value is 1 then this is a DOWN event, 0 is UP and 2 is REPEATED
type KeyEvent struct {
	KeyboardID int
	KeyCode    uint16
	Value      int32
}

//flag for init: if we already started a goroutine for listening on devices set the flag to true
var flagRunning = false

func initKeylogger(conf config) {
	devs, err := evdev.ListInputDevices()
	if err != nil {
		logErr(err)
		os.Exit(-1)
	}
	for _, device := range devs {
		//simple search in the conf.Lsitening slice for the id
		for _, listenID := range conf.Listening {
			if listenID == getID(device) {
				setListeningDevice(device, true)
				break
			}
		}
	}
	//only start routine, if we do not have another running
	if !flagRunning {
		go listenRoutine()
	}
}

func listenRoutine() {
	flagRunning = true
	for {
		time.Sleep(1 * time.Millisecond)
		//read from all devices, that should be read from
		for _, dev := range getListeningDevices() {
			fmt.Println("Test", time.Now())
			listenToDevice(dev)
		}
	}
	flagRunning = false
}

func listenToDevice(device *evdev.InputDevice) {
	for {
		//This is blocking:
		rawEvents, err := device.Read()
		if err != nil {
			/*if os.IsTimeout(err) {
				continue // simply do nothing, if we had a tiemout
			}*/
			//if we encountered an error, stop listening on that device
			//setListeningDevice(device, false)
			return
		}
		for _, rawEvent := range rawEvents {
			//if we had an event, convert it to KEyEvent if possible
			if event := inputEvent2KeyEvent(rawEvent, device); event != nil {
				keyChan <- *event
				setWebsocketEvent(KEY_EVENT, "", event)
				fmt.Println("Event!", *event)
			}
		}
	}
	fmt.Println("listenToDevice stopped")
}

//getID returns the ID of an inputDevice via the first number in the string.
//If the ID could not be determined, the return value is -1, otherwise a value >= 0
func getID(device *evdev.InputDevice) int {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(device.Fn, 1)
	if len(numbers) <= 0 {
		return -1
	}
	numInt, err := strconv.Atoi(numbers[0])
	if err != nil {
		return -1
	}
	return numInt
}

func inputEvent2KeyEvent(inEv evdev.InputEvent, dev *evdev.InputDevice) (event *KeyEvent) {
	if inEv.Type == evdev.EV_KEY {
		event = &KeyEvent{
			KeyboardID: getID(dev),
			KeyCode:    inEv.Code,
			Value:      inEv.Value,
		}
	}
	return
}
