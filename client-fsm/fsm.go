package main

import (
    "fmt"
    "os"
    "os/exec"
    "bufio"
    "strings"
    "syscall"
    "runtime"
    "golang.org/x/crypto/ssh/terminal"
)

type State interface {
	Init()
	LoginSuccess(string, int)
	LoginFailed()
	// FriendInvited()
	Logout()
}

type InitState struct {}
func (this *InitState) Init(){
	username, password := credentials()
	// TODO: perform a TCP call to our central server to fetch the user id
	user_id := 42
	// just a trivial check: actually this should be done on the server
	if (username == "" || password == "") { // if status is success
		fsm.LoginFailed()
	} else {
		fsm.LoginSuccess(username, user_id)
	}
}
func (this *InitState) LoginSuccess(usr_name string, usr_id int){
	fsm = &AuthenticatedState{user_name: usr_name, user_id: usr_id}
	fsm.Init()
}
func (this *InitState) LoginFailed(){
	fmt.Println("\nError: cannot login with the provided credentials.")
}
func (this *InitState) Logout(){}


type AuthenticatedState struct {
	user_name string
	user_id int
}
func (this *AuthenticatedState) Init(){
	clear_cmd()
	fmt.Printf("\nWelcome %s!\n", this.user_name)
}
func (this *AuthenticatedState) LoginSuccess(_ string, _ int){}
func (this *AuthenticatedState) LoginFailed(){}
func (this *AuthenticatedState) Logout(){}

func check_error(err error) {
    if err != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

var fsm State = &InitState{}

func main() {
	fsm.Init()
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

// the below code was borrowed from here: 
// http://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go 
var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["darwin"] = func() { 
        cmd := exec.Command("clear") //os x example, tested by ivan
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cls") //Windows example it is untested, but I think its working 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func clear_cmd() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}