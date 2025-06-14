package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c, x0, x1, n float64
	fmt.Println("Ця програма знаходить наступний член рекурентної послідовності за двома попередніми. \nЩоб знайти n-не значення послідовності вкажіть коефіцієнти формули a, b, c:")
	fmt.Scan(&a, &b, &c)
	fmt.Println("Вкажіть перші значення послідовності x0 та x1 :")
	fmt.Scan(&x0, &x1)
	fmt.Println("Вкажіть натуральне значення n >= 2 для знаходження xn-го цієї послідовності :")
	fmt.Scan(&n)
	for n < 2 || n != math.Trunc(n) {
		fmt.Println("Це число не задовольняє умову, введіть ще раз")
		fmt.Scan(&n)
	}
	fmt.Print("x", n, "=", RecSeries(a, b, c, x0, x1, n))

}

func RecSeries(a, b, c, x0, x1, n float64) float64 {
	var x float64
	for k := 2; k <= int(n); k++ {
		x = a*x1 + b*x0 + c
		x0 = x1
		x1 = x
	}
	return x
}
