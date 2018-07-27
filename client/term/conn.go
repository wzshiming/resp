package term

import (
	"fmt"

	"github.com/wzshiming/resp"
	"github.com/wzshiming/resp/client"
)

func Run(addr string) error {
	cli, err := client.NewConnect(addr)
	if err != nil {
		return err
	}

	conn, err := Conn(cli)
	if err != nil {
		return err
	}
	
	

	return NewTerminal(fmt.Sprintf("\rRESP %s> ", addr), NewExtra(conn).Cmd).Run()
}

func Conn(cli *client.Connect) (CmdFunc, error) {
	return func(cmd ...string) (string, error) {
		if len(cmd) == 0 {
			return "", nil
		}

		val, err := cli.Cmd(resp.Convert(cmd))
		if err != nil {
			return "", err
		}
		return val.Format(0), nil

	}, nil
}
