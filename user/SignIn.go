package user

import (
	"fmt"
	"olx/register"
	"olx/window"
)

var u register.User

func SignIn() error {
	fmt.Print("Enter name >>> ")
	if _, err := fmt.Scan(&u.Name); err != nil {
		return err
	}
	fmt.Print("Enter password >>> ")
	if _, err := fmt.Scan(&u.Password); err != nil {
		return err
	}
	users, err := register.Get()
	if err != nil {
		return err
	}
	var check, p bool = false, false
	for i := range users {
		if users[i].Name == u.Name {
			check = true
			if check && u.Password == users[i].Password {
				u = users[i]
				p = true
				fmt.Println("Successfully logged in")
				window.App(u.Id)
				break
			}
		}
	}
	if check && p {
		return err
	}
	return nil
}
