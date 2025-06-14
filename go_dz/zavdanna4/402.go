// Hrytsenko Iryna Valeriivna Task2
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(79) + 1
	fmt.Println("n = ", n)
	a := r.Float64() * 10000
	a = math.Round(a) / 1000
	fmt.Println("a = ", a)

	x := power(float64(n), a)
	fmt.Println(x)
	x = prodOfSum(float64(n), a)
	fmt.Println(x)
	x = SumofDenom(float64(n), a)
	fmt.Println(x)
	x = DenomInPowOfTwo(float64(n), a)
	fmt.Println(x)
	x = Prod(float64(n), a)
	fmt.Println(x)
}

func power(n, a float64) float64 {
	fmt.Print("Power a of n: ")
	res := 1.0
	for k := 0; float64(k) < n; k++ {
		res *= a
	}
	return res
}

func prodOfSum(n, a float64) float64 {
	fmt.Print("Product of sum: ")
	res := 1.0
	for k := 0; float64(k) < n-1; k++ {
		res *= a + float64(k)
	}
	return res
}

func SumofDenom(n, a float64) float64 {
	fmt.Print("Sum of denoms: ")
	onestep := 1.0
	res := 0.0
	for k := 0; float64(k) < n; k++ {
		onestep *= a + float64(k)
		res += (1 / onestep)
	}
	return res
}
func DenomInPowOfTwo(n, a float64) float64 {
	fmt.Print("Denom In Pow Of Two: ")
	pow := 1.0
	res := 0.0
	for k := 0; float64(k) <= n; k++ {
		pow = math.Pow(a, math.Pow(2, float64(k)))
		res += (1 / pow)
	}
	return res
}

func Prod(n, a float64) float64 {
	fmt.Print("Product of difference: ")
	res := 1.0
	for k := 0; float64(k) <= n; k++ {
		res *= a - float64(k)*n
	}
	return res
}
