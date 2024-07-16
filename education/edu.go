package education

/*
Per meliora
ad astra
et altiora!

========================================
//firstNum := int
//secondNum := int

fmt.Scan(&firstNum)
fmt.Scan(&secondNum)

add := "+"
subtract := "-"
multiply := "*"
divide := "/"
modulo := "%"

========================================

var a, b int
fmt.Scan(&a)
fmt.Scan(&b)
if a > b; a > 1000 {
    fmt.Println("a больше чем b")
} else if a == b {
    fmt.Println("a равно b")
} else {
    fmt.Println("a меньше чем b")
}

========================================

height := 177
if height <= 165 {
     fmt.Println("Человек невысокого роста")
} else if height > 165 && height <= 185 {
     fmt.Println("Человек среднего роста")
} else {
     fmt.Println("Высокий человек")
}
// выведет: Человек среднего роста

========================================

month := "Январь"

switch {
    case month == "Декабрь" || month == "Январь" || month == "Февраль":
        fmt.Println("Зима")
    case month == "Март" || month == "Апрель" || month == "Май":
        fmt.Println("Весна")
}

========================================

day := 3

switch day {
    case 1:
        fmt.Println("Понедельник")
        fallthrough
    case 2:
        fmt.Println("Вторник")
        fallthrough
    case 3:
        fmt.Println("Среда")
        fallthrough
    case 4:
        fmt.Println("Четверг")
        fallthrough
    case 5:
        fmt.Println("Пятница")
    default:
        fmt.Println("Неправильный день")
}

выведет
Среда
Четверг
Пятница

========================================

s := 3

switch s {
    case 1:
        fmt.Println("Один")
        fallthrough
    case 2:
        fmt.Println("Два")
        fallthrough
    case 3:
        fmt.Println("Три")
    case 4:
        fmt.Println("Четыре")
    case 5:
        fmt.Println("Пять")
}

// вывод 3

========================================

if result = add:
	return firstNum + secondNum
	fmt.Println(result)

if result = subtract:
	return firstNum - secondNum
	fmt.Println(result)

if result = multiply:
	return firstNum * secondNum
	fmt.Println(result)

if result = divide:
	return firstNum / secondNum
		if firstNum and secondNum != float32:
			if secondNum := 0:
				Division by zero is not allowed
	fmt.Println(result)

if result = modulo:
	return firstNum % secondNum
	fmt.Println(result)

fmt.Println(result)

var a int = 31
var b int = 14

--------------
a += b
fmt.Println(a)  // выведет 45
--------------
a -= b
fmt.Println(a)  // выведет 17
--------------
a *= b
fmt.Println(a)  // выведет 434
--------------
a /= b
fmt.Println(a)  // выведет 2
--------------
a %= b
fmt.Println(a)  // выведет 3


func mainSecond() {
    var x int
    var y int
    fmt.Scanln(&x)     // считаем первое число в переменную x
    fmt.Scanln(&y)     // считаем второе число в переменную y
    fmt.Println(x + y) // выведем сумму чисел - 12
}


sum_even := 0

for i := 1; i <= 10; i++ {
    if i % 2 != 0 {
        continue             // число нечетное, то пропускаем его
    }                        // и переходим к следующей итерации
    sum_even += i
}
fmt.Println(sum_even)        // выведет: 30

for i := 1; i <= 9; i++{
    if i % 5 == 0 {
        break            // если число кратно 5, то выходим из цикла
    }
    fmt.Println(i)
}

выведет:
1
2
3
4

========================================
package main

import "fmt"

func main() {
    var a, b, c int

    // fmt.Println("Ввод значений a, b, c:")
    fmt.Scanln(&a, &b, &c)

    // Вывод суммы и произведения
    fmt.Println(a + b + c)
    fmt.Println(a * b * c)
}


========================================
ROBOT Number teller (5 - FIVE)

package main

import "fmt"

func main() {
    var a int
    for i := 0; i < 3; i++ {
        // Считываем число с клавиатуры
        fmt.Scanln(&a)
        // Выводим число на русском языке
        switch a {
        case 0:
            fmt.Println("Ноль")
        case 1:
            fmt.Println("Один")
        case 2:
            fmt.Println("Два")
        case 3:
            fmt.Println("Три")
        case 4:
            fmt.Println("Четыре")
        case 5:
            fmt.Println("Пять")
        case 6:
            fmt.Println("Шесть")
        case 7:
            fmt.Println("Семь")
        case 8:
            fmt.Println("Восемь")
        case 9:
            fmt.Println("Девять")
        case 10:
            fmt.Println("Десять")
        }
    }
}

OR PERHAPS:

package main

import "fmt"

func main() {
    var firstNumber, secondNumber, thirdNumber int

    fmt.Scan(&firstNumber)
    fmt.Scan(&secondNumber)
    fmt.Scan(&thirdNumber)

    numbers := map[int]string{
        0: "Ноль",
        1: "Один",
        2: "Два",
        3: "Три",
        4: "Четыре",
        5: "Пять",
        6: "Шесть",
        7: "Семь",
        8: "Восемь",
        9: "Девять",
        10: "Десять",
    }

    fmt.Println(numbers[firstNumber])
    fmt.Println(numbers[secondNumber])
    fmt.Println(numbers[thirdNumber])
}

========================================
SWAPPER

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
========================================

========================================

*/
