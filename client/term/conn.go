package term

import (
	"fmt"

	"github.com/wzshiming/resp"
	"github.com/wzshiming/resp/client"
)

// Run is run service
func Run(addr string) error {
	cli, err := client.NewConnect(addr)
	if err != nil {
		return err
	}

	return NewTerminal(fmt.Sprintf("RESP %s> ", addr), NewExtra(commands(cli)).Cmd).Run()
}

func commands(cli *client.Connect) CmdFunc {
	return func(cmd ...string) (string, error) {
		if len(cmd) == 0 {
			return "", nil
		}

		d, err := resp.ConvertTo(cmd)
		if err != nil {
			return "", err
		}
		val, err := cli.Cmd(d)
		if err != nil {
			return "", err
		}
		return val.Format(0), nil
	}
}
