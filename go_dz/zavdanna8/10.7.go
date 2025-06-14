// Гриценко Ірина Валеріївна Ex 10.7
package main

import (
	"fmt"
)

func main() {
	numbers := []int{123, 111, 333, 777, 222, 999, 555, 888, 666, 444}
	// print each value in the array
	var ctr int
	for ctr = 0; ctr < len(numbers); ctr++ {
		fmt.Println(numbers[ctr])
	}

	var numbersptr [10]*int
	for ctr := 0; ctr < len(numbersptr); ctr++ {
		numbersptr[ctr] = &numbers[ctr]
	}

	for t := 0; t < 9; t++ {
		for i := 0; i < 9; i++ {
			if numbers[i] < numbers[i+1] {
				tmp := *numbersptr[i+1]
				*numbersptr[i+1] = *numbersptr[i]
				*numbersptr[i] = tmp
			}
		}
	}
	fmt.Println(numbers)

}
