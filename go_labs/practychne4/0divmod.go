package main

import (
	"fmt"
	"math"
)

func main() {
	var num, denom float64

	fmt.Println("Введіть 2 натуральні числа, а я скажу вам яка ціла частина та залишок при діленні перше на друге\n Введіть перше:")
	fmt.Scan(&num)
	//Перевірка чи отримали натуральне значення
	for num < 1 || num != math.Trunc(num) {
		fmt.Println("Це не натуральне число, введіть ще раз")
		fmt.Scan(&num)
	}

	fmt.Println("Введіть друге:")
	fmt.Scan(&denom)
	//Перевірка чи отримали натуральне значення
	for denom < 1 || denom != math.Trunc(denom) {
		fmt.Println("Це не натуральне число, введіть ще раз")
		fmt.Scan(&denom)
	}

	x, y := divMod(int(num), int(denom))
	fmt.Println("Ціла частина:", x)
	fmt.Println("Залишок:", y)

}

func divMod(a, b int) (int, int) {
	return a / b, a % b
}
