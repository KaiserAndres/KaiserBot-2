package main

import (
	_ "errors"
	"fmt"
	"github.com/thoj/go-ircevent"
	"regexp"
	"strings"
)

type Command interface {
	Validate(string) []string
	Run([]string, *irc.Event)
}

type rollCmd struct {
	regex *regexp.Regexp
}

func (r rollCmd) Validate(s string) []string {
	return r.regex.FindAllString(s, -1)
}

func (r rollCmd) Run(s []string, e *irc.Event) {
	//do rolling stuff
	return
}

type logCmd string

func (l logCmd) Validate(s string) []string {
	roomList := make([]string, 2, 2)
	names := strings.Split(s, " ")[1:]
	if len(names) < 2 return []string{}
	else roomsList[0], roomList[1] = names[0], names[1]

	if roomList[0][0] == '#' return roomlist
	else return []string{}
}

func (l logCmd) Run(s []string, e *irc.Event) {
	room, ses := s[0], s[1]
	if !inLogPile(room) {
		addRoom(room, ses)
	}
	else {
		terminate(room)
	}
}

var commands map[string]Command = make(map[string]Command)

func setCommands() {
	commands["!roll"] = rollCmd{regexp.MustCompile(
		"(\\d+#)?(\\d+d\\d+)((\\+|-)\\d)?")}
	commands["!log"] = logCmd("!log")
}

func badCommand(s string, e *irc.Event) {
	text := "Sorry but '" + s + "' isn't a valid command :("
	send(e.Arguments[0], text, e)
}

func messageHandler(event *irc.Event) {
	fmt.Println(event.Message())
	if strings.HasPrefix(event.Message(), "!") {
		input := strings.Split(event.Message(), " ")
		cmd := commands[input[0]]
		body := strings.Join(input[1:], " ")
		if cmd != nil {
			cmd.Run(cmd.Validate(body), event)
		} else {
			badCommand(input[0], event)
		}
	}
}
