package udp

import (
    "fmt"
    "net"
    "os"
    "time"
    "strconv"
    "os/signal"
    "sync"
    "syscall"
    "github.com/golang/protobuf/proto"
    "github.com/ivan-golubev/go-chat/model"
)

func CheckError(err error) {
    if err != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func StartUdpServer(port int, wg *sync.WaitGroup) {
    /* a channel for messages and channel for quit */
    c := make(chan *InputMessage)
    quit := make(chan int)
    
    go listen(port, c, quit, wg)
    go process_messages(c, quit, wg)

    // Handle SIGINT and SIGTERM.
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-ch
        close(quit) // send quit signal to all workers
        os.Exit(1)
    }()

    /* wait for all workers to stop */
    wg.Add(2)
}

type InputMessage struct {
    message *model.GenericMessage
    address string
}

func listen(port int, messages chan<- *InputMessage, quit <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    server_addr, err1 := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(port))
    CheckError(err1)
    conn, err2 := net.ListenUDP("udp", server_addr)
    CheckError(err2)
    defer conn.Close()
 
    // TODO: read message size from udp first
    buf := make([]byte, 1024)
 
    for {
        select {
            case <- quit:
                fmt.Println("Stopping the UDP listener...")
                return
            default:
        }
        conn.SetDeadline(time.Now().Add(1e9))

        n, addr, err3 := conn.ReadFromUDP(buf)

        if opErr, ok := err3.(*net.OpError); ok && opErr.Timeout() {
            continue // operation timeout, should not terminate the program
        }

        CheckError(err3) 
        generic_message := &model.GenericMessage{}
        err4 := proto.Unmarshal(buf[0:n], generic_message)        
        CheckError(err4)
        wrapped_message := &InputMessage {
            message: generic_message,
            address: fmt.Sprint(addr),
        }        
        messages <- wrapped_message
    }
}

func process_messages(messages <-chan *InputMessage, quit <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        select { // channel select
                case <- quit:
                    fmt.Println("Stopping the message processor...")
                    return

                case wrapped_message := <- messages:                    
                    message := wrapped_message.message
                    if (message.Type == model.GenericMessage_TEXT) {
                        timestamp := fmt.Sprint(time.Unix(message.TextMessage.Timestamp, 0))
                        fmt.Println("[Sent " + timestamp + "] Received text message: ",
                         message.TextMessage.Text, " from ", wrapped_message.address)
                    }
        }
    }
}