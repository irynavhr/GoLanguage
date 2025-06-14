// Гриценко Ірина Валеріївна Ex 9.5
package main

import (
	"fmt"
)

func main() {
	const k = 10
	var numbers [k]int
	for x := 1; x < 11; x++ {
		numbers[x-1] = x * x
	}
	fmt.Println(numbers)

	var copy [k]int
	for i, _ := range copy {
		copy[i] = numbers[i]
	}

	for i, _ := range numbers {
		numbers[i] = copy[9-i]
	}
	fmt.Println(numbers)

}
