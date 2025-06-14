// Гриценко Ірина Валеріївна Ex 11.1
package main

import (
	"fmt"
)

type address struct {
	firstname    string
	lastname     string
	addressline1 string
	addressline2 string
	city         string
	state        string
	country      string
	zipcode      int64
}

func printer(User *address) {
	fmt.Println(User.firstname, User.lastname)
	fmt.Println(User.addressline1)
	if User.addressline2 != "" {
		fmt.Println(User.addressline2)
	}
	fmt.Println(User.city+", "+User.state, User.zipcode)
	fmt.Println(User.country)
}

func main() {
	user := new(address)
	user.firstname = "Jon"
	user.lastname = "Hover"
	user.addressline1 = "457 Grove Street"
	//user.addressline2 = "Apartment #15"
	user.city = "Chicago"
	user.state = "IL"
	user.country = "USA"
	user.zipcode = 33458

	printer(user)
}
