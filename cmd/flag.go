package main

import "flag"

var search string

func ParseArgs() {
	flag.StringVar(&search, "s", "", "String can be found")

	flag.Parse()
}
