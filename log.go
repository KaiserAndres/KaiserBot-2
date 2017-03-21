package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

type LogRoom struct {
	Name        string
	SessionName string
	Nicks       map[string]string
	Date        time.Time
	File        *os.File
}

var logPile []*LogRoom = []*LogRoom{}

func AddRoom(room, ses string) {
	// Adds a room to the log pile, the log itself will keep going
	// until it's terminated. The log will be saved under the file
	// <room.Name>_<room.SessionName>.log in the directory logs

	var err error

	toListen := &LogRoom{
		Name:        room,
		SessionName: ses,
		Date:        time.Now(),
	}

	fileName, err := filepath.Abs(toListen.Name + toListen.SessionName +
		toListen.Date.Format(time.RFC3339) + ".log")

	if err != nil {
		log.Println(err)
	}

	toListen.File, err = os.Open(fileName)

	defer toListen.File.Close()

	if err != nil {
		log.Println(err)
		return
	}
	logPile = append(logPile, toListen)

	return
}

func Terminate(room string) {
	// Removes a room from the log pile and terminates the loggin
	// of said room.

	return
}

func InLogPile(room string) bool {
	return false
}
