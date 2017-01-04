package hardware

import (
	"strconv"
)

type Command struct {
	Pin    int
	Action string
}

func (c *Command) parse() string {
	var parsed string
	if c.Pin < 10 {
		parsed = "0"
	}
	parsed = parsed + strconv.Itoa(c.Pin)
	parsed = parsed + c.Action
	return parsed
}

func (c *Command) parseBytes() []byte {
	return []byte(c.parse())
}
