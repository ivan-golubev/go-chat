package main

import (
    "fmt"
    "net"
    "os"
    "github.com/golang/protobuf/proto"
    "github.com/ivan-golubev/go-chat/model"
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

    message := &model.TextMessage {
        MessageUid: "message-id",
        SenderId: 42,        
        Timestamp: 100500,
        Text: "This is the message text",
    }
    wrapper := &model.GenericMessage {
        Type: model.GenericMessage_TEXT,
        TextMessage: message,
    }

    payload, err2 := proto.Marshal(wrapper)
    check_error(err2)

	_, err3 := conn.Write(payload)
    check_error(err3)
}