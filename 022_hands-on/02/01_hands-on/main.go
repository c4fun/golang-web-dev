package main

import (
	"log"
	"io"
	"net"
)

func main() {
	// It seems to yield a *TCPListener type
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Cannot Listen: ", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Cannot Accept: ", err)
			continue
		}
		io.WriteString(conn, "I see you connected.")

		conn.Close()
		// 我之前是这下面写错了
		// err = ln.Close()
		// if err != nil {
		// 	log.Fatalln("Cannot Close: ", err)
		// }
	}
}