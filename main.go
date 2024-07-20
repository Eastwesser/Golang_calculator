package main

import (
	"GolangCalculator/calculator"
	"fmt"
)

func main() {
	var phrase = "Благодарим за использование нашего калькулятора! ^w^"
	calculator.MainCalc()
	fmt.Println(phrase)
}
