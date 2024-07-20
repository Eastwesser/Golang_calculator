package calculator

import (
	"GolangCalculator/roman"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// MainCalc обрабатывает основную логику калькулятора и взаимодействие с пользователем
func MainCalc() {
	for {
		var number1, number2, operator string

		fmt.Println("Введите первое число (или введите 'exit' для выхода):")
		fmt.Scanln(&number1)
		number1 = strings.TrimSpace(number1)
		if strings.ToLower(number1) == "exit" {
			break
		}

		fmt.Println("Введите второе число:")
		fmt.Scanln(&number2)
		number2 = strings.TrimSpace(number2)

		fmt.Println("Введите операцию (+, -, *, /):")
		fmt.Scanln(&operator)
		operator = strings.TrimSpace(operator)

		isRoman := roman.RomanToNumeral(number1)
		var num1, num2 int

		if isRoman {
			if !roman.RomanToNumeral(number2) {
				fmt.Println("Ошибка: оба числа должны быть римскими или арабскими числами")
				continue
			}
			num1, _ = roman.RomanToArabic(number1)
			num2, _ = roman.RomanToArabic(number2)
		} else {
			var err1, err2 error
			num1, err1 = strconv.Atoi(number1)
			num2, err2 = strconv.Atoi(number2)
			if err1 != nil || err2 != nil {
				fmt.Println("Ошибка: оба числа должны быть допустимыми целыми числами")
				continue
			}
		}

		result, err := calculate(num1, num2, operator)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		if isRoman {
			if result < 1 {
				fmt.Println("Ошибка: Результат в римских числах меньше I")
				continue
			}
			fmt.Println("Результат:", roman.ArabicToRoman(result))
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

// calculate выполняет указанную арифметическую операцию над двумя целыми числами
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
			return 0, errors.New("деление на ноль")
		}
		return divide(x, y), nil
	default:
		return 0, errors.New("неизвестный оператор")
	}
}

// add возвращает сумму x и y
func add(x int, y int) int {
	return x + y
}

// subtract возвращает разность между x и y
func subtract(x int, y int) int {
	return x - y
}

// multiply возвращает произведение x и y
func multiply(x int, y int) int {
	return x * y
}

// divide возвращает частное от деления x на y
func divide(x int, y int) int {
	return x / y
}
