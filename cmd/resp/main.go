package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wzshiming/resp/client/term"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `
Usage of %s:
	resp [address]
`, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}

	err := term.Run(args[0])
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}
}
