package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"
)

func serve(conn net.Conn) {
	// 因为defer会等到同级都执行完才执行,所以这里可以把关闭作为一个语法糖,写到最前面
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
	}
	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected.")
}



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
	
	go serve(conn)
	}
}
