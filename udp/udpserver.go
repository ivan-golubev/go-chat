package main

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

func check_error(err error) {
    if err != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
    wg := &sync.WaitGroup{}

    /* a channel for messages and channel for quit */
    c := make(chan *model.GenericMessage)
    quit := make(chan int)
    
    go listen(10000, c, quit, wg)
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
    wg.Wait()
}

func listen(port int, messages chan<- *model.GenericMessage, quit <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    server_addr, err1 := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(port))
    check_error(err1)
    conn, err2 := net.ListenUDP("udp", server_addr)
    check_error(err2)
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

        n, _, err3 := conn.ReadFromUDP(buf) // addr is _ for now

        if opErr, ok := err3.(*net.OpError); ok && opErr.Timeout() {
            continue // operation timeout, should not terminate the program
        }

        check_error(err3) 
        message := &model.GenericMessage{}
        err4 := proto.Unmarshal(buf[0:n], message)
        check_error(err4)
        messages <- message        
    }
}

func process_messages(messages <-chan *model.GenericMessage, quit <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    for {
        select { // channel select
                case <- quit:
                fmt.Println("Stopping the message processor...")
                return

                case message := <- messages:
                if (message.Type == model.GenericMessage_TEXT) {
                    timestamp := fmt.Sprint(time.Unix(message.TextMessage.Timestamp, 0))
                    fmt.Println("[Sent " + timestamp + "] Received text message: ", message.TextMessage.Text, " from ")
                }
        }
    }
}