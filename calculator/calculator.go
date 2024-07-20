package calculator

import (
	"GolangCalculator/roman"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MainCalc() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter the first number (or type 'exit' to quit):")
		number1, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		number1 = strings.TrimSpace(number1)
		if strings.ToLower(number1) == "exit" {
			break
		}

		fmt.Println("Enter the second number:")
		number2, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		number2 = strings.TrimSpace(number2)

		fmt.Println("Enter the operation (+, -, *, /):")
		operator, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		operator = strings.TrimSpace(operator)

		isRoman := roman.RomanToNumeral(number1)
		var num1, num2 int

		if isRoman {
			if !roman.RomanToNumeral(number2) {
				fmt.Println("Error: Both numbers must be Roman numerals or Arabic numerals")
				continue
			}
			num1 = roman.RomanToArabic(number1)
			num2 = roman.RomanToArabic(number2)
		} else {
			var err1, err2 error
			num1, err1 = strconv.Atoi(number1)
			num2, err2 = strconv.Atoi(number2)
			if err1 != nil || err2 != nil {
				fmt.Println("Error: Both numbers must be valid integers")
				continue
			}
		}

		result, err := calculate(num1, num2, operator)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if isRoman {
			if result < 1 {
				fmt.Println("Error: Result in Roman numerals is less than I")
				continue
			}
			fmt.Println("Result:", roman.ArabicToRoman(result))
		} else {
			fmt.Println("Result:", result)
		}
	}
}

func calculate(x int, y int, operator string) (int, error) {
	switch operator {
	case "+":
		return add(x, y), nil
	case "-":
		return subtract(x, y), nil
	case "*":
		return multiply(x, y), nil
	case "/":
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return divide(x, y), nil
	default:
		return 0, errors.New("unrecognized operator")
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
