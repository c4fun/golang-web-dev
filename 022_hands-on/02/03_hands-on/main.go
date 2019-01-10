package main

import (
	"io"
	"log"
	"net"
	"bufio"
	"fmt"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// read from connection
		scanner := bufio.NewScanner(conn)
		// my solution
		// fmt.Println(scanner.Bytes())
		// Todd's solution 
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}
		defer conn.Close()

		// write to connection
		// we never get here because of the 29th line scanning loop of open stream
		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}
