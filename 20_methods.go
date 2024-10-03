// Go поддерживает _методы_, определенные для типов структур.

package main

import "fmt"

type rect struct {
	width, height int
}

// Этот метод `area` имеет _тип получателя_ `*rect`.
func (r *rect) area() int {
	return r.width * r.height
}

// Методы могут быть определены как для указателей, так и для значений
// получателей. Вот пример метода с получателем-значением.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Здесь мы вызываем 2 метода, определенные для нашей структуры.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go автоматически обрабатывает преобразование между значениями
	// и указателями для вызовов методов. Вы можете использовать
	// тип указателя-получателя, чтобы избежать копирования при вызовах
	// методов или чтобы метод мог изменять
	// получающую структуру.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}