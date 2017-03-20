package main

import (
	"errors"
	"github.com/thoj/go-ircevent"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	badConfig error       = errors.New("Bad config file.")
	sendMutex *sync.Mutex = &sync.Mutex{}
)

func main() {
	rand.Seed(time.Now().UnixNano())
	settings, err := loadSettings()

	if err != nil {
		panic(err.Error())
	}

	conn := irc.IRC(settings["BotNick"], settings["BotNick"])

	conn.VerboseCallbackHandler = false
	conn.Debug = false
	conn.UseTLS = false

	conn.AddCallback("PING", ping)
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		go messageHandler(event)
	})
	conn.AddCallback("001", func(e *irc.Event) {
		for _, room := range strings.Split(settings["Channels"], ",") {
			e.Connection.Join(room)
			send(room, "This is a testing bot, please ignore", e)
		}
	})

	setCommands()
	err = conn.Connect(settings["Server"])
	if err != nil {
		panic(err.Error())
	}
	conn.Loop()
}

func ping(event *irc.Event) {
	event.Connection.SendRaw("PONG " + event.Message())
}

func roll(size, times int) int {
	res := 0
	for i := 0; i < times; i++ {
		res += rand.Int() % size
	}
	return res
}

func send(room, message string, e *irc.Event) {
	sendMutex.Lock()
	e.Connection.Privmsg(room, message)
	sendMutex.Unlock()
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

	conf := strings.Split(data, "\n")

	for n, line := range conf {
		subs := strings.Split(line, "|")
		if n == len(conf)-1 {
			continue
		}

		if len(subs) != 2 {
			return nil, badConfig
		}

		settings[subs[0]] = subs[1]
	}

	return settings, nil
}
