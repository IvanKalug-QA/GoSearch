package netsrv

import (
	c "GoSearch/pkg/crawler"
	"GoSearch/utils/file"
	"GoSearch/utils/text"
	"bufio"
	"fmt"
	"net"
	"time"
)

func handler(conn net.Conn, crawler c.Crawler) {
	defer conn.Close()
	defer fmt.Println("Connection closed")

	conn.SetDeadline(time.Now().Add(time.Second * 30))

	r := bufio.NewReader(conn)

	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		output := crawler.GetResult(string(msg))
		if text.CheckResult(output) {
			file.Write(output)
		}
		conn.SetDeadline(time.Now().Add(time.Second * 30))
	}
}

func StartServer(crawler c.Crawler) {
	listener, err := net.Listen("tcp4", ":80")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Accept connect to -> %v\n", conn.RemoteAddr().String())
		go handler(conn, crawler)
	}
}
