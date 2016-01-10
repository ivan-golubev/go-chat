package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "strconv"
    "github.com/golang/protobuf/proto"
    "github.com/ivan-golubev/go-chat/model"
    "crypto/rand"
)

func check_error(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Please specify the destination ip and port. ")
        os.Exit(1)
    }

    address := os.Args[1]    
    port, err := strconv.Atoi(os.Args[2])
    check_error(err)


    message_id := send_message(address, port, "This is the message text", 42)
    fmt.Println("Sent a message with id: " + message_id)
}

func send_message(ip string, port int, text string, sender_id int32) (string) {
    conn, err := net.Dial("udp", ip + ":" + strconv.Itoa(port))
    check_error(err)
    defer conn.Close()

    message_id := pseudo_uuid()
    message := &model.TextMessage {
        MessageUid: message_id,
        SenderId: sender_id,        
        Timestamp: int64(time.Now().Unix()),
        Text: text,
    }
    wrapper := &model.GenericMessage {
        Type: model.GenericMessage_TEXT,
        TextMessage: message,
    }

    payload, err2 := proto.Marshal(wrapper)
    check_error(err2)

    _, err3 := conn.Write(payload)
    check_error(err3)
    return message_id
}

// borrowed from here: http://stackoverflow.com/questions/15130321/is-there-a-method-to-generate-a-uuid-with-go-language
func pseudo_uuid() (uuid string) {

    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
    return
}