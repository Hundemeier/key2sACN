package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const configFile = "./config.json"

type config struct {
	KeyMap    []mapType
	Listening map[int]bool
	Outputs   []sACNconf
}

type sACNconf struct {
	Universe     uint16
	Multicast    bool
	Destinations []string
}

func writeConfig() error {
	//get the sACN outputs that are currently active
	sACNlist := make([]sACNconf, 0)
	for _, univ := range trans.GetActivated() {
		sACN := sACNconf{
			Universe:     univ,
			Multicast:    trans.IsMulticast(univ),
			Destinations: getDestinations(univ),
		}
		sACNlist = append(sACNlist, sACN)
	}
	conf := config{
		KeyMap:    getKeyMapAsMapType(),
		Listening: listening,
		Outputs:   sACNlist,
	}
	//Write data to file:
	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(configFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

func readConfig() (conf config) {
	conf.KeyMap = make([]mapType, 0)
	conf.Listening = make(map[int]bool)
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}
	json.Unmarshal(file, &conf)
	return
}
