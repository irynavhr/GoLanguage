// Hrytsenko Iryna Valeriivna Task13
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
	d := r.Intn(10)
	fmt.Println("d = ", d)

	fmt.Println("a) ", IsInnumber(n, d))
	fmt.Println("b) ", InversInnumber(n))
	fmt.Println("c) ", firstLast(n))
	fmt.Println("d) ", firstLastSame(n, d))

}

func QuantityOfDigits(n int) int {
	var Q int = 0
	for n >= 1 {
		Q += 1
		n /= 10
	}

	return Q
}
func FirstDigit(n, e int) int {
	return n / tenInPov(e-1)
}
func tenInPov(u int) int {
	return int(math.Pow(float64(10), float64(u)))
}

func IsInnumber(n, a int) bool {
	for n > 0 {
		if a == n%10 {
			return true
		}
		n /= 10
	}
	return false
}
func InversInnumber(n int) int {
	new := 0
	for n > 0 {
		a := n % 10
		new += a
		new *= 10
		n /= 10
	}
	return new / 10

}
func firstLast(n int) int {
	new := 0
	m := QuantityOfDigits(n)
	l := n % 10
	f := FirstDigit(n, m)
	s := n / 10 % tenInPov(m-2)
	new = (l*tenInPov(m-2)+s)*10 + f
	return new
}
func firstLastSame(n, a int) int {
	new := 0
	m := QuantityOfDigits(n)
	l, f := a, a
	s := n / 10 % tenInPov(m-2)
	new = (l*tenInPov(m-2)+s)*10 + f
	return new
}
