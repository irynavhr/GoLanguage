// Hrytsenko Iryna Valeriivna Task3
package main

import (
	"fmt"
	"math/rand"
)

func def_equaly_of_statements(a int) {
	var sum int
	sum = a/10 + a%10
	if a*a == sum*sum*sum {
		fmt.Println("For", a, "Yes!!! My congratulations!")
	} else {
		fmt.Println("For", a, "unfortunately, no")
	}
}

func main() {

	x := rand.Intn(99) + 1

	fmt.Println("x=", x)
	def_equaly_of_statements(x)
}
