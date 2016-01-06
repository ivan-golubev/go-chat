package main

import (
    "fmt"
    "net"
    "os"
)

func check_error(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
	conn, err := net.Dial("udp", "192.168.1.4:10000")
	check_error(err)
	defer conn.Close()

	_, err2 := conn.Write([]byte("Message"))
    check_error(err2)
}