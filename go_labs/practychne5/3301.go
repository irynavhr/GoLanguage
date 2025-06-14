// Гриценко Ірина Валеріївна 08.11.2023
package main

import "fmt"

func main() {
	var a int // дане натуральне число
	fmt.Println("Введіть натуральне число:")
	fmt.Scan(&a)
	k := 2
	next, prev := 0, 0 //визначимо номери сусідніх чисел фібоначі
	for {
		f := Fibo(k)
		if f >= a {
			next = Fibo(k)
			prev = Fibo(k - 1)
			break
		}
		k++
	}
	// виведемо ближче з них, якщо числа однаково близькі, менше з них
	if (a - prev) > (next - a) {
		fmt.Println("Номер найближчого числа Фібоначі: ", k, ", його значення: ", next)
	} else {
		fmt.Println("Номер найближчого числа Фібоначі: ", k-1, ", його значення: ", prev)
	}
}

func Fibo(i int) int {

	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return Fibo(i-1) + Fibo(i-2)
}
