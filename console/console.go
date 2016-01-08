package console

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

func check_error(err error) {
    if err != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func Credentials() (string, string) {
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

func Clear_cmd() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}