package main

import (
	"os"
)

type LogRoom struct {
	Name        string
	SessionName string
	Nicks       map[string]string
	File        *os.File
}
