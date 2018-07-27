package main

import (
	"flag"
	"fmt"

	"github.com/wzshiming/resp/client/term"
)

func main() {
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
