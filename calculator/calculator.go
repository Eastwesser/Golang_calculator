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
}
