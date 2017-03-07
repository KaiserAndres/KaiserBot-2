package main

import (
	"github.com/thoj/go-ircevent"
	"errors"
	"regexp"
)

type Command struct {
	prefix string
	regex *regexp.Regexp
	run func(string, irc.Connection)
}

var commands map[string]*Command = make(map[string]*Command)

func messageHandler(event *irc.Event) {
	if strings.HasPrefix(event.Message(), "!") {
		input := strings.Split(event.Message(), " ")
		cmd := command[input[0]]
		if cmd != nil {
			*cmd.run(
				strings.Join(input[1:], " "),
				event.Connection)
		} else {
			badCommand(input[0], event.Connection)
		}
	}
}
