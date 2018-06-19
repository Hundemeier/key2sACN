package main

//DMXAddress stroes inforamtion about an dmx address with universe information
type DMXAddress struct {
	Universe uint16
	Channel  uint16 //has to be between 0 and 511
}

//Key stores the information about a key. the keyboard id and the keycode
type Key struct {
	KeyboardID int
	Key        uint16
}

//KeyMap stores the connection between a key and the corresponding DMXaddress
var keyMap = make(map[Key]DMXAddress)

func initMapping(conf config) {
	//convert the slice of mapTypes to key and dmx addresses
	for _, val := range conf.KeyMap {
		setMapping(val.Universe, val.Channel, val.Keycode, val.KeyboardID)
	}
}

func sendViaMap(event KeyEvent) {
	key := Key{
		KeyboardID: event.DeviceID,
		Key:        event.Code,
	}
	addr := keyMap[key]
	if event.Value == 1 {
		sendData(addr.Universe, addr.Channel, 255)
	} else if event.Value == 0 {
		sendData(addr.Universe, addr.Channel, 0)
	}
}

//setMapping sets the mapping of the key on the keyboard to the desired values
func setMapping(universe uint16, channel uint16, keycode uint16, keyboard int) mapType {
	key := Key{
		Key:        keycode,
		KeyboardID: keyboard,
	}
	dmx := DMXAddress{
		Universe: universe,
		Channel:  channel,
	}
	keyMap[key] = dmx
	return convertToMapType(key, dmx)
}

func deleteMapping(keycode uint16, keyboardID int) bool {
	key := Key{
		Key:        keycode,
		KeyboardID: keyboardID,
	}
	if _, ok := keyMap[key]; !ok {
		return false
	}
	delete(keyMap, key)
	return true
}

func convertToMapType(key Key, dmx DMXAddress) mapType {
	return mapType{
		Universe:   dmx.Universe,
		Channel:    dmx.Channel,
		Keycode:    key.Key,
		KeyboardID: key.KeyboardID,
	}
}
func getKeyMapAsMapType() []mapType {
	list := make([]mapType, 0)
	for key, dmx := range keyMap {
		list = append(list, convertToMapType(key, dmx))
	}
	return list
}
