package main

import (
	"reflect"

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

func queryDirty(p graphql.ResolveParams) (interface{}, error) {
	config := readConfig()
	retrnVal := dirtyType{
		KeyMapDirty:    false,
		SacnDirty:      false,
		ListeningDirty: false,
	}
	//Check KeyMap:----------------------------------------------
	//store copy of keyMap
	keyMap := keyMap

	if len(keyMap) != len(config.KeyMap) {
		retrnVal.KeyMapDirty = true
	} else {
		//check content
		for _, val := range config.KeyMap {
			//for every item in config.KEyMap check the item in the keyMap
			//key from keyMap:
			tmp := Key{Key: val.Keycode, KeyboardID: val.KeyboardID}
			dmx, ok := keyMap[tmp]
			if !ok {
				retrnVal.KeyMapDirty = true
				break
			}
			checkObj := convertToMapType(tmp, dmx)
			if val.Universe != checkObj.Universe ||
				val.Channel != checkObj.Channel ||
				val.Keycode != checkObj.Keycode ||
				val.KeyboardID != checkObj.KeyboardID {
				retrnVal.KeyMapDirty = true
				break
			}
		}
	}
	//Check sACN setup:------------------------------------------
	sACNsetup := getSacnCurrentSetup()
	sACNconfig := config.Outputs //store copy of output config
	if len(sACNsetup) != len(sACNconfig) {
		retrnVal.SacnDirty = true
	} else {
		//check every entry
	OuterLoop:
		for _, val := range sACNsetup {
			//check if the entry on the current setting also exists with the same values in the config
		InnerLoop:
			for _, valConf := range sACNconfig {
				if val.Multicast == valConf.Multicast ||
					val.Universe == valConf.Universe {
					if len(val.Destinations) != len(valConf.Destinations) {
						continue InnerLoop
					}
					//check if the destination lists are containing the same
					for _, destVal := range val.Destinations {
						contains := false
						for _, destValConf := range valConf.Destinations {
							if destVal == destValConf {
								contains = true
								break
							}
						}
						if !contains {
							continue InnerLoop
						}
					}

					//if the val valConf are the same:
					continue OuterLoop
				}
			}
			//if we are here, nothing could be found, so -> dirty
			retrnVal.SacnDirty = true
			break
		}
	}
	//Check Listening:-------------------------------------------
	listeningConf := config.Listening
	listeningCurrent := listening
	if !reflect.DeepEqual(listeningConf, listeningCurrent) {
		retrnVal.ListeningDirty = true
	}

	return retrnVal, nil
}
