package calculator

import (
	"GolangCalculator/roman"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func MainCalc() {
	fmt.Println("Enter your calculation:")
	var input string
	fmt.Scanln(&input)

	number1, number2, operator, err := parseInput(input)
	if err != nil {
		panic(err)
	}

	isRoman := roman.IsRomanNumeral(number1)
	if isRoman != roman.IsRomanNumeral(number2) {
		panic("different numeral systems used")
	}

	var num1, num2 int
	if isRoman {
		num1 = roman.RomanToArabic(number1)
		num2 = roman.RomanToArabic(number2)
	} else {
		num1, _ = strconv.Atoi(number1)
		num2, _ = strconv.Atoi(number2)
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("numbers must be between 1 and 10")
	}

	var wg sync.WaitGroup
	resultChan := make(chan int)
	errorChan := make(chan error)

	wg.Add(1)
	go calculate(num1, num2, operator, resultChan, errorChan, &wg)

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case result := <-resultChan:
		if isRoman {
			if result < 1 {
				panic("Result in Roman numerals is less than I")
			}
			fmt.Println("Result:", roman.ArabicToRoman(result))
		} else {
			fmt.Println("Result:", result)
		}
	case err := <-errorChan:
		panic(err)
	}
}

func parseInput(input string) (string, string, string, error) {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		return "", "", "", errors.New("invalid format")
	}
	return parts[0], parts[2], parts[1], nil
}

func calculate(x int, y int, operator string, resultChan chan int, errorChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	switch operator {
	case "+":
		resultChan <- add(x, y)
	case "-":
		resultChan <- subtract(x, y)
	case "*":
		resultChan <- multiply(x, y)
	case "/":
		if y == 0 {
			errorChan <- errors.New("division by zero")
			return
		}
		resultChan <- divide(x, y)
	default:
		errorChan <- errors.New("unrecognized operator")
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
