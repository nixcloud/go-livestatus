package livestatus

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Command represents a Livestatus command instance.
type Command struct {
	name string
	args []string
}

// NewCommand creates a new Livestatus command instance.
func NewCommand(name string, args ...string) *Command {
	return &Command{
		name: name,
		args: args,
	}
}

// Arg appends a new argument to the command.
func (c *Command) Arg(v interface{}) *Command {
	c.args = append(c.args, fmt.Sprintf("%v", v))
	return c
}

// String returns a string representation of the Livestatus command.
func (c Command) String() string {
	s := fmt.Sprintf("COMMAND [%d] %s", time.Now().Unix(), c.name)
	if len(c.args) > 0 {
		s += ";" + strings.Join(c.args, ";")
	}
	s += "\n\n"

	return s
}

func (c Command) handle(conn net.Conn) (*Response, error) {

	cmd := c.String()
	lcmd := len(cmd)

	// Send query data
	n, err := conn.Write([]byte(cmd))
	if err != nil {
		return nil, err
	}

	if n != lcmd {
		return nil, fmt.Errorf("incomplete write to livestatus. Wrote %d bytes while %d were to be written", n, lcmd)
	}

	return nil, nil
}

func (c Command) keepAlive() bool {
	return true
}
