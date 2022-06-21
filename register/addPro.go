package register

import (
	"fmt"
	"olx/data"
)

var id, name string

func New() string {

	p := data.Product{}

	if reg() {
		fmt.Print("Enter product name >>> ")
		if _, err := fmt.Scan(&p.Name); err != nil {
			fmt.Print("can't read pro name")
		}
		fmt.Print("Enter product price >>> ")
		if _, err := fmt.Scan(&p.Price); err != nil {
			fmt.Print("can't read pro price")
		}
		for {
			fmt.Print(`Choose quality to your pro
		 ^- Gold
		 ^- Silver
		 ^- Bronze
			`)
			if _, err := fmt.Scan(&p.Type); err != nil {
				fmt.Print("can't read quality")
			}
			if p.Type == "Bronze" || p.Type == "Gold" || p.Type == "Silver" {
				break
			} else {
				fmt.Print("don't be cheater | type quality name")
			}
		}

		p.TraderId = id
		p.TraderName = name

		if err := data.WritePro(p); err != nil {
			fmt.Print("can't write pro")
			return id
		} else {
			fmt.Print("Successfully added pro")
			return id
		}
	} else {
		fmt.Print("can't add")
		return id
	}
}

func reg() bool {
	var (
		name string
		pass string
	)
	fmt.Print("Enter your name >>> ")
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Print("can't read name")
	}
	fmt.Print("Enter your password >>> ")
	if _, err := fmt.Scan(&pass); err != nil {
		fmt.Print("can't read password")
	}
	if check(name, pass) {
		return true
	} else {
		return false
	}
}

func check(nameUser string, password string) bool {
	users, err := Get()
	if err != nil {
		fmt.Print("can't get user")
		return false
	}
	var (
		check int
		//pass string
	)

	for _, v := range users {
		if v.Name == nameUser {
			if v.Password == password {
				id = v.Id
				name = v.Name
				check++
			}
		}
	}
	if check == 1 {
		return true
	}
	return false
}
