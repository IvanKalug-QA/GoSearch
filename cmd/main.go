package main

import (
	"GoSearch/utils/spider"
)

func main() {
	ParseArgs()
	spider.RunSpiner(searchParam)
}
