package main

import (
	"fmt"
	"math"
	"time"
)

func main1a() {
	var n float64
	fmt.Println("Введіть натуральне число, та отримайте 2 в цьому степені:")
	fmt.Scan(&n)

	var result uint64
	result = 1

	for a := 1; a <= int(n); a++ {
		result *= 2
	}

	if result == 1 || n != math.Trunc(n) {
		fmt.Println("Прочитай уважно умову")
		time.Sleep(5 * time.Second)
		fmt.Println("\nНезнайшов свою помилку? ХА-ХА-ХА\nЙди вивчи поняття натурального числа")
	} else {
		fmt.Println(result)
	}
}

//go run main1a.go
