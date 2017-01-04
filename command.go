package main

import (
	"strconv"
)

type Command struct {
	Pin    uint16
	Action string
}

func (c *Command) parse() string {
	parsed := strconv.FormatUint(uint64(c.Pin), 10)
	parsed = parsed + c.Action
	return parsed
}
