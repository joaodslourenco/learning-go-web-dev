/*
Building upon the code from the previous problem:
Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.

Add this data to your REPONSE so that this data is displayed in the browser.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	var i int
	var reqMethod, reqURI string

	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			xs := strings.Fields(ln)
			reqMethod = xs[0]
			reqURI = xs[1]
			fmt.Println("METHOD:", reqMethod)
			fmt.Println("URI:", reqURI)
		}
		if ln == "" {
			break
		}
		i++
	}
	body := "This is the response body payload"
	body += "\n"
	body += reqMethod
	body += "\n"
	body += reqURI

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
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
