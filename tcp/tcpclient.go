// tcpclient
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/ivan-golubev/go-chat/model"
)

func check_Error(err error) {
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
	check_Error(err)

	conn, err2 := net.DialTCP("tcp", nil, tcpAddr)
	check_Error(err2)

	message := &model.SignInReq{
		UserName: "Goga",
		Password: "Letme1n",
	}
	wrapper := &model.GenericMessage{
		Type:      model.GenericMessage_SIGN_IN_REQ,
		SignInReq: message,
	}

	payload, err3 := proto.Marshal(wrapper)
	check_Error(err3)

	_, err4 := conn.Write(payload)
	check_Error(err4)

	result, err5 := ioutil.ReadAll(conn)
	check_Error(err5)

	response := &model.GenericMessage{}
	err6 := proto.Unmarshal(result, response)
	check_Error(err6)

	if response.SignInResp.Status == true {
		fmt.Println("Authenticated with id:", response.SignInResp.UserId, " and token:", response.SignInResp.Token)
	} else {
		fmt.Println("Authentication failed!")
	}

	os.Exit(0)
}
