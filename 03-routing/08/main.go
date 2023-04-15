/*

Add code to WRITE to the connection.

*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		fmt.Println(ln)
	}
	io.WriteString(conn, "i am writing to the response")
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
			log.Fatalln(err)
		}
		go serve(conn)
	}
}
