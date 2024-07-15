package calculator

import (
	"fmt"
)

/*
Per meliora
ad astra
et altiora!
*/

func MainCalc() {
	fmt.Println("Enter first number: ")
	var number1 int
	fmt.Scanln(&number1)

	fmt.Println("Enter second number: ")
	var number2 int
	fmt.Scanln(&number2)

	fmt.Println("Choose your math option:")
	fmt.Println("A. Sum")
	fmt.Println("B. Subtraction")
	fmt.Println("C. Multiplication")
	fmt.Println("D. Division")

	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "A":
		fmt.Println("Result:", add(number1, number2))
	case "B":
		fmt.Println("Result:", subtract(number1, number2))
	case "C":
		fmt.Println("Result:", multiply(number1, number2))
	case "D":
		fmt.Println("Result:", divide(number1, number2))
	default:
		fmt.Println("Unrecognized choice")
	}
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}
