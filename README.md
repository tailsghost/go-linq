# Go.Linq

**Go.Linq** – это LINQ-подобная библиотека для языка Go.

## Установка

```bash
go get github.com/tailsghost/go-linq@v0.1.2
```

## Импорт

```go
import Enumerable "github.com/tailsghost/go-linq"
```

## Основные методы и функции

- `Range(start int, count int) Enumerable[int]`  
  Создаёт срез чисел от `start` длины `count`.

- `Repeat[T any](elem T, count int) Enumerable[T]`  
  Повторяет элемент `elem` `count` раз.

- `From[T any](slice []T) Enumerable[T]`  
  Клонирует и создаёт Enumerable из существующего среза.

- `Empty[T any]() Enumerable[T]`  
  Возвращает пустую Enumerable.

- `ToSlice() []T`  
  Преобразует Enumerable в срез.

- `ToMap(keySel func(T) K, valSel func(T) V) map[K]V`  
  Преобразует в `map` по ключу и значению.

- `Count() int`  
  Количество элементов в последовательности.

- `Where(pred func(T) bool) Enumerable[T]`  
  Возвращает только те элементы, которые удовлетворяют условию.

- `Select[U any](sel func(T) U) Enumerable[U]`  
  Проецирует элементы в новый тип.

- `Skip(n int) Enumerable[T]`  
  Пропускает первые `n` элементов.

- `Take(n int) Enumerable[T]`  
  Берёт первые `n` элементов.

- `Any(pred func(T) bool) bool`  
  Проверяет, есть ли хотя бы один элемент, удовлетворяющий условию.

- `All(pred func(T) bool) bool`  
  Проверяет, все ли элементы удовлетворяют условию.

- `First(pred func(T) bool) (T, error)`  
  Первый элемент по условию или ошибка

- `FirstOrDefault(pred func(T) bool, defaultValue T) T`  
  Первый элемент или значение по умолчанию.

- `Last(pred func(T) bool) (T, error)`  
  Последний элемент по условию.

- `LastOrDefault(pred func(T) bool, defaultValue T) T`  
  Последний элемент или значение по умолчанию.

- `ForEach(action func(T))`  
  Выполняет действие для каждого элемента.

- `Min(sel func(T) float64) (float64, error)`  
  Минимальное значение по селектору.

- `Max(sel func(T) float64) (float64, error)`  
  Максимальное значение.

- `Sum(sel func(T) float64) float64`  
  Сумма значений.

- `Reverse() Enumerable[T]`  
  Разворачивает Enumerable.

- `Distinct() Enumerable[T]`  
  Убирает дубликаты.

- `Cast[U any](castFunc func(T) U) Enumerable[U]`  
  Приводит элементы к другому типу.

- `OrderBy(key func(T) K) Enumerable[T]`  
  Сортирует по возрастанию.

- `OrderByDescending(key func(T) K) Enumerable[T]`  
  Сортирует по убыванию.

- `ThenBy(key func(T) K) Enumerable[T]`  
  Вторичная сортировка по возрастанию.

- `ThenByDescending(key func(T) K) Enumerable[T]`  
  Вторичная сортировка по убыванию.

- `GroupBy(key func(T) K) map[K][]T`  
  Группирует элементы по ключу.

- `Join(other Enumerable[U], outerKey func(T) K, innerKey func(U) K, resultSel func(T, U) R) Enumerable[R]`  
  Соединяет две последовательности по ключам.

## Примеры
```
package main

import (
	"fmt"

	Enumerable "github.com/tailsghost/go-linq"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	nums := Enumerable.Range(1, 10).
		Where(func(x int) bool { return x%2 == 0 })

	results := Enumerable.Select(nums, func(x int) int { return x * x })

	fmt.Println("Squares of even numbers:", results.ToSlice())

	people := []Person{
		{"Alice", 30},
		{"Bob", 20},
		{"Charlie", 30},
		{"Dave", 20},
	}

	grouped := Enumerable.GroupBy(Enumerable.From(people), func(p Person) int {
		return p.Age
	})
	fmt.Println("Grouped by age:", grouped)
}

```
