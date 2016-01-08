// tcpclient
package main

import (
	"fmt"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/ivan-golubev/go-chat/data-model"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	message := &gochat.SignInReq{
		user_name: "Goga",
		password:  "Letme1n",
	}
	wrapper := &gochat.GenericMessage{
		Type:      gochat.GenericMessage_SIGN_IN_REQ,
		SignInReq: message,
	}

	payload, err := proto.Marshal(wrapper)
	check_error(err)

	_, err := conn.Write(payload)
	check_error(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	response := &gochat.GenericMessage{}
	err := proto.Unmarshal(result, response)
	checkError(err)

	if response.SignInResp.status == true {
		fmt.Println("Authenticated with id: ", response.SignInResp.user_id)
	}
	else {
		fmt.Println("Authentication failed!")
	}

	os.Exit(0)
}
