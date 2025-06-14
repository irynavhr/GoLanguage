// Hrytsenko Iryna Valeriivna Task2
package main

import (
	"fmt"
	"math/rand"
)

func def_hundred(a int) int {
	return a / 100 % 10
}

func main() {
	x := rand.Intn(10000) + 100
	fmt.Println("x=", x, "\nThis number have", def_hundred(x), "hundred")
}
