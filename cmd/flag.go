package main

import (
	"flag"
)

var searchParam string

func ParseArgs() {
	flag.StringVar(&searchParam, "s", "", "String can be found")

	flag.Parse()

}
