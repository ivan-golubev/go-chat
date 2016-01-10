package main

import (
	"fmt"    
	"os"
	"strconv"
	"github.com/ivan-golubev/go-chat/udp"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Please specify the destination ip and port. ")
        os.Exit(1)
    }

    address := os.Args[1]    
    port, err := strconv.Atoi(os.Args[2])
    udp.CheckError(err)

    message_id := udp.SendMessage(address, port, "This is the message text", 42)
    fmt.Println("Sent a message with id: " + message_id)
}