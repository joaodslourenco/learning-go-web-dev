/*
Building upon the code from the previous problem:
Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method, request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in

tags.
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
	body := `
		<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>My Code</title>
				</head>
				<body>
					<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
				</body>
			</html>
	`
	body += "\n"
	body += reqMethod
	body += "\n"
	body += reqURI

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
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
