package main

import (
	"fmt"
	"math"
	"time"
)

func mainv() {

	var a float64
	fmt.Println("Введіть дійсне число а, крім 0")
	fmt.Scan(&a)

	for a == 0 {
		fmt.Println("Число не задовольняє умову\nВведіть ще раз")
		fmt.Scan(&a)
	}

	var n float64
	fmt.Println("Введіть натуральне число n")
	fmt.Scan(&n)

	for n != math.Trunc(n) || n <= 0 {
		fmt.Println("Це не натуральне число\nВведіть ще раз")
		fmt.Scan(&n)
	}

	fmt.Println("Обчислюю...")

	var dodanok float64
	dodanok = 1

	var result float64
	result = 0

	for k := 0; k <= int(n); k++ {

		dodanok /= (a + float64(k))

		result += dodanok
	}

	time.Sleep(1 * time.Second)

	fmt.Println(result)
}

//go run main2v.go
