package main

import (
	"GolangCalculator/calculator"
	"fmt"
)

// Основная функция для запуска калькулятора
func main() {
	var phrase = "Благодарим за использование нашего калькулятора! ^w^"
	calculator.MainCalc()
	fmt.Println(phrase)
}
