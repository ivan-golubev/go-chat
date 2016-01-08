// simple server that listens on all NICs, accepts TCP connections,
// prints client's IP and closes connection
package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Please specify the port to listen on %s", os.Args[0])
		os.Exit(1)
	}

	service := ":" + os.Args[1]
	fmt.Fprintf(os.Stdout, "Starting listening on TCP:%s\n", os.Args[1])

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		clientIP := conn.RemoteAddr().String()
		fmt.Fprintf(os.Stdout, "Received a connection from: %s\n", clientIP)
		conn.Write([]byte(clientIP))
		conn.Close()
	}
}
