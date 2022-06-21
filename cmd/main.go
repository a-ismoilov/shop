package main

import (
	"fmt"
	"olx/user"
	"olx/window"
)

func main() {
	var selection string
	fmt.Print("If you already have an account SIGN IN or create your account on OLX")
	fmt.Print(`
		1 - SIGN UP
		2 - SIGN IN
		3 - APP
		4 - Exit
	`)
	if _, err := fmt.Scan(&selection); err != nil {
		fmt.Println("can't read input")
		return
	}
	switch selection {
	case "1":
		user.SignUp()
	case "2":
		user.SignIn()
	case "3":
		window.App("admin")
	default:
		return
	}
}
