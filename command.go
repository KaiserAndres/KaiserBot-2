package main

import (
	"github.com/thoj/go-ircevent"
	"errors"
	"regexp"
)

type Command struct {
	prefix string
	regex *regexp.Regexp
}

type Runnable interface {
	run(cmd string)
}

var commands map[string]func

func messageHandler(event *irc.Event) {
	if strings.HasPrefix(event.Message(), "!") {
		command[strings.Split(
			event.Message, " ")[0]](strings.Split(
				event.Message(), " "))
	}
}
