package calculator

import "fmt"

/*
Per meliora
ad astra
et altiora!
*/

func MainCalc() {
	fmt.Println("Enter first number: ")
	var numberA int
	fmt.Scanln(&numberA)

	fmt.Println("Enter second number: ")
	var numberB int
	fmt.Scanln(&numberB)

	if numberA < numberB {
		numberA, numberB = numberB, numberA
	}

	fmt.Println("Result:")
	fmt.Println(numberA, numberB)

	//sum := 1
	//for i := 5; i <= 8; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
}

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
