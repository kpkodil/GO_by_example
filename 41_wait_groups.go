package main

import (
	"fmt"
	"sync"
	"time"
)

// Эта функция будет выполняться в каждой горутине.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Симуляция длительной задачи с помощью sleep.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// WaitGroup используется для ожидания завершения всех
	// запущенных здесь горутин. Замечание: если WaitGroup
	// передается в функции, это должно быть сделано *по указателю*.
	var wg sync.WaitGroup

	// Запускаем несколько горутин и увеличиваем счетчик
	// WaitGroup для каждой из них.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Обертываем вызов worker в замыкание, чтобы гарантировать,
		// что WaitGroup уведомляется о завершении работы. Таким образом,
		// сама функция worker не обязана знать о механизмах конкурентности.
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}

	// Ожидаем, пока счетчик WaitGroup вернется к 0;
	// это значит, что все горутины завершили выполнение.
	wg.Wait()
}

// Пояснение:
// WaitGroup:

// В Go пакет sync предоставляет структуру WaitGroup, которая позволяет ожидать завершения набора горутин. WaitGroup ведет счетчик, который увеличивается при запуске горутины и уменьшается при завершении работы горутины.
// Важно отметить, что WaitGroup должен передаваться в функции по указателю, чтобы изменения были видны в основной программе.

// Цикл запуска горутин:
// В цикле for программа запускает 5 горутин. Перед запуском каждой горутины увеличивается счетчик WaitGroup с помощью метода Add(1), который говорит, что добавлена еще одна горутина.
// Горутину можно запустить, передав в нее замыкание (анонимную функцию), где вызов функции worker(i) происходит асинхронно.
// Внутри замыкания используется defer wg.Done(), чтобы гарантировать вызов метода Done() после завершения работы горутины, что уменьшает счетчик WaitGroup.

// Ожидание завершения:
// Метод wg.Wait() блокирует выполнение основной программы до тех пор, пока счетчик WaitGroup не станет равным нулю, что означает завершение всех горутин.

// Замыкание с параметром:
// Замыкание передает значение i в функцию worker(i) при каждом запуске горутины. Обратите внимание, что параметр i передается в анонимную функцию, чтобы избежать ошибки, связанной с общим состоянием переменной цикла i.