package main

import (
	"fmt"
	"math"
)

func main() {
	var n, a float64

	fmt.Println("Ця програма визначає чи є в числі певна цифра\n Введіть натуральне число:")
	fmt.Scan(&n)
	//Перевірка чи отримали натуральне значення
	for n < 1 || n != math.Trunc(n) {
		fmt.Println("Це не натуральне число, введіть ще раз")
		fmt.Scan(&n)
	}

	fmt.Println("Введіть цифру - число від 0 до 9:")
	fmt.Scan(&a)
	//Перевірка чи отримали натуральне значення
	for a < 0 || a > 9 || a != math.Trunc(a) {
		fmt.Println("Це не цифра, введіть ще раз")
		fmt.Scan(&a)
	}
	// p bool
	//p =IsInnumber(int(n), int(a))
	if IsInnumber(int(n), int(a)) == true {
		fmt.Println("В числі", n, "міститься цифра", a)
	} else {
		fmt.Println("В числі", n, " немає цифри", a)
	}
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
