// Go поддерживает <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">указатели</a></em>,
// позволяя вам передавать ссылки на значения и записи
// в вашей программе.

package main

import "fmt"

// Мы покажем, как работают указатели в контексте значений с
// помощью двух функций: `zeroval` и `zeroptr`. `zeroval` имеет
// параметр типа `int`, поэтому аргументы будут переданы в неё
// по значению. `zeroval` получит копию `ival`, отличную
// от той, что была в вызывающей функции.
func zeroval(ival int) {
	ival = 0
}

// В отличие от этого, `zeroptr` имеет параметр типа `*int`,
// что означает, что она принимает указатель на `int`. Код
// `*iptr` в теле функции затем _разыменовывает_ указатель с
// его адреса в памяти на текущее значение по этому адресу.
// Присваивание значения разыменованному указателю изменяет
// значение по адресу ссылки.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// Синтаксис `&i` даёт адрес памяти переменной `i`,
	// то есть указатель на `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Указатели также можно печатать.
	fmt.Println("pointer:", &i)
}