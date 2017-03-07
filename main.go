package main

import (
	"errors"
	"fmt"
	"github.com/thoj/go-ircevent"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	badConfig error = errors.New("Bad config file.")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	settings, err := loadSettings()

	if err != nil {
		panic(err.Error())
	}

	conn := irc.IRC(settings["BotNick"], settings["BotNick"])
	err = conn.Connect(settings["Server"])
	if err != nil {
		panic(err.Error())
	}

	conn.AddCallback("PING", ping)
	conn.AddCallback("PRVMSG", messageHandler)
	return
}

func ping(event *irc.Event) {
	event.Connection.SendRaw("PONG" + event.Message())
}

func roll(size, times int) int {
	res := 0
	for i := 0; i < times; i++ {
		res += rand.Int() % size
	}
	return res
}

func loadSettings() (map[string]string, error) {

	var (
		data     string            = ""
		n        int               = -1
		settings map[string]string = make(map[string]string)
		buffer   []byte            = make([]byte, 64)
	)

	file, err := os.Open("settings.txt")
	if err != nil {
		return nil, err
	}

	defer file.Close()
	for n != 0 {
		n, err = file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
		data += string(buffer)
		buffer = make([]byte, 64)
	}

	for _, line := range strings.Split(data, "\n") {
		subs := strings.Split(line, "|")
		if len(subs) != 2 {
			return nil, badConfig
		}

		settings[subs[0]] = subs[1]
	}

	return settings, nil
}
