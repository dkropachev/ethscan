package main

import (
	"fmt"
	"github.com/dkropachev/ethscan/pkg/cli"
	"os"
)

func main() {
	opts := &cli.Options{}
	opts.Parse()
	if err := opts.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err := opts.Run(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
