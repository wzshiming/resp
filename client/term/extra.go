package term

import (
	"os"
	"strings"
)

// Extra is Additional terminal commands
type Extra struct {
	Map   map[string]CmdFunc
	Other CmdFunc
}

func (e *Extra) Cmd(cmd ...string) (string, error) {
	if len(cmd) == 0 {
		return "", nil
	}
	fun := e.Map[strings.ToLower(cmd[0])]
	if fun != nil {
		return fun(cmd...)
	}
	return e.Other(cmd...)
}

func (e *Extra) AddCmd(name string, fun CmdFunc) {
	e.Map[strings.ToLower(name)] = fun
}

// NewExtra Create a new extra.
func NewExtra(other CmdFunc) *Extra {
	e := &Extra{
		Map:   map[string]CmdFunc{},
		Other: other,
	}
	e.AddCmd("quit", e.quit)
	return e
}

func (e *Extra) quit(cmd ...string) (string, error) {
	defer os.Exit(0)
	return e.Other(cmd...)
}
