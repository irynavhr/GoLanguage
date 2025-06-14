// Hrytsenko Iryna Valeriivna Task1
package main

import (
	"fmt"
	"math/rand"
)

func def_parity(a int) {
	if a%2 == 0 {
		fmt.Println("It is even number")
	} else {
		fmt.Println("It is odd number")
	}
}

func main() {

	x := rand.Intn(9999) + 1
	fmt.Println(x)
	def_parity(x)
}
