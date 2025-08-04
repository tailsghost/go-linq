# Go.Linq

**Go.Linq** – это LINQ-подобная библиотека для языка Go.

## Установка

```bash
go get github.com/tailsghost/go-linq
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
  Количество элементов в Enumerable.

- `Where(pred func(T) bool) Enumerable[T]`  
  Возвращает только те элементы, которые удовлетворяют условию.

- `Select[U any](sel func(T) U) Enumerable[U]`  
  Проецирует элементы в новый тип.

- `func SelectMany[T any, U any](e Enumerable[T], sel func(T) []U) Enumerable[U]`
  Сводит коллекции в одну


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
  Соединяет два Enumerable по ключам.

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

type Company struct {
	Name   string
	Person Person
}

func main() {
	nums := Enumerable.Select(Enumerable.Range(1, 10).
		Where(func(x int) bool { return x%2 == 0 }), func(x int) any { return x * x }).ToSlice()

	fmt.Println("Squares of even numbers:", nums)

	people := []Person{
		{"Alice", 30},
		{"Bob", 20},
		{"Charlie", 30},
		{"Dave", 20},
	}

	selected := Enumerable.
		Select(Enumerable.From(people), func(p Person) Company {
			return Company{Name: "google", Person: p}
		})

	orders := Enumerable.OrderBy(selected, func(a Company) string { return a.Person.Name }, func(a, b string) int { return Enumerable.StringCmp(a, b) })

	groups := Enumerable.GroupBy(orders, func(p Company) int { return p.Person.Age })

	fmt.Println("Grouped by age:", groups)
}

```
