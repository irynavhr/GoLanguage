package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64

	fmt.Println("Введіть натуральне число, а я скажу вам його  суму цифр:")
	fmt.Scan(&number)

	//Перевірка чи отримали натуральне значення
	for number < 1 || number != math.Trunc(number) {
		fmt.Println("Це не натуральне число, введіть ще раз")
		fmt.Scan(&number)
	}

	fmt.Println(SumOfDigits(int(number)))

}

func SumOfDigits(n int) int {
	sum := 0
	for n >= 1 {
		sum += n % 10
		n /= 10
	}

	return sum

}
