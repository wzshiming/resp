package client

import (
	"net"

	"github.com/wzshiming/resp"
)

type Connect struct {
	decoder *resp.Decoder
	encoder *resp.Encoder
}

func NewConnect(address string) (*Connect, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Connect{
		decoder: resp.NewDecoder(conn),
		encoder: resp.NewEncoder(conn),
	}, nil
}

func (c *Connect) Cmd(r resp.Reply) (resp.Reply, error) {
	err := c.encoder.Encode(r)
	if err != nil {
		return nil, err
	}
	return c.decoder.Decode()
}
