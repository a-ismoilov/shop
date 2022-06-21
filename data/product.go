package data

import (
	"bufio"
	"fmt"
	//"olx/user"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Name     string
	TraderId string
	Price    uint
	Type     string
}

var products []Product

func Get(id string) ([]Product, error) {
	file, err := os.Open("data/products.txt")
	defer file.Close()
	if err != nil {
		return products, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		products = append(products, convert(data, id))
	}
	return products, err
}

func convert(data string, id string) Product {
	p := Product{}
	line := strings.Fields(data)
	p.TraderId = id
	p.Name = line[0]
	p.Type = line[3]
	str, err := strconv.Atoi(line[2])
	p.Price = uint(str)
	if err != nil {
		fmt.Println("can't read file")
		return p
	}
	return p
}

func WritePro(pro Product) error {
	file, err := os.OpenFile("data/products.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("can't append")
		return err
	}
	u := pro.Name
	u += " "
	u += pro.TraderId
	u += " "
	str := strconv.Itoa(int(pro.Price))
	u += str
	u += " "
	u += pro.Type
	if _, err := file.WriteString("\n"); err != nil {
		return err
	}
	if _, err := file.WriteString(u); err != nil {
		return err
	}
	return nil
}

//func SignUp() {
//	user.SignUp()
//}
//func SignIn() error {
//	if err := user.SignIn(); err != nil {
//		return err
//	}
//	return nil
//}
