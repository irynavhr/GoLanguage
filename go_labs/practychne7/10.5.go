// Гриценко Ірина Валеріївна Ex 10.5
package main

import "fmt"

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
	tmp := *numbersptr[8]
	*numbersptr[8] = *numbersptr[1]
	*numbersptr[1] = tmp

	fmt.Println(numbers)
}
