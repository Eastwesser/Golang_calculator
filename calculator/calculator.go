package calculator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"GolangCalculator/roman"
)

// Регулярное выражение для проверки римских чисел (до 100)
var romanRegex = regexp.MustCompile(`^(C|XC|L?X{0,9}|IX|V?I{0,3})$`)

// CalculateExpression вычисляет выражение, заданное в строковом формате
func CalculateExpression(input string) (string, error) {
	// Разделяем строку на компоненты, разделяя строку по пробелам
	tokens := strings.Fields(input)

	// Проверяем корректность формата
	if len(tokens) != 3 {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	// Получаем первый операнд, оператор и второй операнд, разбитые по пробелам
	number1, operator, number2 := tokens[0], tokens[1], tokens[2]

	// Проверяем, являются ли оба операнда римскими числами
	isRoman1 := romanRegex.MatchString(number1)
	isRoman2 := romanRegex.MatchString(number2)

	// Проверяем, чтобы оба числа были в одной системе счисления
	if isRoman1 != isRoman2 {
		panic("используются одновременно разные системы счисления")
	}

	var num1, num2 int
	var err error

	// Если числа римские, конвертируем их в арабские
	if isRoman1 {
		num1, err = roman.RomanToArabic(number1)
		if err != nil {
			panic(err)
		}
		num2, err = roman.RomanToArabic(number2)
		if err != nil {
			panic(err)
		}
	} else {
		// Если числа арабские, конвертируем их в целые числа
		num1, err = strconv.Atoi(number1)
		if err != nil {
			panic(err)
		}
		num2, err = strconv.Atoi(number2)
		if err != nil {
			panic(err)
		}
	}

	// Проверяем, чтобы числа были в диапазоне от 1 до 10
	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("числа должны быть от 1 до 10 включительно")
	}

	// Выполняем вычисление
	result, err := calculate(num1, num2, operator)
	if err != nil {
		panic(err)
	}

	// Если числа были римскими, конвертируем результат обратно в римские
	if isRoman1 {
		if result < 1 {
			panic("в римской системе нет отрицательных чисел")
		}
		return roman.ArabicToRoman(result)
	}

	// Возвращаем результат для арабских чисел
	return strconv.Itoa(result)
}

// calculate выполняет указанную арифметическую операцию над двумя целыми числами
func calculate(x int, y int, operator string) (int, error) {
	switch operator {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, errors.New("деление на ноль")
		}
		return x / y, nil
	default:
		return 0, errors.New("неизвестный оператор")
	}
}
