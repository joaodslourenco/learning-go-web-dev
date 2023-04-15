/*

Building upon the code from the previous problem:
Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".

Pass the connection of type net.Conn as an argument into this function.

Add "go" in front of the call to "serve" to enable concurrency and multiple connections.

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
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		fmt.Println(ln)
	}
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
		defer conn.Close()

		go serve(conn)

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

	}
}
