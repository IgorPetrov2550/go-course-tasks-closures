// Задание 1: Фабрика умножителей
//
// Замыкание - это функция, которая "помнит" переменные из той области,
// где она была создана.
//
// Представь: ты хочешь функцию "умножить на 2", отдельно "умножить на 3" и т.д.
// Вместо того чтобы писать три разные функции, ты пишешь одну фабрику:
//
//   double := makeMultiplier(2)  // double "помнит" factor=2
//   triple := makeMultiplier(3)  // triple "помнит" factor=3
//
// Когда функция "помнит" переменную из внешней области - это и есть замыкание.
//
// Напиши функцию makeMultiplier(factor int) func(int) int, которая
// возвращает функцию, умножающую любое число на factor.
//
// Ожидаемый вывод:
//   double(5) = 10
//   double(7) = 14
//   triple(5) = 15
//   triple(7) = 21
//   Это одна и та же фабрика, но разные замыкания!
//
// Запусти: go run main.go

package main

import "fmt"

// TODO: напиши функцию makeMultiplier(factor int) func(int) int

func makeMultiplier(factor int) func(int) int {
	return func(n int) int { return n * factor }
}

// Подсказка: внутри верни func(n int) int { return n * factor }
// factor - это и есть то, что "захватывает" замыкание

func main() {
	// TODO: создай double := makeMultiplier(2) и triple := makeMultiplier(3)
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	// TODO: вызови double(5), double(7), triple(5), triple(7) и выведи результаты
	fmt.Println("double(5) = ", double(5))
	fmt.Println("double(7) = ", double(7))
	fmt.Println("triple(5) = ", triple(5))
	fmt.Println("triple(7) = ", triple(7))
	// TODO: добавь fmt.Println("Это одна и та же фабрика, но разные замыкания!")
}
