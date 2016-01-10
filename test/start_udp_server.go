package main

import (
	"fmt"    
	"os"
	"strconv"
	"github.com/ivan-golubev/go-chat/udp"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Please specify the port to listen: ")
        os.Exit(1)
    }

    port, err := strconv.Atoi(os.Args[1])
    udp.CheckError(err)
    udp.StartUdpServer(port)
}