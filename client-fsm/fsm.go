package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "syscall"
    "golang.org/x/crypto/ssh/terminal"
)

func check_error(err error) {
    if err != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
	username, _ := credentials()
	fmt.Printf("\nWelcome %s!\n", username)
}

func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("login: ")
	username, err1 := reader.ReadString('\n')
	check_error(err1)

	fmt.Print("pass: ")
	bytePassword, err2 := terminal.ReadPassword(int(syscall.Stdin))
	check_error(err2)

	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}