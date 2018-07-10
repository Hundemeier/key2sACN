package main

import (
	"os"
	"regexp"
	"strconv"

	"github.com/Hundemeier/key2sACN/keylogger"

	evdev "github.com/gvalkov/golang-evdev"
)

//KeyEvent is an abstract event for storing the events information
//If Value is 1 then this is a DOWN event, 0 is UP and 2 is REPEATED
type KeyEvent struct {
	KeyboardID int
	KeyCode    uint16
	Value      int32
}

//listening stores wether or not on a given device id is being listened
var listening = make(map[int]bool)

var keylogs = make([]*keylogger.KeyLogger, 0)

//a set of all devices, that are listend on
var listenedDevices = make(map[*evdev.InputDevice]struct{})

//flag for init: if we already started a goroutine for listening on devices set the flag to true
var flagRunning = false

func initKeylogger(conf config) {
	//get all devices and then try to start as man y as possible according to the listener list
	devs, err := keylogger.NewDevices()
	if err != nil {
		logErr(err)
		os.Exit(-1)
	}
	for id, listen := range conf.Listening {
		if listen {
			for _, dev := range devs {
				if dev.Id == id {
					startKeylogger(dev, keyChan)
					break
				}
			}
		}
	}
}

func startKeylogger(device *keylogger.InputDevice, ch chan KeyEvent) (err error) {
	rd := keylogger.NewKeyLogger(device)
	keylogs = append(keylogs, rd)

	in, err := rd.Read()
	if err != nil {
		return
	}

	//only start if we do not have already a listener
	if !listening[device.Id] {
		//function that formattes to the formatted channel
		go func() {
			listening[device.Id] = true
			for i := range in {
				if i.Type == keylogger.EV_KEY {
					event := KeyEvent{
						KeyboardID: device.Id,
						KeyCode:    i.Code,
						Value:      i.Value,
					}
					ch <- event
					setEvent(KEY_EVENT, "", event)
				}
			}
			listening[device.Id] = false
			//delete from list of keyloggers
			deleteKeylogger(rd)
		}()
	}
	return
}

func stopKeylogger(deviceID int) {
	for _, logger := range keylogs {
		if logger.Dev.Id == deviceID {
			logger.Stop()
		}
	}
}

func deleteKeylogger(logger *keylogger.KeyLogger) {
	for i, log := range keylogs {
		if log == logger {
			//delete:
			if i == len(keylogs)-1 {
				keylogs = keylogs[:len(keylogs)]
			} else {
				keylogs = append(keylogs[:i], keylogs[i+1:]...)
			}
		}
	}
}

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
