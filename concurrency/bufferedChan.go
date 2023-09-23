package main

import (
	"fmt"
)

type bufferedChan struct {
}

func NewBufferedChan() bufferedChan {
	return bufferedChan{}
}

func (bufferedChan) SimpleBuffered() {
	fmt.Println("start SimpleBuffered ...")

	c := make(chan int, 3)  // Crear un canal con capacidad para 3 valores
	c2 := make(chan int, 3) // Crear un canal con capacidad para 3 valores

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
		}
		close(c2)
	}()

	// Leer valores del canal hasta que esté cerrado
	for {
		value, ok := <-c
		if !ok {
			fmt.Println("canal cerrado")
			break
		}
		fmt.Println("Valor leído:", value)
	}

	// Intentar leer más valores después de que el canal esté cerrado
	value := <-c
	value = <-c
	fmt.Println("Valor después del cierre del canal:", value)

	fmt.Println("... end SimpleBuffered")

}
