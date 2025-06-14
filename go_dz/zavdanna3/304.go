// Hrytsenko Iryna Valeriivna Task4
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func QuantityOfDigits(n int) int {
	var Q int = 0
	for n >= 1 {
		Q += 1
		n /= 10
	}

	return Q

}
func SumOfDigits(n int) int {
	var sum int = 0
	for n >= 1 {
		sum += n % 10
		n /= 10
	}

	return sum

}
func LastDigit(n int) int {
	return n % 10

}
func FirstDigit(n int) int {
	e := QuantityOfDigits(n)
	return n / int(math.Pow(float64(10), float64(e-1)))
}
func PreLastDigit(n int) int {
	return n / 10 % 10

}

func main() {
	x := rand.Intn(100) + 1
	fmt.Println(x)

	fmt.Println("QuantityOfDigits = ", QuantityOfDigits(x))
	fmt.Println("SumOfDigits = ", SumOfDigits(x))
	fmt.Println("LastDigit = ", LastDigit(x))
	fmt.Println("FirstDigit = ", FirstDigit(x))
	if x > 9 {
		fmt.Println("PreLastDigit = ", PreLastDigit(x))
	}
}
