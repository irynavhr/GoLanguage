// Гриценко Ірина Валеріївна Ex 10.2
package main

import "fmt"

func main() {
	fmt.Print("Enter your name: ")
	var name string
	nameptr := &name
	fmt.Scanln(&*nameptr)

	fmt.Print("Enter your age: ")
	var age string
	ageptr := &age
	fmt.Scanln(&*ageptr)

	fmt.Print("Enter your gender: ")
	var gender string
	genderptr := &gender
	fmt.Scanln(&*genderptr)

	fmt.Print("Your first name is: ")
	fmt.Println(*nameptr)

	fmt.Print("Your age is: ")
	fmt.Println(*ageptr)

	fmt.Print("Your gender is: ")
	fmt.Println(*genderptr)

}
