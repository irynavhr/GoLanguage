// Hrytsenko Iryna Valeriivna Task5
package main

import (
	"fmt"
	"math/rand"
)

func QuantityOfDigitsIs(n int) int {
	var Q int = 0
	for n >= 1 {
		Q += 1
		n /= 10
	}

	return Q
}
func Palindrom(a int) bool {
	last := a % 10
	preLast := a / 10 % 10
	prpreLast := a / 100 % 10
	prprprlast := a / 1000

	if QuantityOfDigitsIs(a) == int(1) {
		return true
	}
	if QuantityOfDigitsIs(a) == int(2) && last == preLast {
		return true
	}
	if QuantityOfDigitsIs(a) == int(3) && last == prpreLast {
		return true
	}
	if QuantityOfDigitsIs(a) == int(4) && last == prprprlast && prpreLast == preLast {
		return true
	}
	return false
}

func ThreeEqual(a int) bool {
	last := a % 10
	preLast := a / 10 % 10
	prpreLast := a / 100 % 10
	prprprlast := a / 1000
	if QuantityOfDigitsIs(a) == int(4) && (last == prprprlast || preLast == prpreLast) {
		if last == preLast || prpreLast == prprprlast {
			return true
		}
	}
	if QuantityOfDigitsIs(a) == int(3) && last == preLast && preLast == prpreLast {
		return true
	}
	return false
}

func DigDiffer(a int) bool {
	last := a % 10
	preLast := a / 10 % 10
	prpreLast := a / 100 % 10
	prprprlast := a / 1000
	if last != preLast && last != prprprlast && preLast != prpreLast && prpreLast != prprprlast {
		return true
	}
	if (QuantityOfDigitsIs(a) == 2 || QuantityOfDigitsIs(a) == 1) && last != preLast {
		return true
	}
	return false
}

func main() {
	x := rand.Intn(9999) + 1
	fmt.Println(x)
	fmt.Println("Is palindrom:", Palindrom(x))
	fmt.Println("Have three equal digits:", ThreeEqual(x))
	fmt.Println("All digits are differen:", DigDiffer(x))
}
