package main

import "fmt"

func main0() {
	var a, b uint16
	fmt.Scan(&a, &b)
	if a > 23 && b > 59 {
		fmt.Println("Нажаль мій циферблат скінченний:(")
		fmt.Println("Не можу відобразити такий час")
	} else {
		switch ost, prdost := a%10, a/10; {
		case a > 23:
			fmt.Print("Я таких годин не вчив, а хвилини такі знаю: ")
		case ost == 1 && prdost != 1:
			fmt.Print(a, " година ")
		case (ost == 2 || ost == 3 || ost == 4) && prdost != 1:
			fmt.Print(a, " години ")
		default: //case a == 0 || a >= 5 && a <= 20
			fmt.Print(a, " годин ")
		}
		switch ost, prdost := b%10, b/10; {
		case b > 59:
			fmt.Print(", а хвилин таких не вчив")
		case ost == 1 && prdost != 1:
			fmt.Print(b, " хвилина")
		case (ost == 2 || ost == 3 || ost == 4) && prdost != 1:
			fmt.Print(b, " хвилини")
		default: //case ost == 0 || ost >= 5 && ost <= 9 || prdost == 1
			fmt.Print(b, " хвилин")
		}
	}

}

//  go run main1.go
