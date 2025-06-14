// Hrytsenko Iryna Valeriivna Task6
package main

import (
	"fmt"
	"math/rand"
)

func exchange(a, b int) int {
	return b*10 + a
}
func doubleDig(a, b int) int {
	return a*int(1100) + b*int(11)
}
func doubleNam(x int) int {
	if x < 10 {
		return x * 11
	}
	return x * 101
}

func main() {
	x := rand.Intn(99) + 1
	fmt.Println(x)
	a := x / 10
	b := x % 10
	fmt.Println("x=", x)
	fmt.Println("digit exchange:", exchange(a, b))
	fmt.Println("double digit:", doubleDig(a, b))
	fmt.Println("double number:", doubleNam(x))
}
