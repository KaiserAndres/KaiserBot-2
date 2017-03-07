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

var commands map[string]Command

func messageHandler(event *irc.Event) {
	if strings.HasPrefix(event.Message(), "!") {
		input := strings.Split(event.Message(), " ")
		command[strings.Split(input[0]].run(
			strings.Join(input[1:], " "))
	}
}
