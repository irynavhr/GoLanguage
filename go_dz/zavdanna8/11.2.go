// Гриценко Ірина Валеріївна Ex 11.2
package main

import (
	"fmt"
)

type Address struct {
	full_name    name
	Addressline1 string
	Addressline2 string
	city         string
	state        string
	country      string
	zipcode      int64
}

type name struct {
	firstname  string
	lastname   string
	middlename string
}

func print_please(User *Address) {
	fmt.Println(User.full_name.firstname, User.full_name.middlename, User.full_name.lastname)
	fmt.Println(User.Addressline1)
	if User.Addressline2 != "" {
		fmt.Println(User.Addressline2)
	}
	fmt.Println(User.city+", "+User.state, User.zipcode)
	fmt.Println(User.country)
}

func main() {
	// n_u := name{"Iryna", "Hrytsenko", "Valeriivna"}
	// u:= Address{}
	// u.full_name = n_u
	// u.Addressline1 = "457 Grove Street"
	// u.Addressline2 = "Apartment #15"
	// u.city = "Chicago"
	// u.state = "IL"
	// u.country = "USA"
	// u.zipcode = 33458

	n_u := new(name)
	n_u.firstname = "Iryna"
	n_u.lastname = "Hrytsenko"
	n_u.middlename = "Valeriivna"

	user := new(Address)
	user.full_name = *n_u
	user.Addressline1 = "457 Grove Street"
	//user.Addressline2 = "Apartment #15"
	user.city = "Chicago"
	user.state = "IL"
	user.country = "USA"
	user.zipcode = 33458

	print_please(user)
}
