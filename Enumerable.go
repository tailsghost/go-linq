package enumerable

import (
	"errors"
	"sort"
)

type Enumerable[T any] struct {
	src           []T
	comparerChain []func(a, b T) int
}

// Создает коллекцию чисел начиная с start, длиной count
func Range(start, count int) Enumerable[int] {
	slice := make([]int, count)
	for i := range count {
		slice[i] = start + i
	}
	return Enumerable[int]{src: slice}
}

// Создает коллекцию elem, которая повторяется count раз
func Repeat[T any](elem T, count int) Enumerable[T] {
	slice := make([]T, count)
	for i := range count {
		slice[i] = elem
	}
	return Enumerable[T]{src: slice}
}

// Создает Enumerable из существующего среза
func From[T any](slice []T) Enumerable[T] {
	clone := make([]T, len(slice))
	copy(clone, slice)
	return Enumerable[T]{src: clone}
}

// Empty возвращает пустой Enumerable
func Empty[T any]() Enumerable[T] {
	return Enumerable[T]{src: []T{}}
}

// Возвращает копию среза
func (e Enumerable[T]) ToSlice() []T {
	out := make([]T, len(e.src))
	copy(out, e.src)
	return out
}

// Преобразует Enumerable в map
func ToMap[T any, K comparable, V any](
	e Enumerable[T],
	keySel func(T) K,
	valSel func(T) V,
) map[K]V {
	m := make(map[K]V, len(e.src))
	for _, x := range e.src {
		m[keySel(x)] = valSel(x)
	}
	return m
}

// Количество элементов в Enumerable
func (e Enumerable[T]) Count() int {
	return len(e.src)
}

// Where фильтрует элементы по условию
func (e Enumerable[T]) Where(pred func(T) bool) Enumerable[T] {
	out := make([]T, 0, len(e.src)/2)
	for _, x := range e.src {
		if pred(x) {
			out = append(out, x)
		}
	}
	return Enumerable[T]{src: out}
}

// Определяет проекцию выбранных значений
func Select[T any, U any](e Enumerable[T], sel func(T) U) Enumerable[U] {
	out := make([]U, 0, len(e.src))
	for _, x := range e.src {
		out = append(out, sel(x))
	}
	return Enumerable[U]{src: out}
}

// Пропускает первые n элементов
func (e Enumerable[T]) Skip(n int) Enumerable[T] {
	if n >= len(e.src) {
		return Empty[T]()
	}
	out := make([]T, len(e.src)-n)
	copy(out, e.src[n:])
	return Enumerable[T]{src: out}
}

// Возвращает первые n элементы
func (e Enumerable[T]) Take(n int) Enumerable[T] {
	if n >= len(e.src) {
		return From(e.src)
	}
	out := make([]T, n)
	copy(out, e.src[:n])
	return Enumerable[T]{src: out}
}

// Возвращает true, если хотя бы один элемент удовлетворяет условию
func (e Enumerable[T]) Any(pred func(T) bool) bool {
	if pred == nil {
		return len(e.src) > 0
	}
	for _, x := range e.src {
		if pred(x) {
			return true
		}
	}
	return false
}

// Возвращает true, если все элементы удовлетворяют условию
func (e Enumerable[T]) All(pred func(T) bool) bool {
	for _, x := range e.src {
		if !pred(x) {
			return false
		}
	}
	return true
}

// Возвращает первый элемент, удовлетворяющий условию
func (e Enumerable[T]) First(pred func(T) bool) (T, error) {
	if pred == nil {
		if len(e.src) == 0 {
			return *new(T), errors.New("Sequence empty")
		}
		return e.src[0], nil
	}

	for _, x := range e.src {
		if pred(x) {
			return x, nil
		}
	}

	return *new(T), errors.New("No match")
}

// Возвращает первый элемент, удовлетворяющий условию иначе default
func (e Enumerable[T]) FirstOrDefault(pred func(T) bool, defaultValue T) T {
	if x, err := e.First(pred); err == nil {
		return x
	}
	return defaultValue
}

// Возвращает последний элемент, удовлетворяющий условию
func (e Enumerable[T]) Last(pred func(T) bool) (T, error) {
	if pred == nil {
		if len(e.src) == 0 {
			return *new(T), errors.New("Sequence empty")
		}
		return e.src[len(e.src)-1], nil
	}

	var found T
	foundAny := false
	for _, x := range e.src {
		if pred(x) {
			found = x
			foundAny = true
		}
	}
	if !foundAny {
		return *new(T), errors.New("No match")
	}
	return found, nil
}

// Возвращает последний элемент, удовлетворяющий условию иначе default
func (e Enumerable[T]) LastOrDefault(pred func(T) bool, defaultValue T) T {
	if x, err := e.Last(pred); err == nil {
		return x
	}
	return defaultValue
}

// Перебирает элементы и с каждым выполняет действие
func (e Enumerable[T]) ForEach(action func(T)) {
	for _, x := range e.src {
		action(x)
	}
}

// Возвращает минимальное значение
func (e Enumerable[T]) Min(sel func(T) float64) (float64, error) {
	if len(e.src) == 0 {
		return 0, errors.New("Empty sequence")
	}
	min := sel(e.src[0])
	for _, x := range e.src[1:] {
		if v := sel(x); v < min {
			min = v
		}
	}

	return min, nil
}

