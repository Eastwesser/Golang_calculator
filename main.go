package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"GolangCalculator/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение (или введите 'exit' для выхода):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Удаляем пробелы в начале и конце строки

		// Отладочная информация
		// fmt.Printf("Получен ввод: %q\n", input)

		if strings.ToLower(input) == "exit" {
			break // Завершаем работу программы, если введено 'exit'
		}

		// Вызываем функцию для вычисления выражения и обрабатываем возможные ошибки
		result, err := calculator.CalculateExpression(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println(result)
		}
	}
	fmt.Println("Благодарим за использование нашего калькулятора! ^w^")
}
