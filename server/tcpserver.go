package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/ivan-golubev/go-chat/model"
	_ "github.com/lib/pq"
)

const DB_CONNECT_STRING = "host=localhost port=5432 user=gochatdbr password=+VHS9q!dle+n dbname=gochat sslmode=disable"

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
			loginid, isAuth, isNew := processauth(message.SignInReq.UserName, message.SignInReq.Password)
			if isAuth == true && isNew == false {
				fmt.Println("Authentication passed. Building sign-in response.")

				resp := &model.SignInResp{
					Status: true,
					UserId: loginid,
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

func processauth(uname, upass string) (userid int32, isAuthenticated, isNewUser bool) {
	// opening db
	db, err := sql.Open("postgres", DB_CONNECT_STRING)
	defer db.Close()
	if err != nil {
		fmt.Printf("Database opening error -->%v\n", err)
		panic("Database error")
	}

	// verifying db connectivity
	err = db.Ping()
	if err != nil {
		fmt.Printf("Database check error -->%v\n", err)
		panic("Database error")
	}

	// checking if user exists
	var uid int32
	err2 := db.QueryRow("SELECT id FROM users WHERE login=$1", uname).Scan(&uid)
	switch {
	case err2 == sql.ErrNoRows:
		fmt.Println("Non-existent user, registration is needed.")
		userid = -1
		isAuthenticated = false
		isNewUser = true
		return
	case err2 != nil:
		checkError(err2)
	default:
		fmt.Println("User is found:%d", uid)
		userid = uid
		isAuthenticated = true
		isNewUser = false
		return
	}
	return
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
