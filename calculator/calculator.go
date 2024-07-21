package calculator

import (
	"regexp"
	"strconv"
	"strings"

	"GolangCalculator/roman"
)

// Регулярное выражение для проверки римских чисел
var romanRegex = regexp.MustCompile(`^(X|IX|IV|V?I{0,3})$`)

// CalculateExpression вычисляет выражение, заданное в строковом формате
func CalculateExpression(input string) (string, error) {
	// Разделяем строку на компоненты, разделяя строку по пробелам
	tokens := strings.Fields(input)

	// Проверяем корректность формата
	if len(tokens) != 3 {
		panic("формат не удовлетворяет заданию — два операнда и один оператор через пробелы! (+, -, /, *)")
		// Раскомментить return для тестов!
		//return "", errors.New("формат не удовлетворяет заданию — два операнда и один оператор через пробелы! (+, -, /, *)")
	}

	// Получаем первый операнд, оператор и второй операнд, разбитые по пробелам
	number1, operator, number2 := tokens[0], tokens[1], tokens[2]

	// Проверяем, являются ли оба операнда римскими числами
	isRoman1 := romanRegex.MatchString(number1)
	isRoman2 := romanRegex.MatchString(number2)
	isRoman := isRoman1 && isRoman2

	// Проверяем, чтобы оба числа были в одной системе счисления
	if (isRoman1 && !isRoman2) || (!isRoman1 && isRoman2) {
		panic("используются одновременно разные системы счисления")
		// Раскомментить return для тестов!
		//return "", errors.New("используются одновременно разные системы счисления")
	}

	var num1, num2 int
	var err error

	// Если числа римские, конвертируем их в арабские
	if isRoman {
		num1, err = roman.RomanToArabic(number1)
		if err != nil {
			panic(err)
			// Раскомментить return для тестов!
			//return "", err
		}
		num2, err = roman.RomanToArabic(number2)
		if err != nil {
			panic(err)
			// Раскомментить return для тестов!
			//return "", err
		}
	} else {
		// Если числа арабские, конвертируем их в целые числа
		num1, err = strconv.Atoi(number1)
		if err != nil {
			panic(err)
			// Раскомментить return для тестов!
			//return "", err
		}
		num2, err = strconv.Atoi(number2)
		if err != nil {
			panic(err)
			// Раскомментить return для тестов!
			//return "", err
		}
	}

	// Проверяем, чтобы числа были в диапазоне от 1 до 10
	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("числа должны быть от 1 до 10 включительно")
		// Раскомментить return для тестов!
		//return "", errors.New("числа должны быть от 1 до 10 включительно")
	}

	result, err := calculate(num1, num2, operator)
	if err != nil {
		return "", err
	}

	if isRoman1 {
		if result < 1 {
			panic("в римской системе нет отрицательных чисел")
			// Раскомментить return для тестов!
			//return "", errors.New("в римской системе нет отрицательных чисел")
		}
		return roman.ArabicToRoman(result)
	}

	return strconv.Itoa(result), nil
}

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
			panic("деление на ноль")
			// Раскомментить return для тестов!
			//return 0, errors.New("деление на ноль")
		}
		return x / y, nil
	default:
		panic("неизвестный оператор")
		// Раскомментить return для тестов!
		//return 0, errors.New("неизвестный оператор")
	}
}
