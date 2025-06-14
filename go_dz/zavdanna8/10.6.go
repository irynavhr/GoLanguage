// Гриценко Ірина Валеріївна Ex 10.6
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
	for i := 0; i < 5; i++ {
		tmp := *numbersptr[9-i]
		*numbersptr[9-i] = *numbersptr[i]
		*numbersptr[i] = tmp
	}

	fmt.Println(numbers)
}
