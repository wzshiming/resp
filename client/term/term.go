package term

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/wzshiming/winseq"
	"golang.org/x/crypto/ssh/terminal"
)

type CmdFunc func(cmd ...string) (string, error)

// Terminal Is a terminal renderer.
type Terminal struct {
	Reader  io.Reader
	Writer  io.Writer
	Prompt  string
	CmdFunc CmdFunc
}

// NewTerminal Create a new Terminal.
func NewTerminal(prompt string, cmd CmdFunc) *Terminal {
	return &Terminal{
		Reader:  os.Stdin,
		Writer:  os.Stdout,
		Prompt:  prompt,
		CmdFunc: cmd,
	}
}

// Run Is run the terminal.
func (c *Terminal) Run() error {
	ter := terminal.NewTerminal(struct {
		io.Reader
		io.Writer
	}{
		c.Reader,
		c.Writer,
	}, "")
	fmt.Fprintln(c.Writer, welcome)
	logger := log.New(c.Writer, "", log.LstdFlags)
	for {
		line, err := ter.ReadPassword(c.Prompt)
		if err != nil {
			if err == io.EOF {
				continue
			}
			return err
		}

		read := csv.NewReader(bytes.NewBufferString(strings.TrimSpace(line)))
		read.Comma = ' '
		read.LazyQuotes = true
		read.TrimLeadingSpace = true
		da, err := read.ReadAll()
		if err != nil {
			logger.Println(err)
			continue
		}
		for _, v := range da {
			beg := time.Now()
			result, err := c.CmdFunc(v...)
			if err != nil {
				logger.Println(err)
				continue
			}
			sub := time.Now().Sub(beg).Truncate(time.Millisecond)
			fmt.Fprintln(ter, result)
			fmt.Fprintf(ter, "(%s)\n", sub)
		}
	}
	return nil
}
