// _Функции_ играют центральную роль в Go. Мы рассмотрим функции
// с помощью нескольких различных примеров.

package main

import "fmt"

// Вот функция, которая принимает два `int` и возвращает
// их сумму в виде `int`.
func plus(a int, b int) int {

	// Go требует явных возвратов, то есть не будет
	// автоматически возвращать значение последнего
	// выражения.
	return a + b
}

// Когда у вас несколько последовательных параметров одного
// типа, вы можете опустить имя типа для параметров того же типа,
// кроме последнего параметра, который указывает тип.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Вызов функции осуществляется так же, как вы и ожидали, с
	// помощью `name(args)`.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}