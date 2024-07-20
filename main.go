package main

import (
	"GolangCalculator/calculator"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()
	calculator.MainCalc()
	fmt.Println("Thank you for using our Golang Calculator! ^w^")
}
