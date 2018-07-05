package main

import (
	"github.com/MarinX/keylogger"
	"github.com/graphql-go/graphql"
)

func queryDevices(p graphql.ResolveParams) (interface{}, error) {
	devs, err := keylogger.NewDevices()
	list := make([]deviceType, 0)
	//strip off the " at the beginning and end of the name (cosmetics)
	for _, val := range devs {
		name := val.Name
		if name[0] == '"' {
			name = name[1:]
		}
		if name[len(name)-1] == '"' {
			name = name[:len(name)-1]
		}
		list = append(list, deviceType{
			Name:      name,
			Id:        val.Id,
			Listening: listening[val.Id],
		})
	}
	return list, err
}

func querySacnOutputs(p graphql.ResolveParams) (interface{}, error) {
	newList := make([]sACNtype, 0)
	list := trans.GetActivated()
	for _, univ := range list {
		item := sACNtype{
			Universe:     univ,
			Multicast:    trans.IsMulticast(univ),
			Destinations: getDestinations(univ),
		}
		newList = append(newList, item)
	}
	return newList, nil
}

func queryMapping(p graphql.ResolveParams) (interface{}, error) {
	list := make([]mapType, 0)
	for key, val := range keyMap {
		item := mapType{
			Universe:   val.Universe,
			Channel:    val.Channel,
			KeyboardID: key.KeyboardID,
			Keycode:    key.Key,
		}
		list = append(list, item)
	}
	return list, nil
}

func queryKeyMapDirty(p graphql.ResolveParams) (interface{}, error) {
	config := readConfig()
	//store copy of keyMap
	keyMap := keyMap

	if len(keyMap) != len(config.KeyMap) {
		return false, nil
	}
	//check content
	for _, val := range config.KeyMap {
		//for every item in config.KEyMap check the item in the keyMap
		//key from keyMap:
		tmp := Key{Key: val.Keycode, KeyboardID: val.KeyboardID}
		dmx, ok := keyMap[tmp]
		if !ok {
			return false, nil
		}
		checkObj := convertToMapType(tmp, dmx)
		if val.Universe != checkObj.Universe ||
			val.Channel != checkObj.Channel ||
			val.Keycode != checkObj.Keycode ||
			val.KeyboardID != checkObj.KeyboardID {
			return false, nil
		}
	}
	return true, nil
}
