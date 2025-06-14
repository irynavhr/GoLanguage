// Гриценко Ірина Валеріївна 08.11.2023
//
//	f  = 0, f  = 1, f  = 3, f  = 6, f  = f    + f    - f
//	 0       1       2       3       n    n-1    n-2    n-3
//
// 0, 1, 3, 6, 8, 11, 13, 16, 18, 21, 23
package main

import "fmt"

func main() {
	var a int // дане натуральне число
	fmt.Println("Введіть натуральне число:")
	fmt.Scan(&a)
	k := 1
	next, prev := 0, 0 //визначимо сусідні чиселa
	for {
		r := rec(k)
		if r >= a {
			next = r
			prev = rec(k - 1)
			break
		}
		k++
	}
	// виведемо ближче з них, якщо числа однаково близькі, менше з них
	if (a - prev) > (next - a) {
		fmt.Println("Номер найближчого числа послідовності: ", k, ", його значення: ", next)
	} else {
		fmt.Println("Номер найближчого числа послідовності: ", k-1, ", його значення: ", prev)
	}
}

func rec(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	if i == 2 {
		return 3
	}
	if i == 3 {
		return 6
	}
	return rec(i-1) + rec(i-2) - rec(i-3)
}
