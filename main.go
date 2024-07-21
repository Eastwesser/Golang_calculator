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

	// Обработка паники
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()

	for {
		fmt.Println("Введите выражение (или введите 'exit' для выхода):")
		input, err := reader.ReadString('\n') // Читаем строку ('_' -> 'err' для обработки ошибок)
		if err != nil {
			fmt.Println("Введите другое выражение в соответствии с инструкцией", err)
			continue
		}
		input = strings.TrimSpace(input) // Удаляем префиксные и постфиксные пробелы в начале и конце строки

		if strings.ToLower(input) == "exit" {
			break // Завершаем работу программы, если введено 'exit', выходим из цикла
		}

		// Вызываем функцию для вычисления выражения и обрабатываем возможные ошибки
		result, err := calculator.CalculateExpression(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		fmt.Println(result)
	}
	fmt.Println("Благодарим за использование нашего калькулятора! ^w^")
}
