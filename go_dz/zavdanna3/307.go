// Hrytsenko Iryna Valeriivna Task7
package main

import (
	"fmt"
	"math/rand"
)

func exchangeThirdAndFirstDig(x, a, c int) int {
	return x - 99*(a-c)
}
func plusExchangeNum(x, a, c int) int {
	return x*1000 + exchangeThirdAndFirstDig(x, a, c)
}

func main() {
	x := rand.Intn(900) + 100
	fmt.Println(x)
	a := x / 100 % 10
	c := x % 10
	fmt.Println("x=", x)
	fmt.Println("exchangeThirdAndFirstDig:", exchangeThirdAndFirstDig(x, a, c))
	fmt.Println("plusExchangeNum:", plusExchangeNum(x, a, c))
	//fmt.Println("double number:", doubleNam(x))
}
