package main

import "fmt"

type generatorPattern struct {
}

func NewGeneratorPattern() generatorPattern {
	return generatorPattern{}
}

func (g generatorPattern) foo() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func (g generatorPattern) GenerateSomething() {
	fmt.Println("start GenerateSomething ...")
	c := g.foo()

	for el := range c {
		fmt.Println(el)
	}

	fmt.Println("... end GenerateSomething")

}
