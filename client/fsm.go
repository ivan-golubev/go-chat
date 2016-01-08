package main

import (
	"fmt"
	"github.com/ivan-golubev/go-chat/console"
)

var fsm State = &InitState{}

func main() {
    fsm.Init()
}

type State interface {
	Init()
	LoginSuccess(string, int)
	LoginFailed()
	// FriendInvited()
	Logout()
}

type InitState struct {}
func (this *InitState) Init(){
	username, password := console.Credentials()
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
	console.Clear_cmd()
	fmt.Printf("\nWelcome %s!\n", this.user_name)
}
func (this *AuthenticatedState) LoginSuccess(_ string, _ int){}
func (this *AuthenticatedState) LoginFailed(){}
func (this *AuthenticatedState) Logout(){}