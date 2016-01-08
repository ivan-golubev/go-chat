package main

import (
	"fmt"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/ivan-golubev/go-chat/model"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	clientIP := conn.RemoteAddr().String()
	fmt.Fprintf(os.Stdout, "Received a connection from: %s\n", clientIP)
	var buf [1024]byte
	for {
		n, err := conn.Read(buf[0:])
		checkError(err)

		message := &model.GenericMessage{}
		err2 := proto.Unmarshal(buf[0:n], message)
		checkError(err2)

		if message.Type == model.GenericMessage_SIGN_IN_REQ {
			fmt.Println("Received sign-in request. User:", message.SignInReq.UserName, " Password:", message.SignInReq.Password)
			if message.SignInReq.UserName == "Goga" {
				if message.SignInReq.Password == "Letme1n" {
					fmt.Println("Authentication passed. Building sign-in response.")

					resp := &model.SignInResp{
						Status: true,
						UserId: 1,
						Token:  "ChatForFreeWithNoSMSNorAds",
					}

					wrapper := &model.GenericMessage{
						Type:       model.GenericMessage_SIGN_IN_RESP,
						SignInResp: resp,
					}

					payload, err3 := proto.Marshal(wrapper)
					checkError(err3)

					_, err4 := conn.Write(payload)
					checkError(err4)
					fmt.Println("Authentication response sent.")
				}
			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Please specify the port to listen on %s", os.Args[0])
		os.Exit(1)
	}

	service := ":" + os.Args[1]
	fmt.Fprintf(os.Stdout, "Starting listening on TCP:%s\n", os.Args[1])

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close()
	}
}
