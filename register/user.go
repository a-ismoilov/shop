package register

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name     string
	Id       string
	Password string
}

func Get() ([]User, error) {
	users := make([]User, 0)
	file, err := os.Open("data/users.txt")
	defer file.Close()
	if err != nil {
		return users, fmt.Errorf("can't open file: %w", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		users = append(users, insert(data))
	}
	return users, nil
}

func insert(data string) User {
	line := strings.Fields(data)
	u := User{
		Name:     line[0],
		Id:       line[1],
		Password: line[2],
	}
	return u
}
