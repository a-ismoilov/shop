package user

import (
	"fmt"
	"github.com/google/uuid"
	"olx/register"
	"olx/window"
	"os"
	"unicode"
)

var per register.User

func SignUp() bool {
	fmt.Print("Enter name >>> ")
	if _, err := fmt.Scan(&per.Name); err != nil {
		return false
	}
	for {
		fmt.Print(`
^must contain latin letters
	or
^must contain arab numbers
	or
^length 8 or more

		Enter password >>> `)
		if _, err := fmt.Scan(&per.Password); err != nil {
			return false
		}
		check := checkPassword(per.Password)
		if check {
			per.Id = uuid.New().String()
			if err := WriteUser(per); err != nil {
				fmt.Print("can't write the user")
				return false
			} else {
				fmt.Print("Successfully added!")
				window.App(per.Id)
				return true
			}
		} else {
			print("Try again ⭯")
		}
	}
}

func checkPassword(password string) bool {
	if len(password) < 7 {
		return false
	}
	for _, val := range password {
		if unicode.IsPunct(val) {
			return false
		} else if unicode.IsSpace(val) {
			return false
		}
	}
	users, err := register.Get()
	if err != nil {
		return false
	}
	for _, v := range users {
		if password == v.Password {
			return false
		}
	}
	return true
}

func WriteUser(per register.User) error {
	file, err := os.OpenFile("data/users.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("can't append")
		return err
	}
	u := per.Name
	u += " "
	u += per.Id
	u += " "
	u += per.Password
	if _, err := file.WriteString("\n"); err != nil {
		return err
	}
	if _, err := file.WriteString(u); err != nil {
		return err
	}
	return nil
}
