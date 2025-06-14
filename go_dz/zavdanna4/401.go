// Hrytsenko Iryna Valeriivna Task1
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(29) + 1
	fmt.Println("n = ", n)
	fmt.Println(powOfTwo(n))
	fmt.Println(Factorial(n))
	fmt.Println(product(n))
	fmt.Println(sinus(n))
	fmt.Println(sqrt(n))
	fmt.Println(sqrt2(n))
}

func powOfTwo(n int) int {
	fmt.Print("Power of two: ")
	res := 1
	for k := 0; k < n; k++ {
		res *= 2
	}
	return res
}

func Factorial(n int) int {
	fmt.Print("Factorial: ")
	res := 1
	for k := 2; k <= n; k++ {
		res *= k
	}
	return res
}

func product(n int) float32 {
	fmt.Print("Expression with products: ")
	var res float32 = 1
	for k := 1; k <= n; k++ {
		res *= (1 + 1/(float32(k*k)))
	}
	return res
}

func sinus(n int) float64 {
	fmt.Print("Expression with sinus: ")
	var res float64 = 0
	var sum float64 = 0
	for k := 1; k <= n; k++ {
		sum += math.Sin(float64(k))
		res += 1 / sum
	}
	return res
}

func sqrt(n int) float64 {
	fmt.Print("Expression with square roots: ")
	var res float64 = 0
	for k := 1; k <= n; k++ {
		res += 2
		res = math.Sqrt(res)
	}
	return res
}

func sqrt2(n int) float64 {
	fmt.Print("Another expression with square roots: ")
	var res float64 = 0
	for k := n; k > 0; k-- {
		res += 3 * float64(k)
		res = math.Sqrt(res)
	}
	return res
}
