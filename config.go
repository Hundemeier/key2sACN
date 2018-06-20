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

func writeConfig() (err error) {
	//function for reporting an error or success on terminating the function
	defer func() {
		if err != nil {
			setEvent(CONFIG_WROTE, err.Error(), false)
		} else {
			setEvent(CONFIG_WROTE, "", true)
		}
	}()

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
		return
	}
	f, err := os.OpenFile(configFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer f.Close()
	f.Truncate(0) //clear existing data in the file
	_, err = f.Write(data)
	return
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

func deleteConfig() (err error) {
	err = os.Remove(configFile)
	if err != nil {
		setEvent(CONFIG_DELETE, err.Error(), false)
		return
	}
	setEvent(CONFIG_DELETE, "", true)
	return
}