// Возвращает максильное значение
func (e Enumerable[T]) Max(sel func(T) float64) (float64, error) {
	if len(e.src) == 0 {
		return 0, errors.New("Empty sequence")
	}
	max := sel(e.src[0])
	for _, x := range e.src[1:] {
		if v := sel(x); v > max {
			max = v
		}
	}
	return max, nil
}

// Возвращает сумму
func (e Enumerable[T]) Sum(sel func(T) float64) float64 {
	var sum float64
	for _, x := range e.src {
		sum += sel(x)
	}
	return sum
}

// Переворачивает Enumerable
func (e Enumerable[T]) Reverse() Enumerable[T] {
	out := make([]T, len(e.src))
	for i, v := range e.src {
		out[len(e.src)-1-i] = v
	}
	return Enumerable[T]{src: out}
}

// Возвращает новый Enumerable с уникальными элементами
func (e Enumerable[T]) Distinct() Enumerable[T] {
	m := make(map[any]struct{})
	out := make([]T, 0, len(e.src))
	for _, v := range e.src {
		if _, exists := m[v]; !exists {
			m[v] = struct{}{}
			out = append(out, v)
		}
	}
	return Enumerable[T]{src: out}
}

// Преобрает T в U
func (e Enumerable[T]) Cast(castFunc func(T) any) Enumerable[any] {
	out := make([]any, 0, len(e.src))
	for _, v := range e.src {
		out = append(out, castFunc(v))
	}
	return Enumerable[any]{src: out}
}

func BuildComparer[T any, K any](
	key func(T) K,
	cmp func(a, b K) int,
	desc bool,
) func(a, b T) int {
	return func(a, b T) int {
		ka, kb := key(a), key(b)
		c := cmp(ka, kb)
		if desc {
			return -c
		}
		return c
	}
}

// OrderBy сортирует элементы по возрастанию
func (e Enumerable[T]) OrderBy(
	key func(T) any,
	cmp func(a, b any) int,
) Enumerable[T] {
	comparer := BuildComparer(key, cmp, false)
	out := e.ToSlice()
	sort.Slice(out, func(i, j int) bool { return comparer(out[i], out[j]) < 0 })
	return Enumerable[T]{src: out, comparerChain: []func(a, b T) int{comparer}}
}

// OrderByDescending сортирует элементы по убыванию
func (e Enumerable[T]) OrderByDescending(
	key func(T) any,
	cmp func(a, b any) int,
) Enumerable[T] {
	comparer := BuildComparer(key, cmp, true)
	out := e.ToSlice()
	sort.Slice(out, func(i, j int) bool { return comparer(out[i], out[j]) < 0 })
	return Enumerable[T]{src: out, comparerChain: []func(a, b T) int{comparer}}
}

// ThenBy добавляет вторичный ключ сортировки по возрастанию
func (e Enumerable[T]) ThenBy(
	key func(T) any,
	cmp func(a, b any) int,
) Enumerable[T] {
	if e.comparerChain == nil {
		panic("ThenBy requires OrderBy first")
	}
	next := BuildComparer(key, cmp, false)
	chain := append([]func(a, b T) int{}, e.comparerChain...)
	chain = append(chain, next)
	out := e.ToSlice()
	sort.Slice(out, func(i, j int) bool {
		for _, c := range chain {
			if r := c(out[i], out[j]); r != 0 {
				return r < 0
			}
		}
		return false
	})
	return Enumerable[T]{src: out, comparerChain: chain}
}

// ThenByDescending добавляет вторичный ключ сортировки по убыванию
func (e Enumerable[T]) ThenByDescending(
	key func(T) any,
	cmp func(a, b any) int,
) Enumerable[T] {
	if e.comparerChain == nil {
		panic("ThenByDescending requires OrderBy first")
	}
	next := BuildComparer(key, cmp, true)
	chain := append([]func(a, b T) int{}, e.comparerChain...)
	chain = append(chain, next)
	out := e.ToSlice()
	sort.Slice(out, func(i, j int) bool {
		for _, c := range chain {
			if r := c(out[i], out[j]); r != 0 {
				return r < 0
			}
		}
		return false
	})
	return Enumerable[T]{src: out, comparerChain: chain}
}

// Группирует элементы по ключу в map.
func (e Enumerable[T]) GroupBy(key func(T) any) map[any][]T {
	m := make(map[any][]T)
	for _, x := range e.src {
		m[key(x)] = append(m[key(x)], x)
	}
	return m
}

// Выполняет соединение двух Enumerable по ключам
func (e Enumerable[T]) Join(
	other Enumerable[any],
	outerKey func(T) any,
	innerKey func(any) any,
	resultSel func(T, any) any,
) Enumerable[any] {
	out := make([]any, 0)
	inSlice := other.ToSlice()
	for _, o := range e.src {
		ok := outerKey(o)
		for _, i := range inSlice {
			if innerKey(i) == ok {
				out = append(out, resultSel(o, i))
			}
		}
	}
	return Enumerable[any]{src: out}
}
