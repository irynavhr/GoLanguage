// Hrytsenko Iryna Valeriivna Task11
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(1000000000)
	fmt.Println("n = ", n)

	fmt.Println("a) ", QuantityOfDigits(n))
	fmt.Println("b) ", SumOfDigits(n))
	fmt.Println("c) ", FirstDigit(n, QuantityOfDigits(n)))
	fmt.Println("d) ", SumDiffOfDigits(n, QuantityOfDigits(n)))

}

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
func FirstDigit(n, e int) int {
	return n / tenInPov(e-1)
}
func tenInPov(u int) int {
	return int(math.Pow(float64(10), float64(u)))
}
func SumDiffOfDigits(n, e int) int {
	res := 0
	a := e
	for i := 1; i <= e; i++ {
		if i%2 == 1 {
			res += FirstDigit(n, a)
		} else {
			res -= FirstDigit(n, a)
		}
		a--
		n %= tenInPov(a)
	}

	return res
}
