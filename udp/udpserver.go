package main

import (
    "fmt"
    "net"
    "os"
    "github.com/golang/protobuf/proto"
    "github.com/ivan-golubev/go-chat/data-model"
)

func check_error(err error) {
    if err != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
	server_addr, err1 := net.ResolveUDPAddr("udp", ":10000")
    check_error(err1)
    conn, err2 := net.ListenUDP("udp", server_addr)
    check_error(err2)
    defer conn.Close()
 
    buf := make([]byte, 1024)
 
    for {
        n, addr, err3 := conn.ReadFromUDP(buf)        
        check_error(err3) 

        unmarshalled_message := &gochat.TextMessage{}
		err4 := proto.Unmarshal(buf[0:n], unmarshalled_message)
		check_error(err4)

		fmt.Println("Message text: ", unmarshalled_message.Text)
    }
}