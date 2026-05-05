// Задание 4: Декоратор - добавить поведение к функции
//
// Замыкание - отличный способ "обернуть" функцию, добавив поведение
// без изменения оригинала. Это паттерн декоратор (decorator).
//
// Напиши функцию withLogging(fn func(int) int, name string) func(int) int,
// которая:
//   - принимает любую функцию func(int) int и её имя
//   - возвращает новую функцию, которая:
//     1. Выводит: "вызов <name>(<аргумент>)"
//     2. Вызывает оригинальную fn
//     3. Выводит: "<name>(<аргумент>) = <результат>"
//     4. Возвращает результат
//
// Напиши функцию withRetry(fn func() error, attempts int) func() error,
// которая:
//   - принимает функцию без аргументов и количество попыток
//   - возвращает новую функцию, которая вызывает fn до attempts раз
//     пока не получит nil (успех) или не исчерпает все попытки
//   - при каждой неудаче выводит: "попытка N не удалась: <ошибка>"
//   - при успехе выводит: "успех на попытке N"
//   - если все попытки исчерпаны - возвращает последнюю ошибку
//
// Ожидаемый вывод:
//   вызов square(5)
//   square(5) = 25
//   вызов square(7)
//   square(7) = 49
//
//   Тестируем retry:
//   попытка 1 не удалась: сервис недоступен
//   попытка 2 не удалась: сервис недоступен
//   успех на попытке 3
//
// Запусти: go run main.go

package main

import (
	"errors"
	"fmt"
)

// TODO: напиши функцию withLogging(fn func(int) int, name string) func(int) int

func withLogging(fn func(int) int, name string) func(int) int {
	return func(arg int) int {
		fmt.Printf("вызов %s(%d)\n", name, arg)
		result := fn(arg)
		fmt.Printf("%s(%d) = %d\n", name, arg, result)
		return result
	}
}

// TODO: напиши функцию withRetry(fn func() error, attempts int) func() error

func withRetry(fn func() error, attempts int) func() error {
	return func() error {
		var err error
		for i := 1; i <= attempts; i++ {
			err = fn()
			if err == nil {
				fmt.Printf("Успех на попытке %d\n", i)
				return nil
			}
			fmt.Printf("Попытка %d не удалась: %v\n", i, err)
		}
		return err
	}
}

func main() {
	// TODO: создай функцию square := func(n int) int { return n * n }
	square := func(n int) int { return n * n }
	// Оберни её через withLogging и вызови несколько раз
	squareWithLogging := withLogging(square, "square")

	squareWithLogging(5)
	squareWithLogging(7)
	squareWithLogging(8)

	// TODO: создай счётчик попыток
	// attempt := 0
	// и нестабильную функцию, которая возвращает ошибку первые 2 раза,
	// потом nil:
	//   unstable := func() error {
	//       attempt++
	//       if attempt < 3 {
	//           return errors.New("сервис недоступен")
	//       }
	//       return nil
	//   }
	// Оберни через withRetry(unstable, 5) и вызови
	attempt := 0
	unstable := func() error {
		attempt++
		if attempt < 3 {
			return errors.New("сервис недоступен")
		}
		return nil
	}
	fmt.Println("Тестируем retry:")
	unstableWithRetry := withRetry(unstable, 5)
	//unstableWithRetry()

	/*err := unstableWithRetry()*/
	if err := unstableWithRetry(); err != nil {
		fmt.Println("Все попытки провалены, финальная ошибка:", err)
	} else {
		fmt.Println("Функция выполнилась без ошибок!")
	}

}
