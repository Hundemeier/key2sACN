package main

import (
	"fmt"

	"github.com/Hundemeier/key2sACN/keylogger"
	"github.com/graphql-go/graphql"
)

func mutateSacnOutput(p graphql.ResolveParams) (sacn interface{}, err error) {
	// marshall and cast the argument value
	univInt, _ := p.Args["universe"].(int)
	univ, err := checkUniverse(univInt)
	if err != nil {
		return
	}
	multicast, _ := p.Args["multicast"].(bool)
	destinations := make([]string, 0) //convert []interface{}
	if p.Args["destinations"] != nil {
		for _, dest := range p.Args["destinations"].([]interface{}) {
			destinations = append(destinations, dest.(string))
		}
	}

	if !trans.IsActivated(univ) {
		err = activateUniverse(univ)
		if err != nil {
			return
		}
	}
	trans.SetMulticast(univ, multicast)
	errs := trans.SetDestinations(univ, destinations) //only get the first error
	if len(errs) != 0 {
		err = errs[0]
	}
	return sACNtype{
		Universe:     univ,
		Multicast:    trans.IsMulticast(univ),
		Destinations: getDestinations(univ),
	}, err
}

func mutateStopSacn(p graphql.ResolveParams) (interface{}, error) {
	univInt, _ := p.Args["universe"].(int)
	univ, err := checkUniverse(univInt)
	if err != nil {
		return false, err
	}

	if _, ok := universeChanMap[univ]; ok {
		close(universeChanMap[univ])
		universeChanMap[univ] = nil
		return true, nil
	}
	return false, nil
}

func mutateKeyMap(p graphql.ResolveParams) (interface{}, error) {
	univInt, _ := p.Args["universe"].(int)
	univ, err := checkUniverse(univInt)
	if err != nil {
		return nil, err
	}
	chanInt, _ := p.Args["channel"].(int)
	channel, err := checkChannel(chanInt)
	if err != nil {
		return nil, err
	}
	keyCodeInt, _ := p.Args["keycode"].(int)
	keyCode := uint16(keyCodeInt)
	keyboardID, _ := p.Args["keyboardID"].(int)
	return setMapping(univ, channel, keyCode, keyboardID), nil
}

func mutateKeyListener(p graphql.ResolveParams) (interface{}, error) {
	listen := p.Args["listen"].(bool)
	deviceID := p.Args["deviceID"].(int)
	if listen {
		//get the device with the id
		devs, err := keylogger.NewDevices()
		if err != nil {
			return nil, err
		}
		var dev *keylogger.InputDevice
		for _, val := range devs {
			if val.Id == deviceID {
				dev = val
			}
		}
		if dev == nil {
			return nil, fmt.Errorf("could not find device with id %v", deviceID)
		}
		startKeylogger(dev, keyChan)
		return deviceType{
			Listening: true,
			Id:        dev.Id,
			Name:      dev.Name,
		}, nil
	}
	stopKeylogger(deviceID)
	return nil, nil
}

func mutateWriteConfig(p graphql.ResolveParams) (interface{}, error) {
	err := writeConfig()
	if err == nil {
		return true, nil
	}
	return false, err
}

func mutateDeleteConfig(p graphql.ResolveParams) (interface{}, error) {
	err := deleteConfig()
	if err == nil {
		return true, nil
	}
	return false, err
}

func mutateDeleteKeyMap(p graphql.ResolveParams) (interface{}, error) {
	keycode := p.Args["keycode"].(int)
	keyboardID := p.Args["keyboardID"].(int)
	ok := deleteMapping(uint16(keycode), keyboardID)
	return ok, nil
}
