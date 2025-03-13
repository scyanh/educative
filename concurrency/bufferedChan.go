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
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	go func() {
		for i := 0; i < 5; i++ {
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
	fmt.Println("canal c1 cerrado")

	for v2 := range c2 {
		fmt.Printf("v2=%d\n", v2)
	}
	fmt.Println("Canal c2 cerrado")

	// Intentar leer más valores después de que el canal esté cerrado
	value := <-c
	value = <-c
	value = <-c
	value = <-c
	fmt.Println("Valor después del cierre del canal:", value)

	fmt.Println("... end SimpleBuffered")

}

func (bufferedChan) SimpleBuffered2() {
	c1 := make(chan int, 3)
	c2 := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			c2 <- i
		}
		close(c2)
	}()

	for {
		val, ok := <-c1
		if ok {
			fmt.Println("c1=", val)
		} else {
			fmt.Println("c1 closed")
			break
		}
	}

	for val := range c2 {
		fmt.Println("c2=", val)
	}

	fmt.Println("c1=", <-c1)
	fmt.Println("c2=", <-c2)

}
