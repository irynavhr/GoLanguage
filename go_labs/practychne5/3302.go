// Гриценко Ірина Валеріївна 08.11.2023
package main

import "fmt"

func main() {
	fmt.Println(Rec(0))
	fmt.Println(Rec(1))
	fmt.Println(Rec(2))
	fmt.Println(Rec(3))
	fmt.Println(Rec(9))
	fmt.Println(Rec(20))

}

func Rec(i int) int {

	if i == 0 {
		return 1
	}
	if i == 1 {
		return 1
	}
	if i == 2 {
		return 1
	}
	return Rec(i-1) - Rec(i-2) + Rec(i-3)
}
