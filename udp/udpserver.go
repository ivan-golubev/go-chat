package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "strconv"
    "github.com/golang/protobuf/proto"
    "github.com/ivan-golubev/go-chat/model"
)

func check_error(err error) {
    if err != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
    listen(10000)
}

func listen(port int) {
    server_addr, err1 := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(port))
    check_error(err1)
    conn, err2 := net.ListenUDP("udp", server_addr)
    check_error(err2)
    defer conn.Close()
 
    // TODO: read message size from udp first
    buf := make([]byte, 1024)
 
    for {
        n, addr, err3 := conn.ReadFromUDP(buf)        
        check_error(err3) 

        message := &model.GenericMessage{}
        err4 := proto.Unmarshal(buf[0:n], message)
        check_error(err4)
        if (message.Type == model.GenericMessage_TEXT) {
            timestamp := fmt.Sprint(time.Unix(message.TextMessage.Timestamp, 0))
            fmt.Println("[Sent " + timestamp + "] Received text message: ", message.TextMessage.Text, " from ", addr)
        }
    }
}