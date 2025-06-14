// Гриценко Ірина Валеріївна Ex 9.4
package main

import (
	"fmt"
)

func main() {
	var numbers [10]int
	for x := 1; x < 11; x++ {
		numbers[x-1] = x * x
	}
	fmt.Println(numbers)

	var reverse [10]int
	fmt.Println(reverse)

	for i, _ := range reverse {
		reverse[i] = numbers[9-i]
	}
	fmt.Println(reverse)

}
