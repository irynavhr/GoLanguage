// Гриценко Ірина Валеріївна 08.11.2023
package main

import "fmt"

const N int16 = 36

func main() {
	var arrQ [2][N]int16 //оголошення масиву
	//заповнення першого рядка сумою цифр
	for i := 0; int16(i) < N; i++ {
		arrQ[0][i] = int16(i + 1)
	}
	//заповнення другого рядка масиву кількістю чисел
	for num := int16(1000); num < 10000; num++ {
		for i := 0; int16(i) < N; i++ {
			if arrQ[0][i] == sumofdigits(num) {
				arrQ[1][i]++
			}
		}
	}
	fmt.Println(arrQ) //виведення
}

func sumofdigits(n int16) int16 {
	sum := int16(0)
	for n >= 1 {
		sum += n % 10
		n /= 10
	}
	return sum
}
