package main

import (
	"fmt"    
	"os"
	"github.com/golang/protobuf/proto"
	"github.com/ivan-golubev/go-chat/data-model"
)

func check_error(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
	message := &gochat.TextMessage {
		MessageUid: "message-id",
		SenderId: 42,
		SenderAddr: "127.0.0.1",
		Timestamp: 100500,
		Text: "This is the message text",
	}
	payload, err := proto.Marshal(message)
	check_error(err)

	unmarshalled_message := &gochat.TextMessage{}
	err = proto.Unmarshal(payload, unmarshalled_message)
	check_error(err)

	fmt.Println("Message text: ", unmarshalled_message.Text)
}