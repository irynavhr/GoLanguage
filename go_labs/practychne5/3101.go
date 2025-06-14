// Гриценко Ірина Валеріївна 08.11.2023
package main

import "fmt"

func main() {
	var arrQ [36]int16 //оголошення масиву
	//заповнення масиву кількістю чисел
	for num := int16(1000); num < 10000; num++ {
		for i, _ := range arrQ {
			if int16(i+1) == sumOfdigits(num) {
				arrQ[i]++
			}
		}
	}

	fmt.Println("Такі суми цифр:")
	maxquantity := arrQ[0]
	for i := 0; i < 35; i++ {
		if arrQ[i+1] > maxquantity {
			maxquantity = arrQ[i+1]
		}
	}
	for i := 0; i < 36; i++ {
		if arrQ[i] == maxquantity {
			fmt.Println(i + 1)
		}
	}
	fmt.Println("мають мають найбільше чисел (з діапазону 1000-10000), а саме: ", maxquantity, " чисел")
}

func sumOfdigits(n int16) int16 {
	sum := int16(0)
	for n >= 1 {
		sum += n % 10
		n /= 10
	}
	return sum

}
