// Задание 2: Аккумулятор
//
// Замыкание может не только ЧИТАТЬ захваченную переменную, но и ИЗМЕНЯТЬ её.
// Переменная живёт пока жива функция-замыкание - между вызовами значение сохраняется.
//
// Напиши функцию makeAdder() func(int) int, которая:
//   - хранит внутри переменную sum = 0
//   - каждый вызов прибавляет переданное число к sum
//   - возвращает новое значение sum
//
// Напиши функцию makeAdderWithReset() (func(int) int, func()), которая:
//   - делает то же самое
//   - НО дополнительно возвращает функцию reset(), которая обнуляет sum
//
// Ожидаемый вывод:
//   Первый аккумулятор:
//   +10 -> 10
//   +5  -> 15
//   +3  -> 18
//
//   Второй аккумулятор (независимый от первого):
//   +100 -> 100
//
//   Первый продолжает с 18:
//   +1 -> 19
//
//   Аккумулятор с ресетом:
//   +7 -> 7
//   +3 -> 10
//   reset!
//   +5 -> 5
//
// Запусти: go run main.go

package main

import "fmt"

// TODO: напиши функцию makeAdder() func(int) int
// Подсказка: объяви sum := 0 внутри makeAdder,
// и верни функцию которая меняет sum и возвращает его
func makeAdder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// TODO: напиши функцию makeAdderWithReset() (func(int) int, func())
func makeAdderWithReset() (func(int) int, func()) {
	sum := 0
	return func(x int) int {
			sum += x
			return sum
		},
		func() {
			sum = 0
		}

}

// Подсказка: та же идея, но верни два значения - add и reset.
// Обе функции захватывают одну и ту же переменную sum.

func main() {
	// TODO: создай два независимых аккумулятора через makeAdder()
	// и проверь что они не мешают друг другу
	add1 := makeAdder()
	add2 := makeAdder()
	add, reset := makeAdderWithReset()
	fmt.Println("Первый аккумулятор:")
	fmt.Printf("+10 -> %d\n", add1(10))
	fmt.Printf("+5  -> %d\n", add1(5))
	fmt.Printf("+3  -> %d\n", add1(3))

	fmt.Println("Второй аккумулятор (независимый от первого):")
	fmt.Printf("+100 -> %d\n", add2(100))

	fmt.Println("Первый продолжает с 18:")
	fmt.Printf("+1  -> %d\n", add1(1))

	fmt.Println("Аккумулятор с ресетом:")
	fmt.Printf("+7  -> %d\n", add(7))
	fmt.Printf("+3  -> %d\n", add(3))
	fmt.Println("reset!")
	reset()
	fmt.Printf("+5  -> %d\n", add(5))

	// TODO: создай аккумулятор с ресетом и проверь reset

}
