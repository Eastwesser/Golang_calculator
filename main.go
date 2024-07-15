package main

import (
	"GolangCalculator/calculator"
	"fmt"
)

func main() {
	var phrase = "Thank you for using our Golang Calculator! ^w^"
	calculator.MainCalc()
	fmt.Println(phrase)
}
