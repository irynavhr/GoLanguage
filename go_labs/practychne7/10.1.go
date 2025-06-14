// Гриценко Ірина Валеріївна Ex 10.1
package main

import (
	"fmt"
)

func main() {
	var name string = "Ada Lovelace"
	nameptr := &name
	fmt.Println(name)
	fmt.Println(*nameptr)

	var age int = 25
	ageptr := &age
	fmt.Println(age)
	fmt.Println(*ageptr)

	var gender bool = true
	genderptr := &gender
	fmt.Println(gender)
	fmt.Println(*genderptr)
}
