package main

import (
	"flag"
	"os"
)

var searchParam string

func ParseArgs() {
	flag.StringVar(&searchParam, "s", "", "String can be found")

	flag.Parse()

	if searchParam == "" {
		os.Exit(1)
	}
}
