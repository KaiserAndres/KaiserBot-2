package main

import (
	"github.com/thoj/go-ircevent"
	_ "errors"
	"fmt"
	"regexp"
	"strings"
)

type Command interface {
	validate(string) []string
	run([]string, *irc.Event)
}

type rollCmd struct {
	regex *regexp.Regexp
}

func (r rollCmd) validate(s string) []string{
	return r.regex.FindAllString(s, -1)
}

func (r rollCmd) run(s []string, e *irc.Event) {
	//do rolling stuff
	return
}

var commands map[string]Command = make(map[string]Command)

func setCommands() {
	commands["!roll"] = rollCmd{regexp.MustCompile(
		"(\\d+#)?(\\d+d\\d+)((\\+|-)\\d)?")}
}

func badCommand(s string, e *irc.Event) {
	text := "Sorry but '"+s+"' isn't a valid command :("
	send(e.Arguments[0], text, e)
}

func messageHandler(event *irc.Event) {
	fmt.Println(event.Message())
	if strings.HasPrefix(event.Message(), "!") {
		input := strings.Split(event.Message(), " ")
		cmd := commands[input[0]]
		body := strings.Join(input[1:], " ")
		if cmd != nil {
			cmd.run(cmd.validate(body), event)
		} else {
			badCommand(input[0], event)
		}
	}
}
